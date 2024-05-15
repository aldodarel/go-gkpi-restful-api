package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// Jemaat is a representation of a congregation
type Jemaat struct {
	ID            int       `gorm:"primary_key" form:"id" json:"id"`
	Nik           string    `gorm:"type:varchar(16)" form:"nik" json:"nik"`
	Username      string    `gorm:"type:varchar(50)" form:"username" json:"username"`
	Nama          string    `gorm:"type:varchar(255)" form:"nama" json:"nama"`
	JenisKelamin  string    `gorm:"type:enum('Laki-laki', 'Perempuan')" form:"jenis_kelamin" json:"jenis_kelamin"`
	Password      string    `gorm:"type:varchar(255)" form:"password" json:"password"`
	Alamat        string    `gorm:"type:varchar(255)" form:"alamat" json:"alamat"`
	TempatLahir   string    `gorm:"type:varchar(255)" form:"tempat_lahir" json:"tempat_lahir"`
	StatusGereja  string    `gorm:"type:enum('Aktif','Pindah','Meninggal')" form:"status_gereja" json:"status_gereja"`
	StatusNikah   string    `gorm:"type:enum('Menikah','Belum Menikah','Cerai Mati')" form:"status_nikah" json:"status_nikah"`
	TanggalLahir  time.Time `gorm:"type:date" form:"tanggal_lahir" json:"tanggal_lahir"`
	StatusBaptis  string    `gorm:"type:enum('Ya','Tidak')" form:"status_baptis" json:"status_baptis"`
	StatusSidi    string    `gorm:"type:enum('Ya','Tidak')" form:"status_sidi" json:"status_sidi"`
	IDSektor      int       `gorm:"type:int(11)" form:"id_sektor" json:"id_sektor"`
	SektorRole    string    `gorm:"type:enum('Penanggung Jawab','Anggota')" form:"sektor_role" json:"sektor_role"`
	GambarProfile string    `gorm:"type:varchar(255)" form:"gambar_profile" json:"gambar_profile"`
	Lampiran      string    `gorm:"type:text" form:"lampiran" json:"lampiran"`
	NoTelepon     string    `gorm:"type:varchar(20)" form:"no_telepon" json:"no_telepon"`
}

// Result is an array of congregation => Respon berhasil/gagal, dll yang dikirim oleh API
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Main
func main() {
	db, err = gorm.Open("mysql", "root:@tcp/go_restapi_jemaat?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	db.AutoMigrate(&Jemaat{})
	handleRequests()
}

// handleRequests handles all API endpoints
func handleRequests() {
	log.Println("Start the development server at http://127.0.0.1:9977")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := Result{Code: 404, Message: "Method not found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := Result{Code: 403, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/jemaats", createJemaats).Methods("POST")
	myRouter.HandleFunc("/api/jemaats", getJemaats).Methods("GET")
	myRouter.HandleFunc("/api/jemaats/{id}", getJemaat).Methods("GET")
	myRouter.HandleFunc("/api/jemaats/{id}", updateJemaat).Methods("PUT")
	myRouter.HandleFunc("/api/jemaats/{id}", deleteJemaat).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9977", myRouter))
}

// homePage handles requests to the root endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// createJemaats handles creating a new congregation
func createJemaats(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jemaatData := r.Form // Access form data

	// Parse and validate date format for TanggalLahir
	tanggalLahirStr := jemaatData.Get("tanggal_lahir")
	tglLahir, err := time.Parse("2006-01-02", tanggalLahirStr)
	if err != nil {
		http.Error(w, "Invalid date format for tanggal_lahir", http.StatusBadRequest)
		return
	}

	// Handle file upload for Gambar Profile
	file, handler, err := r.FormFile("gambar_profile")
	if err == nil {
		defer file.Close()

		// Save uploaded file to server or process it accordingly
		// For simplicity, let's just print the file name here
		fmt.Println("Uploaded file:", handler.Filename)
	}

	// Create Jemaat object with parsed data
	jemaat := Jemaat{
		Nik:           jemaatData.Get("nik"),
		Username:      jemaatData.Get("username"),
		Nama:          jemaatData.Get("nama"),
		JenisKelamin:  jemaatData.Get("jenis_kelamin"),
		Password:      jemaatData.Get("password"),
		Alamat:        jemaatData.Get("alamat"),
		TempatLahir:   jemaatData.Get("tempat_lahir"),
		StatusGereja:  jemaatData.Get("status_gereja"),
		StatusNikah:   jemaatData.Get("status_nikah"),
		TanggalLahir:  tglLahir,
		StatusBaptis:  jemaatData.Get("status_baptis"),
		StatusSidi:    jemaatData.Get("status_sidi"),
		IDSektor:      parseID(jemaatData.Get("id_sektor")),
		SektorRole:    jemaatData.Get("sektor_role"),
		GambarProfile: handler.Filename,
		Lampiran:      jemaatData.Get("lampiran"),
		NoTelepon:     jemaatData.Get("no_telepon"),
	}

	// Save Jemaat object to database
	db.Create(&jemaat)

	// Prepare response
	res := Result{Code: 200, Data: jemaat, Message: "Success create congregation"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getJemaats(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get congregations")

	jemaats := []Jemaat{}
	db.Find(&jemaats)

	res := Result{Code: 200, Data: jemaats, Message: "Success get congregations"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getJemaat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jemaatID := vars["id"]

	var jemaat Jemaat

	db.First(&jemaat, jemaatID)

	res := Result{Code: 200, Data: jemaat, Message: "Success get congregation"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateJemaat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jemaatID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var jemaatUpdates Jemaat
	json.Unmarshal(payloads, &jemaatUpdates)

	var jemaat Jemaat
	db.First(&jemaat, jemaatID)
	db.Model(&jemaat).Updates(jemaatUpdates)

	res := Result{Code: 200, Data: jemaat, Message: "Success update congregation"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteJemaat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jemaatID := vars["id"]

	var jemaat Jemaat

	db.First(&jemaat, jemaatID)
	db.Delete(&jemaat)

	res := Result{Code: 200, Message: "Success delete congregation"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// parseID converts string ID to int
func parseID(idStr string) int {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0 // default value if conversion fails
	}
	return id
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// PelayanGereja is a representation of a family
type PelayanGereja struct {
	ID               int    `gorm:"primary_key" form:"id" json:"id"`
	NIK              string `form:"nik" json:"nik"`
	Peran            string `gorm:"type:enum('Pendeta','Penatua', 'PHJ', 'Pelayan_Ibadah', 'Tata Usaha')" form:"peran" json:"peran"`
	TglTerimaJabatan string `form:"tanggal_terima_jabatan" json:"tanggal_terima_jabatan"`
	TglAkhirJabatan  string `form:"tanggal_akhir_jabatan" json:"tanggal_akhir_jabatan"`
}

// Result is an array of family => Respon berhasil/gagal, dll yang dikirim oleh API
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Main
func main() {
	db, err = gorm.Open("mysql", "root:@tcp/go_restapi_pelayangereja?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	db.AutoMigrate(&PelayanGereja{})
	handleRequests()
}

func handleRequests() {
	log.Println("Start the development server at http://127.0.0.1:9999")
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
	myRouter.HandleFunc("/api/pelayangerejas", createPelayans).Methods("POST")
	myRouter.HandleFunc("/api/pelayangerejas", getPelayans).Methods("GET")
	myRouter.HandleFunc("/api/pelayangerejas/{id}", getPelayan).Methods("GET")
	myRouter.HandleFunc("/api/pelayangerejas/{id}", updatePelayan).Methods("PUT")
	myRouter.HandleFunc("/api/pelayangerejas/{id}", deletePelayan).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// createPelayans handles creating a new congregation
func createPelayans(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pelayanData := r.Form // Access form data

	// Parse and validate date format for TanggalLahir
	// tanggal_terima_jabatanStr := pelayanData.Get("tanggal_terima_jabatan")
	// tgl_terima_jabatan, err := time.Parse("2006-01-02", tanggal_terima_jabatanStr)
	// if err != nil {
	// 	http.Error(w, "Invalid date format for tanggal_terima_jabatan", http.StatusBadRequest)
	// 	return
	// }
	// Parse and validate date format for TanggalLahir
	// tanggal_akhir_jabatanStr := pelayanData.Get("tanggal_akhir_jabatan")
	// tgl_akhir_jabatan, err := time.Parse("2006-01-02", tanggal_akhir_jabatanStr)
	// if err != nil {
	// 	http.Error(w, "Invalid date format for tanggal_akhir_jabatan", http.StatusBadRequest)
	// 	return
	// }

	// Handle file upload for Gambar Profile
	// file, handler, err := r.FormFile("gambar_profile")
	// if err == nil {
	// 	defer file.Close()

	// 	// Save uploaded file to server or process it accordingly
	// 	// For simplicity, let's just print the file name here
	// 	fmt.Println("Uploaded file:", handler.Filename)
	// }

	// Create PelayanGereja object
	pelayan := PelayanGereja{
		NIK:              pelayanData.Get("nik"),
		Peran:            pelayanData.Get("peran"),
		TglTerimaJabatan: pelayanData.Get("tanggal_terima_jabatan"),
		TglAkhirJabatan:  pelayanData.Get("tanggal_akhir_jabatan"),
	}

	// Save PelayanGereja object to database
	db.Create(&pelayan)

	// Prepare response
	res := Result{Code: 200, Data: pelayan, Message: "Success create church minister"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getPelayans(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get church minister")

	pelayans := []PelayanGereja{}
	db.Find(&pelayans)

	res := Result{Code: 200, Data: pelayans, Message: "Success get church minister"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getPelayan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pelayanID := vars["id"]

	var pelayan PelayanGereja

	db.First(&pelayan, pelayanID)

	res := Result{Code: 200, Data: pelayan, Message: "Success get church minister"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updatePelayan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pelayanID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var pelayanUpdates PelayanGereja
	json.Unmarshal(payloads, &pelayanUpdates)

	var pelayan PelayanGereja
	db.First(&pelayan, pelayanID)
	db.Model(&pelayan).Updates(pelayanUpdates)

	res := Result{Code: 200, Data: pelayan, Message: "Success update church minister"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deletePelayan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pelayanID := vars["id"]

	var pelayan PelayanGereja

	db.First(&pelayan, pelayanID)
	db.Delete(&pelayan)

	res := Result{Code: 200, Message: "Success delete church minister"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

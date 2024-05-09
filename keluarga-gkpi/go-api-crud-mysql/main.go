package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// Keluarga is a representation of a family
type Keluarga struct {
	ID           int    `form:"id" json:"id"`
	NoKK         string `form:"no_kk" json:"no_kk"`
	NamaKeluarga string `form:"nama_keluarga" json:"nama_keluarga"`
	Alamat       string `form:"alamat" json:"alamat"`
	Status		 string `gorm:"type:enum('Aktif','Pindah', 'Disabled')" form:"status" json:"status"`
	TglNikah     time.Time `form:"tgl_nikah" json:"tgl_nikah"`
	Lampiran     string `form:"lampiran" json:"lampiran"`
}

// Result is an array of family => Respon berhasil/gagal, dll yang dikirim oleh API
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Main
func main() {
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3308)/go_restapi_gereja?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	db.AutoMigrate(&Keluarga{})
	handleRequests()
}

func handleRequests() {
	log.Println("Start the development server at http://127.0.0.1:9999")
	// log.Println("Benyamin")

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
	myRouter.HandleFunc("/api/keluargas", createKeluarga).Methods("POST")
	myRouter.HandleFunc("/api/keluargas", getKeluargas).Methods("GET")
	myRouter.HandleFunc("/api/keluargas/{id}", getKeluarga).Methods("GET")
	myRouter.HandleFunc("/api/keluargas/{id}", updateKeluarga).Methods("PUT")
	myRouter.HandleFunc("/api/keluargas/{id}", deleteKeluarga).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func createKeluarga(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    keluargaData := r.Form // Access form data

    // Parse and validate date format for TglNikah
    tglNikahStr := keluargaData.Get("tgl_nikah")
    tglNikah, err := time.Parse("2006-01-02", tglNikahStr)
    if err != nil {
        http.Error(w, "Invalid date format for tgl_nikah", http.StatusBadRequest)
        return
    }

    // Handle file upload for Lampiran
    file, handler, err := r.FormFile("lampiran")
    if err == nil {
        defer file.Close()
        
        // Save uploaded file to server or process it accordingly
        // For simplicity, let's just print the file name here
        fmt.Println("Uploaded file:", handler.Filename)
    }

    // Create Keluarga object with parsed data
    keluarga := Keluarga{
        NoKK:         keluargaData.Get("no_kk"),
        NamaKeluarga: keluargaData.Get("nama_keluarga"),
        Alamat:       keluargaData.Get("alamat"),
		Status: 	  keluargaData.Get("status"),	
        TglNikah:     tglNikah,
        Lampiran:     handler.Filename, // You may adjust this based on your file handling logic
    }

	// 	Status		 string `gorm:"type:enum('Aktif','Pindah', 'Disabled')" form:"status" json:"status"`

    // Save Keluarga object to database
    db.Create(&keluarga)

    // Prepare response
    res := Result{Code: 200, Data: keluarga, Message: "Success create family"}
    result, err := json.Marshal(res)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(result)
}


func getKeluargas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get families")

	keluargas := []Keluarga{}
	db.Find(&keluargas)

	res := Result{Code: 200, Data: keluargas, Message: "Success get families"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getKeluarga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keluargaID := vars["id"]

	var keluarga Keluarga

	db.First(&keluarga, keluargaID)

	res := Result{Code: 200, Data: keluarga, Message: "Success get family"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateKeluarga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keluargaID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var keluargaUpdates Keluarga
	json.Unmarshal(payloads, &keluargaUpdates)

	var keluarga Keluarga
	db.First(&keluarga, keluargaID)
	db.Model(&keluarga).Updates(keluargaUpdates)

	res := Result{Code: 200, Data: keluarga, Message: "Success update family"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteKeluarga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keluargaID := vars["id"]

	var keluarga Keluarga

	db.First(&keluarga, keluargaID)
	db.Delete(&keluarga)

	res := Result{Code: 200, Message: "Success delete family"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

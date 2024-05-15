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

// Sektor is a representation of a sector
type Sektor struct {
	ID         int    `form:"id" json:"id"`
	Nama       string `form:"nama" json:"nama"`
	Keterangan string `form:"keterangan" json:"keterangan"`
}

// Result is an array of sector => Respon berhasil/gagal, dll yang dikirim oleh API
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Main
func main() {
	db, err = gorm.Open("mysql", "root:@tcp/go_restapi_sektor?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	db.AutoMigrate(&Sektor{})
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
	myRouter.HandleFunc("/api/sektors", createSektor).Methods("POST")
	myRouter.HandleFunc("/api/sektors", getSektors).Methods("GET")
	myRouter.HandleFunc("/api/sektors/{id}", getSektor).Methods("GET")
	myRouter.HandleFunc("/api/sektors/{id}", updateSektor).Methods("PUT")
	myRouter.HandleFunc("/api/sektors/{id}", deleteSektor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func createSektor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sektorData := r.Form // Access form data

	// Create Sektor object with parsed data
	sektor := Sektor{
		Nama:       sektorData.Get("nama"),
		Keterangan: sektorData.Get("keterangan"),
	}

	// Save Sektor object to database
	db.Create(&sektor)

	// Prepare response
	res := Result{Code: 200, Data: sektor, Message: "Success create sector"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getSektors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get sectors")

	sektors := []Sektor{}
	db.Find(&sektors)

	res := Result{Code: 200, Data: sektors, Message: "Success get sectors"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getSektor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sektorID := vars["id"]

	var sektor Sektor

	db.First(&sektor, sektorID)

	res := Result{Code: 200, Data: sektor, Message: "Success get sector"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateSektor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sektorID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var sektorUpdates Sektor
	json.Unmarshal(payloads, &sektorUpdates)

	var sektor Sektor
	db.First(&sektor, sektorID)
	db.Model(&sektor).Updates(sektorUpdates)

	res := Result{Code: 200, Data: sektor, Message: "Success update sector"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteSektor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sektorID := vars["id"]

	var sektor Sektor

	db.First(&sektor, sektorID)
	db.Delete(&sektor)

	res := Result{Code: 200, Message: "Success delete sector"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

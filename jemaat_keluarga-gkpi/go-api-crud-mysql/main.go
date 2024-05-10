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

// Struct which is a representation of JemaatKeluarga
type JemaatKeluarga struct {
	ID           		int        `gorm:"primary_key" form:"id" json:"id"`
	NIK           		string    `form:"nik" json:"nik"`
	NoKK           		string    `form:"no_kk" json:"no_kk"`
	Status         		string 	  `gorm:"type:enum('Suami','Istri', 'Anak', 'Menikah')" form:"status" json:"status"`

}

// Result is an array of JemaatKeluargas => Respon berhasil/gagal, dll yang dikirim oleh API
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Main
func main() {
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3308)/go_restapi_jemaatkeluarga?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	db.AutoMigrate(&JemaatKeluarga{})
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
	myRouter.HandleFunc("/api/jemaatkeluargas", createJemaatKeluargas).Methods("POST")
	myRouter.HandleFunc("/api/jemaatkeluargas", getJemaatKeluargas).Methods("GET")
	myRouter.HandleFunc("/api/jemaatkeluargas/{id}", getJemaatKeluarga).Methods("GET")
	myRouter.HandleFunc("/api/jemaatkeluargas/{id}", updateJemaatKeluargas).Methods("PUT")
	myRouter.HandleFunc("/api/jemaatkeluargas/{id}", deleteJemaatKeluargas).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}





// createJemaatKeluargas handles creating a new congregation
func createJemaatKeluargas(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jemaatKeluargaData := r.Form // Access form data


	// Create JemaatKeluarga object
	JemaatKeluarga := JemaatKeluarga{
		NIK:     jemaatKeluargaData.Get("nik"),
		NoKK:    jemaatKeluargaData.Get("no_kk"),
		Status:	 jemaatKeluargaData.Get("status"),
	}

	// Save JemaatKeluarga object to database
	db.Create(&JemaatKeluarga)

	// Prepare response
	res := Result{Code: 200, Data: JemaatKeluarga, Message: "Success create JemaatKeluarga data"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


func getJemaatKeluargas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get JemaatKeluarga data")

	JemaatKeluargas := []JemaatKeluarga{}
	db.Find(&JemaatKeluargas)

	res := Result{Code: 200, Data: JemaatKeluargas, Message: "Success get JemaatKeluarga data"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getJemaatKeluarga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keluargaJmtID := vars["id"]

	var JemaatKeluarga JemaatKeluarga

	db.First(&JemaatKeluarga, keluargaJmtID)

	res := Result{Code: 200, Data: JemaatKeluarga, Message: "Success get JemaatKeluarga data"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateJemaatKeluargas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keluargaJmtID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var jemaatKeluargaUpdates JemaatKeluarga
	json.Unmarshal(payloads, &jemaatKeluargaUpdates)

	var JemaatKeluarga JemaatKeluarga
	db.First(&JemaatKeluarga, keluargaJmtID)
	db.Model(&JemaatKeluarga).Updates(jemaatKeluargaUpdates)

	res := Result{Code: 200, Data: JemaatKeluarga, Message: "Success update JemaatKeluarga data"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteJemaatKeluargas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keluargaJmtID := vars["id"]

	var JemaatKeluarga JemaatKeluarga

	db.First(&JemaatKeluarga, keluargaJmtID)
	db.Delete(&JemaatKeluarga)

	res := Result{Code: 200, Message: "Success delete JemaatKeluarga data"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// func parseID(idStr string) int {
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return 0 // default value if conversion fail
// 	}
// 	return id
// }
package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"postgres_RestApi/models"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []models.Barang `json:"data"`
}

func GetAllItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	barangS, err := models.GetAllItem()
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	response.Status = 1
	response.Message = "succes"
	response.Data = barangS

	json.NewEncoder(w).Encode(response)
}

func GetOneItem(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get params dari request
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	baranG, err := models.GetOneItem(int64(id))

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(baranG)

}

func AddItem(w http.ResponseWriter, r *http.Request) {

	// create new variabel barang
	var barang models.Barang

	// decode data
	err := json.NewDecoder(r.Body).Decode(&barang)
	if err != nil {
		log.Fatal(err)
	}

	insert := models.AddItem(barang)

	res := response{
		ID:      insert,
		Message: "data berhasil ditambah",
	}

	json.NewEncoder(w).Encode(res)

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// create new barang
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	// membuat banrang baru
	var barang models.Barang

	err = json.NewDecoder(r.Body).Decode(&barang)
	if err != nil {
		log.Fatal(err)
	}

	insertID := models.UpdateItem(int64(id), barang)

	res := response{
		ID:      insertID,
		Message: "berhasil",
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteOneItem(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	deletId := models.DeleteOneItem(int64(id))

	mes := fmt.Sprintf("buku sukses dihapus .total yang dihapus %v", deletId)
	res := response{
		ID:      int64(id),
		Message: mes,
	}

	json.NewEncoder(w).Encode(res)

}

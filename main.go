package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"receipt-processor-challenge/models"
	"receipt-processor-challenge/utils"
)

var receipts = make(map[string]models.Receipt)
var pointsMap = make(map[string]int)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")

	http.Handle("/", r)
	log.Println("Server started at :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	receipts[id] = receipt
	pointsMap[id] = utils.CalculatePoints(receipt)

	response := models.ProcessReceiptResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := pointsMap[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := models.GetPointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

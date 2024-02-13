package main

import (
    "encoding/json"
    "net/http"
    "log"
    "github.com/google/uuid"
	"github.com/gorilla/mux"

)

var pointsData = make(map[string]int)

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Processing new receipt")

    var receipt Receipt
    w.Header().Set("Content-Type", "application/json") // Set the Content-Type header to application/json
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Printf("Error decoding receipt: %v", err)
        return
    }

    log.Println("Receipt decoded successfully")

    // Validate receipt
    if err := validateReceipt(&receipt); err != nil {
        log.Printf("Receipt validation failed: %v", err)
         http.Error(w, "The receipt is invalid", http.StatusBadRequest)
        return
    }
    id := uuid.New().String()

    // Calculate the points
    points := calculatePoints(receipt)
    pointsData[id] = points

    log.Printf("Receipt processed successfully, ID: %s, Points: %d", id, points)

    json.NewEncoder(w).Encode(struct{ ID string `json:"id"` }{ID: id})
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]

    log.Printf("Retrieving points for receipt ID: %s", id)

    w.Header().Set("Content-Type", "application/json") // Set the Content-Type header to application/json
  
    // Validate that ID is a non-whitespace string
    if !validateStringField(id, `^\S+$`) {
        log.Printf("No points found for receipt ID: %s", id)
        http.Error(w, "invalid ID format", http.StatusBadRequest)
        return
    }

    points, ok := pointsData[id]
    if !ok {
        w.WriteHeader(http.StatusNotFound) // Explicitly set the HTTP status code to 404
        json.NewEncoder(w).Encode(map[string]string{"message": "No receipt found for that id"}) // Send the custom error message as JSON    
        return    
    }
    log.Printf("Points retrieved successfully for receipt ID: %s", id)
    json.NewEncoder(w).Encode(PointsResponse{Points: points})}

// Main function to calculate points by applying all rules
func calculatePoints(receipt Receipt) int {
    points := 0
    points += pointsForRetailerName(receipt.Retailer)
    points += pointsForRoundTotal(receipt.Total)
    points += pointsForMultipleOfQuarter(receipt.Total)
    points += pointsForItemPairs(receipt.Items)
    points += pointsForDescriptionLength(receipt.Items)
    points += pointsForPurchaseTime(receipt.PurchaseTime)
	points += pointsForOddDay(receipt.PurchaseDate) 
    return points
}
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/akshat-o5/shopping-complex-api/models"
	"github.com/gorilla/mux"
)

var db = models.ConnectDB()

// Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := map[string]string{
		"message": "Welcome to home page",
	}
	jsonResponse, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Get All Shops in Shopping Complex
func GetShops(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shops []models.Shop
	db.Find(&shops)
	json.NewEncoder(w).Encode(shops)
}

// Get details of one shop in Shopping Complex on basis of their id
func GetShopById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var shop models.Shop
	db.First(&shop, params["id"])
	json.NewEncoder(w).Encode(shop)
}

// Create Shop
func CreateShop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shop models.Shop
	json.NewDecoder(r.Body).Decode(&shop)
	db.Create(&shop)
	json.NewEncoder(w).Encode(&shop)
}

// Change the Rent Status of the shop based on it's id
func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var shop models.Shop
	db.Where("id = ?", params["id"]).First(&shop)
	shop.Status = !shop.Status
	db.Save(&shop)
	json.NewEncoder(w).Encode(shop)
}

// Change the name of existing shop based uopn it's id
func UpdateName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedShop models.Shop
	json.NewDecoder(r.Body).Decode(&updatedShop)

	// Prepare SQL statement
	query := "UPDATE shops SET name = ? WHERE id = ?"
	db.Exec(query, updatedShop.Name, params["id"])
	var shop models.Shop
	db.First(&shop, params["id"])

	json.NewEncoder(w).Encode(shop)
}

// Delete Shop according to it's id
func DeleteShop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	shopID := params["id"]
	query := "DELETE FROM shops WHERE id = ?"
	db.Exec(query, shopID)

}

// Get all shops who have paid their rent
func RentPaid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shops []models.Shop
	db.Where("status = ?", true).Find(&shops)
	json.NewEncoder(w).Encode(shops)
}

// Get all shops who have not paid their rent
func RentNotPaid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shops []models.Shop
	db.Where("status = ?", false).Find(&shops)
	json.NewEncoder(w).Encode(shops)
}

package routes

import (
	"github.com/akshat-o5/shopping-complex-api/controllers"
	"github.com/gorilla/mux"
)

func SetUpRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/getShops", controllers.GetShops).Methods("GET")
	r.HandleFunc("/getShop/{id}", controllers.GetShopById).Methods("GET")
	r.HandleFunc("/createShop", controllers.CreateShop).Methods("POST")
	r.HandleFunc("/updateStatus/{id}", controllers.UpdateStatus).Methods("PUT")
	r.HandleFunc("/updateName/{id}", controllers.UpdateName).Methods("PUT")
	r.HandleFunc("/deleteShop/{id}", controllers.DeleteShop).Methods("DELETE")
	r.HandleFunc("/statusTrue", controllers.RentPaid).Methods("GET")
	r.HandleFunc("/statusFalse", controllers.RentNotPaid).Methods("GET")

	return r
}

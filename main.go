package main

import (
	"net/http"

	//package car
	"github.com/raogopal5/Go_Crud_MySQL/car"
	//package order
	"github.com/raogopal5/Go_Crud_MySQL/order"
)

func main() {
	//Car package
	http.HandleFunc("/", index)
	http.HandleFunc("/cars", car.Index)
	http.HandleFunc("/cars/show", car.Show)
	http.HandleFunc("/cars/create", car.Create)
	http.HandleFunc("/cars/create/process", car.CreateProcess)
	http.HandleFunc("/cars/update", car.Update)
	http.HandleFunc("/cars/update/process", car.UpdateProcess)
	http.HandleFunc("/cars/delete/process", car.DeleteProcess)
	//Order package
	http.HandleFunc("/order", order.Index)
	http.HandleFunc("/order/update", order.Update)
	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/cars", http.StatusSeeOther)
}

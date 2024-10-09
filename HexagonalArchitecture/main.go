package main

import (
	"net/http"

	"github.com/Aritiaya50217/GolangByCodeBangkok/handler"
	"github.com/Aritiaya50217/GolangByCodeBangkok/repository"
	"github.com/Aritiaya50217/GolangByCodeBangkok/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73:3306)/banking?parseTime=true")
	if err != nil {
		panic(err)
	}
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomers).Methods(http.MethodGet)
	http.ListenAndServe(":8000", router)
}

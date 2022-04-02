package app

import (
	"log"
	"net/http"
	"test/domain"
	"test/service"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()
	dbClient := getDbClient()
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))}
	ah := AccountHandler{service: service.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8585", router))
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

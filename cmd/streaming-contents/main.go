package main

import (
	"github.com/gorilla/mux"
	"log"
	"time"
	"net/http"
	"fmt"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", StreamHandler)
	srv := &http.Server{
		Handler: router,
		Addr: ":8000",
		WriteTimeout: 15 * time.Minute,
		ReadTimeout: 15 * time.Minute,
	}
	fmt.Println("Starting server")
	log.Fatal(srv.ListenAndServe())

}

// Stream a json back with 1 million numbers
func StreamHandler(response http.ResponseWriter, request *http.Request){
	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, "{\"numbers\": [")
	fmt.Fprintf(response, "%d", 1)
	for i := 1; i < 1000000000; i++ {
		fmt.Fprintf(response, ",%d", i)
	}
	fmt.Fprintf(response, "]}")
	return
}

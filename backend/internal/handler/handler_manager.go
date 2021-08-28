package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func HandleRequests() {

	http.HandleFunc("/getCategories", GetterCategories)
	r := mux.NewRouter()
	r.HandleFunc("/getCategories", GetterCategories)
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000", WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	/*go func() {*/
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
	/*}()*/

}

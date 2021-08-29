package handler

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

func HandleRequests() {

	// http.HandleFunc("/getCategories", GetterCategories)
	r := mux.NewRouter()
	r.HandleFunc("/getCategories", GetterCategories)
	r.HandleFunc("/getServices", GetterServices)
	http.Handle("/", r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	h := c.Handler(r)
	srv := &http.Server{
		Handler: h,
		Addr:    "127.0.0.1:8000", WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	/*go func() {*/
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
	/*}()*/

}

package main

import (
	"fmt"
	"net/http"

	"github.com/babylearnstocode/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
) 

func main() {
log.SetReportCaller(true) //file and line number
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println(("Starting GO API Server."))

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil{
		log.Error(err)
	}

}
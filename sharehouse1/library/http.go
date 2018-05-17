package library

import (
	"net/http"
	"sharehouse1/library/route"
)

func Run(port string){
	mux := route.Router
	http.ListenAndServe(port,mux)
}
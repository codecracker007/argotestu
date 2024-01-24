package main

import("fmt"
	   "net/http"
)	

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/api/test",uHandler)
	http.ListenAndServe(":8081",mux)
}

func uHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello uo")
}
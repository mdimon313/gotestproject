package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success", "message": "Welcome to Go Test API on Netlify!"}`))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"server": "online", "language": "Go (Golang)"}`))
}

func main() {
	// http.HandleFunc diye regular routing mesh setup kora holo
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/status", statusHandler)

	// Lambda handle korar jonno net/http standard mux proxy kora holo
	lambda.Start(httpadapter.New(http.DefaultServeMux).Proxy)
}

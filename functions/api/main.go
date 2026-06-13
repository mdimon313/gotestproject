package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

// API Endpoint Handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// JSON Response
	w.Write([]byte(`{"status": "success", "message": "Hello! Your Go API is running on Netlify!"}`))
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	// net/http রাউটারটিকে AWS Lambda (Netlify) এর উপযোগী করে প্রক্সি করা হচ্ছে
	lambda.Start(httpadapter.New(http.DefaultServeMux).Proxy)
}

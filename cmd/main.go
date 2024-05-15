package main

import (
	"fmt"
	"gopipe/internal/dataingestion"
	"gopipe/internal/gopipeauth"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ReceiveData(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Println("Got request from ", r.RemoteAddr, "\nType: ", r.Method, " \nContent: ", string(body[:]))
	userIdStr := r.Header.Get("UserId")
	userId, _ := strconv.Atoi(userIdStr)
	apiKey := r.Header.Get("ApiKey")
	authData := gopipeauth.DataInputAuth{UserId: uint64(userId), ApiKey: apiKey}
	err := gopipeauth.AuthenticateDataInput(authData)
	if err != nil {
		status := http.StatusUnauthorized
		http.Error(w, http.StatusText(status), status)
		return
	}
	if r.Method == "POST" {
		application := r.Header.Get("Application")
		if application == "" {
			application = "Test"
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprintf("Application '%s'", application))
		dataInput := dataingestion.DataIngestionInput{UserId: uint64(userId), Data: body, Application: application, ShouldPipe: r.Header.Get("ShouldPipe")}
		go dataingestion.Ingest(dataInput)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Only accept POST")
		return
	}
}

func main() {
	http.HandleFunc("/data", ReceiveData)
	s := &http.Server{
		Addr:           ":8123",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

// curl -X POST "localhost:8123/data" -H "ApiKey: 123" -H "UserId: 1" -H "Application: River Monitoring" -H "ShouldPipe: false" -d "{\"riverName\":\"MuchWater\",\"riverIdPoint\":123,\"level\":3.5,\"temperature\":30.0}"

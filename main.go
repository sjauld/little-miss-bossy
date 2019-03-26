package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	maximumAllowableCPUUtilisation = 75
	sampleDuration                 = time.Second
)

func main() {
	log.Println("Starting Little Miss Bossy health check server")
	http.HandleFunc("/health", health)
	http.HandleFunc("/coffee", teapot)
	http.ListenAndServe(":65432", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	err, d := newCPUDelta(sampleDuration)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if d.utilisation() > maximumAllowableCPUUtilisation {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("OK: %f", d.utilisation())))
}

func teapot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

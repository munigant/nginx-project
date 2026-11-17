package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
)

func main() {

    hostname, _ := os.Hostname()

    mux := http.NewServeMux()

    mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "application/json")

        json.NewEncoder(w).Encode(map[string]string{
            "status":   "UP",
            "service":  "Go API",
            "instance": hostname,
        })
    })

    mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "application/json")

        json.NewEncoder(w).Encode(map[string]string{
            "instance": hostname,
            "users":    "Employee records",
        })
    })

    log.Printf("Go API running on port 8080 - Instance: %s", hostname)

    log.Fatal(http.ListenAndServe(":8080", mux))
}

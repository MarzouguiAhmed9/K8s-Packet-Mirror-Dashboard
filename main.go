package main

import (
    "fmt"
    "log"
    "net/http"
)

// Logs pour stocker les "paquets"
var logs []string

// Endpoint pour recevoir un "paquet"
func handler(w http.ResponseWriter, r *http.Request) {
    msg := r.URL.Query().Get("msg")
    if msg != "" {
        logs = append(logs, msg)
        fmt.Fprintf(w, "Packet duplicated: %s\n", msg)
    } else {
        fmt.Fprintf(w, "Send a packet using ?msg=your_message\n")
    }
}



// Dashboard pour voir tous les paquets
func dashboard(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h2>Packet Mirror Dashboard</h2>")
    for _, log := range logs {
        fmt.Fprintf(w, "%s<br>", log)
    }
}

func main() {
    http.HandleFunc("/send", handler)
    http.HandleFunc("/dashboard", dashboard)
    fmt.Println("Mirror running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/tarm/serial"
)

type SMSRequest struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

func sendSMS(to, message string) error {
	// Open serial port
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}
	s.Write([]byte("AT\r")) // Initialize modem
	s.Write([]byte(fmt.Sprintf("AT+CMGF=1\r"))) // Set SMS mode
	s.Write([]byte(fmt.Sprintf("AT+CMGS=\"%s\"\r", to))) // Set recipient
	s.Write([]byte(message + string(26))) // Send message with Ctrl+Z to end

	log.Printf("SMS sent to %s: %s", to, message)
	return nil
}

func smsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req SMSRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := sendSMS(req.To, req.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "SMS sent successfully"})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/send-sms", smsHandler)
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

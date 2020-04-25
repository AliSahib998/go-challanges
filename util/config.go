package util

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)
import "github.com/joho/godotenv"

type ErrorResponse struct {
	Code            int
	Message         string
	ValidationError map[string]interface{}
}

func InitLogger() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func InitEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error in loading environment variables from: .env")
	}
}

func ErrorHandler(w http.ResponseWriter, err ErrorResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}

func DecodeJson(r io.Reader, data interface{}) ErrorResponse {
	err := json.NewDecoder(r).Decode(data)
	if err != nil {
		log.Error("json parse exception", err)
		return ErrorResponse{Code: 400, Message: "Parse Exception"}
	}

	return ErrorResponse{}
}

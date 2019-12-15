package globals

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		//panic(err)
		code = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}
func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file GAAN")
	}
	return os.Getenv(key)
}

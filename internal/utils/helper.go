package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}

func ParseDate(input string) (time.Time, error) {
	layout := "2006-01-02"
	parsed, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %w", err)
	}
	return parsed, nil
}

func GetDefaultLimitAmount(tenor int) float64 {
	switch tenor {
	case 1:
		return 100000
	case 2:
		return 200000
	case 3:
		return 500000
	case 6:
		return 700000
	default:
		return 0
	}
}

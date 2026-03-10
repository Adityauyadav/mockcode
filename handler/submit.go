package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"mockcode/db"
)

type SubmitRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type SubmitResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func Submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
	var req SubmitRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	if req.Language == "" || req.Code == "" {
		http.Error(w, "Language and Code is required", http.StatusBadRequest)
		return
	}
	id := uuid.New().String()

	_, err = db.DB.Exec(
		context.Background(),
		"INSERT INTO submissions (id,language,code,status) VALUES ($1,$2,$3,$4)",
		id, req.Language, req.Code, "pending",
	)

	if err != nil {
		http.Error(w, "Failed to store submission", http.StatusInternalServerError)
		return
	}

	resp := SubmitResponse{
		ID:     id,
		Status: "pending",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

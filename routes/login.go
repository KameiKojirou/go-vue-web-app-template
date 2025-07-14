package routes

import (
	"main/utils"
	"net/http"
	"github.com/google/uuid"
)

func Login(w http.ResponseWriter, r *http.Request){
	uuid, _ := uuid.NewV7()
	utils.SetCookie(w, "token", uuid.String())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"token": "` + uuid.String() + `"}`))
}
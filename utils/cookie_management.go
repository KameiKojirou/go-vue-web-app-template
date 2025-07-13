package utils

import (
	"log"
	"net/http"
	"time"
)


func SetCookie(w http.ResponseWriter, name, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: time.Now().Add(24 * time.Hour),
	})
}


func GetCookie(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		log.Println(err)
		return ""
	}
	return cookie.Value
}


func DeleteCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   "",
		Expires: time.Now(),
	})
}
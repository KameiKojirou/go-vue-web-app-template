package routes

import (
	"net/http"
	"main/utils"
)

func Logout(w http.ResponseWriter, r *http.Request){
	utils.DeleteCookie(w, "token")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
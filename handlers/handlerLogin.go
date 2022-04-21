package handlers

import (
	"net/http"
	"packs/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	models.GetUserByCookie(r, w)
}

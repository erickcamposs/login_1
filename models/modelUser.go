package models

import (
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	IdUser    int
	Username  string
	FirstName string
	Lastname  string
	Password  []byte
}

func GetUserByCookie(r *http.Request, w http.ResponseWriter) {
	c, err := r.Cookie("session")
	if err != nil {

		idCookie, err := uuid.NewUUID()
		if err != nil {
			http.Error(w, "Cannot create the cookie", http.StatusInternalServerError)
			return
		}
		c = &http.Cookie{
			Name:  "session",
			Value: idCookie.String(),
		}

		http.SetCookie(w, c)
	}

}

package main

import "packs/db"

func main() {
	//cargar archivos estaticos
	// http.Handle("/static", http.StripPrefix("/static", nil))

	// //Rutas para el login
	// http.HandleFunc("/login", handlers.Login)
	// http.ListenAndServe(":8080", nil)
	db.Open()
	db.Close()
}

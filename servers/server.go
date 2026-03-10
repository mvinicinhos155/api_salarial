package servers

import (
	"log"
	"os"
	"net/http"


)


func Server() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // porta padrão
	}

	log.Println("Servidor rodando 🚀 na porta", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
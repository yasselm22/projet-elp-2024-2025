package main

import (
	"fmt"
	"net"
)

func main() {

	//Connexion au serveur
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Envoi de l'image au serveur
	_, err = conn.Write([]byte("IMAGE"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fermer la connexion
	conn.Close()
}

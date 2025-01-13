package main

import (
	"fmt"
	"net"
)

func main() {

	// Ce serveur écoute sur le port 8000 pour recevoir des requetes TCP
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accepte les connexions
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Lance une go routine pour chaque nouvelle connexion
		go Connection(conn)
	}
}

func Connection(conn net.Conn) {
	// Ferme la connexion à la fin de la go routine

	defer conn.Close()

	// Récupère l'image envoyé par le client
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Appelle la fonction main de notre programme pour traiter l'image
	// Renvoie l'image résultat au client
}

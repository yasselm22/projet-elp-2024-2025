package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	imagePath := "chat.jpg"

	// Lire l'image  // Lire le contenu de l'image depuis le disque dur. Charger ce contenu dans une variable en mémoire, sous forme de tableau d'octets ([]byte), pour qu'il puisse être transmis via la connexion TCP.
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Erreur de lecture de l'image:", err)
		return
	}

	//Connexion au serveur
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Erreur de connexion au serveur", err)
		return
	}
	// Fermer la connexion
	defer conn.Close() // La connexion TCP sera toujours fermée, même en cas d'erreurs, mais à la fin de la fonction main, donc pas de problème pour la reception de l'image

	// Envoi de l'image au serveur
	_, err = conn.Write(imageData) // []byte("IMAGE") au lieu de imageData
	if err != nil {
		fmt.Println("Erreur d'envoi de l'image", err)
		return
	}

	// Réception de l'image traitée
	processedImage := "processed_image.jpg"
	file, err := os.Create(processedImage)
	if err != nil {
		fmt.Println("Erreur de création du fichier local", err)
		return
	}
	defer file.Close() // "Garantie que toutes les données sont écrites sur le disque" : les données en mémoire tampon sont réellement sauvegardées sur le disque. "Le fichier n'est plus verrouillé" : le fichier redevient accessible à d'autres programmes ou parties du code
}

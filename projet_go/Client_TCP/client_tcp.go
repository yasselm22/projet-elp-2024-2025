package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {
	var NomImg string
	var N int

	// Demander au client quelle image il souhaite filtrer
	fmt.Println("Le nom de l'image que vous voulez filtrer (entrez le nom seul sans l'extension) : ")
	fmt.Scanln(&NomImg)
	imagePath := NomImg + ".jpg"

	fmt.Println("Combien de go routines voulez-vous exécuter (veuillez entrer un nombre entier) ? : ")
	fmt.Scanln(&N)

	// Lire le contenu de l'image depuis le disque dur. Charger ce contenu dans une variable en mémoire, sous forme de tableau d'octets ([]byte), pour qu'il puisse être transmis via la connexion TCP.
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Erreur de lecture de l'image:", err)
		return
	}

	// Récupérer la taille de l'image qu'on va envoyer au serveur TCP pour qu'il s'assure que l'image a été reçue dans son entièreté
	imageSize := int64(len(imageData))

	//Connexion au serveur
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Erreur de connexion au serveur", err)
		return
	}
	// Fermer la connexion
	//defer conn.Close() // La connexion TCP sera toujours fermée, même en cas d'erreurs, mais à la fin de la fonction main, donc pas de problème pour la reception de l'image

	//Envoyer la taille de l'image au serveur
	err = binary.Write(conn, binary.BigEndian, imageSize)
	if err != nil {
		fmt.Println("Erreur d'envoi de la taille de l'image : ", err)
		return
	}

	// Envoi du nombres de go routines au serveur
	N_routine := int64(N)
	err = binary.Write(conn, binary.BigEndian, N_routine)
	if err != nil {
		fmt.Println("Erreur d'envoi du nombre de go routines au serveur", err)
		return
	}

	// Envoi de l'image au serveur
	_, err = conn.Write(imageData) // []byte("IMAGE") au lieu de imageData
	if err != nil {
		fmt.Println("Erreur d'envoi de l'image", err)
		return
	}

	// Réception de l'image traitée sous le nom de "processed_image"
	processedImage := "processed_image.jpg"
	file, err := os.Create(processedImage)
	if err != nil {
		fmt.Println("Erreur de création du fichier local", err)
		return
	}
	defer file.Close() // "Garantie que toutes les données sont écrites sur le disque" : les données en mémoire tampon sont réellement sauvegardées sur le disque. "Le fichier n'est plus verrouillé" : le fichier redevient accessible à d'autres programmes ou parties du code

	buffer := make([]byte, 1024) // Buffer pour stocker les données reçues
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break // On arrête de lire le contenu du fichier quand on arrive à sa fin
			}
			fmt.Println("Erreur de lecture de l'image traitée:", err)
			return
		}
		// Ecriture des données reçues dans le fichier "processed_image"
		_, err = file.Write(buffer[:n])
		if err != nil {
			fmt.Println("Erreur d'écriture du fichier:", err)
			return
		}
	}
	fmt.Println("Image traitée reçue et sauvegardée sous", processedImage)
}

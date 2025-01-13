package main

import (
	f "filtre"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	// Ce serveur écoute sur le port 8000 pour recevoir des requetes TCP
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Erreur de création du serveur:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Serveur démarré sur le port 8080")

	// Accepte les connexions
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur de connexion:", err)
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
	img_recue := make([]byte, 1024*1024) // Le tableau img_recue est initialisé avec une taille de 1024*1024 (1Mo).
	n, err := conn.Read(img_recue)       // n=nombre de bytes dans l'image. Si l'image envoyée par le client fait moins de 1Mo, le tableau img_recue contient des données inutiles à la fin
	if err != nil {
		fmt.Println("Erreur de lecture:", err)
		return
	}

	//Supprime les données inutiles à la fin du tableau img_recue si l'image envoyée par le client est inférieure à la taille initialement allouée à img_recue
	img_recue = img_recue[:n]

	// Détecter le format de l'image et sauvegarder l'image dans un fichier local
	imgFormat := detectImageFormat(img_recue)
	imgName := "received_image" + imgFormat

	err = ioutil.WriteFile(imgName, img_recue, 0644)
	if err != nil {
		fmt.Println("Erreur d'enregistrement de l'image:", err)
		return
	}
	fmt.Println("Image sauvegardée sous:", imgName)

	// Appelle la fonction filtre pour traiter l'image envoyée par le client
	f.Filtre(imgName)

	// Lire l'image filtrée
	processedImageData, err := ioutil.ReadFile(imgName)
	if err != nil {
		fmt.Println("Erreur de lecture de l'image filtrée:", err)
		return
	}

	// Renvoie l'image filtrée au client
	_, err = conn.Write(processedImageData)
	if err != nil {
		fmt.Println("Erreur d'envoi de l'image:", err)
	}

	// Supprime l'image originale envoyée par le client
	supImage(err, imgName)

	// Supprime l'image filtrée après l'avoir envoyée au client
	supImage(err, "processed_"+imgName)
}

func detectImageFormat(imgData []byte) string {
	// Vérifie les premiers octets de l'image pour déterminer le format
	if len(imgData) > 4 {
		if imgData[0] == 0xFF && imgData[1] == 0xD8 {
			// JPEG (début d'un fichier JPEG)
			return ".jpg"
		} else if imgData[0] == 0x89 && imgData[1] == 0x50 && imgData[2] == 0x4E && imgData[3] == 0x47 {
			// PNG (début d'un fichier PNG)
			return ".png"
		}
	}
	// Si l'image est dans un format inconnu ou non pris en charge
	return ""
}

func supImage(err error, imgName string) {
	err = os.Remove(imgName)
	if err != nil {
		fmt.Println("Erreur lors de la suppression de l'image :", err)
	} else {
		fmt.Println("Image supprimée:", imgName)
	}

}

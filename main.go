package main

import (
	"encoding/binary"
	f "filtre"
	"fmt"
	"log"
	"net"
	"os"
)

func main() { // Serveur

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
	var expectedImageSize int64
	var N_routines int64
	// Ferme la connexion à la fin de la go routine
	defer conn.Close()

	// Récupère la taille de l'image envoyée par le client et s'assure que l'image entiere a été récupérée
	err1 := binary.Read(conn, binary.BigEndian, &expectedImageSize)
	if err1 != nil {
		log.Printf("Erreur pendant la lecture de la taille de l'image : %v", err1)
		return
	}

	// Récupère nombre de go routines à exécuter
	err2 := binary.Read(conn, binary.BigEndian, &N_routines)
	if err2 != nil {
		fmt.Println("Erreur récupération nombre de goroutines", err2)
		return
	}

	// Récupère l'image envoyée par le client
	img_recue := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Erreur pendant lecture de la data reçue par le serveur :", err)
			break
		}
		img_recue = append(img_recue, buffer[:n]...)
		// Si la taille de ce qui est reçu correspond à la taille de l'image, on arrête de lire
		if int64(len(img_recue)) >= expectedImageSize {
			break
		}
	}
	// Détecte le format de l'image et sauvegarde l'image dans un fichier local sous le nom "received_image"
	imgFormat := detectImageFormat(img_recue)
	fmt.Printf("Format de l'image détecté: %s\n", imgFormat)

	imgName := "received_image" + imgFormat

	err := os.WriteFile(imgName, img_recue, 0644)
	if err != nil {
		fmt.Println("Erreur d'enregistrement de l'image:", err)
		return
	}
	fmt.Println("Image sauvegardée sous:", imgName)

	// Appelle la fonction filtre pour traiter l'image envoyée par le client (création de l'image "resultat")
	N := int(N_routines)
	fmt.Println("Nb go routines : ", N)
	f.Filtre(imgName, N)

	imgResult := "resultat.jpeg"
	// Lit l'image filtrée
	processedImageData, err := os.ReadFile(imgResult)
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
	supImage(err, imgResult)
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

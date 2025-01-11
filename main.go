package main

import (
	"fmt" // package utilisé pour les entrées/sorties en Go
	"image"
	"image/color"
)

/* Fonction main qui lance une boucle for exécutant N fois la routine go, N est en paramètre de la
fonction et pour avoir une exécution optimale, il doit être égal au nombre de threads de
notre ordi  Marine:16 threads*/

func main() {
	var N int
	var img image.Image
	var format string
	var err error
	var matrice_img [][]color.Color

	// Demander N à l'utilisateur
	fmt.Println("Combien de threads possède votre ordinateur? \n = Nombre de routines en concurrence \n = Nombre de bandes découpées dans l'image")

	fmt.Scanln(&N) // Cette fonction attend que l'utilisateur saisisse des données au clavier et appuie sur Entrée.
	// On utilise l'addresse de N &N en argument de la fonction pour que Scanln modifie directement la valeur de N avec la valeur donnée par l'utilisateur.

	//Transforme l'image en une image de format image.Image
	img, format, err = DecodeImage("koala.jpg")
	if err != nil {
		fmt.Println("Erreur lors de l'encodage de l'image:", err)
		return
	}

	// Afficher le format de l'image
	fmt.Println("Format de l'image:", format)

	// Transforme l'image de type image.Image en matrice
	matrice_img = ImageToColorMatrix(img)

	//Découpage de l'image en N bandes
	liste_hauteurs := Decoupe_image(N, matrice_img)

	// Lancement des routines Go pour détecter les contours
	for i := 0; i < N; i++ {
		go EdgeDetection(&matrice_img, liste_hauteurs[i], liste_hauteurs[i+1])
	}
}

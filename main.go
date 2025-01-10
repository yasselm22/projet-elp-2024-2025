package main

import (
	sobel "_/home/marine/Documents/GO/projet-elp-2024-2025/edgeDetection"
	image_coder "_/home/marine/Documents/GO/projet-elp-2024-2025/image_encoder_decoder"
	"fmt"
)

/* Fonction main qui lance une boucle for exécutant N fois la routine go, N est en paramètre de la
fonction et pour avoir une exécution optimale, il doit être égal au nombre de threads de
notre ordi  Marine:16 threads*/

func main() {
	// Demander N à l'utilisateur
	fmt.Println("Combien de threads possède votre ordinateur? \n = Nombre de routines en concurrence \n = Nombre de bandes découpées dans l'image")
	var N int
	fmt.Scanln(&N)

	//Récupérer l'image
	img, format, err = image_coder.EncodeImage()
	if err != nil {
		fmt.Println("Erreur lors de l'encodage de l'image:", err)
		return
	}
	//Découpage de l'image en N bandes
	var liste_hauteurs = [2 * N]int{0}
	liste_hauteurs = sobel.Decoupe_image(N, img)

	// Lancement des routines Go pour détecter les contours
	for i := 0; i < N; i++ {
		go E.EdgeDetection(img, N, liste_hauteurs[i], liste_hauteurs[i+1])
	}
}

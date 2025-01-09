package main

import (
	image_coder "_/C_/Users/elmou/OneDrive/Documents/INSA/3A/ELP/github/projet-elp-2024-2025"
	sobel "_/home/marine/Documents/GO/projet-elp-2024-2025"
	"fmt"
	"image_coder"
	"sobel"
)

/* Fonction main qui lance une boucle for exécutant N fois la routine go, N est en paramètre de la
fonction et pour avoir une exécution optimale, il doit être égal au nombre de threads de
notre ordi  Marine:16 threads*/

func main() {
	print("Combien de threads possède votre ordinateur? \n = Nombre de routines en concurrence \n = Nombre de bandes découpées dans l'image")
	var N int
	fmt.Scanln(&N)

	img, format, nil = image_coder.EncodeImage()
	var liste_hauteurs = [2 * N]int{0}
	liste_hauteurs = sobel.Decoupe_image()
	for i := 0; i < N; i++ {
		go sobel.EdgeDetection(img, N, liste_hauteurs[i], liste_hauteurs[i+1])
	}
}

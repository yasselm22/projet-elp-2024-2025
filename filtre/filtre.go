package filtre

/*POUR LANCER LE PROGRAMME :
  go build -o monprogramme
  ./monprogramme */

import (
	"fmt" // package utilisé pour les entrées/sorties en Go
	"image"
	"image/color"
	"sync"
	"time"
)

/* Fonction main qui lance une boucle for exécutant N fois la routine go, N est en paramètre de la
fonction et pour avoir une exécution optimale, il doit être égal au nombre de threads de
notre ordi  Marine:16 threads*/

func Filtre(filename string) {
	var N int
	var img image.Image
	var newImage image.Image
	var format string
	var err error
	var matrice_img [][]color.Color
	var waitgr sync.WaitGroup

	// Demander N à l'utilisateur
	fmt.Println("Combien de threads possède votre ordinateur? \n = Nombre de routines en concurrence \n = Nombre de bandes découpées dans l'image")

	fmt.Scanln(&N) // Cette fonction attend que l'utilisateur saisisse des données au clavier et appuie sur Entrée.
	// On utilise l'addresse de N &N en argument de la fonction pour que Scanln modifie directement la valeur de N avec la valeur donnée par l'utilisateur.

	//Transforme l'image en une image de format image.Image
	img, format, err = DecodeImage(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'encodage de l'image:", err)
		return
	}

	// Afficher le format de l'image
	fmt.Println("Format de l'image:", format)

	// Transforme l'image de type image.Image en matrice
	matrice_img = ImageToColorMatrix(img)

	//Découpage de l'image en N bandes
	sousMatrices := Decoupe_image(N, matrice_img)

	// Création d'une matrice résultat de la taille de l'image
	result := make([][][]color.Color, N)

	timer1 := time.Now()
	// Lancement des routines Go pour détecter les contours
	for i := 0; i < N; i++ {
		waitgr.Add(1) // On ajoute une tâche au Wait group
		go func(i int) {
			defer waitgr.Done()
			result[i] = EdgeDetection(sousMatrices[i])
		}(i)
	}
	waitgr.Wait() // On attend que toutes les go routines se terminent

	timer2 := time.Now()
	duree := timer2.Sub(timer1)

	// Assembler les sous-matrices traitées
	assembledImage := assemblerSousMatrices(result)

	newImage = MatrixToImage(assembledImage)

	err2 := EncodeImage("resultat."+format, newImage, format)
	if err2 != nil {
		fmt.Println("Error encoding image:", err2)
		return
	}

	fmt.Println("Nouvelle image créée en", duree, "secondes.")
}

func assemblerSousMatrices(sousMatrices [][][]color.Color) [][]color.Color {
	hauteur := 0
	// largeur := len(sousMatrices[0][0])
	for _, sousMat := range sousMatrices {
		hauteur += len(sousMat)
	}

	result := make([][]color.Color, hauteur)
	offset := 0
	for _, sousMat := range sousMatrices {
		copy(result[offset:], sousMat)
		offset += len(sousMat)
	}
	return result
}

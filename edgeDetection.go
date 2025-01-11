package main

import (
	"image/color"
	"math"

	"gonum.org/v1/gonum/mat"
)

// https://medium.com/@damithadayananda/image-processing-with-golang-8f20d2d243a2

func Decoupe_image(N int, img [][]color.Color) []int {

	/* Fonction qui découpe l'image en N bandes horizontales. N est passé en argument et pour
	   une exécution optimale, il doit être égale au nombre de thread de notre ordi.
	   Renvoie les coordonnées de chaque bande.
	   Notre fonction edgeDetection utilise les coordonnées retournées par cette fonction
	   L'image prise en argument doit être une slice de slices de color.Color et non pas de
	   type image.Image
	   liste_hauteurs renvoyée est de taille N+1
	*/

	pimg := img

	hauteur_bande := len(pimg) / N
	liste_hauteurs := make([]int, N+1) // // Créé un tableau dynamique, qui nous permet d'utiliser N
	print("test 11")
	for i := 0; i <= N; i++ {
		print("test 666")
		liste_hauteurs[i] = i * hauteur_bande
	}
	print("test 12")
	return liste_hauteurs
}

// pixels est équivalent à img, &pixels est l'adresse de l'image à traiter

/*Ce processus utilise les filtres de Sobel pour détecter les contours dans une image en niveaux de gris, en
calculant les gradients dans les directions horizontales et verticales, puis en combinant ces gradients pour
déterminer la présence de contours. */

func EdgeDetection(pixels *[][]color.Color, x_haut int, x_bas int) {
	ppixels := *pixels

	//make image grey scale
	for x := x_haut; x <= x_bas; x++ { // On parcourt tous les pixels de limage
		for y := 0; y < len(ppixels[0]); y++ {
			r, g, b, a := ppixels[x][y].RGBA()                             // Pour chaque pixel, on récupère ses composantes rouge, verte, bleue et alpha
			grey := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b) // On calcule une valeur de gris, qui correspond à une pondération des composantes RGB pour refléter la perception humaine des couleurs.
			c := uint8(grey + 0.5)
			ppixels[x][y] = color.RGBA{ // On crée un nouveau pixel avec les trois composantes (R, G, B) égales à cette valeur de gris, ce qui rend l'image en niveaux de gris.
				R: c,
				G: c,
				B: c,
				A: uint8(a),
			}
		}
	}
	// Deux matrices 3×3, kernelx et kernely, sont définies pour détecter les contours horizontaux et verticaux respectivement. Ces matrices sont utilisées pour appliquer des filtres de Sobel, un algorithme courant pour la détection de contours.
	kernelx := mat.NewDense(3, 3, []float64{
		1, 0, 1,
		-2, 0, 2,
		-1, 0, 1,
	})
	kernely := mat.NewDense(3, 3, []float64{
		-1, -2, -1,
		0, 0, 0,
		1, 2, 1,
	})
	//create two dimensional array to store intensities of each pixels
	intensity := make([][]int, len(ppixels))
	for y := 0; y < len(intensity); y++ { // On parcourt l'image pour remplir ce tableau avec les valeurs de gris calculées.
		intensity[y] = make([]int, len(ppixels[0]))
	}
	//calculate intensities : calcule et stocke les valeurs de gris dans un tableau séparé sans modifier l'image, en préparation pour l'application des filtres de Sobel.
	for i := 0; i < len(ppixels); i++ {
		for j := 0; j < len(ppixels[0]); j++ {
			colors := color.RGBAModel.Convert(ppixels[i][j]).(color.RGBA)
			r := colors.R
			g := colors.G
			b := colors.B
			v := int(float64(float64(0.299)*float64(r) + float64(0.587)*float64(g) + float64(0.114)*float64(b)))
			intensity[i][j] = v
		}
	}
	//create new image
	newImage := make([][]color.Color, len(ppixels)) // Un tableau newImage est créé pour stocker l'image après l'application des filtres de Sobel.
	for i := 0; i < len(newImage); i++ {
		newImage[i] = make([]color.Color, len(ppixels[0]))
	}

	for x := 1; x < len(ppixels)-1; x++ { // On parcourt chaque pixel de l'image, sauf ceux sur les bords (pour éviter de sortir des limites lors de la convolution).
		for y := 1; y < len(ppixels[0])-1; y++ {
			var magx, magy int
			for a := 0; a < 3; a++ { // Pour chaque pixel, on applique les noyaux kernelx et kernely pour calculer les gradients horizontal (magx) et vertical (magy).
				for b := 0; b < 3; b++ {
					xn := x + a - 1
					yn := y + b - 1

					// On vérifie que les indices soient dans les limites de l'image
					if xn >= 0 && xn < len(ppixels) && yn >= 0 && yn < len(ppixels[0]) {
						magx += intensity[xn][yn] * int(kernelx.At(a, b))
						magy += intensity[xn][yn] * int(kernely.At(a, b))
					}
				}
			}
			p := int(math.Sqrt(float64(magx*magx + magy*magy))) // La magnitude du gradient est calculée en utilisant la racine carrée de la somme des carrés des gradients horizontal et vertical.
			newImage[x][y] = color.RGBA{                        // Un nouveau pixel est créé avec la valeur de la magnitude comme composante RGB (valeur du contour détecté), et l'alpha est mis à zéro.
				R: uint8(p),
				G: uint8(p),
				B: uint8(p),
				A: 0,
			}
		}
	}
	*pixels = newImage // Le pointeur pixels est mis à jour pour pointer vers la nouvelle image contenant les contours détectés.
}

package filtre

import (
	"image/color"
	"math"

	"gonum.org/v1/gonum/mat"
)

// https://medium.com/@damithadayananda/image-processing-with-golang-8f20d2d243a2

func Decoupe_image(N int, img [][]color.Color) [][][]color.Color {

	/* Fonction qui découpe l'image en N bandes horizontales. N est passé en argument et pour
	   une exécution optimale, il doit être égal au nombre de thread de notre ordi pour un fonctionnement optimal.
	   Renvoie les sous-matrices qui correspondent aux bandes découpées dans l'image.
	   Notre fonction edgeDetection utilise les coordonnées retournées par cette fonction
	   L'image prise en argument doit être une slice de color.Color et non pas de
	   type image.Image

	*/

	hauteur_bande := len(img) / N
	sousMatrices := make([][][]color.Color, N) // Crée un tableau dynamique, qui nous permet d'utiliser N

	for i := 0; i < N; i++ {
		debut := i * hauteur_bande
		fin := debut + hauteur_bande
		if i == N-1 {
			fin = len(img) // Assure que la dernière bande inclut le reste de l'image
		}
		sousMatrices[i] = img[debut:fin]
	}
	return sousMatrices
}

/*Ce processus utilise les filtres de Sobel pour détecter les contours dans une image en niveaux de gris, en
calculant les gradients dans les directions horizontales et verticales, puis en combinant ces gradients pour
déterminer la présence de contours. */

func EdgeDetection(pixels [][]color.Color) [][]color.Color {

	hauteur := len(pixels)
	largeur := len(pixels[0])

	// On crée la matrice résultat
	newImage := make([][]color.Color, hauteur)
	for i := 0; i < hauteur; i++ {
		newImage[i] = make([]color.Color, largeur)
	}

	// Transforme l'image en noir et blanc (gray scale)
	for x := 0; x < hauteur; x++ { // On parcourt tous les pixels de limage
		for y := 0; y < len(pixels[0]); y++ {
			r, g, b, a := pixels[x][y].RGBA()                              // Pour chaque pixel, on récupère ses composantes rouge, verte, bleue et alpha
			grey := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b) // On calcule une valeur de gris, qui correspond à une pondération des composantes RGB pour refléter la perception humaine des couleurs.
			c := uint8(grey + 0.5)
			pixels[x][y] = color.RGBA{ // On crée un nouveau pixel avec les trois composantes (R, G, B) égales à cette valeur de gris, ce qui rend l'image en niveaux de gris.
				R: c,
				G: c,
				B: c,
				A: uint8(a),
			}
		}
	}
	// Deux matrices 3×3, kernelx et kernely, sont définies pour détecter les contours horizontaux et verticaux respectivement. Ces matrices sont utilisées pour appliquer le filtre de Sobel
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
	// Crée deux matrices pour stocker l'intensité de chaque pixel
	intensity := make([][]int, len(pixels))
	for y := 0; y < len(intensity); y++ { // On parcourt l'image pour remplir ce tableau avec les valeurs de gris calculées.
		intensity[y] = make([]int, len(pixels[0]))
	}
	// Calcule les intensités : calcule et stocke les valeurs de gris dans un tableau séparé sans modifier l'image, en préparation pour l'application des filtres de Sobel.
	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[0]); j++ {
			if i >= 0 && i < len(pixels) && j >= 0 && j < len(pixels[0]) {
				colors := color.RGBAModel.Convert(pixels[i][j]).(color.RGBA)
				r := colors.R
				g := colors.G
				b := colors.B
				v := int(float64(float64(0.299)*float64(r) + float64(0.587)*float64(g) + float64(0.114)*float64(b)))
				intensity[i][j] = v
			}
		}
	}

	for x := 1; x < len(pixels)-1; x++ { // On parcourt chaque pixel de l'image, sauf ceux sur les bords (pour éviter de sortir des limites lors de la convolution).
		for y := 1; y < len(pixels[0])-1; y++ {
			var magx, magy int
			for a := 0; a < 3; a++ { // Pour chaque pixel, on applique les noyaux kernelx et kernely pour calculer les gradients horizontal (magx) et vertical (magy).
				for b := 0; b < 3; b++ {
					xn := x + a - 1
					yn := y + b - 1

					// On vérifie que les indices soient dans les limites de l'image
					if xn >= 0 && xn < len(pixels) && yn >= 0 && yn < len(pixels[0]) {
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
	return newImage

}

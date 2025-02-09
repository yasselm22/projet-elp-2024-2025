package filtre

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func ImageToColorMatrix(img image.Image) [][]color.Color {
	/* Fonction qui prend en paramètre un objet de type image.Image
	Extrait la couleur de chaque pixel et la stocke dans une matrice de type color.Color */

	// Obtenir les dimensions de l'image
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Créer une matrice 2D pour stocker les couleurs des pixels
	colorMatrix := make([][]color.Color, height)
	for y := 0; y < height; y++ {
		colorMatrix[y] = make([]color.Color, width)
		for x := 0; x < width; x++ {
			// Extraire la couleur du pixel à la position (x, y)
			colorMatrix[y][x] = img.At(x, y)
		}
	}

	return colorMatrix
}

// MatrixToImage prend une matrice de couleurs ([][]color.Color) et la convertit en une image.Image
func MatrixToImage(matrix [][]color.Color) image.Image {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		fmt.Println("Erreur: la matrice est vide ou mal initialisée")
		return nil
	}

	// Obtenir les dimensions de la matrice
	height := len(matrix)
	width := len(matrix[0])

	// Créer une nouvelle image avec la même taille que la matrice
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Remplir l'image avec les couleurs de la matrice
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			if matrix[y][x] == nil {
				img.Set(x, y, color.Black) // Utilise une couleur par défaut (ici le noir)
				continue
			}

			// On prend la couleur de la matrice à la position (x, y)
			c := matrix[y][x]
			// On définit le pixel de l'image avec la couleur
			img.Set(x, y, c)
		}
	}

	return img
}

func DecodeImage(filename string) (image.Image, string, error) {
	/* Cette fonction prend en argument une image dans un format jpg, png, ...
	et renvoie cette image en type image.Image (qui est un type propre à go)
	ainsi que le format de l'image décodée (jpeg, png,...) */
	file, err1 := os.Open(filename)
	if err1 != nil {
		return nil, "", err1
	}
	defer file.Close()

	img, format, err1 := image.Decode(file)
	if err1 != nil {
		return nil, "", err1
	}

	return img, format, nil
}

func EncodeImage(filename string, img image.Image, format string) error {
	/* Cette fonction prend en argument une image de type image.Image
	(qui est un type propre à go) et renvoie cette image dans un format jpg, png, ...
	en fonction de la valeur de "format" passé en argument. */

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "jpeg", "jpg":
		return jpeg.Encode(file, img, nil)
	case "png":
		return png.Encode(file, img)
	case "gif":
		return gif.Encode(file, img, nil)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

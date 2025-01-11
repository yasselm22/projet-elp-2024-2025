package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

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

func ImageToColorMatrix(img image.Image) [][]color.Color {
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

/*	*********************************Pour décoder l'image***********************
	img, format, err1 := decodeImage("path/to/image.jpg")
	if err1 != nil {
		fmt.Println("Error decoding image:", err1)
		return
	}
	fmt.Printf("Decoded image format: %s\n", format)

	*********************************Pour encoder l'image*****************
	err := encodeImage("path/to/output.jpg", img, format)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	fmt.Println("Image saved successfully") */

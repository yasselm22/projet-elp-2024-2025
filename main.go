package main

import (
	sobel "_/home/marine/Documents/GO/projet-elp-2024-2025"
	"sobel"
)

/* Fonction main qui lance une boucle for exécutant N fois la routine go, N est en paramètre de la
fonction et pour avoir une exécution optimale, il doit être égal au nombre de threads de
notre ordi  Marine:16 threads*/

func main() {
	sobel.Decoupe_image()
}

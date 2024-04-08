package main

import (
    "fmt"
    "os"
    "log"
    "image"
    "image/jpeg"
)

func main() {
    // Ouvrir le fichier vidéo
    file, err := os.Open("sample.mp4")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Créer un décodeur d'image JPEG
    decoder := jpeg.NewDecoder(file)

    // Lecture de la première image
    prevImg, _, err := image.Decode(decoder)
    if err != nil {
        log.Fatal(err)
    }

    // Variables pour stocker l'histogramme de l'image précédente
    prevHist := makeHistogram(prevImg)

    // Variable pour compter le numéro d'image
    frameNumber := 1

    // Boucle pour lire toutes les images de la vidéo
    for {
        // Lecture de l'image suivante
        img, _, err := image.Decode(decoder)
        if err != nil {
            // Si erreur de fin de fichier, arrêter la boucle
            if err == io.EOF {
                break
            }
            log.Fatal(err)
        }

        // Calcul de l'histogramme de l'image actuelle
        hist := makeHistogram(img)

        // Comparaison de l'histogramme avec l'image précédente
        if !compareHistograms(prevHist, hist) {
            fmt.Printf("Changement détecté à l'image %d\n", frameNumber)
        }

        // Mise à jour de l'image et de l'histogramme précédents
        prevImg = img
        prevHist = hist

        // Incrémenter le numéro d'image
        frameNumber++
    }
}

// Fonction pour créer un histogramme de couleurs à partir d'une image
func makeHistogram(img image.Image) [256][256][256]int {
    histogram := [256][256][256]int{}

    // Parcourir chaque pixel de l'image
    bounds := img.Bounds()
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            // Obtenir la couleur du pixel
            r, g, b, _ := img.At(x, y).RGBA()

            // Mettre à jour l'histogramme avec la couleur du pixel
            histogram[int(r>>8)][int(g>>8)][int(b>>8)]++
        }
    }

    return histogram
}

// Fonction pour comparer deux histogrammes de couleurs
func compareHistograms(hist1, hist2 [256][256][256]int) bool {
    // Comparer les valeurs de chaque bin d'histogramme
    for r := 0; r < 256; r++ {
        for g := 0; g < 256; g++ {
            for b := 0; b < 256; b++ {
                if hist1[r][g][b] != hist2[r][g][b] {
                    return false
                }
            }
        }
    }

    return true
}

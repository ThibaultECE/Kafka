package main

import (
    "fmt"
    "os/exec"
)

func main() {
    // Chemin de l'entrée
    inputFile := "input.mp4"

    // Répertoire de sortie
    outputDir := "output/"

    // Nombre total d'images à extraire
    numImages := 10

    // Boucle pour générer chaque image
    for i := 1; i <= numImages; i++ {
        // Nom du fichier de sortie avec un nombre incrémenté
        outputFile := fmt.Sprintf("playlist%03d.png", i)

        // Commande ffmpeg pour extraire une image à partir de la vidéo
        ffmpegCmd := exec.Command("ffmpeg", "-i", inputFile, "-vf", fmt.Sprintf("fps=1", i), fmt.Sprintf("%s/%s", outputDir, outputFile))

        // Exécution de la commande ffmpeg
        err := ffmpegCmd.Run()
        if err != nil {
            fmt.Printf("Erreur lors de l'exécution de ffmpeg: %v\n", err)
            return
        }
    }
}

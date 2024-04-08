package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"io/ioutil"
	"os"
    "os/exec"
    "path/filepath"
    "strings"
)

func ConvertirVideoEnImage(chemin string) ([]string, error) {

	tempDir := ioutil.TempDir("", "video_frames")
	defer os.RemoveAll(tempDir)

	imagePattern := filepath.Join(tempDir, "image-%03d.jpg")
	cmd := exec.Command("ffmpeg", "-i", chemin, "-vf", "fps=1/1", imagePattern)

	err = cmd.Run()
    if err != nil {
        return nil, err
    }

	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("fra")

	var livres []string
	for i := 1; ; i ++ {
		cheminImage := fmt.Sprintf(imagePattern,i)
		alpha := os.Stat(cheminImage)

		textBrut := extractTextFromImage(client,cheminImage)
		livreBrut = append(livreBrut,textBrut)
	}

	return texts, nil
}

func extractTextFromImage(client *gosseract.Client, cheminImage string) (string, error) {

	client.SetImage(cheminImage)
	page := client.Text()

	return page, nil

}

func main() {

	cheminVideo := "sample.mp4"
	transcription, err := ConvertirVideoEnImage(cheminVideo)
	if err != nil {
		fmt.Printf("Une erreur est survenue: ", err)
		return
	}

	for i, text := range transcription	{
        fmt.Printf("Texte extrait de l'image %d:\n%s\n", i+1, text)	
	}


	// client := gosseract.NewClient()
	// defer client.Close()
	// client.SetLanguage("fra")
	// // client.SetImage("001-helloworld.png")
	// client.SetImage("sample.jpg")
	// text, _ := client.Text()
	// fmt.Println(text)
	// fmt.Println("text")
	// // Hello, World!
}
package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"io/ioutil"
	"os"
    "os/exec"
    "path/filepath"
)

func DemoConvertir(chemin string) ([]string, error) {
	tempDir,err := ioutil.TempDir("", "video_frames")
	if err != nil {
        return nil, err
    }
	defer os.RemoveAll(tempDir)

	imagePattern := filepath.Join(tempDir, fmt.Sprintf("image-%03d.jpg",100))
	fmt.Printf(fmt.Sprintf("image-%03d.jpg",100))
	cmd := exec.Command("ffmpeg", "-i", chemin, "-vf", "fps=1/1", imagePattern)
	cmd.Run()
	return nil, nil
}

func ConvertirVideoEnImage(chemin string) ([]string, error) {

	// tempDir,err := ioutil.TempDir("", "video_frames")
	// if err != nil {
    //     return nil, err
    // }
	tempDir := "/temp"
	// defer os.RemoveAll(tempDir)

	imagePattern := filepath.Join(tempDir, fmt.Sprintf("%03d.jpg",100))
	fmt.Printf("ffmpeg", "-i", chemin, "-vf", "fps=1 %03d.jpg")
	cmd := exec.Command("ffmpeg -i "+chemin + " fps=1 %03d.jpg")
	// cmd := exec.Command("ffmpeg", "-i", chemin, "-vf", "fps=1", imagePattern)
	// cmd := exec.Command("ffmpeg", "-i", chemin, "-vf", "fps=1 %03d.jpg")

	err := cmd.Run()
    if err != nil {
		fmt.Printf("erreur dans l exec de la commande")
        return nil, err
    }
	fmt.Printf("Début de la traduction\n")
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("fra")

	var livreBrut []string
	for i := 1; ; i ++ {
		cheminImage := fmt.Sprintf(imagePattern,i)
        _, err := os.Stat(chemin)
        if os.IsNotExist(err) {
            break 
        }

		textBrut, err := extractTextFromImage(client,cheminImage)
		if err != nil {
            return nil, err
        }
		livreBrut = append(livreBrut,textBrut)
	}

	return livreBrut, nil
}

func extractTextFromImage(client *gosseract.Client, cheminImage string) (string, error) {

	err := client.SetImage(cheminImage)
    if err != nil {
        return "", err
    }

	page,err := client.Text()
	if err != nil {
        return "", err
    }

	return page, nil

}

func main() {

	cheminVideo := "sample.mp4"
	transcription, err := ConvertirVideoEnImage(cheminVideo)
	// transcription, err := DemoConvertir(cheminVideo)
	if err != nil {
		fmt.Printf("Une erreur est survenue: ", err)
		return
	}
	fmt.Printf("Le traitement c'est bien passé\n")
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
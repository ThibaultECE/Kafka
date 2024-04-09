package main

import (
	"fmt"
	"log"
	"github.com/otiai10/gosseract/v2"
	"os"
    "os/exec"
)

func CreerFichiersImage(inputFile string, outputDir string, segmentDuration int) error {

	if err := os.MkdirAll(outputDir, 0755); err != nil {
	 return fmt.Errorf("failed to create output directory: %v", err)
	}

	ffmpegCmd := exec.Command(
		"ffmpeg", 
		"-i",
		inputFile, 
		"-vf",
		"fps=1/5",
		"temp/image%d.png",
	)
   
	output, err := ffmpegCmd.CombinedOutput()
	if err != nil {
	 return fmt.Errorf("failed to create HLS: %v\nOutput: %s", err, string(output))
	}
   
	return nil
   }

func Traducteur(chemin string,frames int) ([]string, error) {
	
	fmt.Printf("\nDébut de la traduction")
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("fra")

	imagePattern := "temp/image%d.png"

	var livreBrut []string
	for i := 1; i <= frames ; i ++ {
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
	inputFile := "sample.mp4"
	outputDir := "temp/"
	segmentDuration := 10 
   
	if err := CreerFichiersImage(inputFile, outputDir, segmentDuration); err != nil {
	 log.Fatalf("Erreur dans le découpage: %v", err)
	}
   
	log.Println("\nSuccès du découpage")

	transcription, err := Traducteur(inputFile,4)
	if err != nil {
		fmt.Printf("\nUne erreur est survenue: ", err)
		return
	}
	fmt.Printf("\nLe traitement c'est bien passé")
	for i, text := range transcription	{
        fmt.Printf("Texte extrait de l'image %d:\n%s\n", i+1, text)	
	}

	os.RemoveAll("temp")
	os.MkdirAll("temp/",0750)
   }

// func main() {

// 	cheminVideo := "sample.mp4"
// 	transcription, err := ConvertirVideoEnImage(cheminVideo)
// 	// transcription, err := DemoConvertir(cheminVideo)
// 	if err != nil {
// 		fmt.Printf("Une erreur est survenue: ", err)
// 		return
// 	}
// 	fmt.Printf("Le traitement c'est bien passé\n")
// 	for i, text := range transcription	{
//         fmt.Printf("Texte extrait de l'image %d:\n%s\n", i+1, text)	
// 	}


// 	// client := gosseract.NewClient()
// 	// defer client.Close()
// 	// client.SetLanguage("fra")
// 	// // client.SetImage("001-helloworld.png")
// 	// client.SetImage("sample.jpg")
// 	// text, _ := client.Text()
// 	// fmt.Println(text)
// 	// fmt.Println("text")
// 	// // Hello, World!
// }
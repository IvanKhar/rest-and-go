package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFileFromGDrive(fileName string) {
	url := "https://drive.google.com/uc?id=1C9peO58IhTVAEVO3BBCTnk_wpKAJ08tq&export=download"
	fmt.Println("Downloading file...")

	output, err := os.Create(fileName)
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)

	fmt.Println(n, "bytes downloaded")
}

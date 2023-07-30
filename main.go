package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func createMultipartFormData(fieldName, fileName string) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		log.Fatal(err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		log.Fatal(err)
	}
	w.Close()
	return b, w
}

func main() {
	serverAddress := "192.168.1.133"
	serverPort := 80
	url := fmt.Sprintf("http://%s:%d/upload/image.png", serverAddress, serverPort)

	b, w := createMultipartFormData("image", "/dev/stdin")
	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		fmt.Print(res.Status)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(resBody))
}

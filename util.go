package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetWebPic(url string) (path string, err error) {
	imgPath := "./data/img/"
	picname := strconv.FormatInt(time.Now().Unix(), 10)
	picpath := imgPath + picname + ".jpg"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Dwonload pic err!")
		return "", err
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(picpath)
	if err != nil {
		return "", err
	}
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Printf("Download pic succ. Total pic length: %d", written)
	return picpath, err
}

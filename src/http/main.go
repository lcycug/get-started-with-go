package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	response, err := http.Get("https://baidu.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 999999)
	// response.Body.Read(bs)
	// fmt.Println(string(bs))
	io.Copy(os.Stdout, response.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	return 1, nil
}

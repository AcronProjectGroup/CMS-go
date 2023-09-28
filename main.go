// package main

// import (
// 	"io"
// 	"net/http"
// )

// type StringHandler struct {
// 	message string
// }

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	Printfln("Method: %v", request.Method)
// 	Printfln("URL: %v", request.URL)
// 	Printfln("HTTP Version: %v", request.Proto)
// 	Printfln("Host: %v", request.Host)
// 	for name, val := range request.Header {
// 		Printfln("Header: %v, Value: %v", name, val)
// 	}
// 	Printfln("---")
// 	io.WriteString(writer, sh.message)
// }
// func main() {
// 	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	}
// }
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type StringHandler struct {
	message string
}



func main() {


	MyServer()

}

func MyServer() {
	logFile, err := os.OpenFile("LogFile/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	
	err = http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Method: %v", request.Method)
	log.Printf("URL: %v", request.URL)
	log.Printf("HTTP Version: %v", request.Proto)
	log.Printf("Host: %v", request.Host)
	for name, val := range request.Header {
		log.Printf("Header: %v, Value: %v", name, val)
	}
	io.WriteString(writer, sh.message)
	log.Printf("============================================◉🧭🧭🧭🧭🧭🧭🧭◉==========================================")
}
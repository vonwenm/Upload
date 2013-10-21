package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Upload(rw http.ResponseWriter, req *http.Request) {

	fmt.Println("Incoming message...")
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Fprintln(rw, "error")
		fmt.Println("Error while starting server...")
	}

	fmt.Println("Name of file incoming ", handler.Filename)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintln(rw, "ERROR")
		fmt.Println("Error reading data...", err)
	}
	err = ioutil.WriteFile(handler.Filename, data, 0700)
	fmt.Println("Writing file to disc")
	if err != nil {
		fmt.Fprintln(rw, "ERROR")
		fmt.Println("Error writing file")
	}
	timeNow := time.Now()
	fmt.Fprintf(rw, "Successfull uploading file: %s\n", handler.Filename)
	fmt.Fprintf(rw, "Time : %s", timeNow.Format(time.Kitchen))
	fmt.Println("Successfull upload ", handler.Filename)
}

func SayDate(rw http.ResponseWriter, req *http.Request) {
	timeNow := time.Now()
	fmt.Fprintf(rw, "Time now%s", timeNow.Format(time.Kitchen))

}

func SayName(rw http.ResponseWriter, req *http.Request) {

	remPartOfURL := req.URL.Path[len("/name/"):]
	fmt.Fprintf(rw, "Hello %s", remPartOfURL)
}

func main() {

	http.HandleFunc("/name/", SayName)
	http.HandleFunc("/date/", SayDate)
	http.HandleFunc("/upload/", Upload)
	http.Handle("/file/", http.StripPrefix("/file", http.FileServer(http.Dir("/Users/Christopher/Documents/Programmering/go"))))
	err := http.ListenAndServe("localhost:4000", nil)

	if err != nil {
		log.Fatal("Error ListenAndServe", err)
		fmt.Println("ERROR")
	}
}

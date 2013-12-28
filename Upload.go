/* Copyright (C) 2013 Christopher
 * This file is part of Upload.
 *
 * Upload is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Upload is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Upload.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var uploadTemplate, _ = template.ParseFiles("/Users/Christopher/Documents/" +
	"Programmering/go/libs/src/github.com/christopherL91/Upload/Upload.html")

func fileserve(rw http.ResponseWriter, req *http.Request) {
	uploadTemplate.Execute(rw, nil)
}

func upload(rw http.ResponseWriter, req *http.Request) {

	fmt.Println("Incoming message...")
	if req.Method != "POST" {
		fmt.Println("ERROR not POST")
		uploadTemplate.Execute(rw, nil)
		return
	}
	file, handler, err := req.FormFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println("Something happended")
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
	fmt.Fprintf(rw, "Time : %s\n", timeNow.Format(time.Kitchen))
	fmt.Println("Successfull upload ", handler.Filename)
}

func sayDate(rw http.ResponseWriter, req *http.Request) {
	timeNow := time.Now()
	fmt.Fprintf(rw, "Time now %s", timeNow.Format(time.Kitchen))

}

func sayName(rw http.ResponseWriter, req *http.Request) {

	remPartOfURL := req.URL.Path[len("/name/"):]
	fmt.Fprintf(rw, "Hello %s", remPartOfURL)
}

func main() {
	cores := flag.Int("cores", 1, "The number of cores used")
	port := flag.Int("port", 4000, "The port number that the server will use")
	flag.Parse()
	runtime.GOMAXPROCS(*cores)

	fmt.Println("Server started on port:", *port)
	http.HandleFunc("/name/", sayName)
	http.HandleFunc("/date/", sayDate)
	http.HandleFunc("/upload/", upload)
	http.HandleFunc("/html/", fileserve)
	http.Handle("/look/", http.StripPrefix("/file", http.FileServer(http.Dir("/Users/Christopher/Documents/Programmering/go"))))
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)

	if err != nil {
		log.Fatal("Error ListenAndServe", err)
		fmt.Println("ERROR")
	}
}

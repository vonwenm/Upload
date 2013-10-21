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
	fmt.Fprintf(rw, "Time : %s\n", timeNow.Format(time.Kitchen))
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

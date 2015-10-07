package main

import (
	"flag"
	"github.com/cross-dev/script-server/processor"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

var baseUrl, listenAt string

func init() {
	flag.StringVar(&baseUrl, "b", "/", "Base URL pathname")
	flag.StringVar(&listenAt, "l", ":41267", "Interface and port to listen")
}

func main() {
	flag.Parse()
	rest := flag.Args()
	if len(rest) == 0 {
		log.Fatal("Pass file name or - to read from STDIN")
	}
	var file *os.File
	if rest[0] == "-" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(rest[0])
		defer file.Close()
		if err != nil {
			log.Fatal("Error opening file ", rest[0], ": ", err)
		}
	}
	input, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Fetching template from STDIN: ", err)
	}
	template := string(input)
	scriptService := func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Serving: ", request.URL.String())
		processor.Process(request.URL.Query(), template, writer)
	}

	http.HandleFunc(path.Join(baseUrl, "get"), scriptService)
	log.Fatal(http.ListenAndServe(listenAt, nil))
}

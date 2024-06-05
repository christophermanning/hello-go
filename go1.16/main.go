package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/christophermanning/hello-go/go1.16/embed"
	"github.com/christophermanning/hello-go/go1.16/filesystem"
	"github.com/christophermanning/hello-go/go1.16/util"
)

var portFlag = flag.Int("p", 8000, "port")

func main() {
	flag.Parse()

	util.ForLoopWithShared()

	html := embed.New()

	parsedTemplate, err := template.New("html").Parse(html)
	if err != nil {
		log.Fatalf("Error parsing template : %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// the "/" path HandleFunc captures everything by default, so create a case for the root path
		if r.URL.Path == "/" {
			err := parsedTemplate.Execute(w, struct {
				StartTime string
			}{
				time.Now().Format(time.Stamp),
			})
			if err != nil {
				log.Println("Error executing template :", err)
				http.Error(w, "Server Error", 500)
			}
		} else {
			http.NotFound(w, r)
		}
	})

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot Getwd: %s", err)
	}

	dfs := os.DirFS(path)
	count, err := filesystem.CountFiles(dfs)
	if err != nil {
		log.Fatalf("cannot CountFiles: %s", err)
	}

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, count)
	})

	fmt.Printf("Listening on %d\n", *portFlag)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), nil))
}

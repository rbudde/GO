package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"fmt"
)

func loadFile(name string) ([]byte, error) {
	filename := name + ".html"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func staticHandler(w http.ResponseWriter, r *http.Request, requestParam string) {
	content, err := loadFile(requestParam)
	if err != nil {
		http.Redirect(w, r, "/static/index", http.StatusFound)
		return
	}
	w.Write(content)
}

func runHandler(w http.ResponseWriter, r *http.Request, requestParam string) {
	fmt.Fprintf(w, "I got param %s\n", requestParam)
}

var validPath = regexp.MustCompile("^/(static|run)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("path is: " + r.URL.Path + "\n")
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/static/", makeHandler(staticHandler))
	http.HandleFunc("/run/", makeHandler(runHandler))

	http.ListenAndServe(":8080", nil)
}

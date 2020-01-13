package main

import (
	"net/http"
	"regexp"
	"fmt"
)

var validPath = regexp.MustCompile("^/(r1|r2|r3)/(.+)$")

func runHandler(w http.ResponseWriter, r *http.Request, pathPrefix string, requestParam string) {
	fmt.Fprintf(w, "I am handler of %s and got param %s\n", pathPrefix, requestParam)
}

func runRedirect(w http.ResponseWriter, r *http.Request, pathPrefix string, requestParam string) {
	fmt.Fprintf(w, "{\"from\":\"%s\"}", pathPrefix)
	http.Redirect(w, r, "/r1/redirect/" + pathPrefix + "/" + requestParam, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("path is: " + r.URL.Path + "\n")
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[1], m[2])
	}
}

func main() {
	http.HandleFunc("/r1/", makeHandler(runHandler))
	http.HandleFunc("/r2/", makeHandler(runHandler))
	http.HandleFunc("/r3/", makeHandler(runRedirect))
	http.ListenAndServe(":8181", nil)
}

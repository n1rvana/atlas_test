package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "os"
	// "unicode/utf8"
	// ed "github.com/agl/ed25519"
)

var err error

func main() {
	myStatus = Status{ChainHeight: 100, CurrentBlock: 0, PercentComplete: 0.0}
	http.HandleFunc("/", statusHandler)
	http.ListenAndServe(":80", nil)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love go!")
}

// This should probably be migrated to something common
func errorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, a...)
}

// Queru Simple Server
// A simple HTTP command line server
// 2018 GPL3
// https://www.gnu.org/licenses/gpl-3.0.en.html

package main

import (
	"fmt"
	"flag"
	"net/http"
)

var port int
var docroot string
func init() {
	flag.IntVar(&port, "p", 8000, "Listening port")
	flag.StringVar(&docroot, "r", ".", "Document Root: Folder to serve")
}

func main() {
	fmt.Println("Queru Simple Server")
	flag.Parse()
	fmt.Printf("Document root at: %q\n", docroot)
	fmt.Printf("Listening at: http://localhost:%d\n", port)
	fmt.Println("Ready")
	http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(docroot)))
}


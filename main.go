package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:next/dist
var nextFS embed.FS

func main() {
	// setup cache path flag
	port := *flag.Int("port", 80, "port to serve")
	flag.Parse()

	// Root at the `dist` folder generated by the Next.js app.
	distFS, err := fs.Sub(nextFS, "next/dist")
	if err != nil {
		log.Fatal(err)
	}

	// The static Next.js app will be served under `/`.
	http.Handle("/", http.FileServer(http.FS(distFS)))
	// The API will be served under `/api`.
	http.HandleFunc("/api/", handleAPI)

	// Start HTTP server at port.
	log.Printf("Starting HTTP server at http://localhost:%d ...\n", port)
	err = http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func handleAPI(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, from api!")
}

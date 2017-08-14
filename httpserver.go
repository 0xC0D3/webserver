package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	rootDir := flag.String("root", "", "HTTP Server root directory.")
	httpPort := flag.Int("port", 8080, "HTTP Server port.")
	flag.Parse()

	if strings.TrimSpace(*rootDir) == "" {
		flag.Usage()
		log.Fatal(fmt.Errorf("%v\n", "Root directory is required."))
	}

	serve(*httpPort, *rootDir)
}

func serve(port int, dir string) {
	host := fmt.Sprintf(":%v", port)
	log.Printf("Listen at %v, serving \"%v\"", host, dir)

	err := http.ListenAndServe(host, http.FileServer(http.Dir(dir)))
	if err != nil {
		panic(err)
	}
}

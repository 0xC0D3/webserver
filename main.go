package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	rootDir := flag.String("root", "", "Content root directory.")
	httpPort := flag.Int("port", 8080, "HTTP/HTTPS Server port.")
	tls := flag.Bool("tls", false, "TLS Enable.")
	pKey := flag.String("pkey", "", "TLS Private Key file path.")
	cert := flag.String("cert", "", "TLS Public Cert file path.")

	flag.Parse()

	if strings.TrimSpace(*rootDir) == "" {
		flag.Usage()
		log.Fatal(fmt.Errorf("%v\n", "Root directory is required."))
	}

	if *tls {
		if *pKey == "" || *cert == "" {
			log.Fatalln("When TLS mode is enabled, Private Key and Public Cert are required.")
		}
	} else {
		log.Println("TLS mode disabled, 'pkey' and 'cert' params will be ignored.")
	}

	serve(*httpPort, *rootDir, *tls, *pKey, *cert)
}

func serve(port int, dir string, tls bool, private, cert string) {
	host := fmt.Sprintf(":%v", port)
	log.Printf("Listen at %v, serving \"%v\"", host, dir)

	if tls {
		err := http.ListenAndServeTLS(host, cert, private, http.FileServer(http.Dir(dir)))
		if err != nil {
			panic(err)
		}
	} else {
		err := http.ListenAndServe(host, http.FileServer(http.Dir(dir)))
		if err != nil {
			panic(err)
		}
	}
}

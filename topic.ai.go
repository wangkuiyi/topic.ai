package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":80", "Listing address")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html>
<head>
<meta name="google-site-verification" content="mn7BAIOajioNPWcDm8PfERRXV8-gcLx_byW06oRIM9s" />
</head>
<body>
Hello!
</body></html>`)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}

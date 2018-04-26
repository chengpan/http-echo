// Copyright (C) 2018 Robert A. Wallis, All Rights Reserved

// http-echo is an HTTP server that just outputs all the requests directly to the console
// this is useful for debugging the output of your clients that are posting data, to see what the server sees.
package main

import "net/http"
import "fmt"
import "os"
import "io"
import "time"
import "flag"

var address = flag.String("address", "localhost:8000", "where to serve HTTP requests from")

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.RemoteAddr)
		fmt.Println(r.Method, r.URL.RequestURI(), r.Proto)
		for k, v := range r.Header {
			for _, innerV := range v {
				fmt.Println(k+":", innerV)
			}
		}
		io.Copy(os.Stdout, r.Body)
		fmt.Println()
		fmt.Println()
		w.Write([]byte("See console for output.\n"))
	})
	fmt.Printf("Echoing to console HTTP requests to http://%s\n", *address)
	fmt.Println(http.ListenAndServe(*address, nil))
}

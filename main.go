package main

// run with: env PORT=8081 go run http-server.go

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {

    port := os.Getenv("PORT")
    if port == "" {
      log.Fatal("Please specify the HTTP port as environment variable, e.g. env PORT=8081 go run http-server.go")
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello World")
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))

}

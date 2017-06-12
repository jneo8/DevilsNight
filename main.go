package main 

import (

    "flag"
    "net/http"
    "log"
)

func main() {

    port := flag.String("p", "19480", "port")
    dir := flag.String("d", ".", "dir")
    flag.Parse()

    http.Handle("/", http.FileServer(http.Dir(*dir)))
    log.Printf("Serving %s on Http port: %s\n ", *dir, *port)
    log.Fatal(http.ListenAndServe("0.0.0.0:"+*port, nil))

}

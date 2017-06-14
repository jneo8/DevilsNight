//routes.go
package main


import (
    "os"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/op/go-logging"
)

var log = logging.MustGetLogger("Web")
var format = logging.MustStringFormatter(
    `%{color}%{time:15:04:05.000} %{longfile} ▶ %{color:bold}%{level:.4s} %{id:03x} %{message}%{color:reset}`,
)

func main() {
    //SetUp Log
    backend1 := logging.NewLogBackend(os.Stderr, "", 0)
    backend2 := logging.NewLogBackend(os.Stderr, "", 0)
    backend2Formatter := logging.NewBackendFormatter(backend2, format)
    backend1Leveled := logging.AddModuleLevel(backend1)
    backend1Leveled.SetLevel(logging.ERROR, "")
    logging.SetBackend(backend1Leveled, backend2Formatter)

    log.Notice("Start mux Router!!")
    userAges := map[string]int{
        "Alice": 25,
        "Bob": 30,
        "Claire": 29,
    }
    r := mux.NewRouter()
    r.HandleFunc("/users/{name}", func(w http.ResponseWriter, r *http.Request) {    
        vars := mux.Vars(r)
        name := vars["name"]
        log.Debug(name)
        age := userAges[name]
        fmt.Fprintf(w, "%s is %d years old!", name, age)
    }).Methods("GET")
    
    http.ListenAndServe(":8080", r)
}

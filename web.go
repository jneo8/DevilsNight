//routes.go
package main


import (
    "os"
    "fmt"
    "net/http"
    "github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(
    `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {
    //SetUp Log
    backend1 := logging.NewLogBackend(os.Stderr, "", 0)
    backend2 := logging.NewLogBackend(os.Stderr, "", 0)
    backend2Formatter := logging.NewBackendFormatter(backend2, format)
    backend1Leveled := logging.AddModuleLevel(backend1)
    backend1Leveled.SetLevel(logging.ERROR, "")
    logging.SetBackend(backend1Leveled, backend2Formatter)

    userAges := map[string]int{
        "Alice": 25,
        "Bob": 30,
        "Claire": 29,
    }

    http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request){    
        name := r.URL.Path[len("/users/"):]
        log.Debugf(name)
        age := userAges[name]
        fmt.Fprintf(w, "%s is %d years old!", name, age)
    })
    
    http.ListenAndServe(":8080", nil)
}

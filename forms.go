//forms.go
package main

import (
    "os"
    "github.com/op/go-logging"
    "html/template"
    "net/http"
)

var log = logging.MustGetLogger("Web")
var format = logging.MustStringFormatter(
    `%{time:15:04:05.000} %{color}%{longfile} %{color:reset}▶ %{color}%{level:.4s} %{id:03x} %{color:bold}%{message}%{color:reset}`,
)

type ContactDetails struct {
    Email string
    Subjuct string
    Message string 
}

func init() {
    backend1 := logging.NewLogBackend(os.Stderr, "", 0)
    backend2 := logging.NewLogBackend(os.Stderr, "", 0)
    backend2Formatter := logging.NewBackendFormatter(backend2, format)
    backend1Leveled := logging.AddModuleLevel(backend1)
    backend1Leveled.SetLevel(logging.ERROR, "")
    logging.SetBackend(backend1Leveled, backend2Formatter)
}

func main() {
    //SetUp Log
    tmpl := template.Must(template.ParseFiles("htmls/forms.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        log.Debug(r)

        details := ContactDetails{
            Email: r.FormValue("email"),
            Subjuct: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }
        
        //do something with details

        log.Info(details.Email)
        log.Info(details.Subjuct)
        log.Info(details.Message)


        tmpl.Execute(w, struct{ Success bool }{true})
    })
    http.ListenAndServe("0.0.0.0:8000", nil)
}

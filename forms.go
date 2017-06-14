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
    `%{color:reset}%{time:15:04:05.000} %{color:bold}%{longfile} ▶ %{level:.4s} %{id:03x} %{color:reset}%{color}%{message}%{color:reset}`,
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
        log.Info(r)
        log.Notice(r)

        details := ContactDetails{
            Email: r.FormValue("email"),
            Subjuct: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }
        
        //do something with details
        _ = details
        tmpl.Execute(w, struct{ Success bool }{true})
    })
    http.ListenAndServe("0.0.0.0:19480", nil)
}

package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This is a routing calculator.   It works by blah blah blah blah do this\n"))
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
    reqs := mux.Vars(r)
    x, y := reqs["x"] ,reqs["y"]
    xf, _ := strconv.ParseFloat(x, 64)
    yf, _ := strconv.ParseFloat(y, 64)
    ans := strconv.FormatFloat(xf + yf, 'f', -1, 64)


    w.Write([]byte(ans))
}

func SubHandler(w http.ResponseWriter, r *http.Request) {
    reqs := mux.Vars(r)
    x, y := reqs["x"] ,reqs["y"]
    xf, _ := strconv.ParseFloat(x, 64)
    yf, _ := strconv.ParseFloat(y, 64)
    ans := strconv.FormatFloat(xf - yf, 'f', -1, 64)


    w.Write([]byte(ans))
}

func MultHandler(w http.ResponseWriter, r *http.Request) {
    reqs := mux.Vars(r)
    x, y := reqs["x"] ,reqs["y"]
    xf, _ := strconv.ParseFloat(x, 64)
    yf, _ := strconv.ParseFloat(y, 64)
    ans := strconv.FormatFloat((xf * yf), 'f', -1, 64)


    w.Write([]byte(ans))
}

func DivHandler(w http.ResponseWriter, r *http.Request) {
    reqs := mux.Vars(r)
    x, y := reqs["x"] ,reqs["y"]
    xf, _ := strconv.ParseFloat(x, 64)
    yf, _ := strconv.ParseFloat(y, 64)
    ans := strconv.FormatFloat((xf / yf), 'f', -1, 64)


    w.Write([]byte(ans))
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", IndexHandler)
    r.HandleFunc("/add/{x}/{y}", AddHandler)
    r.HandleFunc("/sub/{x}/{y}", SubHandler)
    r.HandleFunc("/mult/{x}/{y}", MultHandler)
    r.HandleFunc("/div/{x}/{y}", DivHandler)
    //Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}

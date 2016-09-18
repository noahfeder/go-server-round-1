package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "strconv"
    "fmt"
    "os"
    "io/ioutil"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    index, _ := os.Open("index.html")
    data, _ := ioutil.ReadAll(index)
    w.Write([]byte(data))
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
    reqs := mux.Vars(r)
    x, y := reqs["x"] ,reqs["y"]
    xf, _ := strconv.ParseFloat(x, 64)
    yf, _ := strconv.ParseFloat(y, 64)
    ans := strconv.FormatFloat(xf + yf, 'f', -1, 64)
    printer := x + " + " + y + " = " + ans + "\nThanks for using our calculator!"
    w.Write([]byte(printer))
}

func SubHandler(w http.ResponseWriter, r *http.Request) {
    reqs := mux.Vars(r)
    x, y := reqs["x"] ,reqs["y"]
    xf, _ := strconv.ParseFloat(x, 64)
    yf, _ := strconv.ParseFloat(y, 64)
    ans := strconv.FormatFloat(xf - yf, 'f', -1, 64)
    printer := x + " - " + y + " = " + ans + "\nThanks for using our calculator!"
    w.Write([]byte(printer))
}

func MultHandler(w http.ResponseWriter, r *http.Request) {
    reqs := mux.Vars(r)
    x, y := reqs["x"] ,reqs["y"]
    xf, _ := strconv.ParseFloat(x, 64)
    yf, _ := strconv.ParseFloat(y, 64)
    ans := strconv.FormatFloat((xf * yf), 'f', -1, 64)
    printer := x + " * " + y + " = " + ans + "\nThanks for using our calculator!"
    w.Write([]byte(printer))
}

func DivHandler(w http.ResponseWriter, r *http.Request) {
    reqs := mux.Vars(r)
    x, y := reqs["x"] ,reqs["y"]
    xf, _ := strconv.ParseFloat(x, 64)
    yf, _ := strconv.ParseFloat(y, 64)
    ans := strconv.FormatFloat((xf / yf), 'f', -1, 64)
    printer := x + " / " + y + " = " + ans + "\nThanks for using our calculator!"
    w.Write([]byte(printer))
}

func main() {
    r := mux.NewRouter()
    fmt.Println("Listening on port 8000\n")
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", IndexHandler)
    r.HandleFunc("/add/{x}/{y}", AddHandler)
    r.HandleFunc("/sub/{x}/{y}", SubHandler)
    r.HandleFunc("/mult/{x}/{y}", MultHandler)
    r.HandleFunc("/div/{x}/{y}", DivHandler)
    //Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}

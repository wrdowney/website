package main

import (
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "log"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got / request\n")
    io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got /hello request\n")
    io.WriteString(w, "Hello, HTTP!\n")
}

func main() {

    resp, err := http.Get("https://medium.com/@wdowney20/fee")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))

    http.HandleFunc("/", getRoot)
    http.HandleFunc("/hello", getHello)

    err = http.ListenAndServe(":3333", nil)

    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed\n")
    } else if err != nil {
        fmt.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}

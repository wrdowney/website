package main

import (
    "errors"
    "fmt"
    "io"
    "net/http"
    "os"
    "log"
    "encoding/xml"
)

type Item struct {
    Title     string    `xml:"title"`
    Link      string    `xml:"link"`
    Guid      string    `xml:"guid"`
    Creator   string    `xml:"creator"`
    PubDate   string    `xml:"pubDate"`
    Updated   string    `xml:"updated"`
    Content   string    `xml:"encoded"`
}

type Channel struct {
    Title string `xml:"title"`
    Link  string `xml:"link"`
    Desc  string `xml:"description"`
    Items []Item `xml:"item"`
}

type Rss struct {
    Channel Channel `xml:"channel"`
}

var rss Rss

func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got / request\n")
    io.WriteString(w, rss.Channel.Items[0].Content)
}

func getHello(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got /hello request\n")
    io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
    resp, err := http.Get("https://medium.com/@wdowney20/feed")
    if err != nil {
        log.Fatalf("Error Get: %v\n", err)
    }
    defer resp.Body.Close()

    rss = Rss{}

    decoder := xml.NewDecoder(resp.Body)
    err = decoder.Decode(&rss)
    if err != nil {
        log.Fatalf("Error Decode: %v\n", err)
        return
    }

    for _, item := range rss.Channel.Items {
        log.Printf("title: %v pub: %v test: %v\n", item.Title, item.PubDate, item.Content)
    }

    http.HandleFunc("/", getRoot)
    http.HandleFunc("/hello", getHello)

    err = http.ListenAndServe(":3333", nil)

    if errors.Is(err, http.ErrServerClosed) {
        log.Print("server closed\n")
    } else if err != nil {
        log.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}

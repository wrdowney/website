package main

import (
    "errors"
    "io"
    "net/http"
    "log"
    "encoding/xml"
    "encoding/json"
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

var feed []byte

func getFeed(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    io.WriteString(w, string(feed))
}

func main() {
    // send GET request to my medium RSS feed
    resp, err := http.Get("https://medium.com/@wdowney20/feed")
    if err != nil {
        log.Fatalf("Error Get: %v\n", err)
    }
    defer resp.Body.Close()

    rss := Rss{}

    // decode the response from xml format into a Rss structure
    decoder := xml.NewDecoder(resp.Body)
    err = decoder.Decode(&rss)
    if err != nil {
        log.Fatalf("Error Decode: %v\n", err)
    }

    for _, item := range rss.Channel.Items {
        log.Printf("title: %v pub: %v test: %v\n", item.Title, item.PubDate, item.Content)
    }

    // convert from Rss structure to JSON format
    feed, err = json.Marshal(rss)
    if err != nil {
        log.Fatalf("Error JSON: %v\n", err)
    }
    log.Println(string(feed))

    http.HandleFunc("/feed", getFeed)

    err = http.ListenAndServe(":3333", nil)

    if errors.Is(err, http.ErrServerClosed) {
        log.Print("server closed\n")
    } else if err != nil {
        log.Fatalf("error starting server: %s\n", err)
    }
}

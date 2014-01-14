package main

import (
    "flag"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

const (
    API         = "https://api.datasift.com"
    API_VERSION = "v1"
)

// Note: returns pointers
var user = flag.String("user", "", "Your datasift user name")
var key = flag.String("key", "", "Your datasift api key")

func main() {

    flag.Parse()

    parts := []string{API, API_VERSION, "push", "get"}
    endpoint := strings.Join(parts, "/")

    client := &http.Client{}
    req, err := http.NewRequest("GET", endpoint, nil)

    if err != nil {
        log.Fatal(err)
    }

    req.Header.Add("Authorization", strings.Join([]string{*user, *key}, ":"))

    resp, err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%v", string(body[:]))

}

package main

import (
    "fmt"
    "io/ioutil"
    "sifty/datasift"

    "github.com/docopt/docopt.go"
)

const usage = `
Sifty

Usage:
    sifty USER KEY
    sifty -h | --help
    sifty --version

Arguments:
    USER    Datasift API Username
    KEY     Datasift API Key

Options:
    -h --help     Show this screen.
    --version     Show version.`

func main() {
    // Parse CLI arguments
    args, _ := docopt.Parse(usage, nil, true, "Sifty 0.1", false)

    // Convert user credentials to strings
    user := args["USER"].(string)
    key := args["KEY"].(string)

    // Create client
    s := datasift.NewClient(user, key, nil)

    // Issue a raw Get request to push/get endpoint
    response, _ := s.Get("push/get")

    // Read the body
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)

    // Print the output
    fmt.Println(string(body[:]))
}

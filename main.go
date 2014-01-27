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
    args, _ := docopt.Parse(usage, nil, true, "Sifty 0.1", false)
    s := datasift.NewClient(args["USER"], args["KEY"], nil)
    response, _ := s.Get("push/get")
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(body[:]))
}

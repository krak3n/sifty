/**
 * Main Sifty Package
 *
 * Provides command line integration for using the Datasift API.
 */

package main

import (
    "flag"
    "log"

    "github.com/krak3n/sifty/datasift"
)

var user = flag.String("user", "", "Your datasift user name")
var key = flag.String("key", "", "Your datasift api key")

func main() {

    flag.Parse()

    // API Endpoint
    endpoint := []string{"push", "get"}

    // Pointer to credentials
    client := new(datasift.Client)
    (*client).User = *user
    (*client).Key = *key

    query := datasift.Query(credentials, datasift.BuildEndpoint(endpoint))

    log.Printf(query)

}

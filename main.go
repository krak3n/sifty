/**
 * Main Sifty Package
 *
 * Provides command line integration for using the Datasift API.
 */

package main

import (
    "flag"
    "log"

    "sifty/datasift"
)

var (
    user = flag.String("user", "", "Your datasift user name")
    key  = flag.String("key", "", "Your datasift api key")
)

func main() {

    flag.Parse()

    // API Endpoint
    endpoint := []string{"push", "get"}

    // Pointer to credentials
    client := new(datasift.Client)
    (*client).User = *user
    (*client).Key = *key

    query := client.Query(endpoint)

    log.Printf(query)

}

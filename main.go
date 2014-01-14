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

var user = flag.String("user", "", "Your datasift user name")
var key = flag.String("key", "", "Your datasift api key")

func main() {

    flag.Parse()

    subscriptions := datasift.Client(user, key)

    log.Printf(subscriptions)

}

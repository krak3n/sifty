/**
 * Client Package
 *
 * Provides standard connection to Datasift API
 */

package datasift

import (
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

const (
    API         = "https://api.datasift.com"
    API_VERSION = "v1"
)

func Client(user *string, key *string) string {

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

    return string(body[:])

}

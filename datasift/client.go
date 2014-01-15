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
    API_ROOT    = "https://api.datasift.com"
    API_VERSION = "v1"
)

type Credentials struct {
    User string
    Key  string
}

func BuildEndpoint(parts []string) string {
    base := []string{
        API_ROOT,
        API_VERSION,
    }
    parts = append(base, parts...)
    return strings.Join(parts, "/")
}

func Query(credentials *Credentials, url string) string {

    log.Println(credentials)
    log.Println(url)

    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        log.Fatal(err)
    }

    req.Header.Add("Authorization", strings.Join([]string{credentials.User, credentials.Key}, ":"))

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

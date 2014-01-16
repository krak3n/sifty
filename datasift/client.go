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

func BuildEndpoint(parts []string) string {
    base := []string{
        API_ROOT,
        API_VERSION,
    }
    parts = append(base, parts...)
    return strings.Join(parts, "/")
}

type Credentials struct {
    User string
    Key  string
}

func (c Credentials) authorizationHeaderValue() (value string) {
    credentials := []string{
        c.User,
        c.Key,
    }
    return strings.Join(credentials, ":")
}

func (c Credentials) addHttpHeaders(request *http.Request) *http.Request {
    headers := make(map[string]string)
    headers["Authorization"] = c.authorizationHeaderValue()
    for key, value := range headers {
        request.Header.Add(key, value)
    }
    return request
}

func (c Credentials) request(url string) *http.Response {
    client := &http.Client{}
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
    }
    request = c.addHttpHeaders(request)
    response, err := client.Do(request)
    if err != nil {
        log.Fatal(err)
    }
    return response
}

func Query(c *Credentials, url string) string {
    response := c.request(url)
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    return string(body[:])
}

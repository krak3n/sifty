/**
 * Datasift Package
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

type Client struct {
    User string
    Key  string
}

func (c Client) authorizationHeaderValue() (value string) {
    credentials := []string{
        c.User,
        c.Key,
    }
    return strings.Join(credentials, ":")
}

func (c Client) addHttpHeaders(request *http.Request) *http.Request {
    headers := make(map[string]string)
    headers["Authorization"] = c.authorizationHeaderValue()
    for key, value := range headers {
        request.Header.Add(key, value)
    }
    return request
}

func (c Client) request(endpoint string) *http.Response {
    client := &http.Client{}
    request, err := http.NewRequest("GET", endpoint, nil)
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

func (c Client) makeEndpoint(parts []string) string {
    base := []string{
        API_ROOT,
        API_VERSION,
    }
    parts = append(base, parts...)
    return strings.Join(parts, "/")
}

func (c Client) Query(parts []string) string {
    endpoint := c.makeEndpoint(parts)
    response := c.request(endpoint)
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    return string(body[:])
}

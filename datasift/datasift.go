package datasift

import (
    "net/http"
    "net/url"
    "strings"
)

// Defaults
const (
    version        = "0.1"
    defaultAPIRoot = "https://api.datasift.com/v1/"
    userAgent      = "sifty/" + version
)

// Client which handles HTTP requests to the Datasift API
type Client struct {

    // Root API url used for all requests to the API. Default is
    // https://api.datasift.com/
    APIRoot *url.URL

    // Http Client used to send requests to Datasift API
    Client *http.Client

    // The HTTP user agent to use when sending requests to Datasift API
    UserAgent string

    // Datasift API credentials
    User string
    Key  string
}

// Return a new Datasift API Client.
func NewClient(httpClient *http.Client) *Client {

    if httpClient == nil {
        httpClient = http.DefaultClient
    }

    // Use the default API root url
    apiRoot, _ := url.Parse(defaultAPIRoot)

    // Create a new Datasift Client
    client := &Client{
        APIRoot:   apiRoot,
        UserAgent: userAgent,
        Client:    httpClient,
    }

    return client
}

// Generate the value required for the Datasift API Authorization
// HTTP header, this is based on the username and API Key found
// on the Datasift web UI
func (c *Client) authorizationHeaderValue() (value string) {
    credentials := []string{
        c.User,
        c.Key,
    }
    return strings.Join(credentials, ":")
}

// Build a new Datasift API request adding all the required HTTP
// headers for authentication with the API.
func (c *Client) Request(method string, endpoint string) (*http.Request, error) {

    // Parse the endpoint into a valid URL structure
    rel, err := url.Parse(endpoint)
    if err != nil {
        return nil, err
    }

    // Build absolute URL using the url reference above with the API root as the base
    e := c.APIRoot.ResolveReference(rel)

    // Create a new HTTP Request object
    req, err := http.NewRequest(method, e.String(), nil)
    if err != nil {
        return nil, err
    }

    // Add required HTTP headers to request
    req.Header.Add("User-Agent", c.UserAgent)
    req.Header.Add("Authorization", c.authorizationHeaderValue())
    req.Header.Add("Connection", "keep-alive")

    return req, nil
}

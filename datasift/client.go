package datasift

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
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

// Error response representation
type ErrorResponse struct {
    Response *http.Response // Http Response
    Message  string         `json:"error"` // The error message
}

func (r *ErrorResponse) Error() string {
    return fmt.Sprintf(
        "%v %v: %d %v",
        r.Response.Request.Method,
        r.Response.Request.URL,
        r.Response.StatusCode,
        r.Message,
    )
}

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
func NewClient(user string, key string, httpClient *http.Client) *Client {

    if httpClient == nil {
        httpClient = http.DefaultClient
    }

    // Use the default API root url
    apiRoot, _ := url.Parse(defaultAPIRoot)

    // Create a new Datasift Client
    client := &Client{
        APIRoot:   apiRoot,
        Client:    httpClient,
        UserAgent: userAgent,
        User:      user,
        Key:       key,
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

    // Create a new HTTP Request object, we don't care about the error
    // as we have already sanitized the url, http.NewRequest only
    // errors on url.Parse erros
    req, _ := http.NewRequest(method, e.String(), nil)

    // Add required HTTP headers to request
    req.Header.Add("User-Agent", c.UserAgent)
    req.Header.Add("Authorization", c.authorizationHeaderValue())
    req.Header.Add("Connection", "keep-alive")

    return req, nil
}

// Response queries the API and returns the API's respinse.
func (c *Client) Response(req *http.Request) (*http.Response, error) {

    // Get response from API
    response, err := c.Client.Do(req)
    if err != nil {
        return nil, err
    }

    // Validate the response from the API
    err = c.ValidateResponse(response)
    if err != nil {
        return response, err
    }

    return response, nil
}

// ValidateResponse enusres the response from the API is a valid one.
// For example if authentication fails Datasifts API will just respond
// with a simple {"error": "the messsage"} json object.
func (c *Client) ValidateResponse(r *http.Response) error {
    // Did we get a valid status code
    if status := r.StatusCode; 200 <= status && status <= 299 {
        return nil
    }
    // Create an error response
    errorResponse := &ErrorResponse{Response: r}
    // Read the response body
    data, err := ioutil.ReadAll(r.Body)
    // If we have data and no error, load the json into the
    // errorResponse
    if err == nil && data != nil {
        json.Unmarshal(data, errorResponse)
    }

    return errorResponse
}

func (c *Client) Get(endpoint string) (*http.Response, error) {

    // Generate Request
    request, err := c.Request("GET", endpoint)
    if err != nil {
        return nil, err
    }

    // Get the response
    response, err := c.Response(request)
    if err != nil {
        return nil, err
    }

    return response, nil

}

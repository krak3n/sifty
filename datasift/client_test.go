package datasift

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "net/url"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

/*
  Test setup
*/

type DatasiftTestSuite struct {
    suite.Suite

    mux    *http.ServeMux
    client *Client
    server *httptest.Server
}

// Setup Client Test Suite
func (s *DatasiftTestSuite) SetupTest() {
    // Create our fake HTTP server
    s.mux = http.NewServeMux()
    s.server = httptest.NewServer(s.mux)
    s.client = NewClient("foo", "bar", nil)

    // Override the client API Root to use the
    // HTTP test serves URL
    url, _ := url.Parse(s.server.URL)
    s.client.APIRoot = url
}

/*
   Tests
*/

func (s *DatasiftTestSuite) TestAuthroizationHeaderValue() {
    c := &Client{
        User: "foo",
        Key:  "bar",
    }

    assert.Equal(s.T(), c.authorizationHeaderValue(), "foo:bar")
}

func (s *DatasiftTestSuite) TestRequestInvalidURl() {
    _, err := s.client.Request("GET", ":")

    assert.NotEqual(s.T(), err, nil)
}

func (s *DatasiftTestSuite) TestAPIResponse() {
    s.mux.HandleFunc("/push/get", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, `{"id":1}`)
    })

    req, _ := s.client.Request("GET", "push/get")

    response, _ := s.client.Client.Do(req)
    defer response.Body.Close()

    body, _ := ioutil.ReadAll(response.Body)

    assert.Equal(s.T(), string(body[:]), `{"id":1}`)
}

/*
   Test Runner
*/

func TestRunDatasiftTestSuite(t *testing.T) {
    suite.Run(t, new(DatasiftTestSuite))
}

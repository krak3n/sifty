package datasift

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/httptest"
    "net/url"
    "testing"

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
func (suite *DatasiftTestSuite) SetupTest() {
    // Create our fake HTTP server
    suite.mux = http.NewServeMux()
    suite.server = httptest.NewServer(suite.mux)
    suite.client = NewClient(nil)

    // Override the client API Root to use the
    // HTTP test serves URL
    url, _ := url.Parse(suite.server.URL)
    suite.client.APIRoot = url
}

/*
   Tests
*/

func (suite *DatasiftTestSuite) TestAPIResponse() {
    suite.mux.HandleFunc("/push/get", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, `{"id":1}`)
    })

    req, _ := suite.client.Request("GET", "push/get")

    response, _ := suite.client.Client.Do(req)
    defer response.Body.Close()

    body, _ := ioutil.ReadAll(response.Body)
    log.Printf(string(body[:]))
}

/*
   Test Runner
*/

func TestRunDatasiftTestSuite(t *testing.T) {
    suite.Run(t, new(DatasiftTestSuite))
}

/*
 * Datasift Package
 *
 * Tests for client.go
 */

package datasift

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

/*
 * Test Suite Setup
 */

// Client Test Suite
type ClientTestSuite struct {
    suite.Suite

    client  *Client
    headers *map[string]string
}

// Setup Client Test Suite
func (suite *ClientTestSuite) SetupTest() {
    // Client with fake API Credentials
    suite.client = &Client{
        "foo",
        "bar",
    }
    // Test http server default headers
    suite.headers = &map[string]string{
        "Content-Type": "application/json",
    }
}

// A HTTP Test Server for use when requesting the Datasift API to fake
// responses from the API
func (suite *ClientTestSuite) HttpTestServer(
    body string,
    headers *map[string]string) *httptest.Server {

    handler := func(w http.ResponseWriter, r *http.Request) {
        for key, value := range *suite.headers {
            w.Header().Set(key, value)
        }
        fmt.Fprintln(w, body)
    }
    server := httptest.NewServer(http.HandlerFunc(handler))

    return server
}

/*
 * Tests
 */

// Ensure correct authorization header values are returned
func (suite *ClientTestSuite) TestAuthorizationHeaderValue() {
    value := suite.client.authorizationHeaderValue()

    assert.Equal(suite.T(), value, "foo:bar")
}

// Ensure correct API endpoint is returned
func (suite *ClientTestSuite) TestMakeEndpoint() {
    endpoint := suite.client.makeEndpoint([]string{"foo", "bar"})
    expected := "https://api.datasift.com/v1/foo/bar"

    assert.Equal(suite.T(), endpoint, expected)
}

func (suite *ClientTestSuite) TestSuccessfulRequest() {
    ts := suite.HttpTestServer(`{"foo": "bar"}`, suite.headers)
    defer ts.Close()
    response := suite.client.request(ts.URL)

    assert.IsType(suite.T(), &http.Response{}, response)
}

/*
 * Test Suite Runner
 */

func TestClientTestSuite(t *testing.T) {
    suite.Run(t, new(ClientTestSuite))
}

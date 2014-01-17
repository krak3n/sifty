/*
 * Datasift Package
 *
 * Tests for client.go
 */

package datasift

import (
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
    client *Client
}

// Setup Client Test Suite with a basic Client struct
// with fake api credentials
func (suite *ClientTestSuite) SetupTest() {
    suite.client = &Client{
        "foo",
        "bar",
    }
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

/*
 * Test Suite Runner
 */

func TestClientTestSuite(t *testing.T) {
    suite.Run(t, new(ClientTestSuite))
}

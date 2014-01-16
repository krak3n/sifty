/*
 * Datasift Package
 *
 * Tests for client.go
 */

package datasift

import "testing"

func TestAuthorizationHeaderValue(t *testing.T) {
    client := &Client{
        "foo",
        "bar",
    }
    value := client.authorizationHeaderValue()
    if value != "foo:bar" {
        t.Errorf("Expected foo:bar, got %v", value)
    }
}

func TestMakeEndpoint(t *testing.T) {
    client := &Client{
        "foo",
        "bar",
    }
    endpoint := client.makeEndpoint([]string{"foo", "bar"})
    expected := "https://api.datasift.com/v1/foo/bar"
    if endpoint != expected {
        t.Errorf("Expected %v, got %v", expected, endpoint)
    }
}

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

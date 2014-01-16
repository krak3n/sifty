/*
 * Datasift Package
 *
 * Tests for client.go
 */

package datasift

import "testing"

func TestAuthorizationHeaderValue(t *testing.T) {
    credentials := &Credentials{
        "foo",
        "bar",
    }
    value := credentials.authorizationHeaderValue()
    if value != "foo:bar" {
        t.Errorf("Expected foo:bar, got %v", value)
    }
}

// Test the httpstress.Test() function.
package httpstress

import "testing"

func TestTest(t *testing.T) {
	up := []string{"https://google.com", "http://google.com"} // These URLs should pass.
	down := []string{"http://test.invalid"}                   // This should fail.
	invalid := []string{"ftp://localhost"}                    // Error. Non HTTP/HTTPS URL.

	if _, err := Test(1, 1, invalid); err == nil {
		t.Errorf("%s is ok (should be an error)", invalid)
	}

	if err, _ := Test(1, 1, up); len(err) > 0 {
		t.Errorf("%s is down (should be up)", up)
	}

	if err, _ := Test(1, 1, down); len(err) == 0 {
		t.Errorf("%s is up (should be down)", down)
	}
}

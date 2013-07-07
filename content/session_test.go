package content

import (
	"testing"
)

func TestSession(t *testing.T) {
	var (
		args    Args
		session = args.SessionOrCreate("some.driver.default")
	)
	args.SessionId = ""
	session2 := args.SessionOrCreate("some.driver.default")

	session.Set("cache", "a")
	a, b := session.Get("cache"), session2.Get("cache")
	if a != b {
		t.Errorf("a and b aren't equal: %s, %s", a, b)
	}
}

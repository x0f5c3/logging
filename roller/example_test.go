package roller_test

import (
	"github.com/x0f5c3/logging/roller"
	"log"
	"time"
)

// To use roller with the standard library's log package, just pass it into
// the SetOutput function when your application starts.
func Example() {
	l, _ := roller.NewRoller(
		"/var/log/myapp/foo.log",
		500*1024*1024, // 500 megabytes
		&roller.Options{
			MaxBackups: 3,
			MaxAge:     28 * time.Hour * 24, // 28 days
			Compress:   true,
		})
	log.SetOutput(l)
}

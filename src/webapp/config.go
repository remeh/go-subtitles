// Webapp configuration.
//
// Copyright © 2015 - Rémy MATHIEU

package webapp

const (
	DEFAULT_ADDR = ":9000"
)

// The webapp configuration.
type Config struct {
	Addr            string // The addres to listen to
	StaticDirectory string // Directory containing the static files

	Username  string
	Password  string
	Language  string
	UserAgent string
}

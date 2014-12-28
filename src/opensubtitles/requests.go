// Client for OpenSubtitles which use // XML-RPC calls (sic.)
//
// Declaration of every used requests.
//
// Copyright © 2014 - Rémy MATHIEU

package opensubtitles

// Call to log in.
type LogInRequest struct {
	Username  string `xmlrpc:"username"`
	Password  string `xmlrpc:"password"`
	Language  string `xmlrpc:"language"`  // ISO639 2 letter code
	UserAgent string `xmlrpc:"useragent"` // API key to obtain from OpenSubtitles.org
}

// Call to log out.
type LogOutRequest struct {
	Token string
}

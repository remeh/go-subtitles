// Client for OpenSubtitles which use
// XML-RPC calls (sic.)
//
// Declaration of every used responses.
//
// Copyright © 2014 - Rémy MATHIEU

package model

type LogInResponse struct {
	Token   string `xmlrpc:"token"`
	Status  string `xmlrpc:"status"`
	Seconds string `xmlrpc:"seconds"`
}

type LogOutResponse struct {
	Status  string `xmlrpc:"status"`
	Seconds string `xmlrpc:"seconds"`
}

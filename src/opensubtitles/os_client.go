// Client for OpenSubtitles which use
// XML-RPC calls (sic.)
//
// Copyright © 2014 - Rémy MATHIEU

package opensubtitles

import (
	"fmt"

	"github.com/kolo/xmlrpc"
)

const (
	OPENSUBTITLES_API_URL = "http://api.opensubtitles.org/xml-rpc"
)

// A connected client to the OpenSubtitles platform.
type OSClient struct {
	Token    string // Identification token
	Language string // Language given during the opening of the connection

	rpcClient *xmlrpc.Client // The XML-RPC client used.
}

// Log in to the OpenSubtitles platform.
func OpenSubtitlesLogIn(username string, password string, language string, userAgent string) (OSClient, error) {
	rpcClient, err := xmlrpc.NewClient(OPENSUBTITLES_API_URL, nil)

	if err != nil {
		return OSClient{}, nil
	}

	osClient := OSClient{
		rpcClient: rpcClient,
	}

	// Request parameters
	loginRequest := LogInRequest{
		Username:  username,
		Password:  password,
		Language:  language,
		UserAgent: userAgent,
	}

	// Request response
	loginResponse := LogInResponse{}

	// Let's login and analyze the response.
	rpcClient.Call("LogIn", loginRequest, &loginResponse)

	if loginResponse.Status != "200 OK" {
		return osClient, fmt.Errorf("Error code while logging to the OpenSubtitles API : %s\n", loginResponse.Status)
	}

	return osClient, nil
}

// Log out an user. Returns whether or not a 200 has been returned.
func (c *OSClient) LogOut() bool {
	logoutRequest := LogOutRequest{
		Token: c.Token,
	}

	logoutResponse := LogOutResponse{}

	c.rpcClient.Call("LogOut", logoutRequest, &logoutResponse)

	if logoutResponse.Status == "200 OK" {
		return true
	}

	return false
}

// Looks for a subtitle given the video filename.
func (c *OSClient) Search(fileName string) {
}

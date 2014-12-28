// Client for OpenSubtitles which use
// XML-RPC calls (sic.)
//
// Copyright © 2014 - Rémy MATHIEU

package opensubtitles

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"analyzer"
	"opensubtitles/model"

	"github.com/kolo/xmlrpc"
)

const (
	OPENSUBTITLES_API_URL = "http://api.opensubtitles.org/xml-rpc"
)

// A connected client to the OpenSubtitles platform.
type OSClient struct {
	Token     string // Identification token
	UserAgent string // User agent used for identification to OpenSubtitles
	Language  string // Language given during the opening of the connection

	httpClient http.Client // HTTP Client
}

func NewOSClient(language string, userAgent string) OSClient {
	return OSClient{
		Language:  language,
		UserAgent: userAgent,
	}
}

// Log in to the OpenSubtitles platform.
func (c *OSClient) LogIn(username string, password string) error {
	resp, err := c.httpCall("LogIn", username, password, c.Language, c.UserAgent)

	if err != nil {
		return fmt.Errorf("Error code while logging to the OpenSubtitles API : %s\n", err)
	}

	var loginResponse model.LogInResponse
	resp.Unmarshal(&loginResponse)

	c.Token = loginResponse.Token

	return nil
}

// Log out an user. Returns whether or not a 200 has been returned.
func (c *OSClient) LogOut() error {
	resp, err := c.httpCall("LogOut", c.Token)

	if err != nil {
		return fmt.Errorf("Error code while logging to the OpenSubtitles API : %s\n", err)
	}

	var logoutResponse model.LogOutResponse
	resp.Unmarshal(&logoutResponse)

	if logoutResponse.Status != "200 OK" {
		return fmt.Errorf("Bad status code returned during log out :%s\n", logoutResponse.Status)
	}

	return nil
}

// Looks for a subtitle given the video filename.
func (c *OSClient) Search(filename string, language string, limit int) error {
	// Builds the query
	result := analyzer.AnalyzeFilename(filename)

	// Builds the query with the analysis result.
	filters := make([]map[string]string, 0)
	filter := make(map[string]string)
	filter["query"] = result.Name
    filter["sublanguageid"] = language
	filters = append(filters, filter)

	// Query options, currently, we just put a limit.
	options := make(map[string]int)
	options["limit"] = limit

	resp, err := c.httpCall("SearchSubtitles", c.Token, filters, options)

	if err != nil {
		return fmt.Errorf("Error code while logging to the OpenSubtitles API : %s\n", err)
	}

	var searchResponse model.SearchSubtitlesResponse
	err = resp.Unmarshal(&searchResponse)

    if err != nil {
        return err
    }

	if searchResponse.Status != "200 OK" {
		return fmt.Errorf("Bad status code returned during search query :%s\n", searchResponse.Status)
	}

    for _, v := range searchResponse.SubtitleEntries {
        fmt.Println(v)
    }

	return nil
}

// Does the XML-RPC over HTTP call.
func (c *OSClient) httpCall(method string, parameters ...interface{}) (*xmlrpc.Response, error) {
	request, err := xmlrpc.NewRequest(OPENSUBTITLES_API_URL, method, parameters)
	request.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Bad HTTP code : %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return xmlrpc.NewResponse(data), nil
}

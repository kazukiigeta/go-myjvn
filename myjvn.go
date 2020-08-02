// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

// Parameter represents all the parameters of API except method and feed.
type Parameter struct {
	StartItem                uint
	MaxCountItem             uint8
	DatePublished            uint16
	DateFirstPublished       uint16
	CPEName                  string
	Format                   string
	Keyword                  string
	Language                 string
	VendorID                 string
	ProductID                string
	Severity                 string
	Vector                   string
	RangeDatePublic          string
	RangeDatePublished       string
	RangeDateFirstPublished  string
	DatePublicStartY         uint16
	DatePublicStartM         uint8
	DatePublicStartD         uint8
	DatePublicEndY           uint16
	DatePublicEndM           uint8
	DatePublicEndD           uint8
	DateFirstPublishedStartY uint16
	DateFirstPublishedStartM uint8
	DateFirstPublishedStartD uint8
	DateFirstPublishedEndY   uint16
	DateFirstPublishedEndM   uint8
	DateFirstPublishedEndD   uint8
}

// Status stores the data from API response.
type Status struct {
	Version     string `xml:"version,attr"`
	Method      string `xml:"method,attr"`
	Language    string `xml:"lang,attr"`
	RetCd       uint   `xml:"retCd,attr"`
	RetMax      string `xml:"retMax,attr"`
	ErrCd       string `xml:"errCd,attr"`
	ErrMsg      string `xml:"errMsg,attr"`
	TotalRes    string `xml:"totalRes,attr"`
	TotalResRet string `xml:"totalResRet,attr"`
	FirstRes    string `xml:"firstRes,attr"`
	Feed        string `xml:"feed,attr"`
}

// Default settings of REST API
const (
	defaultBaseURL string = "https://jvndb.jvn.jp/"
	defaultAPIPath string = "/myjvn"
)

// A Client manages communication with the MyJVN API.
type Client struct {
	BaseURL *url.URL

	httpClient *http.Client
}

// NewClient creates an instance of Client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		BaseURL:    baseURL,
		httpClient: httpClient,
	}

	return c
}

// newRequest create an API request.
func (c *Client) newRequest(method, path string) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

var strJSON string = "json"
var strXML string = "xml"

// do decodes HTTP response to store the data into the struct given as v.
func (c *Client) do(ctx context.Context, req *http.Request, format *string, v interface{}) error {
	if v == nil || reflect.ValueOf(v).IsNil() {
		return fmt.Errorf("v must not be nil")
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var decoder interface {
		Decode(interface{}) error
	}
	if format == nil || *format == strXML {
		d := xml.NewDecoder(resp.Body)
		d.Strict = false
		decoder = d
	} else if *format == strJSON {
		decoder = json.NewDecoder(resp.Body)
	} else {
		return fmt.Errorf(`format must be either nil, "xml" or "json"`)
	}

	decErr := decoder.Decode(v)
	if decErr == io.EOF {
		decErr = nil
	}
	if decErr != nil {
		err = decErr
	}

	return err
}

// addOptions returnes a string added query strings to URL path given as s.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}
	u.RawQuery = qs.Encode()

	return u.String(), nil
}

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

// parameter represents all the parameters of API except method and feed.
type parameter struct {
	Method                   string `url:"method"`
	Feed                     string `url:"feed"`
	StartItem                uint   `url:"startItem,omitempty"`
	MaxCountItem             uint8  `url:"maxCountItem,omitempty"`
	DatePublished            uint16 `url:"datePublished,omitempty"`
	DateFirstPublished       uint16 `url:"dateFirstPublished,omitempty"`
	CPEName                  string `url:"cpeName,omitempty"`
	Format                   string `url:"ft,omitempty"`
	Keyword                  string `url:"keyword,omitempty"`
	Language                 string `url:"language,omitempty"`
	VendorID                 string `url:"vendorId,omitempty"`
	ProductID                string `url:"productId,omitempty"`
	VulnID                   string `url:"vulnId,omitempty"`
	Severity                 string `url:"severity,omitempty"`
	Vector                   string `url:"vector,omitempty"`
	RangeDatePublic          string `url:"rangeDatePublic,omitempty"`
	RangeDatePublished       string `url:"rangeDatePublished,omitempty"`
	RangeDateFirstPublished  string `url:"rangeDateFirstPublished,omitempty"`
	DatePublicStartY         uint16 `url:"datePublicStartY,omitempty"`
	DatePublicStartM         uint8  `url:"datePublicStartM,omitempty"`
	DatePublicStartD         uint8  `url:"datePublicStartD,omitempty"`
	DatePublicEndY           uint16 `url:"datePublicEndY,omitempty"`
	DatePublicEndM           uint8  `url:"datePublicEndM,omitempty"`
	DatePublicEndD           uint8  `url:"datePublicEndD,omitempty"`
	DateFirstPublishedStartY uint16 `url:"dateFirstPublicStartY,omitempty"`
	DateFirstPublishedStartM uint8  `url:"dateFirstPublicStartM,omitempty"`
	DateFirstPublishedStartD uint8  `url:"dateFirstPublicStartD,omitempty"`
	DateFirstPublishedEndY   uint16 `url:"dateFirstPublicEndY,omitempty"`
	DateFirstPublishedEndM   uint8  `url:"dateFirstPublicEndM,omitempty"`
	DateFirstPublishedEndD   uint8  `url:"dateFirstPublicEndD,omitempty"`
	Theme                    string `url:"theme,omitempty"`
	AggrType                 string `url:"type,omitempty"`
	CWEID                    string `url:"cweId,omitempty"`
	PID                      uint   `url:"pid,omitempty"`
}

// Option AAA
type Option func(p *parameter)

// SetMethod BBB
func SetMethod(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Method = s
		}
	}
}

// SetFeed BBB
func SetFeed(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Feed = s
		}
	}
}

// SetStartItem BBB
func SetStartItem(u uint) Option {
	return func(p *parameter) {
		if p != nil {
			p.StartItem = u
		}
	}
}

// SetMaxCountItem BBB
func SetMaxCountItem(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.MaxCountItem = u
		}
	}
}

// SetDatePublished BBB
func SetDatePublished(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublished = u
		}
	}
}

// SetDateFirstPublished BBB
func SetDateFirstPublished(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublished = u
		}
	}
}

// SetCPEName BBB
func SetCPEName(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.CPEName = s
		}
	}
}

// SetFormat BBB
func SetFormat(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Format = s
		}
	}
}

// SetKeyword BBB
func SetKeyword(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Keyword = url.QueryEscape(s)
		}
	}
}

// SetLanguage BBB
func SetLanguage(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Language = s
		}
	}
}

// SetVendorID BBB
func SetVendorID(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.VendorID = s
		}
	}
}

// SetProductID BBB
func SetProductID(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.ProductID = s
		}
	}
}

// SetVulnID BBB
func SetVulnID(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.VulnID = s
		}
	}
}

// SetSeverity BBB
func SetSeverity(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Severity = s
		}
	}
}

// SetVector BBB
func SetVector(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Vector = s
		}
	}
}

// SetRangeDatePublic BBB
func SetRangeDatePublic(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.RangeDatePublic = s
		}
	}
}

// SetRangeDatePublished BBB
func SetRangeDatePublished(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.RangeDatePublished = s
		}
	}
}

// SetRangeDateFirstPublished BBB
func SetRangeDateFirstPublished(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.RangeDateFirstPublished = s
		}
	}
}

// SetDatePublicStartY BBB
func SetDatePublicStartY(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicStartY = u
		}
	}
}

// SetDatePublicStartM BBB
func SetDatePublicStartM(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicStartM = u
		}
	}
}

// SetDatePublicStartD BBB
func SetDatePublicStartD(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicStartD = u
		}
	}
}

// SetDatePublicEndY BBB
func SetDatePublicEndY(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicEndY = u
		}
	}
}

// SetDatePublicEndM BBB
func SetDatePublicEndM(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicEndM = u
		}
	}
}

// SetDatePublicEndD BBB
func SetDatePublicEndD(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicEndD = u
		}
	}
}

// SetDateFirstPublishedStartY BBB
func SetDateFirstPublishedStartY(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedStartY = u
		}
	}
}

// SetDateFirstPublishedStartM BBB
func SetDateFirstPublishedStartM(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedStartM = u
		}
	}
}

// SetDateFirstPublishedStartD BBB
func SetDateFirstPublishedStartD(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedStartD = u
		}
	}
}

// SetDateFirstPublishedEndY BBB
func SetDateFirstPublishedEndY(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedEndY = u
		}
	}
}

// SetDateFirstPublishedEndM BBB
func SetDateFirstPublishedEndM(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedEndM = u
		}
	}
}

// SetDateFirstPublishedEndD BBB
func SetDateFirstPublishedEndD(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedEndD = u
		}
	}
}

// SetTheme BBB
func SetTheme(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Theme = s
		}
	}
}

// SetAggrType BBB
func SetAggrType(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.AggrType = s
		}
	}
}

// SetCWEID BBB
func SetCWEID(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.CWEID = s
		}
	}
}

// SetPID BBB
func SetPID(u uint) Option {
	return func(p *parameter) {
		if p != nil {
			p.PID = u
		}
	}
}

// Status stores the data from API response.
type Status struct {
	Version                  string `xml:"version,attr"`
	Method                   string `xml:"method,attr"`
	Language                 string `xml:"lang,attr"`
	RetCd                    uint   `xml:"retCd,attr"`
	RetMax                   string `xml:"retMax,attr"`
	RetMaxCnt                string `xml:"retMaxCnt,attr"`
	ErrCd                    string `xml:"errCd,attr"`
	ErrMsg                   string `xml:"errMsg,attr"`
	TotalRes                 string `xml:"totalRes,attr"`
	TotalResRet              string `xml:"totalResRet,attr"`
	FirstRes                 string `xml:"firstRes,attr"`
	Feed                     string `xml:"feed,attr"`
	StartItem                uint   `xml:"startItem,attr"`
	MaxCountItem             uint8  `xml:"maxCountItem,attr"`
	DatePublished            uint16 `xml:"datePublished,attr"`
	DateFirstPublished       uint16 `xml:"dateFirstPublished,attr"`
	CPEName                  string `xml:"cpeName,attr"`
	Format                   string `xml:"format,attr"`
	Keyword                  string `xml:"keyword,attr"`
	VendorID                 string `xml:"vendorId,attr"`
	ProductID                string `xml:"productId,attr"`
	VulnID                   string `xml:"vulnId,attr"`
	Severity                 string `xml:"severity,attr"`
	Vector                   string `xml:"vector,attr"`
	RangeDatePublic          string `xml:"rangeDatePublic,attr"`
	RangeDatePublished       string `xml:"rangeDatePublished,attr"`
	RangeDateFirstPublished  string `xml:"rangeDateFirstPublished,attr"`
	DatePublicStartY         uint16 `xml:"datePublicStartY,attr"`
	DatePublicStartM         uint8  `xml:"datePublicStartM,attr"`
	DatePublicStartD         uint8  `xml:"datePublicStartD,attr"`
	DatePublicEndY           uint16 `xml:"datePublicEndY,attr"`
	DatePublicEndM           uint8  `xml:"datePublicEndM,attr"`
	DatePublicEndD           uint8  `xml:"datePublicEndD,attr"`
	DateFirstPublishedStartY uint16 `xml:"dateFirstPublishedStartY,attr"`
	DateFirstPublishedStartM uint8  `xml:"dateFirstPublishedStartM,attr"`
	DateFirstPublishedStartD uint8  `xml:"dateFirstPublishedStartD,attr"`
	DateFirstPublishedEndY   uint16 `xml:"dateFirstPublishedEndY,attr"`
	DateFirstPublishedEndM   uint8  `xml:"dateFirstPublishedEndM,attr"`
	DateFirstPublishedEndD   uint8  `xml:"dateFirstPublishedEndD,attr"`
	Theme                    string `xml:"theme,attr"`
	AggrType                 string `xml:"type,attr"`
	CWEID                    string `xml:"cweId,attr"`
	PID                      uint   `xml:"pid,attr"`
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

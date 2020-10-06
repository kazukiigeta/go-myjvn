// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/cenkalti/backoff"
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

// Option represents a function of which a parameter is REST API parameter.
type Option func(p *parameter)

// SetMethod returns Option which sets Method as REST API parameter.
func SetMethod(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Method = s
		}
	}
}

// SetFeed returns Option which sets Feed as REST API parameter.
func SetFeed(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Feed = s
		}
	}
}

// SetStartItem returns Option which sets StartItem as REST API parameter.
func SetStartItem(u uint) Option {
	return func(p *parameter) {
		if p != nil {
			p.StartItem = u
		}
	}
}

// SetMaxCountItem returns Option which sets MaxCountItem as REST API parameter.
func SetMaxCountItem(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.MaxCountItem = u
		}
	}
}

// SetDatePublished returns Option which sets DatePublished as REST API parameter.
func SetDatePublished(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublished = u
		}
	}
}

// SetDateFirstPublished returns Option which sets DateFirstPublished as REST API parameter.
func SetDateFirstPublished(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublished = u
		}
	}
}

// SetCPEName returns Option which sets CPEName as REST API parameter.
func SetCPEName(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.CPEName = s
		}
	}
}

// SetFormat returns Option which sets Format as REST API parameter.
func SetFormat(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Format = s
		}
	}
}

// SetKeyword returns Option which sets Keyword as REST API parameter.
func SetKeyword(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Keyword = s
		}
	}
}

// SetLanguage returns Option which sets Language as REST API parameter.
func SetLanguage(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Language = s
		}
	}
}

// SetVendorID returns Option which sets VendorID as REST API parameter.
func SetVendorID(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.VendorID = s
		}
	}
}

// SetProductID returns Option which sets ProductID as REST API parameter.
func SetProductID(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.ProductID = s
		}
	}
}

// SetVulnID returns Option which sets VulnID as REST API parameter.
func SetVulnID(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.VulnID = s
		}
	}
}

// SetSeverity returns Option which sets Severity as REST API parameter.
func SetSeverity(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Severity = s
		}
	}
}

// SetVector returns Option which sets Vector as REST API parameter.
func SetVector(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Vector = s
		}
	}
}

// SetRangeDatePublic returns Option which sets RangeDatePublic as REST API parameter.
func SetRangeDatePublic(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.RangeDatePublic = s
		}
	}
}

// SetRangeDatePublished returns Option which sets RangeDatePublished as REST API parameter.
func SetRangeDatePublished(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.RangeDatePublished = s
		}
	}
}

// SetRangeDateFirstPublished returns Option which sets RangeDateFirstPublished as REST API parameter.
func SetRangeDateFirstPublished(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.RangeDateFirstPublished = s
		}
	}
}

// SetDatePublicStartY returns Option which sets DatePublicStartY as REST API parameter.
func SetDatePublicStartY(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicStartY = u
		}
	}
}

// SetDatePublicStartM returns Option which sets DatePublicStartM as REST API parameter.
func SetDatePublicStartM(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicStartM = u
		}
	}
}

// SetDatePublicStartD returns Option which sets DatePublicStartD as REST API parameter.
func SetDatePublicStartD(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicStartD = u
		}
	}
}

// SetDatePublicEndY returns Option which sets DatePublicEndY as REST API parameter.
func SetDatePublicEndY(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicEndY = u
		}
	}
}

// SetDatePublicEndM returns Option which sets DatePublicEndM as REST API parameter.
func SetDatePublicEndM(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicEndM = u
		}
	}
}

// SetDatePublicEndD returns Option which sets DatePublicEndD as REST API parameter.
func SetDatePublicEndD(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DatePublicEndD = u
		}
	}
}

// SetDateFirstPublishedStartY returns Option which sets DatePublishedStartY as REST API parameter.
func SetDateFirstPublishedStartY(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedStartY = u
		}
	}
}

// SetDateFirstPublishedStartM returns Option which sets DatePublishedStartM as REST API parameter.
func SetDateFirstPublishedStartM(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedStartM = u
		}
	}
}

// SetDateFirstPublishedStartD returns Option which sets DatePublishedStartD as REST API parameter.
func SetDateFirstPublishedStartD(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedStartD = u
		}
	}
}

// SetDateFirstPublishedEndY returns Option which sets DatePublishedEndY as REST API parameter.
func SetDateFirstPublishedEndY(u uint16) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedEndY = u
		}
	}
}

// SetDateFirstPublishedEndM returns Option which sets DatePublishedEndM as REST API parameter.
func SetDateFirstPublishedEndM(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedEndM = u
		}
	}
}

// SetDateFirstPublishedEndD returns Option which sets DatePublishedEndD as REST API parameter.
func SetDateFirstPublishedEndD(u uint8) Option {
	return func(p *parameter) {
		if p != nil {
			p.DateFirstPublishedEndD = u
		}
	}
}

// SetTheme returns Option which sets Theme as REST API parameter.
func SetTheme(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.Theme = s
		}
	}
}

// SetAggrType returns Option which sets AggrType as REST API parameter.
func SetAggrType(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.AggrType = s
		}
	}
}

// SetCWEID returns Option which sets CWEID as REST API parameter.
func SetCWEID(s string) Option {
	return func(p *parameter) {
		if p != nil {
			p.CWEID = s
		}
	}
}

// SetPID returns Option which sets PID as REST API parameter.
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

// IsErrorRetryable shows HTTP response error is retryable or not.
func IsErrorRetryable(resp *http.Response) bool {
	return resp.StatusCode >= http.StatusInternalServerError
}

var strJSON string = "json"
var strXML string = "xml"

// DoWithRetry execute do with retry functionality.
func (c *Client) DoWithRetry(req *http.Request) (*http.Response, error) {
	maxRetryNumber := uint64(7)
	var resp *http.Response
	var err error

	operationWithRetry := func() error {
		resp, err = c.httpClient.Do(req)
		if err == nil && IsErrorRetryable(resp) {
			err = errors.New("retryable")
		}
		return err
	}

	bo := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetryNumber)
	err = backoff.Retry(operationWithRetry, bo)

	return resp, err
}

// do decodes HTTP response to store the data into the struct given as v.
func (c *Client) do(ctx context.Context, req *http.Request, format *string, v interface{}) error {
	if v == nil || reflect.ValueOf(v).IsNil() {
		return fmt.Errorf("v must not be nil")
	}

	req = req.WithContext(ctx)

	// resp, err := c.httpClient.Do(req)
	resp, err := c.DoWithRetry(req)
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

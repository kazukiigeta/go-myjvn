// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSetFunc(t *testing.T) {
	type testCase struct {
		description string
		setFunc     Option
		before      *parameter
		after       *parameter
	}

	var testcases = []testCase{
		{
			description: "SetMethod",
			setFunc:     SetMethod("getAlertList"),
			before:      &parameter{},
			after:       &parameter{Method: "getAlertList"},
		},
		{
			description: "SetMethod without parameter",
			setFunc:     SetMethod("getAlertList"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetFeed",
			setFunc:     SetFeed("itm"),
			before:      &parameter{},
			after:       &parameter{Feed: "itm"},
		},
		{
			description: "SetFeed without parameter",
			setFunc:     SetFeed("itm"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetStartItem",
			setFunc:     SetStartItem(1),
			before:      &parameter{},
			after:       &parameter{StartItem: 1},
		},
		{
			description: "SetStartItem without parameter",
			setFunc:     SetStartItem(1),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetMaxCountItem",
			setFunc:     SetMaxCountItem(1),
			before:      &parameter{},
			after:       &parameter{MaxCountItem: 1},
		},
		{
			description: "SetMaxCountItem without parameter",
			setFunc:     SetMaxCountItem(1),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDatePublished",
			setFunc:     SetDatePublished(2020),
			before:      &parameter{},
			after:       &parameter{DatePublished: 2020},
		},
		{
			description: "SetDatePublished without parameter",
			setFunc:     SetDatePublished(2020),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDateFirstPublished",
			setFunc:     SetDateFirstPublished(2020),
			before:      &parameter{},
			after:       &parameter{DateFirstPublished: 2020},
		},
		{
			description: "SetDateFirstPublished without parameter",
			setFunc:     SetDateFirstPublished(2020),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetCPEName",
			setFunc:     SetCPEName("cpe:/o:google:android 2.2"),
			before:      &parameter{},
			after:       &parameter{CPEName: "cpe:/o:google:android 2.2"},
		},
		{
			description: "SetCPEName without parameter",
			setFunc:     SetCPEName("cpe:/o:google:android 2.2"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetFormat",
			setFunc:     SetFormat("json"),
			before:      &parameter{},
			after:       &parameter{Format: "json"},
		},
		{
			description: "SetFormat without parameter",
			setFunc:     SetFormat("json"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetKeyword",
			setFunc:     SetKeyword("android"),
			before:      &parameter{},
			after:       &parameter{Keyword: "android"},
		},
		{
			description: "SetKeyword without parameter",
			setFunc:     SetKeyword("android"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetVendorID",
			setFunc:     SetVendorID("4499"),
			before:      &parameter{},
			after:       &parameter{VendorID: "4499"},
		},
		{
			description: "SetVendorID without parameter",
			setFunc:     SetVendorID("4499"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetProductID",
			setFunc:     SetProductID("100"),
			before:      &parameter{},
			after:       &parameter{ProductID: "100"},
		},
		{
			description: "SetProductID without parameter",
			setFunc:     SetProductID("100"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetVulnID",
			setFunc:     SetVulnID("JVNDB-2020-008081"),
			before:      &parameter{},
			after:       &parameter{VulnID: "JVNDB-2020-008081"},
		},
		{
			description: "SetVulnID without parameter",
			setFunc:     SetVulnID("JVNDB-2020-008081"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetSeverity",
			setFunc:     SetSeverity("m"),
			before:      &parameter{},
			after:       &parameter{Severity: "m"},
		},
		{
			description: "SetSeverity without parameter",
			setFunc:     SetSeverity("m"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetVector",
			setFunc:     SetVector("AV:N"),
			before:      &parameter{},
			after:       &parameter{Vector: "AV:N"},
		},
		{
			description: "SetVector without parameter",
			setFunc:     SetVector("AV:N"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetRangeDatePublic",
			setFunc:     SetRangeDatePublic("n"),
			before:      &parameter{},
			after:       &parameter{RangeDatePublic: "n"},
		},
		{
			description: "SetRangeDatePublic without parameter",
			setFunc:     SetRangeDatePublic("n"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetRangeDatePublished",
			setFunc:     SetRangeDatePublished("n"),
			before:      &parameter{},
			after:       &parameter{RangeDatePublished: "n"},
		},
		{
			description: "SetRangeDatePublished without parameter",
			setFunc:     SetRangeDatePublished("n"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetRangeDateFirstPublished",
			setFunc:     SetRangeDateFirstPublished("n"),
			before:      &parameter{},
			after:       &parameter{RangeDateFirstPublished: "n"},
		},
		{
			description: "SetRangeDateFirstPublished without parameter",
			setFunc:     SetRangeDateFirstPublished("n"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDatePublicStartY",
			setFunc:     SetDatePublicStartY(2020),
			before:      &parameter{},
			after:       &parameter{DatePublicStartY: 2020},
		},
		{
			description: "SetDatePublicStartY without parameter",
			setFunc:     SetDatePublicStartY(2020),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDatePublicStartM",
			setFunc:     SetDatePublicStartM(12),
			before:      &parameter{},
			after:       &parameter{DatePublicStartM: 12},
		},
		{
			description: "SetDatePublicStartM without parameter",
			setFunc:     SetDatePublicStartM(12),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDatePublicStartD",
			setFunc:     SetDatePublicStartD(2),
			before:      &parameter{},
			after:       &parameter{DatePublicStartD: 2},
		},
		{
			description: "SetDatePublicStartD without parameter",
			setFunc:     SetDatePublicStartD(2),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDatePublicEndY",
			setFunc:     SetDatePublicEndY(2020),
			before:      &parameter{},
			after:       &parameter{DatePublicEndY: 2020},
		},
		{
			description: "SetDatePublicEndY without parameter",
			setFunc:     SetDatePublicEndY(2020),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDatePublicEndM",
			setFunc:     SetDatePublicEndM(12),
			before:      &parameter{},
			after:       &parameter{DatePublicEndM: 12},
		},
		{
			description: "SetDatePublicEndM without parameter",
			setFunc:     SetDatePublicEndM(12),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDatePublicEndD",
			setFunc:     SetDatePublicEndD(2),
			before:      &parameter{},
			after:       &parameter{DatePublicEndD: 2},
		},
		{
			description: "SetDatePublicEndD without parameter",
			setFunc:     SetDatePublicEndD(2),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDateFirstPublishedStartY",
			setFunc:     SetDateFirstPublishedStartY(2020),
			before:      &parameter{},
			after:       &parameter{DateFirstPublishedStartY: 2020},
		},
		{
			description: "SetDateFirstPublishedStartY without parameter",
			setFunc:     SetDateFirstPublishedStartY(2020),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDateFirstPublishedStartM",
			setFunc:     SetDateFirstPublishedStartM(12),
			before:      &parameter{},
			after:       &parameter{DateFirstPublishedStartM: 12},
		},
		{
			description: "SetDateFirstPublishedStartM without parameter",
			setFunc:     SetDateFirstPublishedStartM(12),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDateFirstPublishedStartD",
			setFunc:     SetDateFirstPublishedStartD(2),
			before:      &parameter{},
			after:       &parameter{DateFirstPublishedStartD: 2},
		},
		{
			description: "SetDateFirstPublishedStartD without parameter",
			setFunc:     SetDateFirstPublishedStartD(2),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDateFirstPublishedEndY",
			setFunc:     SetDateFirstPublishedEndY(2020),
			before:      &parameter{},
			after:       &parameter{DateFirstPublishedEndY: 2020},
		},
		{
			description: "SetDateFirstPublishedEndY without parameter",
			setFunc:     SetDateFirstPublishedEndY(2020),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDateFirstPublishedEndM",
			setFunc:     SetDateFirstPublishedEndM(12),
			before:      &parameter{},
			after:       &parameter{DateFirstPublishedEndM: 12},
		},
		{
			description: "SetDateFirstPublishedEndM without parameter",
			setFunc:     SetDateFirstPublishedEndM(12),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetDateFirstPublishedEndD",
			setFunc:     SetDateFirstPublishedEndD(2),
			before:      &parameter{},
			after:       &parameter{DateFirstPublishedEndD: 2},
		},
		{
			description: "SetDateFirstPublishedEndD without parameter",
			setFunc:     SetDateFirstPublishedEndD(2),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetTheme",
			setFunc:     SetTheme("sumCvss"),
			before:      &parameter{},
			after:       &parameter{Theme: "sumCvss"},
		},
		{
			description: "SetTheme without parameter",
			setFunc:     SetTheme("sumCvss"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetAggrType",
			setFunc:     SetAggrType("m"),
			before:      &parameter{},
			after:       &parameter{AggrType: "m"},
		},
		{
			description: "SetAggrType without parameter",
			setFunc:     SetAggrType("m"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetCWEID",
			setFunc:     SetCWEID("CWE-20"),
			before:      &parameter{},
			after:       &parameter{CWEID: "CWE-20"},
		},
		{
			description: "SetCWEID without parameter",
			setFunc:     SetCWEID("CWE-20"),
			before:      nil,
			after:       nil,
		},
		{
			description: "SetPID",
			setFunc:     SetPID(4499),
			before:      &parameter{},
			after:       &parameter{PID: 4499},
		},
		{
			description: "SetPID without parameter",
			setFunc:     SetPID(4499),
			before:      nil,
			after:       nil,
		},
	}

	for _, c := range testcases {
		t.Run(c.description, func(t *testing.T) {
			p := c.before
			c.setFunc(p)

			want, got := c.after, p
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

		})
	}
}

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(defaultAPIPath+"/", http.StripPrefix(defaultAPIPath, mux))

	server := httptest.NewServer(apiHandler)

	client = NewClient(nil)
	url, _ := url.Parse(server.URL + defaultAPIPath + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func TestDo(t *testing.T) {
	type sampleStruct struct {
		A string `json:"a"`
		B string `json:"b"`
	}

	type testCase struct {
		description string
		format      *string
		structured  *sampleStruct
		serialized  string
		err         error
	}

	var testcases = []testCase{
		{
			description: "Not Passing a struct without specifying format",
			format:      nil,
			structured:  nil,
			err:         errors.New("v must not be nil"),
		},
		{
			description: "Passing a struct with specifying JSON format",
			format:      &strJSON,
			structured: &sampleStruct{
				A: "a",
				B: "b",
			},
			serialized: `{"A":"a", "B":"b"}`,
			err:        nil,
		},
	}

	for _, c := range testcases {
		t.Run(c.description, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, c.serialized)
			})

			req, _ := client.newRequest("GET", ".")

			err := client.do(context.Background(), req, c.format, c.structured)
			if c.err == nil && err != c.err {
				t.Fatalf("do returns unexpected error: %s", err)
			}
			if err != nil && err.Error() != c.err.Error() {
				t.Fatalf("do returns unexpected error: %s", err)
			}

			want, got := c.structured, c.structured
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestDo_noContent(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.newRequest("GET", ".")
	err := client.do(context.Background(), req, nil, &body)
	if err != nil {
		t.Fatalf("do returned unexpected error: %v", err)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}

	c2 := NewClient(nil)
	if c.httpClient == c2.httpClient {
		t.Error("NewClient returned same http.Clients, but they should differ")
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	req, _ := c.newRequest("GET", inURL)

	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("newRequest(%q) URL is %v, want %v", inURL, got, want)
	}
}

func TestNewRequest_InvalidPath(t *testing.T) {
	c := NewClient(nil)

	inURL := "%invalidPath%"
	_, err := c.newRequest("GET", inURL)
	if err == nil {
		t.Fatalf("newRequest must return an error of invalid URL escape")
	}
}

func TestAddOptions(t *testing.T) {
	type option struct {
		Param1 string `url:"str"`
		Param2 int    `url:"int"`
		Param3 bool   `url:"bool"`
	}
	opt := &option{"a", 0, false}
	u, err := addOptions("path", opt)
	if err != nil {
		t.Fail()
	}

	inPath := "path"
	outQuery := inPath + "?bool=false&int=0&str=a"
	if got, want := u, outQuery; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

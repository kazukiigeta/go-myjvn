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

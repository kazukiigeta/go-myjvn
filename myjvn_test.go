// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
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
	client, mux, _, teardown := setup()
	defer teardown()

	type foo struct {
		A string `json:"A"`
		B string `json:"B"`
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"A": "a", "B":"b"}`)
	})

	req, _ := client.newRequest("GET", ".")
	body := new(foo)
	if _, err := client.do(context.Background(), req, body); err != nil {
		t.Errorf("do returned error: %v", err)
	}

	want := &foo{"a", "b"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
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
	_, err := client.do(context.Background(), req, &body)
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

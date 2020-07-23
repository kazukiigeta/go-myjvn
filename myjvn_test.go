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

func TestNewParamsGetAlertList(t *testing.T) {
	var startItem uint = 10
	var maxCountItem uint8 = 3
	var datePublished uint16 = 2020
	var dateFirstPublished uint16 = 2020
	var cpeName string = "cpe:/*"

	got := NewParamsGetAlertList(
		&startItem, &maxCountItem, &datePublished, &dateFirstPublished, &cpeName,
	)

	want := &ParamsGetAlertList{
		Method:             "getAlertList",
		Feed:               "hnd",
		StartItem:          startItem,
		MaxCountItem:       maxCountItem,
		DatePublished:      datePublished,
		DateFirstPublished: dateFirstPublished,
		CpeName:            cpeName,
		Format:             "json",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}

func TestGetAlertList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	var expectedHTTPResp = `{"feed":{"title":"title","id":"id","author":{"name":"name",
"uri":"http://www.example.com/"},"updated":"2020-07-22T16:00:40+09:00","link":"http://www.example.com/",
"sec:handling":{"marking:Marking":{"marking:Marking_Structure":{"xsi:type":"xsitype","marking_model_name":"TLP",
"marking_model_ref":"http://www.example.com/","color":"WHITE"}}},"entry":[{"title":"title",
"id":"MYJVN-ALT-0000-0000","link":"http://www.example.com/","summary":"summary"}],
"status:Status":{"version":"3.3","method":"getAlertList","retCd":0,"retMax":"50","errCd":"0","errMsg":"errmsg",
"totalRes":"28","totalResRet":"0","firstRes":"1","maxCountItem":"2","cpeName":"cpe:/*","ft":"json","feed":"hnd"}}}`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedHTTPResp)
	})

	p := NewParamsGetAlertList(nil, nil, nil, nil, nil)
	alertList, _, err := client.GetAlertList(context.Background(), p)
	if err != nil {
		t.Fatalf("GetAlertList returned error: %v", err)
	}

	a := &AlertList{
		Feed: Feed{
			Title: "title",
			ID:    "id",
			Author: Author{
				Name: "name",
				URI:  "http://www.example.com/",
			},
			Updated: "2020-07-22T16:00:40+09:00",
			Link:    "http://www.example.com/",
			SecHandling: SecHandling{
				Marking: Marking{
					MarkingStructure: MarkingStructure{
						XSIType:   "xsitype",
						ModelName: "TLP",
						ModelRef:  "http://www.example.com/",
						Color:     "WHITE",
					},
				},
			},
			Entry: []*Entry{
				&Entry{
					Title:   "title",
					ID:      "MYJVN-ALT-0000-0000",
					Link:    "http://www.example.com/",
					Summary: "summary",
				},
			},
			Status: Status{
				Version:     "3.3",
				Method:      "getAlertList",
				RetCd:       0,
				RetMax:      "50",
				ErrCd:       "0",
				ErrMsg:      "errmsg",
				TotalRes:    "28",
				TotalResRet: "0",
				FirstRes:    "1",
				Format:      "json",
				Feed:        "hnd",
			},
		},
	}

	got, want := alertList, a
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}

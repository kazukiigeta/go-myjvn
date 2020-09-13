// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAlertList(t *testing.T) {
	var expectedHTTPResp = `<feed xsi:schemaLocation="http://www.w3.org/2005/Atom https://jvndb.jvn.jp/schema/atom.xsd http://jvn.jp/rss/mod_sec/3.0/ https://jvndb.jvn.jp/schema/mod_sec_3.0.xsd http://data-marking.mitre.org/extensions/MarkingStructure#TLP-1 https://jvndb.jvn.jp/schema/tlp_marking.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd" xml:lang="ja"><title type="text">title</title><updated>2020-07-22T16:00:40+09:00</updated><id>id</id><link rel="alternate" type="text/html" hreflang="ja" href="http://example.com/"/><author><name>name</name><uri>http://example.com/</uri></author><sec:handling><marking:Marking><marking:Marking_Structure xsi:type="xsitype" marking_model_name="TLP" marking_model_ref="http://example.com/" color="WHITE"/></marking:Marking></sec:handling><entry><title>title</title><id>id</id><published>2020-01-15T12:04:14+09:00</published><updated>2020-01-15T12:04:14+09:00</updated><category label="label" term="term"/><sec:items><sec:item><sec:title>Microsoft 製品の脆弱性対策について(2020年1月)</sec:title><sec:identifier>MYJVN-ALT-2020-0002-0001</sec:identifier><sec:link href="https://www.ipa.go.jp/security/ciadr/vul/20200115-ms.html"/><sec:published>2020-01-15T00:00:00+09:00</sec:published><sec:updated>2020-01-15T12:04:14+09:00</sec:updated></sec:item><sec:item><sec:title>2020 年 1 月のセキュリティ更新プログラム (月例)</sec:title><sec:identifier>MYJVN-ALT-2020-0002-0002</sec:identifier><sec:link href="https://msrc-blog.microsoft.com/2020/01/14/202001-security-updates/"/><sec:published>2020-01-15T00:00:00+09:00</sec:published><sec:updated>2020-01-15T12:04:14+09:00</sec:updated></sec:item></sec:items></entry><status:Status version="3.3" method="getAlertList" retCd="0" retMax="50" errCd="errcd" errMsg="errmsg" totalRes="54" totalResRet="1" firstRes="1" feed="hnd" maxCountItem="1"/></feed>`

	expectedAlertList := &AlertList{
		XMLName:        xml.Name{Local: "feed"},
		Text:           "",
		SchemaLocation: "http://www.w3.org/2005/Atom https://jvndb.jvn.jp/schema/atom.xsd http://jvn.jp/rss/mod_sec/3.0/ https://jvndb.jvn.jp/schema/mod_sec_3.0.xsd http://data-marking.mitre.org/extensions/MarkingStructure#TLP-1 https://jvndb.jvn.jp/schema/tlp_marking.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd",
		Lang:           "ja",
		Title: ALTitle{
			Text: "title",
			Type: "text",
		},
		Updated: "2020-07-22T16:00:40+09:00",
		ID:      "id",
		Link: ALLink{
			Text:     "",
			Rel:      "alternate",
			Type:     "text/html",
			HRefLang: "ja",
			HRef:     "http://example.com/",
		},
		Author: ALAuthor{
			Text: "",
			Name: "name",
			URI:  "http://example.com/",
		},
		Handling: ALHandling{
			Text: "",
			Marking: ALMarking{
				MarkingStruct: ALMarkingStruct{
					Text:             "",
					Type:             "xsitype",
					MarkingModelName: "TLP",
					MarkingModelRef:  "http://example.com/",
					Color:            "WHITE",
				},
			},
		},
		Entries: []*ALEntry{
			&ALEntry{
				Text:      "",
				Title:     "title",
				ID:        "id",
				Published: "2020-01-15T12:04:14+09:00",
				Updated:   "2020-01-15T12:04:14+09:00",
				Category: ALCategory{
					Text:  "",
					Label: "label",
					Term:  "term",
				},
				Items: ALItems{
					Text: "",
					Items: []*ALItem{
						&ALItem{
							Text:       "",
							Title:      "Microsoft 製品の脆弱性対策について(2020年1月)",
							Identifier: "MYJVN-ALT-2020-0002-0001",
							Link: ALItemLink{
								Text: "",
								Href: "https://www.ipa.go.jp/security/ciadr/vul/20200115-ms.html",
							},
							Published: "2020-01-15T00:00:00+09:00",
							Updated:   "2020-01-15T12:04:14+09:00",
						},
						&ALItem{
							Text:       "",
							Title:      "2020 年 1 月のセキュリティ更新プログラム (月例)",
							Identifier: "MYJVN-ALT-2020-0002-0002",
							Link: ALItemLink{
								Text: "",
								Href: "https://msrc-blog.microsoft.com/2020/01/14/202001-security-updates/",
							},
							Published: "2020-01-15T00:00:00+09:00",
							Updated:   "2020-01-15T12:04:14+09:00",
						},
					},
				},
			},
		},
		Status: Status{
			Version:      "3.3",
			Method:       "getAlertList",
			RetCd:        0,
			RetMax:       "50",
			ErrCd:        "errcd",
			ErrMsg:       "errmsg",
			TotalRes:     "54",
			TotalResRet:  "1",
			FirstRes:     "1",
			Feed:         "hnd",
			MaxCountItem: 1,
		},
	}

	type testCase struct {
		description string
		httpResp    string
		respFormat  string
		structured  *AlertList
	}
	var testcases = []testCase{
		{
			description: "Not specifying response format",
			httpResp:    expectedHTTPResp,
			structured:  expectedAlertList,
		},
		{
			description: "Specifying XML as the response format",
			httpResp:    expectedHTTPResp,
			respFormat:  "xml",
			structured:  expectedAlertList,
		},
		{
			description: "Specifying JSON as the response format",
			httpResp:    expectedHTTPResp,
			respFormat:  "json",
			structured:  expectedAlertList,
		},
	}

	for _, c := range testcases {
		t.Run(c.description, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, c.httpResp)
			})

			alertList, err := client.GetAlertList(context.Background(), SetFormat(c.respFormat))
			if err != nil {
				t.Fatalf("GetAlertList returned error: %v", err)
			}

			want, got := c.structured, alertList
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

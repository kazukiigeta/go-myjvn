// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewParamsGetProductList(t *testing.T) {
	var startItem uint = 10
	var maxCountItem uint8 = 3
	var cpeName string = "cpe:/*"
	var keyword = "test & result"

	params := &Parameter{
		StartItem:    startItem,
		MaxCountItem: maxCountItem,
		CPEName:      cpeName,
		Keyword:      keyword,
	}

	got := NewParamsGetProductList(params)

	want := &ParamsGetProductList{
		Method:       "getProductList",
		Feed:         "hnd",
		StartItem:    startItem,
		MaxCountItem: maxCountItem,
		CPEName:      cpeName,
		Keyword:      url.QueryEscape(keyword),
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}

func TestGetProductList(t *testing.T) {
	var expectedHTTPResp = `<Result version="3.3" xsi:schemaLocation="http://jvndb.jvn.jp/myjvn/Results https://jvndb.jvn.jp/schema/results_3.3.xsd"><VendorInfo xml:lang="ja"><Vendor vname="#1 deals and maps app" cpe="cpe:/:pointinside" vid="10133"><Product pname="Point Inside Shopping & Travel" cpe="cpe:/a:pointinside:point_inside_shopping_%26_travel" pid="21248"/></Vendor><Vendor vname="#sysPass" cpe="cpe:/:syspass" vid="12776"><Product pname="#sysPass" cpe="cpe:/a:syspass:syspass" pid="29385"/></Vendor></VendorInfo><status:Status version="3.3" method="getProductList" lang="ja" retCd="0" retMax="10000" errCd="errcd" errMsg="errmsg" totalRes="43019" totalResRet="2" firstRes="1" feed="hnd" maxCountItem="2"/></Result>`

	var expectedProductList = &ProductList{
		XMLName:        xml.Name{Local: "Result"},
		Text:           "",
		Version:        "3.3",
		XSI:            "",
		XMLNS:          "",
		MJRes:          "",
		AttrStatus:     "",
		SchemaLocation: "http://jvndb.jvn.jp/myjvn/Results https://jvndb.jvn.jp/schema/results_3.3.xsd",
		VendorInfo: PLVendorInfo{
			Text: "",
			Lang: "ja",
			Vendors: []*PLVendor{
				&PLVendor{
					Text:  "",
					VName: "#1 deals and maps app",
					CPE:   "cpe:/:pointinside",
					VID:   "10133",
					Products: []*PLProduct{
						&PLProduct{
							Text:  "",
							PName: "Point Inside Shopping & Travel",
							CPE:   "cpe:/a:pointinside:point_inside_shopping_%26_travel",
							PID:   "21248",
						},
					},
				},
				&PLVendor{
					Text:  "",
					VName: "#sysPass",
					CPE:   "cpe:/:syspass",
					VID:   "12776",
					Products: []*PLProduct{
						&PLProduct{
							Text:  "",
							PName: "#sysPass",
							CPE:   "cpe:/a:syspass:syspass",
							PID:   "29385",
						},
					},
				},
			},
		},
		Status: Status{
			Version:      "3.3",
			Method:       "getProductList",
			Language:     "ja",
			RetCd:        0,
			RetMax:       "10000",
			ErrCd:        "errcd",
			ErrMsg:       "errmsg",
			TotalRes:     "43019",
			TotalResRet:  "2",
			FirstRes:     "1",
			Feed:         "hnd",
			MaxCountItem: 2,
		},
	}

	type testCase struct {
		description string
		httpResp    string
		respFormat  string
		structured  *ProductList
	}
	var testcases = []testCase{
		{
			description: "Not specifying optional parameters",
			httpResp:    expectedHTTPResp,
			structured:  expectedProductList,
		},
	}

	for _, c := range testcases {
		t.Run(c.description, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, c.httpResp)
			})

			params := &Parameter{Format: c.respFormat}
			p := NewParamsGetProductList(params)

			productList, err := client.GetProductList(context.Background(), p)
			if err != nil {
				t.Fatalf("GetProductList returned error: %v", err)
			}

			want, got := c.structured, productList
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

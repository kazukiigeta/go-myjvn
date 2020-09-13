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

func TestGetVendorList(t *testing.T) {
	var expectedHTTPResp = `<Result version="3.3" xsi:schemaLocation="http://jvndb.jvn.jp/myjvn/Results https://jvndb.jvn.jp/schema/results_3.3.xsd"><VendorInfo xml:lang="ja"><Vendor vname="#1 deals and maps app" cpe="cpe:/:pointinside" vid="10133"/><Vendor vname="#sysPass" cpe="cpe:/:syspass" vid="12776"/></VendorInfo><status:Status version="3.3" method="getVendorList" lang="ja" retCd="0" retMax="10000" errCd="errcd" errMsg="errmsg" totalRes="54" totalResRet="1" firstRes="1" feed="hnd"/></Result>`

	var expectedVendorList = &VendorList{
		XMLName:        xml.Name{Local: "Result"},
		Text:           "",
		Version:        "3.3",
		XSI:            "",
		XMLNS:          "",
		MJRes:          "",
		AttrStatus:     "",
		SchemaLocation: "http://jvndb.jvn.jp/myjvn/Results https://jvndb.jvn.jp/schema/results_3.3.xsd",
		VendorInfo: VLVendorInfo{
			Text: "",
			Lang: "ja",
			Vendors: []*VLVendor{
				&VLVendor{
					Text:  "",
					VName: "#1 deals and maps app",
					CPE:   "cpe:/:pointinside",
					VID:   "10133",
				},
				&VLVendor{
					Text:  "",
					VName: "#sysPass",
					CPE:   "cpe:/:syspass",
					VID:   "12776",
				},
			},
		},
		Status: Status{
			Version:     "3.3",
			Method:      "getVendorList",
			Language:    "ja",
			RetCd:       0,
			RetMax:      "10000",
			ErrCd:       "errcd",
			ErrMsg:      "errmsg",
			TotalRes:    "54",
			TotalResRet: "1",
			FirstRes:    "1",
			Feed:        "hnd",
		},
	}

	type testCase struct {
		description string
		httpResp    string
		structured  *VendorList
	}
	var testcases = []testCase{
		{
			description: "Not specifying optional parameters",
			httpResp:    expectedHTTPResp,
			structured:  expectedVendorList,
		},
	}

	for _, c := range testcases {
		t.Run(c.description, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, c.httpResp)
			})

			vendorList, err := client.GetVendorList(context.Background())
			if err != nil {
				t.Fatalf("GetVendorList returned error: %v", err)
			}

			want, got := c.structured, vendorList
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

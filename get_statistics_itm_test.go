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

func TestNewParamsGetStatisticsITM(t *testing.T) {
	var theme string = "sumCvss"
	var dataType string = "m"
	var cweID string = "CWE-20"
	var pid uint = 43096
	var cpeName string = "cpe:/*"
	var datePublicStartY uint16 = 2020
	var datePublicStartM uint8 = 12
	var datePublicEndY uint16 = 2020
	var datePublicEndM uint8 = 12

	params := &Parameter{
		Theme:            theme,
		DataType:         dataType,
		CWEID:            cweID,
		PID:              pid,
		CPEName:          cpeName,
		DatePublicStartY: datePublicStartY,
		DatePublicStartM: datePublicStartM,
		DatePublicEndY:   datePublicEndY,
		DatePublicEndM:   datePublicEndM,
	}

	got := NewParamsGetStatisticsITM(params)

	want := &ParamsGetStatisticsITM{
		Method:           "getStatistics",
		Feed:             "itm",
		Theme:            theme,
		DataType:         dataType,
		CWEID:            cweID,
		PID:              pid,
		CPEName:          cpeName,
		DatePublicStartY: datePublicStartY,
		DatePublicStartM: datePublicStartM,
		DatePublicEndY:   datePublicEndY,
		DatePublicEndM:   datePublicEndM,
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}

func TestGetStatisticsITM(t *testing.T) {
	var expectedHTTPResp = `
<Result version="3.2" xsi:schemaLocation="http://jvndb.jvn.jp/myjvn/Results https://jvndb.jvn.jp/schema/results_3.2.xsd">
	<mjstat:sumCvss>
		<mjstat:title xml:lang="ja">CVSSスコア</mjstat:title>
		<mjstat:title xml:lang="en-US">CVSS Score</mjstat:title>
		<mjstat:resData date="2015" cntAll="7820" cntH="2717" cntM="4323" cntL="775"/>
		<mjstat:resData date="2016" cntAll="8913" cntH="3017" cntM="4923" cntL="965"/>
	</mjstat:sumCvss>
	<status:Status version="3.2" method="getStatistics" lang="ja" retCd="0" retMax="10000" retMaxCnt="15579" errCd="errcd" errMsg="errmsg" totalRes="6" totalResRet="6" firstRes="1" feed="itm" theme="sumCvss" cweId="CWE-20" datePublicStartY="2015"/>
</Result>
`

	var expectedStatisticsITM = &StatisticsITM{
		XMLName: xml.Name{Local: "Result"},
		SumCVSS: SumCVSS{
			Title: []*Title{
				&Title{
					Text:     "CVSSスコア",
					Language: "ja",
				},
				&Title{
					Text:     "CVSS Score",
					Language: "en-US",
				},
			},
			ResData: []*ResData{
				&ResData{
					Date:   "2015",
					CntAll: "7820",
					CntH:   "2717",
					CntM:   "4323",
					CntL:   "775",
				},
				&ResData{
					Date:   "2016",
					CntAll: "8913",
					CntH:   "3017",
					CntM:   "4923",
					CntL:   "965",
				},
			},
		},
		Status: Status{
			Version:          "3.2",
			Method:           "getStatistics",
			Language:         "ja",
			RetCd:            0,
			RetMax:           "10000",
			RetMaxCnt:        "15579",
			ErrCd:            "errcd",
			ErrMsg:           "errmsg",
			TotalRes:         "6",
			TotalResRet:      "6",
			FirstRes:         "1",
			Feed:             "itm",
			Theme:            "sumCvss",
			CWEID:            "CWE-20",
			DatePublicStartY: 2015,
		},
	}

	type testCase struct {
		description string
		httpResp    string
		respFormat  string
		structured  *StatisticsITM
	}
	var testcases = []testCase{
		{
			description: "Not specifying optional parameters",
			httpResp:    expectedHTTPResp,
			structured:  expectedStatisticsITM,
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
			p := NewParamsGetStatisticsITM(params)

			productList, err := client.GetStatisticsITM(context.Background(), p)
			if err != nil {
				t.Fatalf("GetStatisticsITM returned error: %v", err)
			}

			want, got := c.structured, productList
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

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

func TestNewParamsGetStatisticsHND(t *testing.T) {
	var theme string = "sumAll"
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

	got := NewParamsGetStatisticsHND(params)

	want := &ParamsGetStatisticsHND{
		Method:           "getStatistics",
		Feed:             "hnd",
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

func TestGetStatisticsHND(t *testing.T) {
	var expectedHTTPResp = `
<Result version="3.3" xsi:schemaLocation="http://jvndb.jvn.jp/myjvn/Results https://jvndb.jvn.jp/schema/results_3.3.xsd">
	<mjstat:sumJvnDb>
		<mjstat:title xml:lang="ja">脆弱性統計情報</mjstat:title>
		<mjstat:title xml:lang="en-US">Statistics Vulnerability Count</mjstat:title>
		<mjstat:resDataTotal vulinfo="122180" vendor="17739" product="43096"/>
		<mjstat:resData date="2015" cntAll="1665"/>
		<mjstat:resData date="2016" cntAll="8002"/>
	</mjstat:sumJvnDb>
	<mjstat:sumCvss>
		<mjstat:title xml:lang="ja">CVSSスコア</mjstat:title>
		<mjstat:title xml:lang="en-US">CVSS Score</mjstat:title>
		<mjstat:resDataTotal vulinfo="122180" vendor="17739" product="43096"/>
		<mjstat:resData date="2015" cntAll="1665" cntC="286" cntH="653" cntM="702" cntL="24" cntN="0"/>
		<mjstat:resData date="2016" cntAll="8002" cntC="1269" cntH="3392" cntM="3111" cntL="230" cntN="0"/>
	</mjstat:sumCvss>
	<mjstat:sumCwe cweId="CWE-20">
		<mjstat:title xml:lang="ja">不適切な入力確認</mjstat:title>
		<mjstat:title xml:lang="en-US">Improper Input Validation</mjstat:title>
		<mjstat:resDataTotal vulinfo="122180" vendor="17739" product="43096"/>
		<mjstat:resData date="2015" cntAll="131"/>
		<mjstat:resData date="2016" cntAll="644"/>
	</mjstat:sumCwe>
	<status:Status version="3.3" method="getStatistics" lang="ja" retCd="0" retMax="10000" retMaxCnt="15558" errCd="errcd" errMsg="errmsg" totalRes="6" totalResRet="6" firstRes="1" feed="hnd" theme="sumAll" cweId="CWE-20" datePublicStartY="2015"/>
</Result>
`

	var expectedStatisticsHND = &StatisticsHND{
		XMLName: xml.Name{Local: "Result"},
		SumJVNDB: SumJVNDB{
			Title: []*Title{
				&Title{
					Text:     "脆弱性統計情報",
					Language: "ja",
				},
				&Title{
					Text:     "Statistics Vulnerability Count",
					Language: "en-US",
				},
			},
			ResDataTotal: ResDataTotal{
				VulInfo: "122180",
				Vendor:  "17739",
				Product: "43096",
			},
			ResData: []*ResData{
				&ResData{
					Date:   "2015",
					CntAll: "1665",
				},
				&ResData{
					Date:   "2016",
					CntAll: "8002",
				},
			},
		},
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
			ResDataTotal: ResDataTotal{
				VulInfo: "122180",
				Vendor:  "17739",
				Product: "43096",
			},
			ResData: []*ResData{
				&ResData{
					Date:   "2015",
					CntAll: "1665",
					CntC:   "286",
					CntH:   "653",
					CntM:   "702",
					CntL:   "24",
					CntN:   "0",
				},
				&ResData{
					Date:   "2016",
					CntAll: "8002",
					CntC:   "1269",
					CntH:   "3392",
					CntM:   "3111",
					CntL:   "230",
					CntN:   "0",
				},
			},
		},
		SumCWE: SumCWE{
			Title: []*Title{
				&Title{
					Text:     "不適切な入力確認",
					Language: "ja",
				},
				&Title{
					Text:     "Improper Input Validation",
					Language: "en-US",
				},
			},
			ResDataTotal: ResDataTotal{
				VulInfo: "122180",
				Vendor:  "17739",
				Product: "43096",
			},
			ResData: []*ResData{
				&ResData{
					Date:   "2015",
					CntAll: "131",
				},
				&ResData{
					Date:   "2016",
					CntAll: "644",
				},
			},
		},
		Status: Status{
			Version:          "3.3",
			Method:           "getStatistics",
			Language:         "ja",
			RetCd:            0,
			RetMax:           "10000",
			RetMaxCnt:        "15558",
			ErrCd:            "errcd",
			ErrMsg:           "errmsg",
			TotalRes:         "6",
			TotalResRet:      "6",
			FirstRes:         "1",
			Feed:             "hnd",
			Theme:            "sumAll",
			CWEID:            "CWE-20",
			DatePublicStartY: 2015,
		},
	}

	type testCase struct {
		description string
		httpResp    string
		respFormat  string
		structured  *StatisticsHND
	}
	var testcases = []testCase{
		{
			description: "Not specifying optional parameters",
			httpResp:    expectedHTTPResp,
			structured:  expectedStatisticsHND,
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
			p := NewParamsGetStatisticsHND(params)

			productList, err := client.GetStatisticsHND(context.Background(), p)
			if err != nil {
				t.Fatalf("GetStatisticsHND returned error: %v", err)
			}

			want, got := c.structured, productList
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

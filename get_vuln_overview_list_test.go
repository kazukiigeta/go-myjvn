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

func TestGetVulnOverviewList(t *testing.T) {
	var expectedHTTPResp = `<rdf:RDF xsi:schemaLocation="http://purl.org/rss/1.0/ https://jvndb.jvn.jp/schema/jvnrss_3.2.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd" xml:lang="en"><channel rdf:about="https://jvndb.jvn.jp/apis/myjvn"><title>JVNDB Vulnerability countermeasure information</title><link>https://jvndb.jvn.jp/apis/myjvn</link><description>JVNDB Vulnerability countermeasure information</description><dc:date>2020-07-31T08:17:20+09:00</dc:date><dcterms:issued/><dcterms:modified>2020-07-31T08:17:20+09:00</dcterms:modified><sec:handling><marking:Marking><marking:Marking_Structure xsi:type="tlpMarking:TLPMarkingStructureType" marking_model_name="TLP" marking_model_ref="http://www.us-cert.gov/tlp/" color="WHITE"/></marking:Marking></sec:handling><items><rdf:Seq><rdf:li rdf:resource="https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html"/></rdf:Seq></items></channel><item rdf:about="https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html"><title>TOYOTA MOTOR's Global TechStream vulnerable to buffer overflow</title><link>https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html</link><description>Global TechStream (GTS) is a diagnostic tool that Toyota Motor Corporation provides for Toyota dealers and independent workshops technicians to utilize. Global TechStream (GTS) contains a buffer overflow vulnerability (CWE-121). Tomoya Kitagawa of LAC Co., Ltd. reported this vulnerability to IPA. JPCERT/CC coordinated with the developer under Information Security Early Warning Partnership.</description><dc:creator>Information-technology Promotion Agency, Japan</dc:creator><sec:identifier>JVNDB-2020-000049</sec:identifier><sec:references source="CVE" id="CVE-2020-5610">https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-5610</sec:references><sec:references source="JVN" id="JVN#40400577">https://jvn.jp/en/jp/JVN40400577/index.html</sec:references><sec:references id="CWE-119" title="Buffer Errors(CWE-119)">http://www.ipa.go.jp/security/english/vuln/CWE_en.html#CWE119</sec:references><sec:cpe version="2.2" vendor="TOYOTA MOTOR CORPORATION" product="Global TechStream">cpe:/a:toyota:global_tech_stream</sec:cpe><sec:cvss score="4.1" severity="Medium" vector="CVSS:3.0/AV:P/AC:L/PR:N/UI:R/S:U/C:L/I:L/A:L" version="3.0" type="Base"/><sec:cvss score="4.4" severity="Medium" vector="AV:L/AC:M/Au:N/C:P/I:P/A:P" version="2.0" type="Base"/><dc:date>2020-07-29T14:48:07+09:00</dc:date><dcterms:issued>2020-07-29T14:48:07+09:00</dcterms:issued><dcterms:modified>2020-07-29T14:48:07+09:00</dcterms:modified></item><status:Status version="3.3" method="getVulnOverviewList" lang="en" retCd="0" retMax="50" errCd="errcd" errMsg="errmsg" totalRes="3" totalResRet="1" firstRes="1" feed="hnd" maxCountItem="1"/></rdf:RDF>`

	var expectedVulnOverviewList = &VulnOverviewList{
		XMLName:        xml.Name{Space: "rdf", Local: "RDF"},
		Text:           "",
		XSI:            "",
		XMLNS:          "",
		RSS:            "",
		RDF:            "",
		DC:             "",
		DCTerms:        "",
		Sec:            "",
		Marking:        "",
		TLPMarking:     "",
		AttrStatus:     "",
		SchemaLocation: "http://purl.org/rss/1.0/ https://jvndb.jvn.jp/schema/jvnrss_3.2.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd",
		Lang:           "en",
		Channel: VOLChannel{
			Text:        "",
			About:       "https://jvndb.jvn.jp/apis/myjvn",
			Title:       "JVNDB Vulnerability countermeasure information",
			Link:        "https://jvndb.jvn.jp/apis/myjvn",
			Description: "JVNDB Vulnerability countermeasure information",
			Date:        "2020-07-31T08:17:20+09:00",
			Issued:      "",
			Modified:    "2020-07-31T08:17:20+09:00",
			Handling: VOLHandling{
				Text: "",
				Marking: VOLMarking{
					Text: "",
					MarkingStruct: VOLMarkingStruct{
						Text:             "",
						Type:             "tlpMarking:TLPMarkingStructureType",
						MarkingModelName: "TLP",
						MarkingModelRef:  "http://www.us-cert.gov/tlp/",
						Color:            "WHITE",
					},
				},
			},
			Items: VOLChannelItems{
				Text: "",
				Seq: VOLSeq{
					Text: "",
					LI: []*VOLLI{
						&VOLLI{
							Resource: "https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html",
						},
					},
				},
			},
		},
		Items: []*VOLItem{
			&VOLItem{
				Text:        "",
				About:       "https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html",
				Title:       "TOYOTA MOTOR's Global TechStream vulnerable to buffer overflow",
				Link:        "https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html",
				Description: "Global TechStream (GTS) is a diagnostic tool that Toyota Motor Corporation provides for Toyota dealers and independent workshops technicians to utilize. Global TechStream (GTS) contains a buffer overflow vulnerability (CWE-121). Tomoya Kitagawa of LAC Co., Ltd. reported this vulnerability to IPA. JPCERT/CC coordinated with the developer under Information Security Early Warning Partnership.",
				Creator:     "Information-technology Promotion Agency, Japan",
				Identifier:  "JVNDB-2020-000049",
				References: []*VOLReferences{
					&VOLReferences{
						Text:   "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-5610",
						Source: "CVE",
						ID:     "CVE-2020-5610",
						Title:  "",
					},
					&VOLReferences{
						Text:   "https://jvn.jp/en/jp/JVN40400577/index.html",
						Source: "JVN",
						ID:     "JVN#40400577",
						Title:  "",
					},
					&VOLReferences{
						Text:   "http://www.ipa.go.jp/security/english/vuln/CWE_en.html#CWE119",
						Source: "",
						ID:     "CWE-119",
						Title:  "Buffer Errors(CWE-119)",
					},
				},
				CPEs: []*VOLCPE{
					&VOLCPE{
						Text:    "cpe:/a:toyota:global_tech_stream",
						Vendor:  "TOYOTA MOTOR CORPORATION",
						Product: "Global TechStream",
						Version: "2.2",
					},
				},
				CVSSes: []*VOLCVSS{
					&VOLCVSS{
						Text:     "",
						Score:    "4.1",
						Severity: "Medium",
						Vector:   "CVSS:3.0/AV:P/AC:L/PR:N/UI:R/S:U/C:L/I:L/A:L",
						Version:  "3.0",
						Type:     "Base",
					},
					&VOLCVSS{
						Text:     "",
						Score:    "4.4",
						Severity: "Medium",
						Vector:   "AV:L/AC:M/Au:N/C:P/I:P/A:P",
						Version:  "2.0",
						Type:     "Base",
					},
				},
				Date:     "2020-07-29T14:48:07+09:00",
				Issued:   "2020-07-29T14:48:07+09:00",
				Modified: "2020-07-29T14:48:07+09:00",
			},
		},
		Status: Status{
			Version:      "3.3",
			Method:       "getVulnOverviewList",
			Language:     "en",
			RetCd:        0,
			RetMax:       "50",
			ErrCd:        "errcd",
			ErrMsg:       "errmsg",
			TotalRes:     "3",
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
		structured  *VulnOverviewList
	}
	var testcases = []testCase{
		{
			description: "Not specifying optional parameters",
			httpResp:    expectedHTTPResp,
			structured:  expectedVulnOverviewList,
		},
	}

	for _, c := range testcases {
		t.Run(c.description, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, c.httpResp)
			})

			productList, err := client.GetVulnOverviewList(context.Background(), SetFormat(c.respFormat))
			if err != nil {
				t.Fatalf("GetVulnOverviewList returned error: %v", err)
			}

			want, got := c.structured, productList
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

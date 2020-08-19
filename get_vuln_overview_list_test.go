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

func TestNewParamsGetVulnOverviewList(t *testing.T) {
	var startItem uint = 10
	var maxCountItem uint8 = 3
	var cpeName string = "cpe:/*"
	var vendorID string = "1"
	var productID string = "1"
	var keyword string = "test & result"
	var severity string = "m"
	var vector string = "CVSS:3.0/AV:[N,A,L,P]/AC:[L,H]"
	var rangeDatePublic string = "m"
	var rangeDatePublished string = "m"
	var rangeDateFirstPublished string = "m"
	var datePublicStartY uint16 = 2020
	var datePublicStartM uint8 = 12
	var datePublicStartD uint8 = 1
	var datePublicEndY uint16 = 2020
	var datePublicEndM uint8 = 12
	var datePublicEndD uint8 = 1
	var dateFirstPublishedStartY uint16 = 2020
	var dateFirstPublishedStartM uint8 = 12
	var dateFirstPublishedStartD uint8 = 1
	var dateFirstPublishedEndY uint16 = 2020
	var dateFirstPublishedEndM uint8 = 12
	var dateFirstPublishedEndD uint8 = 1
	var language string = "ja"

	params := &Parameter{
		StartItem:                startItem,
		MaxCountItem:             maxCountItem,
		CPEName:                  cpeName,
		VendorID:                 vendorID,
		ProductID:                productID,
		Keyword:                  keyword,
		Severity:                 severity,
		Vector:                   vector,
		RangeDatePublic:          rangeDatePublic,
		RangeDatePublished:       rangeDatePublished,
		RangeDateFirstPublished:  rangeDateFirstPublished,
		DatePublicStartY:         datePublicStartY,
		DatePublicStartM:         datePublicStartM,
		DatePublicStartD:         datePublicStartD,
		DatePublicEndY:           datePublicEndY,
		DatePublicEndM:           datePublicEndM,
		DatePublicEndD:           datePublicEndD,
		DateFirstPublishedStartY: dateFirstPublishedStartY,
		DateFirstPublishedStartM: dateFirstPublishedStartM,
		DateFirstPublishedStartD: dateFirstPublishedStartD,
		DateFirstPublishedEndY:   dateFirstPublishedEndY,
		DateFirstPublishedEndM:   dateFirstPublishedEndM,
		DateFirstPublishedEndD:   dateFirstPublishedEndD,
		Language:                 language,
	}

	got := NewParamsGetVulnOverviewList(params)

	want := &ParamsGetVulnOverviewList{
		Method:                   "getVulnOverviewList",
		Feed:                     "hnd",
		StartItem:                startItem,
		MaxCountItem:             maxCountItem,
		CPEName:                  cpeName,
		Keyword:                  url.QueryEscape(keyword),
		VendorID:                 vendorID,
		ProductID:                productID,
		Severity:                 severity,
		Vector:                   vector,
		RangeDatePublic:          rangeDatePublic,
		RangeDatePublished:       rangeDatePublished,
		RangeDateFirstPublished:  rangeDateFirstPublished,
		DatePublicStartY:         datePublicStartY,
		DatePublicStartM:         datePublicStartM,
		DatePublicStartD:         datePublicStartD,
		DatePublicEndY:           datePublicEndY,
		DatePublicEndM:           datePublicEndM,
		DatePublicEndD:           datePublicEndD,
		DateFirstPublishedStartY: dateFirstPublishedStartY,
		DateFirstPublishedStartM: dateFirstPublishedStartM,
		DateFirstPublishedStartD: dateFirstPublishedStartD,
		DateFirstPublishedEndY:   dateFirstPublishedEndY,
		DateFirstPublishedEndM:   dateFirstPublishedEndM,
		DateFirstPublishedEndD:   dateFirstPublishedEndD,
		Language:                 language,
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}

func TestGetVulnOverviewList(t *testing.T) {
	var expectedHTTPResp = `<rdf:RDF xsi:schemaLocation="http://purl.org/rss/1.0/ https://jvndb.jvn.jp/schema/jvnrss_3.2.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd" xml:lang="en">
	<channel rdf:about="https://jvndb.jvn.jp/apis/myjvn">
		<title>JVNDB Vulnerability countermeasure information</title>
		<link>https://jvndb.jvn.jp/apis/myjvn</link>
		<description>JVNDB Vulnerability countermeasure information</description>
		<dc:date>2020-07-31T08:17:20+09:00</dc:date>
		<dcterms:issued/>
		<dcterms:modified>2020-07-31T08:17:20+09:00</dcterms:modified>
		<sec:handling>
			<marking:Marking>
				<marking:Marking_Structure xsi:type="tlpMarking:TLPMarkingStructureType" marking_model_name="TLP" marking_model_ref="http://www.us-cert.gov/tlp/" color="WHITE"/>
			</marking:Marking>
		</sec:handling>
		<items>
			<rdf:Seq>
				<rdf:li rdf:resource="https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html"/>
			</rdf:Seq>
		</items>
	</channel>
	<item rdf:about="https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html">
		<title>TOYOTA MOTOR's Global TechStream vulnerable to buffer overflow</title>
		<link>https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html</link>
		<description>Global TechStream (GTS) is a diagnostic tool that Toyota Motor Corporation provides for Toyota dealers and independent workshops technicians to utilize. Global TechStream (GTS) contains a buffer overflow vulnerability (CWE-121). Tomoya Kitagawa of LAC Co., Ltd. reported this vulnerability to IPA. JPCERT/CC coordinated with the developer under Information Security Early Warning Partnership.</description>
		<dc:creator>Information-technology Promotion Agency, Japan</dc:creator>
		<sec:identifier>JVNDB-2020-000049</sec:identifier>
		<sec:references source="CVE" id="CVE-2020-5610">https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-5610</sec:references>
		<sec:references source="JVN" id="JVN#40400577">https://jvn.jp/en/jp/JVN40400577/index.html</sec:references>
		<sec:references id="CWE-119" title="Buffer Errors(CWE-119)">http://www.ipa.go.jp/security/english/vuln/CWE_en.html#CWE119</sec:references>
		<sec:cpe version="2.2" vendor="TOYOTA MOTOR CORPORATION" product="Global TechStream">cpe:/a:toyota:global_tech_stream</sec:cpe>
		<sec:cvss score="4.1" severity="Medium" vector="CVSS:3.0/AV:P/AC:L/PR:N/UI:R/S:U/C:L/I:L/A:L" version="3.0" type="Base"/>
		<sec:cvss score="4.4" severity="Medium" vector="AV:L/AC:M/Au:N/C:P/I:P/A:P" version="2.0" type="Base"/>
		<dc:date>2020-07-29T14:48:07+09:00</dc:date>
		<dcterms:issued>2020-07-29T14:48:07+09:00</dcterms:issued>
		<dcterms:modified>2020-07-29T14:48:07+09:00</dcterms:modified>
	</item>
	<status:Status version="3.3" method="getVulnOverviewList" lang="en" retCd="0" retMax="50" errCd="errcd" errMsg="errmsg" totalRes="3" totalResRet="1" firstRes="1" feed="hnd" maxCountItem="1"/>
</rdf:RDF>
`

	var expectedVulnOverviewList = &VulnOverviewList{
		XMLName: xml.Name{Space: "rdf", Local: "RDF"},
		Channel: Channel{
			Title:       "JVNDB Vulnerability countermeasure information",
			Link:        "https://jvndb.jvn.jp/apis/myjvn",
			Description: "JVNDB Vulnerability countermeasure information",
			Date:        "2020-07-31T08:17:20+09:00",
			Issued:      "",
			Modified:    "2020-07-31T08:17:20+09:00",
			SecHandling: SecHandling{
				Marking: Marking{
					MarkingStructure: MarkingStructure{
						XSIType:          "tlpMarking:TLPMarkingStructureType",
						MarkingModelName: "TLP",
						MarkingModelRef:  "http://www.us-cert.gov/tlp/",
						Color:            "WHITE",
					},
				},
			},
			Items: Items{
				Seq: Seq{
					LI: LI{
						Resource: "https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html",
					},
				},
			},
		},
		Item: []*Item{
			&Item{
				Title:       "TOYOTA MOTOR's Global TechStream vulnerable to buffer overflow",
				Link:        "https://jvndb.jvn.jp/en/contents/2020/JVNDB-2020-000049.html",
				Description: "Global TechStream (GTS) is a diagnostic tool that Toyota Motor Corporation provides for Toyota dealers and independent workshops technicians to utilize. Global TechStream (GTS) contains a buffer overflow vulnerability (CWE-121). Tomoya Kitagawa of LAC Co., Ltd. reported this vulnerability to IPA. JPCERT/CC coordinated with the developer under Information Security Early Warning Partnership.",
				Creator:     "Information-technology Promotion Agency, Japan",
				Identifier:  "JVNDB-2020-000049",
				References: []*References{
					&References{
						Text:   "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-5610",
						Source: "CVE",
						ID:     "CVE-2020-5610",
					},
					&References{
						Text:   "https://jvn.jp/en/jp/JVN40400577/index.html",
						Source: "JVN",
						ID:     "JVN#40400577",
					},
					&References{
						Text:  "http://www.ipa.go.jp/security/english/vuln/CWE_en.html#CWE119",
						ID:    "CWE-119",
						Title: "Buffer Errors(CWE-119)",
					},
				},
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

			params := &Parameter{Format: c.respFormat}
			p := NewParamsGetVulnOverviewList(params)

			productList, err := client.GetVulnOverviewList(context.Background(), p)
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

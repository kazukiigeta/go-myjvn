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

func TestNewParamsGetVulnDetailInfo(t *testing.T) {
	var startItem uint = 10
	var maxCountItem uint8 = 3
	var vulnID string = "JVNDB-2020-000001"
	var language string = "en"

	params := &Parameter{
		StartItem:    startItem,
		MaxCountItem: maxCountItem,
		VulnID:       vulnID,
		Language:     language,
	}

	got := NewParamsGetVulnDetailInfo(params)

	want := &ParamsGetVulnDetailInfo{
		Method:       "getVulnDetailInfo",
		Feed:         "hnd",
		StartItem:    startItem,
		MaxCountItem: maxCountItem,
		VulnID:       vulnID,
		Language:     language,
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}

func TestGetVulnDetailInfo(t *testing.T) {
	var expectedHTTPResp = `<VULDEF-Document version="3.2" xsi:schemaLocation="http://jvn.jp/vuldef/ https://jvndb.jvn.jp/schema/vuldef_3.2.xsd http://jvn.jp/rss/mod_sec/3.0/ https://jvndb.jvn.jp/schema/mod_sec_3.0.xsd http://data-marking.mitre.org/extensions/MarkingStructure#TLP-1 https://jvndb.jvn.jp/schema/tlp_marking.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd" xml:lang="ja"><Vulinfo><VulinfoID>JVNDB-2020-006469</VulinfoID><VulinfoData><Title>三菱電機製 GOT2000 シリーズの TCP/IP 機能における複数の脆弱性</Title><VulinfoDescription><Overview>三菱電機株式会社が提供する GOT2000 シリーズの GT27、GT25、GT23 のファームウェアに組込まれている TCP/IP 機能には、次の複数の脆弱性が存在します。 * バッファエラー (CWE-119) - CVE-2020-5595 * セッションの固定化 (CWE-384) - CVE-2020-5596 * NULL ポインタデリファレンス (CWE-476) - CVE-2020-5597 * 不適切なアクセス制御 (CWE-284) - CVE-2020-5598 * 引数の挿入または変更 (CWE-88) - CVE-2020-5599 * リソース管理の問題 (CWE-399) - CVE-2020-5600</Overview></VulinfoDescription><Affected><AffectedItem><Name>三菱電機</Name><ProductName>GT23 モデル</ProductName><Cpe version="2.2">cpe:/o:mitsubishielectric:gt23_model</Cpe><VersionNumber/></AffectedItem><AffectedItem><Name>三菱電機</Name><ProductName>GT25 モデル</ProductName><Cpe version="2.2">cpe:/o:mitsubishielectric:gt25_model</Cpe><VersionNumber/></AffectedItem><AffectedItem><Name>三菱電機</Name><ProductName>GT27 モデル</ProductName><Cpe version="2.2">cpe:/o:mitsubishielectric:gt27_model</Cpe><VersionNumber/></AffectedItem></Affected><Impact><Cvss version="3.0"><Severity type="Base">Critical</Severity><Base>9.8</Base><Vector>CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H</Vector></Cvss><ImpactItem><Description>第三者によって細工されたパケットを受信することで、製品のネットワーク機能が停止したり悪意あるプログラムが実行されたりする可能性があります。</Description></ImpactItem></Impact><Solution><SolutionItem><Description>[アップデートする] 開発者が提供する次の手順に従って、CoreOS を最新版にアップデートしてください。 1.MELSOFT GT Designer3(2000) 1.240A 以降のバージョンを、三菱電機FAサイトのソフトウェアダウンロードコーナーよりダウンロードしインストールする 2.MELSOFT GT Designer3(2000) を起動し、バージョン Z 以降の CoreOS を SD カードに作成する 3.SD カードを対象製品に挿入し、CoreOS を最新バージョンへアップデートする [ワークアラウンドを実施する] 信頼できないネットワークやホストからのアクセスを制限することで、本脆弱性の影響を軽減できます。 詳しくは、開発者が提供する情報をご確認ください。</Description></SolutionItem></Solution><Related><RelatedItem type="vendor"><Name>三菱電機株式会社</Name><VulinfoID>GOT2000シリーズにおけるTCP/IPスタックの複数の脆弱性</VulinfoID><URL>https://www.mitsubishielectric.co.jp/psirt/vulnerability/pdf/2020-005.pdf</URL></RelatedItem><RelatedItem type="advisory"><Name>Common Vulnerabilities and Exposures (CVE)</Name><VulinfoID>CVE-2020-5595</VulinfoID><URL>https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-5595</URL></RelatedItem></Related><History><HistoryItem><HistoryNo>1</HistoryNo><DateTime>2020-07-09T14:43:53+09:00</DateTime><Description>[2020年07月09日]\n 掲載</Description></HistoryItem></History><DateFirstPublished>2020-07-09T14:43:53+09:00</DateFirstPublished><DateLastUpdated>2020-07-09T14:43:53+09:00</DateLastUpdated><DatePublic>2020-07-03T00:00:00+09:00</DatePublic></VulinfoData></Vulinfo><sec:handling><marking:Marking><marking:Marking_Structure xsi:type="tlpMarking:TLPMarkingStructureType" marking_model_name="TLP" marking_model_ref="http://www.us-cert.gov/tlp/" color="WHITE"/></marking:Marking></sec:handling><status:Status version="3.3" method="getVulnDetailInfo" lang="ja" retCd="0" retMax="10" errCd="errcd" errMsg="errmsg" totalRes="1" totalResRet="1" firstRes="1" feed="hnd" vulnId="JVNDB-2020-006469"/></VULDEF-Document>`

	var expectedVulnDetailInfo = &VulnDetailInfo{
		XMLName:        xml.Name{Local: "VULDEF-Document"},
		Text:           "",
		Version:        "3.2",
		XSI:            "",
		XMLNS:          "",
		VulDef:         "",
		AttrStatus:     "",
		Sec:            "",
		Marking:        "",
		TLPMarking:     "",
		SchemaLocation: "http://jvn.jp/vuldef/ https://jvndb.jvn.jp/schema/vuldef_3.2.xsd http://jvn.jp/rss/mod_sec/3.0/ https://jvndb.jvn.jp/schema/mod_sec_3.0.xsd http://data-marking.mitre.org/extensions/MarkingStructure#TLP-1 https://jvndb.jvn.jp/schema/tlp_marking.xsd http://jvndb.jvn.jp/myjvn/Status https://jvndb.jvn.jp/schema/status_3.3.xsd",
		Lang:           "ja",
		VulInfo: VDIVulInfo{
			Text:      "",
			VulInfoID: "JVNDB-2020-006469",
			VulInfoData: VDIVulInfoData{
				Text:  "",
				Title: "三菱電機製 GOT2000 シリーズの TCP/IP 機能における複数の脆弱性",
				VulInfoDesc: VDIVulInfoDesc{
					Text:     "",
					Overview: "三菱電機株式会社が提供する GOT2000 シリーズの GT27、GT25、GT23 のファームウェアに組込まれている TCP/IP 機能には、次の複数の脆弱性が存在します。 * バッファエラー (CWE-119) - CVE-2020-5595 * セッションの固定化 (CWE-384) - CVE-2020-5596 * NULL ポインタデリファレンス (CWE-476) - CVE-2020-5597 * 不適切なアクセス制御 (CWE-284) - CVE-2020-5598 * 引数の挿入または変更 (CWE-88) - CVE-2020-5599 * リソース管理の問題 (CWE-399) - CVE-2020-5600",
				},
				Affected: VDIAffected{
					Text: "",
					AffectedItem: []*VDIAffectedItem{
						&VDIAffectedItem{
							Text:        "",
							Name:        "三菱電機",
							ProductName: "GT23 モデル",
							CPE: VDICPE{
								Text:    "cpe:/o:mitsubishielectric:gt23_model",
								Version: "2.2",
							},
							VersionNum: "",
						},
						&VDIAffectedItem{
							Text:        "",
							Name:        "三菱電機",
							ProductName: "GT25 モデル",
							CPE: VDICPE{
								Text:    "cpe:/o:mitsubishielectric:gt25_model",
								Version: "2.2",
							},
							VersionNum: "",
						},
						&VDIAffectedItem{
							Text:        "",
							Name:        "三菱電機",
							ProductName: "GT27 モデル",
							CPE: VDICPE{
								Text:    "cpe:/o:mitsubishielectric:gt27_model",
								Version: "2.2",
							},
							VersionNum: "",
						},
					},
				},
				Impact: VDIImpact{
					Text: "",
					CVSS: VDICVSS{
						Text:    "",
						Version: "3.0",
						Severity: VDISeverity{
							Text: "Critical",
							Type: "Base",
						},
						Base:   "9.8",
						Vector: "CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H",
					},
					ImpactItem: VDIImpactItem{
						Text:        "",
						Description: "第三者によって細工されたパケットを受信することで、製品のネットワーク機能が停止したり悪意あるプログラムが実行されたりする可能性があります。",
					},
				},
				Solution: VDISolution{
					SolutionItem: VDISolutionItem{
						Text:        "",
						Description: "[アップデートする] 開発者が提供する次の手順に従って、CoreOS を最新版にアップデートしてください。 1.MELSOFT GT Designer3(2000) 1.240A 以降のバージョンを、三菱電機FAサイトのソフトウェアダウンロードコーナーよりダウンロードしインストールする 2.MELSOFT GT Designer3(2000) を起動し、バージョン Z 以降の CoreOS を SD カードに作成する 3.SD カードを対象製品に挿入し、CoreOS を最新バージョンへアップデートする [ワークアラウンドを実施する] 信頼できないネットワークやホストからのアクセスを制限することで、本脆弱性の影響を軽減できます。 詳しくは、開発者が提供する情報をご確認ください。",
					},
				},
				Related: VDIRelated{
					Text: "",
					RelatedItems: []*VDIRelatedItem{
						&VDIRelatedItem{
							Text:      "",
							Type:      "vendor",
							Name:      "三菱電機株式会社",
							VulInfoID: "GOT2000シリーズにおけるTCP/IPスタックの複数の脆弱性",
							URL:       "https://www.mitsubishielectric.co.jp/psirt/vulnerability/pdf/2020-005.pdf",
							Title:     "",
						},
						&VDIRelatedItem{
							Text:      "",
							Type:      "advisory",
							Name:      "Common Vulnerabilities and Exposures (CVE)",
							VulInfoID: "CVE-2020-5595",
							URL:       "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-5595",
							Title:     "",
						},
					},
				},
				History: VDIHistory{
					Text: "",
					HistoryItem: VDIHistoryItem{
						Text:        "",
						HistoryNo:   "1",
						DateTime:    "2020-07-09T14:43:53+09:00",
						Description: `[2020年07月09日]\n 掲載`,
					},
				},
				DateFirstPublished: "2020-07-09T14:43:53+09:00",
				DateLastUpdated:    "2020-07-09T14:43:53+09:00",
				DatePublic:         "2020-07-03T00:00:00+09:00",
			},
		},
		Handling: VDIHandling{
			Marking: VDIMarking{
				MarkingStruct: VDIMarkingStruct{
					Text:             "",
					Type:             "tlpMarking:TLPMarkingStructureType",
					MarkingModelName: "TLP",
					MarkingModelRef:  "http://www.us-cert.gov/tlp/",
					Color:            "WHITE",
				},
			},
		},
		Status: Status{
			Version:     "3.3",
			Method:      "getVulnDetailInfo",
			Language:    "ja",
			RetCd:       0,
			RetMax:      "10",
			ErrCd:       "errcd",
			ErrMsg:      "errmsg",
			TotalRes:    "1",
			TotalResRet: "1",
			FirstRes:    "1",
			Feed:        "hnd",
			VulnID:      "JVNDB-2020-006469",
		},
	}

	type testCase struct {
		description string
		httpResp    string
		respFormat  string
		structured  *VulnDetailInfo
	}
	var testcases = []testCase{
		{
			description: "Not specifying optional parameters",
			httpResp:    expectedHTTPResp,
			structured:  expectedVulnDetailInfo,
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
			p := NewParamsGetVulnDetailInfo(params)

			vendorList, err := client.GetVulnDetailInfo(context.Background(), p)
			if err != nil {
				t.Fatalf("GetVulnDetailInfo returned error: %v", err)
			}

			want, got := c.structured, vendorList
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

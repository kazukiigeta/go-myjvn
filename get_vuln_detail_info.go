// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
)

// VDIVulInfoDesc stores the data from API response.
type VDIVulInfoDesc struct {
	Text     string `xml:",chardata"`
	Overview string `xml:"Overview"`
}

// VDICPE stores the data from API response.
type VDICPE struct {
	Text    string `xml:",chardata"`
	Version string `xml:"version,attr"`
}

// VDIAffectedItem stores the data from API response.
type VDIAffectedItem struct {
	Text        string `xml:",chardata"`
	Name        string `xml:"Name"`
	ProductName string `xml:"ProductName"`
	CPE         VDICPE `xml:"Cpe"`
	VersionNum  string `xml:"VersionNumber"`
}

// VDIAffected stores the data from API response.
type VDIAffected struct {
	Text         string             `xml:",chardata"`
	AffectedItem []*VDIAffectedItem `xml:"AffectedItem"`
}

// VDISeverity stores the data from API response.
type VDISeverity struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

// VDICVSS stores the data from API response.
type VDICVSS struct {
	Text     string      `xml:",chardata"`
	Version  string      `xml:"version,attr"`
	Severity VDISeverity `xml:"Severity"`
	Base     string      `xml:"Base"`
	Vector   string      `xml:"Vector"`
}

// VDIImpactItem stores the data from API response.
type VDIImpactItem struct {
	Text        string `xml:",chardata"`
	Description string `xml:"Description"`
}

// VDIImpact stores the data from API response.
type VDIImpact struct {
	Text       string        `xml:",chardata"`
	CVSS       VDICVSS       `xml:"Cvss"`
	ImpactItem VDIImpactItem `xml:"ImpactItem"`
}

// VDISolutionItem stores the data from API response.
type VDISolutionItem struct {
	Text        string `xml:",chardata"`
	Description string `xml:"Description"`
}

// VDISolution stores the data from API response.
type VDISolution struct {
	Text         string          `xml:",chardata"`
	SolutionItem VDISolutionItem `xml:"SolutionItem"`
}

// VDIRelatedItem stores the data from API response.
type VDIRelatedItem struct {
	Text      string `xml:",chardata"`
	Type      string `xml:"type,attr"`
	Name      string `xml:"Name"`
	VulInfoID string `xml:"VulinfoID"`
	URL       string `xml:"URL"`
	Title     string `xml:"Title"`
}

// VDIRelated stores the data from API response.
type VDIRelated struct {
	Text         string            `xml:",chardata"`
	RelatedItems []*VDIRelatedItem `xml:"RelatedItem"`
}

// VDIHistoryItem stores the data from API response.
type VDIHistoryItem struct {
	Text        string `xml:",chardata"`
	HistoryNo   string `xml:"HistoryNo"`
	DateTime    string `xml:"DateTime"`
	Description string `xml:"Description"`
}

// VDIHistory stores the data from API response.
type VDIHistory struct {
	Text        string         `xml:",chardata"`
	HistoryItem VDIHistoryItem `xml:"HistoryItem"`
}

// VDIVulInfoData stores the data from API response.
type VDIVulInfoData struct {
	Text               string         `xml:",chardata"`
	Title              string         `xml:"Title"`
	VulInfoDesc        VDIVulInfoDesc `xml:"VulinfoDescription"`
	Affected           VDIAffected    `xml:"Affected"`
	Impact             VDIImpact      `xml:"Impact"`
	Solution           VDISolution    `xml:"Solution"`
	Related            VDIRelated     `xml:"Related"`
	History            VDIHistory     `xml:"History"`
	DateFirstPublished string         `xml:"DateFirstPublished"`
	DateLastUpdated    string         `xml:"DateLastUpdated"`
	DatePublic         string         `xml:"DatePublic"`
}

// VDIVulInfo stores the data from API response.
type VDIVulInfo struct {
	Text        string         `xml:",chardata"`
	VulInfoID   string         `xml:"VulinfoID"`
	VulInfoData VDIVulInfoData `xml:"VulinfoData"`
}

// VDIMarkingStruct stores the data from API response.
type VDIMarkingStruct struct {
	Text             string `xml:",chardata"`
	Type             string `xml:"type,attr"`
	MarkingModelName string `xml:"marking_model_name,attr"`
	MarkingModelRef  string `xml:"marking_model_ref,attr"`
	Color            string `xml:"color,attr"`
}

// VDIMarking stores the data from API response.
type VDIMarking struct {
	Text          string           `xml:",chardata"`
	MarkingStruct VDIMarkingStruct `xml:"Marking_Structure"`
}

// VDIHandling stores the data from API response.
type VDIHandling struct {
	Text    string     `xml:",chardata"`
	Marking VDIMarking `xml:"Marking"`
}

// VulnDetailInfo stores the data from API response.
type VulnDetailInfo struct {
	XMLName        xml.Name    `xml:"VULDEF-Document"`
	Text           string      `xml:",chardata"`
	Version        string      `xml:"version,attr"`
	XSI            string      `xml:"xsi,attr"`
	XMLNS          string      `xml:"xmlns,attr"`
	VulDef         string      `xml:"vuldef,attr"`
	AttrStatus     string      `xml:"status,attr"`
	Sec            string      `xml:"sec,attr"`
	Marking        string      `xml:"marking,attr"`
	TLPMarking     string      `xml:"tlpMarking,attr"`
	SchemaLocation string      `xml:"schemaLocation,attr"`
	Lang           string      `xml:"lang,attr"`
	VulInfo        VDIVulInfo  `xml:"Vulinfo"`
	Handling       VDIHandling `xml:"handling"`
	Status         Status      `xml:"Status"`
}

// GetVulnDetailInfo downloads a vendor list.
// See: https://jvndb.jvn.jp/apis/getVulnDetailInfo_api_hnd.html
func (c *Client) GetVulnDetailInfo(ctx context.Context, opts ...Option) (*VulnDetailInfo, error) {
	p := &parameter{
		Method: "getVulnDetailInfo",
		Feed:   "hnd",
	}

	for _, opt := range opts {
		opt(p)
	}

	u, err := addOptions(defaultAPIPath, p)
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("GET", u)
	if err != nil {
		return nil, err
	}

	vendorList := new(VulnDetailInfo)
	err = c.do(ctx, req, nil, vendorList)
	if err != nil {
		return nil, err
	}

	return vendorList, nil
}

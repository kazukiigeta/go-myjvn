// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
)

// ParamsGetVulnDetailInfo specifies the parameters of a HTTP request for GetVulnDetailInfo.
type ParamsGetVulnDetailInfo struct {
	Method       string `url:"method"`
	Feed         string `url:"feed"`
	StartItem    uint   `url:"startItem,omitempty"`
	MaxCountItem uint8  `url:"maxCountItem,omitempty"`
	VulnID       string `url:"vulnId"`
	Language     string `url:"lang,omitempty"`
}

// VulInfoDescription stores the data from API response.
type VulInfoDescription struct {
	Overview string `xml:"Overview"`
}

// CPE stores the data from API response.
type CPE struct {
	Text     string `xml:",chardata"`
	Score    string `xml:"score,attr"`
	Severity string `xml:"severity,attr"`
	Vector   string `xml:"vector,attr"`
	Version  string `xml:"version,attr"`
	Type     string `xml:"type,attr"`
}

// AffectedItem stores the data from API response.
type AffectedItem struct {
	Name        string `xml:"Name"`
	ProductName string `xml:"ProductName"`
	CPE         CPE    `xml:"Cpe"`
}

// Affected stores the data from API response.
type Affected struct {
	AffectedItem []*AffectedItem `xml:"AffectedItem"`
}

// Severity stores the data from API response.
type Severity struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

// CVSS stores the data from API response.
type CVSS struct {
	Version  string   `xml:"version,attr"`
	Score    string   `xml:"score,attr"`
	Severity Severity `xml:"Severity"`
	Base     string   `xml:"Base"`
	Vector   string   `xml:"Vector"`
}

// ImpactItem stores the data from API response.
type ImpactItem struct {
	Description string `xml:"Description"`
}

// Impact stores the data from API response.
type Impact struct {
	CVSS       CVSS       `xml:"Cvss"`
	ImpactItem ImpactItem `xml:"ImpactItem"`
}

// SolutionItem stores the data from API response.
type SolutionItem struct {
	Description string `xml:"Description"`
}

// Solution stores the data from API response.
type Solution struct {
	SolutionItem SolutionItem `xml:"SolutionItem"`
}

// RelatedItem stores the data from API response.
type RelatedItem struct {
	Type      string `xml:"type,attr"`
	Name      string `xml:"Name"`
	VulInfoID string `xml:"VulinfoID"`
	URL       string `xml:"URL"`
	Title     string `xml:"Title"`
}

// Related stores the data from API response.
type Related struct {
	RelatedItem []*RelatedItem `xml:"RelatedItem"`
}

// HistoryItem stores the data from API response.
type HistoryItem struct {
	HistoryNo   string `xml:"HistoryNo"`
	DateTime    string `xml:"DateTime"`
	Description string `xml:"Description"`
}

// History stores the data from API response.
type History struct {
	HistoryItem HistoryItem `xml:"HistoryItem"`
}

// VulInfoData stores the data from API response.
type VulInfoData struct {
	Title              string             `xml:"Title"`
	VulInfoDescription VulInfoDescription `xml:"VulinfoDescription"`
	Affected           Affected           `xml:"Affected"`
	Impact             Impact             `xml:"Impact"`
	Solution           Solution           `xml:"Solution"`
	Related            Related            `xml:"Related"`
	History            History            `xml:"History"`
	DateFirstPublished string             `xml:"DateFirstPublished"`
	DateLastUpdated    string             `xml:"DateLastUpdated"`
	DatePublic         string             `xml:"DatePublic"`
}

// VulInfo stores the data from API response.
type VulInfo struct {
	VulInfoID   string      `xml:"VulinfoID"`
	VulInfoData VulInfoData `xml:"VulinfoData"`
}

// VulnDetailInfo stores the data from API response.
type VulnDetailInfo struct {
	XMLName     xml.Name    `xml:"VULDEF-Document"`
	VulInfo     VulInfo     `xml:"Vulinfo"`
	SecHandling SecHandling `xml:"handling"`
	Status      Status      `xml:"Status"`
}

// NewParamsGetVulnDetailInfo creates an instance of ParamsGetVulnDetailInfo.
func NewParamsGetVulnDetailInfo(params *Parameter) *ParamsGetVulnDetailInfo {
	if params == nil {
		params = &Parameter{}
	}

	p := &ParamsGetVulnDetailInfo{
		Method: "getVulnDetailInfo",
		Feed:   "hnd",
	}

	p.StartItem = params.StartItem
	p.MaxCountItem = params.MaxCountItem
	p.VulnID = params.VulnID
	p.Language = params.Language

	return p
}

// GetVulnDetailInfo downloads a vendor list.
// See: https://jvndb.jvn.jp/apis/getVulnDetailInfo_api_hnd.html
func (c *Client) GetVulnDetailInfo(ctx context.Context, params *ParamsGetVulnDetailInfo) (*VulnDetailInfo, error) {
	if params == nil {
		params = NewParamsGetVulnDetailInfo(nil)
	}

	u, err := addOptions(defaultAPIPath, params)
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

// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
	"net/url"
)

// VOLMarkingStruct stores the data from API response.
type VOLMarkingStruct struct {
	Text             string `xml:",chardata"`
	Type             string `xml:"type,attr"`
	MarkingModelName string `xml:"marking_model_name,attr"`
	MarkingModelRef  string `xml:"marking_model_ref,attr"`
	Color            string `xml:"color,attr"`
}

// VOLMarking stores the data from API response.
type VOLMarking struct {
	Text          string           `xml:",chardata"`
	MarkingStruct VOLMarkingStruct `xml:"Marking_Structure"`
}

// VOLHandling stores the data from API response.
type VOLHandling struct {
	Text    string     `xml:",chardata"`
	Marking VOLMarking `xml:"Marking"`
}

// VOLLI stores the data from API response.
type VOLLI struct {
	Text     string `xml:",chardata"`
	Resource string `xml:"resource,attr"`
}

// VOLSeq stores the data from API response.
type VOLSeq struct {
	Text string   `xml:",chardata"`
	LI   []*VOLLI `xml:"li"`
}

// VOLChannelItems stores the data from API response.
type VOLChannelItems struct {
	Text string `xml:",chardata"`
	Seq  VOLSeq `xml:"Seq"`
}

// VOLChannel stores the data from API response.
type VOLChannel struct {
	Text        string          `xml:",chardata"`
	About       string          `xml:"about,attr"`
	Title       string          `xml:"title"`
	Link        string          `xml:"link"`
	Description string          `xml:"description"`
	Date        string          `xml:"date"`
	Issued      string          `xml:"issued"`
	Modified    string          `xml:"modified"`
	Handling    VOLHandling     `xml:"handling"`
	Items       VOLChannelItems `xml:"items"`
}

// VOLReferences stores the data from API response.
type VOLReferences struct {
	Text   string `xml:",chardata"`
	Source string `xml:"source,attr"`
	ID     string `xml:"id,attr"`
	Title  string `xml:"title,attr"`
}

// VOLCPE stores the data from API response.
type VOLCPE struct {
	Text    string `xml:",chardata"`
	Version string `xml:"version,attr"`
	Vendor  string `xml:"vendor,attr"`
	Product string `xml:"product,attr"`
}

// VOLCVSS stores the data from API response.
type VOLCVSS struct {
	Text     string `xml:",chardata"`
	Score    string `xml:"score,attr"`
	Severity string `xml:"severity,attr"`
	Vector   string `xml:"vector,attr"`
	Version  string `xml:"version,attr"`
	Type     string `xml:"type,attr"`
}

// VOLItem stores the data from API response.
type VOLItem struct {
	Text        string           `xml:",chardata"`
	About       string           `xml:"about,attr"`
	Title       string           `xml:"title"`
	Link        string           `xml:"link"`
	Description string           `xml:"description"`
	Creator     string           `xml:"creator"`
	Identifier  string           `xml:"identifier"`
	References  []*VOLReferences `xml:"references"`
	CPEs        []*VOLCPE        `xml:"cpe"`
	CVSSes      []*VOLCVSS       `xml:"cvss"`
	Date        string           `xml:"date"`
	Issued      string           `xml:"issued"`
	Modified    string           `xml:"modified"`
}

// VulnOverviewList stores the data from API response.
type VulnOverviewList struct {
	XMLName        xml.Name   `xml:"RDF"`
	Text           string     `xml:",chardata"`
	XSI            string     `xml:"xsi,attr"`
	XMLNS          string     `xml:"xmlns,attr"`
	RSS            string     `xml:"rss,attr"`
	RDF            string     `xml:"rdf,attr"`
	DC             string     `xml:"dc,attr"`
	DCTerms        string     `xml:"dcterms,attr"`
	Sec            string     `xml:"sec,attr"`
	Marking        string     `xml:"marking,attr"`
	TLPMarking     string     `xml:"tlpMarking,attr"`
	AttrStatus     string     `xml:"status,attr"`
	SchemaLocation string     `xml:"schemaLocation,attr"`
	Lang           string     `xml:"lang,attr"`
	Channel        VOLChannel `xml:"channel"`
	Items          []*VOLItem `xml:"item"`
	Status         Status     `xml:"Status"`
}

// ParamsGetVulnOverviewList specifies the parameters of a HTTP request for GetVulnOverviewList.
type ParamsGetVulnOverviewList struct {
	Method                   string `url:"method"`
	Feed                     string `url:"feed"`
	StartItem                uint   `url:"startItem,omitempty"`
	MaxCountItem             uint8  `url:"maxCountItem,omitempty"`
	CPEName                  string `url:"cpeName,omitempty"`
	VendorID                 string `url:"vendorId,omitempty"`
	ProductID                string `url:"productId,omitempty"`
	Keyword                  string `url:"keyword,omitempty"`
	Severity                 string `url:"severity,omitempty"`
	Vector                   string `url:"vector,omitempty"`
	RangeDatePublic          string `url:"rangeDatePublic,omitempty"`
	RangeDatePublished       string `url:"rangeDatePublished,omitempty"`
	RangeDateFirstPublished  string `url:"rangeDateFirstPublished,omitempty"`
	DatePublicStartY         uint16 `xml:"datePublicStartY,omitempty"`
	DatePublicStartM         uint8  `xml:"datePublicStartM,omitempty"`
	DatePublicStartD         uint8  `xml:"datePublicStartD,omitempty"`
	DatePublicEndY           uint16 `xml:"datePublicEndY,omitempty"`
	DatePublicEndM           uint8  `xml:"datePublicEndM,omitempty"`
	DatePublicEndD           uint8  `xml:"datePublicEndD,omitempty"`
	DateFirstPublishedStartY uint16 `xml:"dateFirstPublishedStartY,omitempty"`
	DateFirstPublishedStartM uint8  `xml:"dateFirstPublishedStartM,omitempty"`
	DateFirstPublishedStartD uint8  `xml:"dateFirstPublishedStartD,omitempty"`
	DateFirstPublishedEndY   uint16 `xml:"dateFirstPublishedEndY,omitempty"`
	DateFirstPublishedEndM   uint8  `xml:"dateFirstPublishedEndM,omitempty"`
	DateFirstPublishedEndD   uint8  `xml:"dateFirstPublishedEndD,omitempty"`
	Language                 string `url:"lang,omitempty"`
}

// NewParamsGetVulnOverviewList creates an instance of ParamsGetVulnOverviewList.
func NewParamsGetVulnOverviewList(params *Parameter) *ParamsGetVulnOverviewList {
	if params == nil {
		params = &Parameter{}
	}

	p := &ParamsGetVulnOverviewList{
		Method: "getVulnOverviewList",
		Feed:   "hnd",
	}

	p.StartItem = params.StartItem
	p.MaxCountItem = params.MaxCountItem
	p.CPEName = params.CPEName
	p.VendorID = params.VendorID
	p.ProductID = params.ProductID
	p.Keyword = url.QueryEscape(params.Keyword)
	p.Severity = params.Severity
	p.Vector = params.Vector
	p.RangeDatePublic = params.RangeDatePublic
	p.RangeDatePublished = params.RangeDatePublished
	p.RangeDateFirstPublished = params.RangeDateFirstPublished
	p.DatePublicStartY = params.DatePublicStartY
	p.DatePublicStartM = params.DatePublicStartM
	p.DatePublicStartD = params.DatePublicStartD
	p.DatePublicEndY = params.DatePublicEndY
	p.DatePublicEndM = params.DatePublicEndM
	p.DatePublicEndD = params.DatePublicEndD
	p.DateFirstPublishedStartY = params.DateFirstPublishedStartY
	p.DateFirstPublishedStartM = params.DateFirstPublishedStartM
	p.DateFirstPublishedStartD = params.DateFirstPublishedStartD
	p.DateFirstPublishedEndY = params.DateFirstPublishedEndY
	p.DateFirstPublishedEndM = params.DateFirstPublishedEndM
	p.DateFirstPublishedEndD = params.DateFirstPublishedEndD
	p.Language = params.Language

	return p
}

// GetVulnOverviewList downloads a product list.
// See: https://jvndb.jvn.jp/apis/getVulnOverviewList_api_hnd.html
func (c *Client) GetVulnOverviewList(ctx context.Context, params *ParamsGetVulnOverviewList) (*VulnOverviewList, error) {
	if params == nil {
		params = NewParamsGetVulnOverviewList(nil)
	}

	u, err := addOptions(defaultAPIPath, params)
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("GET", u)
	if err != nil {
		return nil, err
	}

	productList := new(VulnOverviewList)
	err = c.do(ctx, req, nil, productList)
	if err != nil {
		return nil, err
	}

	return productList, nil
}

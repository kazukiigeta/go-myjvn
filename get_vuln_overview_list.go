// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
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

// GetVulnOverviewList downloads a product list.
// See: https://jvndb.jvn.jp/apis/getVulnOverviewList_api_hnd.html
func (c *Client) GetVulnOverviewList(ctx context.Context, opts ...Option) (*VulnOverviewList, error) {
	p := &parameter{
		Method: "getVulnOverviewList",
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

	productList := new(VulnOverviewList)
	err = c.do(ctx, req, nil, productList)
	if err != nil {
		return nil, err
	}

	return productList, nil
}

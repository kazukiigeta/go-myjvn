// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
)

// ALTitle stores the data from API response.
type ALTitle struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

//ALLink stores the data from API response.
type ALLink struct {
	Text     string `xml:",chardata"`
	Rel      string `xml:"rel,attr"`
	Type     string `xml:"type,attr"`
	HRefLang string `xml:"hreflang,attr"`
	HRef     string `xml:"href,attr"`
}

// ALAuthor stores the data from API response.
type ALAuthor struct {
	Text string `xml:",chardata"`
	Name string `xml:"name"`
	URI  string `xml:"uri"`
}

// ALMarkingStruct stores the data from API response.
type ALMarkingStruct struct {
	Text             string `xml:",chardata"`
	Type             string `xml:"type,attr"`
	MarkingModelName string `xml:"marking_model_name,attr"`
	MarkingModelRef  string `xml:"marking_model_ref,attr"`
	Color            string `xml:"color,attr"`
}

// ALMarking stores the data from API response.
type ALMarking struct {
	Text          string          `xml:",chardata"`
	MarkingStruct ALMarkingStruct `xml:"Marking_Structure"`
}

// ALHandling stores the data from API response.
type ALHandling struct {
	Text    string    `xml:",chardata"`
	Marking ALMarking `xml:"Marking"`
}

// ALCategory stores the data from API response.
type ALCategory struct {
	Text  string `xml:",chardata"`
	Label string `xml:"label,attr"`
	Term  string `xml:"term,attr"`
}

// ALItemLink stores the data from API response.
type ALItemLink struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
}

// ALItem stores the data from API response.
type ALItem struct {
	Text       string     `xml:",chardata"`
	Title      string     `xml:"title"`
	Identifier string     `xml:"identifier"`
	Link       ALItemLink `xml:"link"`
	Published  string     `xml:"published"`
	Updated    string     `xml:"updated"`
}

// ALItems stores the data from API response.
type ALItems struct {
	Text  string    `xml:",chardata"`
	Items []*ALItem `xml:"item"`
}

// ALEntry stores the data from API response.
type ALEntry struct {
	Text      string     `xml:",chardata"`
	Title     string     `xml:"title"`
	ID        string     `xml:"id"`
	Published string     `xml:"published"`
	Updated   string     `xml:"updated"`
	Category  ALCategory `xml:"category"`
	Items     ALItems    `xml:"items"`
}

// AlertList stores the data from API response.
type AlertList struct {
	XMLName xml.Name `xml:"feed"`
	//Text           string     `xml:",chardata"`
	//SchemaLocation string     `xml:"schemaLocation,attr"`
	//Lang           string     `xml:"lang,attr"`
	Title ALTitle `xml:"title"`
	//Updated        string     `xml:"updated"`
	//ID             string     `xml:"id"`
	//Link           ALLink     `xml:"link"`
	//Author         ALAuthor   `xml:"author"`
	//Handling       ALHandling `xml:"handling"`
	Entries []*ALEntry `xml:"entry"`
	//Status         Status     `xml:"Status"`
}

// GetAlertList downloads an alert list.
// See: https://jvndb.jvn.jp/apis/getAlertList_api_hnd.html
func (c *Client) GetAlertList(ctx context.Context, opts ...Option) (*AlertList, error) {
	p := &Parameter{
		Method: "getAlertList",
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

	alertList := new(AlertList)
	err = c.do(ctx, req, nil, alertList)
	if err != nil {
		return nil, err
	}

	return alertList, nil
}

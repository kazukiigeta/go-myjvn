// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
)

// StatisticsITM stores the data from API response.
type StatisticsITM struct {
	XMLName  xml.Name `xml:"Result"`
	SumJVNDB SumJVNDB `xml:"sumJvnDb"`
	SumCVSS  SumCVSS  `xml:"sumCvss"`
	SumCWE   SumCWE   `xml:"sumCwe"`
	Status   Status   `xml:"Status"`
}

// ParamsGetStatisticsITM specifies the parameters of a HTTP request for GetStatisticsITM.
type ParamsGetStatisticsITM struct {
	Method           string `url:"method"`
	Feed             string `url:"feed"`
	Theme            string `url:"theme"`
	DataType         string `url:"type,omitempty"`
	CWEID            string `url:"cweId"`
	PID              uint   `url:"pid,omitempty"`
	CPEName          string `url:"cpeName,omitempty"`
	DatePublicStartY uint16 `url:"datePublicStartY,omitempty"`
	DatePublicStartM uint8  `url:"datePublicStartM,omitempty"`
	DatePublicEndY   uint16 `url:"datePublicEndY,omitempty"`
	DatePublicEndM   uint8  `url:"datePublicEndM,omitempty"`
}

// NewParamsGetStatisticsITM creates an instance of ParamsGetStatisticsITM.
func NewParamsGetStatisticsITM(params *Parameter) *ParamsGetStatisticsITM {
	if params == nil {
		params = &Parameter{}
	}

	p := &ParamsGetStatisticsITM{
		Method: "getStatistics",
		Feed:   "itm",
	}

	p.Theme = params.Theme
	p.DataType = params.DataType
	p.CWEID = params.CWEID
	p.PID = params.PID
	p.CPEName = params.CPEName
	p.DatePublicStartY = params.DatePublicStartY
	p.DatePublicStartM = params.DatePublicStartM
	p.DatePublicEndY = params.DatePublicEndY
	p.DatePublicEndM = params.DatePublicEndM

	return p
}

// GetStatisticsITM downloads a product list.
// See: https://jvndb.jvn.jp/apis/getStatistics_api_itm.html
func (c *Client) GetStatisticsITM(ctx context.Context, params *ParamsGetStatisticsITM) (*StatisticsITM, error) {
	if params == nil {
		params = NewParamsGetStatisticsITM(nil)
	}

	u, err := addOptions(defaultAPIPath, params)
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("GET", u)
	if err != nil {
		return nil, err
	}

	statisticsITM := new(StatisticsITM)
	err = c.do(ctx, req, nil, statisticsITM)
	if err != nil {
		return nil, err
	}

	return statisticsITM, nil
}

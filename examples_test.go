// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn_test

import (
	"context"
	"fmt"

	"github.com/kazukiigeta/go-myjvn"
)

func ExampleClient_GetAlertList() {
	c := myjvn.NewClient(nil)
	alertList, err := c.GetAlertList(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(alertList.Title)
}

func ExampleClient_GetVendorList() {
	c := myjvn.NewClient(nil)
	vendorList, err := c.GetVendorList(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range vendorList.VendorInfo.Vendors {
		fmt.Println(*v)
	}
}

func ExampleClient_GetProductList() {
	c := myjvn.NewClient(nil)
	productList, err := c.GetProductList(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	for _, product := range productList.VendorInfo.Vendors[100].Products {
		fmt.Println(*product)
	}
}

func ExampleClient_GetVulnOverviewList() {
	c := myjvn.NewClient(nil)
	vulnOverviewList, err := c.GetVulnOverviewList(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(vulnOverviewList.Items[0].Title)
}

func ExampleClient_GetVulnDetailInfo() {
	c := myjvn.NewClient(nil)
	vulnDetailInfo, err := c.GetVulnDetailInfo(context.Background(), myjvn.SetVulnID("JVNDB-2020-006469"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(vulnDetailInfo.VulInfo.VulInfoID)
}

func ExampleClient_GetStatistics() {
	c := myjvn.NewClient(nil)

	statisticsHND, err := c.GetStatistics(context.Background(),
		myjvn.SetFeed("hnd"),
		myjvn.SetTheme("sumCvss"),
		myjvn.SetCWEID("CWE-20"),
		myjvn.SetDatePublicStartY(2015),
	)

	if err != nil {
		fmt.Println(err)
	}

	statisticsITM, err := c.GetStatistics(context.Background(),
		myjvn.SetFeed("itm"),
		myjvn.SetTheme("sumCvss"),
		myjvn.SetCWEID("CWE-20"),
		myjvn.SetDatePublicStartY(2015),
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(statisticsHND.SumCVSS.Titles)
	fmt.Println(statisticsITM.SumCVSS.Titles)
}

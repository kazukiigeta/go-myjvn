// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kazukiigeta/go-myjvn"
)

func main() {
	var (
		vulnid = flag.String("vulnid", "", "vulnID for vulnDetailInfo search")
	)
	flag.Parse()

	if *vulnid == "" {
		fmt.Println("vulnid must be specified.")
		os.Exit(1)
	} else if !strings.HasPrefix(*vulnid, "JVNDB") {
		fmt.Println("Invalid vulnID: vulnID must start with JVNDB")
		os.Exit(1)
	}

	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{
		VulnID: *vulnid,
	}
	p := myjvn.NewParamsGetVulnDetailInfo(params)
	vulnDetailInfo, err := c.GetVulnDetailInfo(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Result of the getVulnDetailInfo command")
	fmt.Println("------------------------------------------")

	n := len(vulnDetailInfo.VulInfo.VulInfoData.Affected.AffectedItem)
	var s string = strconv.Itoa(n)
	switch n {
	case 0:
		fmt.Println("vulnDetailInfo.VulInfo.VulInfoData.Affected.AffectedItem is nil")
		os.Exit(1)
	case 1:
		s += " AffectedItem is found."
	default:
		s += " AffectedItems are found."
	}

	fmt.Printf("%s\n\n", s)
	fmt.Printf("Here is the results.\n\n")
	for _, v := range vulnDetailInfo.VulInfo.VulInfoData.Affected.AffectedItem {
		fmt.Printf("%v\n", v.ProductName)
	}
	fmt.Println("---------------------------------------")
}

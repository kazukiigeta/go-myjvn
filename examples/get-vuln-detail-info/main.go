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
		vulnID = flag.String("vulnID", "", "vulnID for vulnDetailInfo search")
	)
	flag.Parse()

	if *vulnID == "" {
		fmt.Println("vulnID must be specified.")
		os.Exit(1)
	} else if !strings.HasPrefix(*vulnID, "JVNDB") {
		fmt.Println("Invalid vulnID: vulnID must start with JVNDB")
		os.Exit(1)
	}

	c := myjvn.NewClient(nil)
	vulnDetailInfo, err := c.GetVulnDetailInfo(context.Background(), myjvn.SetVulnID(*vulnID))
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
		fmt.Printf("ProductName: %v\n", v.ProductName)
		fmt.Printf("Version: %v\n", v.VersionNum)
		fmt.Printf("CPE: %v\n", v.CPE)
	}
	fmt.Println("---------------------------------------")
}

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

	"github.com/kazukiigeta/go-myjvn"
)

func main() {
	var (
		keyword = flag.String("keyword", "", "Keyword for vulnOverview search")
	)
	flag.Parse()

	if *keyword == "" {
		fmt.Println("keyword must be specified")
		os.Exit(1)
	}

	c := myjvn.NewClient(nil)
	vulnOverviewList, err := c.GetVulnOverviewList(context.Background(),
		myjvn.SetKeyword(*keyword),
		myjvn.SetRangeDatePublic("n"),
		myjvn.SetRangeDatePublished("n"),
		myjvn.SetRangeDateFirstPublished("n"),
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Result of the getVulnOverviewList command")
	fmt.Println("------------------------------------------")

	if vulnOverviewList == nil {
		fmt.Println("no items")
		os.Exit(1)
	}

	n := len(vulnOverviewList.Items)
	var s string = strconv.Itoa(n)
	switch n {
	case 0:
		fmt.Println("vulnOverviewList.Item is nil")
		os.Exit(1)
	case 1:
		s += " vulnOverview is found."
	case 50:
		s += " vulnOverview are found. (Reached upper limit of displaying)"
	default:
		s += " vulnOverviews are found."
	}

	fmt.Printf("%s\n\n", s)
	fmt.Printf("Here is the results.\n\n")
	for _, v := range vulnOverviewList.Items {
		fmt.Printf("%v\n", v.Title)
	}
	fmt.Println("---------------------------------------")
}

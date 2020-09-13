// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/kazukiigeta/go-myjvn"
)

func main() {
	var (
		theme            = flag.String("theme", "", "Theme for statistics search")
		cweID            = flag.String("cweID", "", "CWE ID for statistics search")
		datePublicStartY = flag.Uint("datePublicStartY", 0, "DatePublicStartY for statistics search")
	)
	flag.Parse()

	if *theme == "" {
		fmt.Println("theme must be specified.")
		os.Exit(1)
	}
	if *cweID == "" {
		fmt.Println("cweID must be specified.")
		os.Exit(1)
	}
	if *datePublicStartY == 0 {
		fmt.Println("datePublicStartY must be specified.")
		os.Exit(1)
	}

	c := myjvn.NewClient(nil)
	statisticsITM, err := c.GetStatistics(context.Background(),
		myjvn.SetFeed("itm"),
		myjvn.SetTheme(*theme),
		myjvn.SetCWEID(*cweID),
		myjvn.SetDatePublicStartY(uint16(*datePublicStartY)),
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Result of the getStatisticsITM command")
	fmt.Println("------------------------------------------")

	fmt.Printf("Here is the results.\n\n")
	fmt.Printf("%+v\n", *statisticsITM)
	fmt.Println("---------------------------------------")

}

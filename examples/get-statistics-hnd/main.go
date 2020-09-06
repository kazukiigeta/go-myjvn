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
		theme            = flag.String("theme", "sumAll", "Theme for statistics search")
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
	params := &myjvn.Parameter{
		Theme:            *theme,
		CWEID:            *cweID,
		DatePublicStartY: uint16(*datePublicStartY),
	}
	p := myjvn.NewParamsGetStatisticsHND(params)
	statisticsHND, err := c.GetStatistics(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Result of the getStatisticsHND command")
	fmt.Println("------------------------------------------")

	fmt.Printf("Here is the results.\n\n")
	fmt.Printf("%+v\n", *statisticsHND)
	fmt.Println("---------------------------------------")

}

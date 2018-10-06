package main

import (
	"flag"
	"fmt"

	"github.com/danielkvist/inspopular/pkg/hashtag"
)

func main() {
	flag.Parse()
	tags := flag.Args()

	l := hashtag.CreateList(tags)

	for _, item := range *l {
		fmt.Printf("%+v\n", *item)
	}
}

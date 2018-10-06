package main

import (
	"flag"
	"fmt"

	"github.com/danielkvist/inspopular/pkg/hashtag"
)

func main() {
	flag.Parse()
	tags := flag.Args()

	for _, t := range tags {
		h := hashtag.New(t)
		h.Popularity()

		fmt.Printf("%+v\n", h)
	}
}

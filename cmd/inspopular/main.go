package main

import (
	"flag"

	"github.com/danielkvist/inspopular/pkg/hashtag"
)

func main() {
	flag.Parse()
	tags := flag.Args()

	l := hashtag.CreateList(tags)
	l.Print()
}

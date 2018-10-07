package hashtag

import (
	"fmt"
	"log"
	"os"
	"sync"
	"text/tabwriter"
)

// List type is a slice of *hashtag.
type List []*hashtag

// CreateList is a method of the List type to create a new List
// with the items contained in the slice that receives as parameter
// and and to obtain concurrently the popularity of each item.
func CreateList(tags []string) *List {
	var w sync.WaitGroup
	var list List

	for _, t := range tags {
		list = append(list, newhashtag(t))
	}

	w.Add(len(tags))
	for i := range list {
		go func(i int) {
			defer w.Done()
			if err := list[i].popularity(); err != nil {
				log.Println(err)
			}
		}(i)
	}

	w.Wait()
	return &list
}

// Print method of the List type provides a pretty format
// to print the results.
func (l *List) Print() {
	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 4, 4, ' ', 0)
	fmt.Fprintf(tw, format, "Hashtag", "URL", "Posts")
	fmt.Fprintf(tw, format, "-------", "---", "-----")

	fmt.Println()
	for _, item := range *l {
		fmt.Fprintf(tw, format, item.tag, item.url, item.posts)
	}
	fmt.Println()
	tw.Flush()
}

// OrderedList type is a slice of *hashtag that
// implements the sort interface.
type OrderedList []*hashtag

// Len is a method of the OrderedList type to satisfy
// the sort interface.
func (ol OrderedList) Len() int {
	return len(ol)
}

// Less is a method of the OrderedList type to satisfy
// the sort interface.
func (ol OrderedList) Less(i, j int) bool {
	return ol[i].posts > ol[j].posts
}

// Swap is a method of the OrderedList type to satisfy
// the sort interface.
func (ol OrderedList) Swap(i, j int) {
	ol[i], ol[j] = ol[j], ol[i]
}

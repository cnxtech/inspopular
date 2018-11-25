# INSPOPULAR

Inspopular is a CLI to get easily the popularity of a list of hashtags in Instagram. No API is used in this tool. The information is extracted using HTTP requests.

## Example

The tool is actually simple and you just need to pass the list of hashtags you want to check.

```
$ ./inspopular go golang dev developer
```

The output of the command above would be something like this:

```output
Hashtag      URL                                                 Posts
-------      ---                                                 -----
go           https://www.instagram.com/explore/tags/go           14818638
developer    https://www.instagram.com/explore/tags/developer    1237858
dev          https://www.instagram.com/explore/tags/dev          428791
golang       https://www.instagram.com/explore/tags/golang       11694

#go #developer #dev #golang
```

As you can see the results are sorted by popularity. You can change this by using the flag ```-s```.

```shell
$ ./inspopular -s=false go golang dev developer
```

Or

```shell
$ ./inspopular -s="false" go golang dev developer
```

## Get & Run

First of all you need to have the Go programming language installed. Then:

```shell
$ go get github.com/danielkvist/inspopular
```

> Obviously, you can also use git to clone the repository.

Once the project has been downloaded. You only need to execute the following command into it.

```shell
$ go build /cmd/inspopular/main.go
```

You should also be able to install it with the following command:

```shell
$ go install /cmd/inspopular/main.go
```

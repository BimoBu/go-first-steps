package main

import (
	"fmt"

	binaryTreeChannels "example.com/a-tour-of-go/binary-tree-channels"
	"example.com/a-tour-of-go/channel"
	channel2 "example.com/a-tour-of-go/channel-2"
	"example.com/a-tour-of-go/crawler"
	"example.com/a-tour-of-go/errors"
	fibonacciClosure "example.com/a-tour-of-go/fibonacci-closure"
	"example.com/a-tour-of-go/image"
	implementingStringer "example.com/a-tour-of-go/implementing-stringer"
	"example.com/a-tour-of-go/readers"
	rot13Reader "example.com/a-tour-of-go/rot-13-reader"
	"example.com/a-tour-of-go/slices"
	"example.com/a-tour-of-go/sqrt"
	wordCount "example.com/a-tour-of-go/word-count"
)

func main() {
	fmt.Println("### binaryTreeChannels ###")
	binaryTreeChannels.Test()
	fmt.Println("")

	fmt.Println("### channel ###")
	channel.Test()
	fmt.Println("")

	fmt.Println("### channel2 ###")
	channel2.Test()
	fmt.Println("")

	fmt.Println("### crawler ###")
	crawler.Test()
	fmt.Println("")

	fmt.Println("### errors ###")
	errors.Test()
	fmt.Println("")

	fmt.Println("### fibonacciClosure ###")
	fibonacciClosure.Test()
	fmt.Println("")

	fmt.Println("### image ###")
	image.Test()
	fmt.Println("")

	fmt.Println("### implementingStringer ###")
	implementingStringer.Test()
	fmt.Println("")

	fmt.Println("### readers ###")
	readers.Test()
	fmt.Println("")

	fmt.Println("### rot13Reader ###")
	rot13Reader.Test()
	fmt.Println("")

	fmt.Println("### slices ###")
	slices.Test()
	fmt.Println("")

	fmt.Println("### sqrt ###")
	sqrt.Test()
	fmt.Println("")

	fmt.Println("### wordCount ###")
	wordCount.Test()
	fmt.Println("")
}

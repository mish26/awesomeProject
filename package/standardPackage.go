package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Clock())

	match, _ := regexp.MatchString("a([a-z]+)","apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)")
	m := r.MatchString("apple")
	fmt.Println(m)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	f2 := r2.FindString("/view/test")
	fmt.Println(f2)
	fs := r2.FindStringSubmatch("/view/tes")
	fmt.Println(fs[0], fs[1], fs[2])



}



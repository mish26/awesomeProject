package main

import (
	"fmt"
	"regexp"
	"sort"
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

	fmt.Println("---about sort---")
	intSlice := []int{5,6,10,100,1,5,4,32}
	stringSlice:= []string{"j","a","f"}
	p := []struct{
		Name string
		Age int
	}{
		{"Nancy",20},
		{"Nike",21},
		{"Dik",29},
		{"Bob",5},
	}
	fmt.Println(intSlice,stringSlice,p)

	sort.Ints(intSlice)
	fmt.Println(intSlice)
	sort.Strings(stringSlice)
	fmt.Println(stringSlice)

	sort.Slice(p,
		func(i, j int) bool {
			return p[i].Name < p[j].Name
		})
	sort.Slice(p,
		func(i, j int) bool {
			return p[i].Age < p[j].Age
		})
	fmt.Println(p)
}



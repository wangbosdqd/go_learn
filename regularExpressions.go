package main

import (
	"regexp"
	"fmt"
	"bytes"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")//charcter regular match p+char+ch
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")//give reg to r  p+char+ch

	fmt.Println(r.MatchString("peach"))// same to line10

	fmt.Println(r.FindString("peach punch"))//find p+char+ch

	fmt.Println(r.FindStringIndex("peach punch"))//find p+char+ch location

	fmt.Println(r.FindStringSubmatch("peach punch"))//sub p+char+ch char

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))//after sub back location

	fmt.Println(r.FindAllString("peach punch pinch", -1))//-1 means all find all string

	fmt.Println(r.FindAllStringSubmatchIndex(//find all string location
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))//find 2 string

	fmt.Println(r.Match([]byte("peach")))//find p+str+ch in peach

	r = regexp.MustCompile("p([a-z]+)ch")//put the value p+str+ch to r

	fmt.Println(r)//output

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))//find p+str+ch and replace by <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))//replace "a peach"by upper
}





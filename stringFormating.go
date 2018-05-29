package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p) //only value

	fmt.Printf("%+v\n", p)//{key:value}

	fmt.Printf("%#v\n", p)//function name+strcut name+{}
	fmt.Printf("%T\n", p)//function name+struce name

	fmt.Printf("%t\n", true)//bool

	fmt.Printf("%d\n", 123)//digital value

	fmt.Printf("%b\n", 14)//digital 2 ascII

	fmt.Printf("%c\n", 33)//ascII 2 real
	fmt.Printf("%x\n", 456)//real 2 16

	fmt.Printf("%E\n", 78.9)//real 2 point6

	fmt.Printf("%e\n", 123400000.0)//real 2 science way
	fmt.Printf("%E\n", 123400000.0)//real 2 upper science way

	fmt.Printf("%s\n", "\"string\"")//format way
	fmt.Printf("%q\n", "\"string\"")//not formate way
	fmt.Printf("%x\n", "hex this")//hex way

	fmt.Printf("%p\n", &p)//adress

	fmt.Printf("|%6d|%6d|\n", 12, 345)//right alignment

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)//right alignment

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)//left alignment

	fmt.Printf("|%6s|%6s|\n", "foo", "b")//right alignment char

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")//left alignment char

	s := fmt.Sprintf("a %s", "string")//only back data not out

	fmt.Println(s)//print output

	fmt.Fprintf(os.Stderr, "an %s\n", "error")


}
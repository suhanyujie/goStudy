package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.CommandLine = flag.NewFlagSet()
}

// 入口函数
func main() {
	var name string;
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,"Usage of %s:\n", "'cmd Greeeting'")
		flag.PrintDefaults()
	}
	//var name = flag.String("name", "everyone","the greeting object")
	flag.StringVar(&name, "name", "everyone", "this is for greetting")
	flag.Parse()
	fmt.Printf("hello,%s!\n", name);
}

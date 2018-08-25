package main

import (
	"flag"
	"fmt"
	"os"
	"practice/lessonHao/command/lib"
)

func init() {
	//flag.CommandLine = flag.NewFlagSet("",flag.ExitOnError)
	//flag.CommandLine.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
}

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

// 入口函数
func main() {
	var name string;

	lib.ShowLib()

	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr,"Usage of %s:\n", "'cmd Greeeting'")
	//	flag.PrintDefaults()
	//}

	cmdLine.StringVar(&name, "name", "everyone", "this is for greetting")
	//flag.Parse()
	cmdLine.Parse(os.Args[1:])


	fmt.Printf("hello,%s!\n", name);
}

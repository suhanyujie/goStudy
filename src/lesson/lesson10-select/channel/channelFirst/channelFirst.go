package channelFirst

import (
	"flag"
	"runtime"
	"fmt"
)

func Test1() {
	var numCores = flag.Int("n", 2, "num of CPU cores to use")
	flag.Parse()
	runtime.GOMAXPROCS(*numCores)
	fmt.Println(*numCores)
}

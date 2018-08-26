package read

import (
	"bytes"
	"encoding/gob"
	"log"
	"fmt"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

// 入口函数
func TestGob1() {
	var netwrok bytes.Buffer

	enc := gob.NewEncoder(&netwrok)
	dec := gob.NewDecoder(&netwrok)
	err := enc.Encode(P{3, 4, 5, "Suhanyu"})
	if err != nil {
		log.Fatal(err)
	}
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q:{%d,%d}\n", q.Name, *q.X, *q.Y)
}

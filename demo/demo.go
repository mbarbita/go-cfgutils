package main

import (
	"fmt"

	cfgutils "github.com/mbarbita/golib-cfgutils"
)

func main() {

	cfgMap := cfgutils.ReadCfgFile("cfg.ini")
	// fmt.Println(cfgMap)

	fmt.Println("Result:")

	for k, v := range cfgMap {
		fmt.Printf("[%v]=[%v]\n", k, v)
	}
}

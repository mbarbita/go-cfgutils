package main

import (
	"fmt"

	"github.com/mbarbita/go-lib/cfg-utils"
)

func main() {

	cfgMap := cfgutils.ReadCfgFile("cfg.ini")
	// fmt.Println(cfgMap)

	fmt.Println("Result:")

	for k, v := range cfgMap {
		fmt.Printf("[%v]=[%v]\n", k, v)
	}
}

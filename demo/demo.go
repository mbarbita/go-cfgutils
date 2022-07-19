package main

import (
	"fmt"

	cfgutils "github.com/mbarbita/go-cfgutils"
)

func main() {

	//display sensitive data:
	fmt.Println("display sensitive data:")
	cfgMap := cfgutils.ReadCfgFile("cfg.ini", false)

	fmt.Println("Result:")

	for k, v := range cfgMap {
		fmt.Printf("[%v]=[%v]\n", k, v)
	}

	//hide sensitive data:
	fmt.Println("hide sensitive data:")
	cfgMap = cfgutils.ReadCfgFile("cfg.ini", true)

	fmt.Println("Result:")

	for k, v := range cfgMap {
		fmt.Printf("[%v]=[%v]\n", k, v)
	}

}

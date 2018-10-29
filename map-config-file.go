package cfgutils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadCfgFile(fileName string) map[string]string {
	fmt.Println("loading config file:", fileName, "...")

	cfgMap := make(map[string]string)

	// Open cfg file
	cfgFile, err := os.Open(fileName)
	if err != nil {
		log.Println("cfg file error:")
		log.Fatal(err)
	}
	defer cfgFile.Close()

	// prepare scanner, scan each line
	scanner := bufio.NewScanner(cfgFile)
	index := 1       // index
	cfgErrors := 0   // errors
	cfgWarnings := 0 // warnings

	for scanner.Scan() {

		// fmt.Printf("processing line: %v [%v]\n", index, scanner.Text())

		// skip comment lines
		if strings.HasPrefix(scanner.Text(), "#") {
			fmt.Println("SKIPPING: comment line", index)
			index++
			continue
		}

		// get firtst =, count =
		eqi := strings.Index(scanner.Text(), "=")
		eqc := strings.Count(scanner.Text(), "=")

		// skip empty lines
		if len(strings.TrimSpace(scanner.Text())) == 0 {
			fmt.Println("SKIPPING: empty line", index)
			index++
			continue
		}

		// check for multiple =
		if eqc > 1 {
			// fmt.Println("WARNING:", eqc, "\"=\" at line", index)
			fmt.Printf("++ WARNING: %v \"=\" at line %v\n", eqc, index)
			fmt.Printf("[%v]\n\n", scanner.Text())
			cfgWarnings++
		}

		//check for no =
		if eqi == -1 {
			fmt.Println("++ WARNING: no \"=\" at line", index)
			fmt.Printf("[%v]\n\n", scanner.Text())
			cfgWarnings++
			index++
			continue
		}

		// trim spaces and convert the key to lower case,
		f1 := strings.TrimSpace(strings.ToLower(scanner.Text()[:eqi]))
		// trim spaces on value
		f2 := strings.TrimSpace(scanner.Text()[eqi+1:])

		// check for duplicate keys
		_, ok := cfgMap[f1]
		if ok {
			fmt.Println("++ WARNING: duplicated (overwriten) key at line", index)
			fmt.Printf("[%v]\n\n", scanner.Text())
			cfgWarnings++
		}

		// add key=value to map
		cfgMap[f1] = f2
		index++

	} //end for

	fmt.Println("loading config file", fileName, "done",
		"\nwarnings:", cfgWarnings, "\nerrors:", cfgErrors)

	fmt.Println()

	return cfgMap
}

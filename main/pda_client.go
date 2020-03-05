package main

import (
	. "../src"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	pdaProcessor := PdaProcessor{}

	if len(os.Args) == 3 {
		if os.Args[1] == "" {
			print("Give PDA specifications ")
		} else {

			trannsactionCount := 1
			specJSON, _ := ioutil.ReadFile(os.Args[1])
			if pdaProcessor.Open(specJSON) {
				if os.Args[2] != "" {
					inputString := ""
					inputBytes, err := ioutil.ReadFile(os.Args[2])
					if err != nil {
						inputString = os.Args[2]
					} else {
						inputString = string(inputBytes)
					}
					pdaProcessor.Reset()
					fmt.Printf("PDA Name=%s ,Token=START ,Transaction Number=%d \n", pdaProcessor.PdaConf.Name, trannsactionCount)
					for _, alphabet := range inputString {
						trannsactionCount = pdaProcessor.Put(string(alphabet))
						fmt.Printf("PDA Name=%s ,Token=%s ,Transaction Number=%d \n", pdaProcessor.PdaConf.Name, string(alphabet), trannsactionCount)
					}
					pdaProcessor.Eos()
					fmt.Printf("PDA Name=%s ,Token=%s ,Transaction Number=%d \n", pdaProcessor.PdaConf.Name, pdaProcessor.PdaConf.Eos, trannsactionCount+1)

				} else {
					print("\nPDA no input stream specified")
				}

			} else {
				print("\nPDA open (spec) API failed!!")
			}
		}
	} else {
		print("Wrong number of inputs ", len(os.Args))
	}
}

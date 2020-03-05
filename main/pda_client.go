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
			println("Give PDA specifications ")
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
					fmt.Printf("PDA Name=%s ,\tToken=START ,\tTransition Number=%d \n", pdaProcessor.GetPDAName(), trannsactionCount)
					for _, alphabet := range inputString {
						trannsactionCount = pdaProcessor.Put(string(alphabet))
						fmt.Printf("PDA Name=%s ,Token=%s ,\tTransition Number=%d \n", pdaProcessor.GetPDAName(), string(alphabet), trannsactionCount)
					}
					pdaProcessor.Eos()
					fmt.Printf("PDA Name=%s ,\tToken=EOS ,\tTransition Number=%d \n", pdaProcessor.GetPDAName(), trannsactionCount+1)

				} else {
					println("PDA no input stream specified")
				}

			} else {
				println("PDA open (spec) API failed!!")
			}
		}
	} else {
		println("Wrong number of inputs provided, Input length", len(os.Args))
	}
}

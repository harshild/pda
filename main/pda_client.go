package main

import (
	. "../src"
	"io/ioutil"
	"os"
)

func main() {
	pdaProcessor := PdaProcessor{}

	if len(os.Args) == 3 {
		if os.Args[1] == "" {
			print("Give PDA specifications ")
		} else {

			specJSON, _ := ioutil.ReadFile(os.Args[1])
			if pdaProcessor.Open(specJSON) {
				if os.Args[2] != "" {
					inputBytes, err := ioutil.ReadFile(os.Args[2])
					inputString := string(inputBytes)
					if err != nil {
						inputString := os.Args[2]
						count := pdaProcessor.Put(inputString)
						print("\nNumber of transitons are ", count, "...\nDone!!")
					} else {
						count := pdaProcessor.Put(inputString)
						print("\nNumber of transitons are ", count, "...\nDone!!")
					}
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

package main

import (
	. "../src"
	"io/ioutil"
	"os"
)

func main() {
	pda := PdaController{}

	if len(os.Args) == 2 {
		if os.Args[1] == "" {
			print("Give PDA specifications ")
		} else {

			specJSON, _ := ioutil.ReadFile(os.Args[1])
			if pda.Open(specJSON) {
				if os.Args[2] != "" {
					inputBytes, err := ioutil.ReadFile(os.Args[2])
					inputString := string(inputBytes)
					if err != nil {
						inputString := os.Args[2]
						count := pda.Put(inputString)
						print("\nNumber of transitons are ", count, "...\nDone!!")
					} else {
						count := pda.Put(inputString)
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
		print("Wrong number of inputs")
	}
}

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

			transition := 1
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
					fmt.Printf("PDA Name=%s \tMethod=Is_Accepted =%t \n", pdaProcessor.GetPDAName(), pdaProcessor.Is_accepted())

					pdaProcessor.Reset()
					fmt.Printf("PDA Name=%s \tToken=START \t Transitions Took=%d\tClock Ticks=%d \n", pdaProcessor.GetPDAName(), pdaProcessor.GetClock(), pdaProcessor.GetClock())

					for _, alphabet := range inputString {
						transition = pdaProcessor.Put(string(alphabet))
						fmt.Printf("PDA Name=%s \tToken=%s \t Transitions Took=%d\tClock Ticks=%d \n", pdaProcessor.GetPDAName(), string(alphabet), transition, pdaProcessor.GetClock())
					}
					prevClock := pdaProcessor.GetClock()
					pdaProcessor.Eos()
					fmt.Printf("PDA Name=%s \tToken=EOS \t Transitions Took=%d\tClock Ticks=%d \n", pdaProcessor.GetPDAName(), pdaProcessor.GetClock()-prevClock, pdaProcessor.GetClock())

					fmt.Print("For PDA=" + pdaProcessor.GetPDAName() + "\t Current State=" + pdaProcessor.Current_state() + "\n")
					pdaProcessor.Peek(5)

					pdaProcessor.Close()

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

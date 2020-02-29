package src

import (
	"encoding/json"
	"io/ioutil"
	"os"
)
var input_string string
var file1 []byte
var count int
func main() {
	pda := PdaController{}

	if os.Args[1] == "" {
		print("Give PDA specifications ")
	}  else{

		file1, _ = ioutil.ReadFile(os.Args[1])
		var ready = pda.Open(file1)
		if ready == false { print("\nPDA open (spec) API failed!!") }
		if os.Args[2] != ""{
			input_string, _ := ioutil.ReadFile(os.Args[2])
			count = pda.Put(string(input_string))
			print("\nNumber of transitons are ", count,"...\nDone!!")
			}

		}
	}
}


func read_token(in []byte) bool {
	var err = json.Unmarshal(in, &a)
	if err != nil {
		return false
	}
	return true
}

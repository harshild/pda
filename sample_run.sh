echo "================================="
echo "Input = 0011"
echo "================================="
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-1.json 0011
echo "================================="
echo "Input = aabcc"
echo "================================="
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-2.json aabcc
echo "================================="
echo "Input = 0101220101"
echo "================================="
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json 0101220101
echo "================================="
echo "Input = 01010101   "
echo "Should fail because, endup in no final state"
echo "================================="
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json 01010101   #endup in no final state
echo "================================="
echo "Input = 0101330101  "
echo "Should fail because, stuck in middle of the PDA because of an undefined transition occurence"
echo "================================="
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json 0101330101  #stuck in middle of the PDA because of an undefined transition occurence
echo "================================="
echo "Input =  "
echo "Should fail because, no input string found"
echo "================================="
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json #no input string found
echo "================================="
echo "Input = abc "
echo "Should fail because, getting an unknown input transition or undefined transition"
echo "================================="
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json abc #getting an unknown input transition or undefined transition



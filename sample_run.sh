go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-1.json 0011
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-2.json aabcc
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json 0101220101
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json 01010101   #endup in no final state
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json 0101330101  #stuck in middle of the PDA because of an undefined transition occurence
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json #no input string found
go build ./main/pda_client.go && ./pda_client ./sample/samplePDA-3.json abc #getting an unknown input transition or undefined transition



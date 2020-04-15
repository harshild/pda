package usecase

import (
	"core"
	"db"
	"encoding/json"
	"fmt"
)

type PDAManager struct {
	//PdaProcessor core.PdaProcessor
	PdaStore db.InMemoryStore
}

//func (pdaManager *PDAManager) NewPDA(id int, json string) {
//
//	if pdaProcessor.Open([]byte(json)) {
//		pdaStore.Save(id, pdaProcessor.PdaConf)
//	}
//}

func (pdaManager *PDAManager) ListAllPDAs() []string {
	pdas := pdaManager.PdaStore.GetAllPDA()
	pdaProcessors := MapStringToPDAProcessor(pdas)
	pdaNames := core.GetAllPDANames(pdaProcessors)
	return pdaNames
}

func (pdaManager *PDAManager) CreateNewPDA(id string, conf string) {
	pdaProcessor := core.PdaProcessor{}
	if pdaProcessor.Open([]byte(conf)) {
		pdaManager.PdaStore.Save(id, pdaProcessor)
	}
}

func MapStringToPDAProcessor(pdas []string) []core.PdaProcessor {
	pdaProcessors := make([]core.PdaProcessor, 0)
	for _, pdaProcessorString := range pdas {
		pdaProcessor := parsePdaProcessor(pdaProcessorString)
		pdaProcessors = append(pdaProcessors, pdaProcessor)
	}
	return pdaProcessors
}

func parsePdaProcessor(pdaProcessorString string) core.PdaProcessor {
	pdaProcessor := core.PdaProcessor{}
	json.Unmarshal([]byte(pdaProcessorString), &(pdaProcessor))
	return pdaProcessor
}

func (pdaManager *PDAManager) Reset(id string) {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.Reset()
	fmt.Printf("PDA Name=%s \tToken=START \t Transitions Took=%d\tClock Ticks=%d \n", pdaProcessor.GetPDAName(), pdaProcessor.GetClock(), pdaProcessor.GetClock())
	pdaManager.PdaStore.Update(id, pdaProcessor)
}

func (pdaManager *PDAManager) Puts(id string, token string, position int) {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.Puts(position, token)
	pdaManager.PdaStore.Update(id, pdaProcessor)
}

func (pdaManager *PDAManager) Is_accepted(id string) bool {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaProcessor := parsePdaProcessor(get)
	isAccepted := pdaProcessor.Is_accepted()
	fmt.Printf("PDA Name=%s \tMethod=Is_Accepted =%t \n", pdaProcessor.GetPDAName(), isAccepted)
	return isAccepted
}

func (pdaManager *PDAManager) Peek(id string, k int) []string {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaProcessor := parsePdaProcessor(get)
	return pdaProcessor.Peek(k)

}

func (pdaManager *PDAManager) Size(id string) int {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaProcessor := parsePdaProcessor(get)
	return pdaProcessor.Stack.Size()
}

func (pdaManager *PDAManager) Currentstate(id string) string {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaProcessor := parsePdaProcessor(get)
	return pdaProcessor.Current_state()
}

func (pdaManager *PDAManager) Queued_token(id string) []string {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaProcessor := parsePdaProcessor(get)
	return pdaProcessor.Queued_tokens()
}

func (pdaManager *PDAManager) Close(id string) {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.Close()
	pdaManager.PdaStore.Update(id, pdaProcessor)
}

func (pdaManager *PDAManager) Deletepda(id string) {
	//get, _ := pdaManager.PdaStore.Get(id)

	//  TODO: a query to delete the pda

}

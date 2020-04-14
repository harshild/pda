package usecase

import (
	"core"
	"db"
	"encoding/json"
)

type PDAManager struct {
	PdaProcessor core.PdaProcessor
	PdaStore     db.InMemoryStore
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
	if pdaManager.PdaProcessor.Open([]byte(conf)) {
		pdaManager.PdaStore.Save(id, pdaManager.PdaProcessor)
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

func (pdaManager *PDAManager) PdaProcessorcallsreset(id string) {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaManager.PdaProcessor = parsePdaProcessor(get)
	pdaManager.PdaProcessor.Reset()
	pdaManager.PdaStore.Update(id, pdaManager.PdaProcessor)
}

func (pdaManager *PDAManager) PdaProcessorcallputs(id string, token string, position int) {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaManager.PdaProcessor = parsePdaProcessor(get)
	pdaManager.PdaProcessor.Puts(position, token)
	pdaManager.PdaStore.Update(id, pdaManager.PdaProcessor)
}

func (pdaManager *PDAManager) PdaProcessorcallis_accepted(id string) bool {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaManager.PdaProcessor = parsePdaProcessor(get)
	return pdaManager.PdaProcessor.Is_accepted()
}

func (pdaManager *PDAManager) Peek(id string, k int) []string {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaManager.PdaProcessor = parsePdaProcessor(get)
	return pdaManager.PdaProcessor.Peek(k)

}

func (pdaManager *PDAManager) Callsize(id string) int {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaManager.PdaProcessor = parsePdaProcessor(get)
	return pdaManager.PdaProcessor.Stack.Size()
}

func (pdaManager *PDAManager) Currentstate(id string) string {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaManager.PdaProcessor = parsePdaProcessor(get)
	return pdaManager.PdaProcessor.Current_state()
}

func (pdaManager *PDAManager) Q_token(id string) []string {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaManager.PdaProcessor = parsePdaProcessor(get)
	return pdaManager.PdaProcessor.Queued_tokens()
}

func (pdaManager *PDAManager) Callclose(id string) {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaManager.PdaProcessor = parsePdaProcessor(get)
	pdaManager.PdaProcessor.Close()
	pdaManager.PdaStore.Update(id, pdaManager.PdaProcessor)
}

func (pdaManager *PDAManager) Deletepda(id string) {
	get, _ := pdaManager.PdaStore.Get(id)
	pdaManager.PdaProcessor = parsePdaProcessor(get)
	//  TODO: a query to delete the pda

}

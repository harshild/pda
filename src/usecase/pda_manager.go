package usecase

import (
	"core"
	"db"
	"encoding/json"
)

type PDAManager struct {
	PdaProcessor core.PdaProcessor
	PdaStore     db.PDAStore
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

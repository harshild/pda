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

func (pdaManager *PDAManager) CreateNewPDA(id string, conf string) error {
	pdaProcessor := core.PdaProcessor{}
	if pdaProcessor.Open([]byte(conf)) {
		pdaManager.PdaStore.Save(id, pdaProcessor)
	}
	return nil
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

func (pdaManager *PDAManager) Reset(id string) error {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return err
	}
	pdaProcessor := parsePdaProcessor(get)
	fmt.Printf(" Name:%s  Token:%s Position: N/A  \n", pdaProcessor.PdaConf.Name, "START")
	pdaProcessor.Reset()
	pdaManager.PdaStore.Update(id, pdaProcessor)
	return err
}

func (pdaManager *PDAManager) Puts(id string, token string, position int) error {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return err
	}
	pdaProcessor := parsePdaProcessor(get)
	fmt.Printf(" Name:%s  Token:%s Position:%d  \n", pdaProcessor.PdaConf.Name, token, position)
	err = pdaProcessor.Puts(position, token)
	if err != nil {
		return err
	}
	pdaManager.PdaStore.Update(id, pdaProcessor)
	return err
}

func (pdaManager *PDAManager) Is_accepted(id string) (bool, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return false, err
	}
	pdaProcessor := parsePdaProcessor(get)
	isAccepted := pdaProcessor.Is_accepted()
	fmt.Printf("PDA Name=%s \tMethod=Is_Accepted =%t \n", pdaProcessor.GetPDAName(), isAccepted)
	return isAccepted, err
}

func (pdaManager *PDAManager) Peek(id string, k int) ([]string, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return nil, err
	}
	pdaProcessor := parsePdaProcessor(get)
	return pdaProcessor.Peek(k)

}

func (pdaManager *PDAManager) Size(id string) (int, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return -1, err
	}
	pdaProcessor := parsePdaProcessor(get)
	return pdaProcessor.Stack.Size(), nil
}

func (pdaManager *PDAManager) Currentstate(id string) (string, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return "", err
	}
	pdaProcessor := parsePdaProcessor(get)
	return pdaProcessor.Current_state(), err
}

func (pdaManager *PDAManager) Queued_token(id string) ([]string, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return nil, err
	}
	pdaProcessor := parsePdaProcessor(get)
	return pdaProcessor.Queued_tokens(), nil
}

func (pdaManager *PDAManager) Close(id string) error {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return err
	}
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.Close()
	pdaManager.PdaStore.Update(id, pdaProcessor)
	return nil
}

func (pdaManager *PDAManager) Deletepda(id string) error {
	_, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return err
	}
	pdaManager.PdaStore.Delete(id)
	return nil
}

func (pdaManager *PDAManager) PutsEOS(id string, position int) error {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return err
	}
	pdaProcessor := parsePdaProcessor(get)
	fmt.Printf(" Name:%s  Token:%s Position:%d  \n", pdaProcessor.PdaConf.Name, pdaProcessor.PdaConf.Eos, position)
	pdaProcessor.Puts(position, " ")
	pdaManager.PdaStore.Update(id, pdaProcessor)
	return nil
}

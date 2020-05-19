package usecase

import (
	"core"
	"db"
	"encoding/json"
	"entity"
	"fmt"
	"math/rand"
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

func (pdaManager *PDAManager) Puts(id string, token string, position int, pdaStatus entity.PDAStatus) error {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return err
	}
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.UpdateStatus(pdaStatus)
	fmt.Printf(" Name:%s  Token:%s Position:%d  \n", pdaProcessor.PdaConf.Name, token, position)
	err = pdaProcessor.Puts(position, token)
	if err != nil {
		return err
	}
	pdaManager.PdaStore.Update(id, pdaProcessor)
	return err
}

func (pdaManager *PDAManager) Is_accepted(id string, pdaStatus entity.PDAStatus) (bool, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return false, err
	}
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.UpdateStatus(pdaStatus)
	isAccepted := pdaProcessor.Is_accepted()
	fmt.Printf("PDA Name=%s \tMethod=Is_Accepted =%t \n", pdaProcessor.GetPDAName(), isAccepted)
	return isAccepted, err
}

func (pdaManager *PDAManager) Peek(id string, k int, pdaStatus entity.PDAStatus) ([]string, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return nil, err
	}
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.UpdateStatus(pdaStatus)
	return pdaProcessor.Peek(k)

}

func (pdaManager *PDAManager) Size(id string, pdaStatus entity.PDAStatus) (int, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return -1, err
	}
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.UpdateStatus(pdaStatus)

	return pdaProcessor.Stack.Size(), nil
}

func (pdaManager *PDAManager) Currentstate(id string, pdaStatus entity.PDAStatus) (string, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return "", err
	}
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.UpdateStatus(pdaStatus)

	return pdaProcessor.Current_state(), err
}

func (pdaManager *PDAManager) Queued_token(id string, pdaStatus entity.PDAStatus) ([]string, error) {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return nil, err
	}
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.UpdateStatus(pdaStatus)

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

func (pdaManager *PDAManager) PutsEOS(id string, position int, pdaStatus entity.PDAStatus) error {
	get, err := pdaManager.PdaStore.Get(id)

	if err != nil {
		return err
	}
	pdaProcessor := parsePdaProcessor(get)
	pdaProcessor.UpdateStatus(pdaStatus)
	fmt.Printf(" Name:%s  Token:%s Position:%d  \n", pdaProcessor.PdaConf.Name, pdaProcessor.PdaConf.Eos, position)
	pdaProcessor.Puts(position, " ")
	pdaManager.PdaStore.Update(id, pdaProcessor)
	return nil
}

func (pdaManager *PDAManager) CreateNewReplicaGroup(gid int, conf string) error {
	var Replica entity.ReplicaConf
	err := json.Unmarshal([]byte(conf), &Replica)
	if err != nil {
		return err
	}
	Replica.Gid = gid
	//fmt.Printf("%+v\n", Replica.Group_members)

	pdaProcessor := core.PdaProcessor{}
	marshal, _ := json.Marshal(Replica.Pda_code)
	if pdaProcessor.Open(marshal) {
		pdaManager.PdaStore.SaveReplica(Replica.Gid, pdaProcessor, Replica.Group_members)
	}
	return nil
}

func (pdaManager *PDAManager) GetAllReplicaIds() []int {
	return pdaManager.PdaStore.GetAllReplicaIds()
}

func (pdaManager *PDAManager) ResetReplicaMembers(gid int) error {
	members := pdaManager.PdaStore.GetAllMembers(gid)
	//var status entity.PDAStatus

	for i := range members {
		pdaStr, _ := pdaManager.PdaStore.Get(members[i])
		processor := parsePdaProcessor(pdaStr)
		processor.Reset()
		//status.Clock = processor.Clock
		//status.InputQueue = processor.InputQueue
		//status.LastConsumedIndex = processor.LastConsumedIndex
		//status.State = processor.State
		//status.Stack = processor.Stack
		pdaManager.PdaStore.Update(members[i], processor)
	}
	//fmt.Printf(" Name:%s  Token:%s Position: N/A  \n", pdaProcessor.PdaConf.Name, "START")
	return nil
}

func (pdaManager *PDAManager) GetMemberAddress(id int) []string {
	return pdaManager.PdaStore.GetAllMembers(id)
}

func (pdaManager *PDAManager) GetRandomMemberAddress(id int) string {
	members := pdaManager.PdaStore.GetAllMembers(id)
	i := len(members)
	if i < 1 {
		return ""
	}
	return members[rand.Intn(i)]
}

func (pdaManager *PDAManager) GetCookieFor(gid int, memberId string) entity.PDAStatus {
	pda := pdaManager.PdaStore.GetPDA(gid, memberId)
	return entity.PDAStatus{
		Stack:             pda.Stack,
		State:             pda.State,
		Clock:             pda.Clock,
		InputQueue:        pda.InputQueue,
		LastConsumedIndex: pda.LastConsumedIndex,
		PdaId:             memberId,
		ReplicaId:         gid,
	}
}

func (pdaManager *PDAManager) CloseReplicaGrpAndMembers(gid int) {
	all_members := pdaManager.PdaStore.GetAllMembers(gid)

	for _, pdaid := range all_members {
		pda := pdaManager.PdaStore.GetPDA(gid, pdaid)
		pda.Close()
	}
}

func (pdaManager *PDAManager) JoinAReplicaGrp(pdaId string, replicaId int) {
	processor := pdaManager.PdaStore.GetReplicaConf(replicaId)
	pdaManager.PdaStore.JoinReplicaGroup(replicaId, pdaId, processor)
	//CALl update
}

//func (pdaManager *PDAManager) ListAllPDAs() []string {
//	return pdaManager.PdaStore.GetAllPDANames()
//}

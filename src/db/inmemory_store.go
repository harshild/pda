package db

import (
	"core"
	"encoding/json"
)

type InMemoryStore struct {
	PdaProcessors   map[string]core.PdaProcessor
	Replica_members map[string][]string
}

func (inMemoryStore *InMemoryStore) InitStore() {
	inMemoryStore.PdaProcessors = make(map[string]core.PdaProcessor, 0)
}

func (inMemoryStore *InMemoryStore) Save(pdaId string, processor core.PdaProcessor) {
	inMemoryStore.PdaProcessors[pdaId] = processor
}

func (inMemoryStore *InMemoryStore) Update(pdaId string, processor core.PdaProcessor) {
	inMemoryStore.PdaProcessors[pdaId] = processor
}

func (inMemoryStore *InMemoryStore) Get(pdaId string) (string, error) {
	if inMemoryStore.idExists(pdaId) {
		jsonVal, _ := json.Marshal(inMemoryStore.PdaProcessors[pdaId])
		return string(jsonVal), nil
	}
	return "", &core.PDARuntimeError{Message: "No PDA found with id " + pdaId}
}

func (inMemoryStore *InMemoryStore) idExists(pdaId string) bool {
	_, ok := inMemoryStore.PdaProcessors[pdaId]
	return ok
}

func (inMemoryStore *InMemoryStore) GetAllPDA() []string {
	pdaStr := make([]string, 0)
	for _, value := range inMemoryStore.PdaProcessors {
		jsonVal, _ := json.Marshal(value)
		pdaStr = append(pdaStr, string(jsonVal))
	}

	return pdaStr
}

func (inMemoryStore *InMemoryStore) Delete(pdaId string) {
	delete(inMemoryStore.PdaProcessors, pdaId)
}

func (inMemoryStore *InMemoryStore) SaveReplica(gid string, processor core.PdaProcessor, add []string) {
	inMemoryStore.Replica_members[gid] = add
	inMemoryStore.PdaProcessors[gid] = processor
	print("replica saved!!!!!!!!")
}

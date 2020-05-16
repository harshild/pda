package db

import (
	"core"
	"encoding/json"
)

type ReplicaInMemoryStore struct {
	PdaProcessors  map[string]core.PdaProcessor
	ReplicaMembers map[string][]string
}

func (replicaInMemoryStore *ReplicaInMemoryStore) InitStore() {
	replicaInMemoryStore.PdaProcessors = make(map[string]core.PdaProcessor, 0)
	replicaInMemoryStore.ReplicaMembers = make(map[string][]string, 0)
}

func (replicaInMemoryStore *ReplicaInMemoryStore) Save(pdaId string, processor core.PdaProcessor) {
	replicaInMemoryStore.PdaProcessors[pdaId] = processor
}

func (replicaInMemoryStore *ReplicaInMemoryStore) Update(pdaId string, processor core.PdaProcessor) {
	replicaInMemoryStore.PdaProcessors[pdaId] = processor
}

func (replicaInMemoryStore *ReplicaInMemoryStore) Get(pdaId string) (string, error) {
	if replicaInMemoryStore.idExists(pdaId) {
		jsonVal, _ := json.Marshal(replicaInMemoryStore.PdaProcessors[pdaId])
		return string(jsonVal), nil
	}
	return "", &core.PDARuntimeError{Message: "No PDA found with id " + pdaId}
}

func (replicaInMemoryStore *ReplicaInMemoryStore) idExists(pdaId string) bool {
	_, ok := replicaInMemoryStore.PdaProcessors[pdaId]
	return ok
}

func (replicaInMemoryStore *ReplicaInMemoryStore) GetAllPDA() []string {
	pdaStr := make([]string, 0)
	for _, value := range replicaInMemoryStore.PdaProcessors {
		jsonVal, _ := json.Marshal(value)
		pdaStr = append(pdaStr, string(jsonVal))
	}

	return pdaStr
}

func (replicaInMemoryStore *ReplicaInMemoryStore) Delete(pdaId string) {
	delete(replicaInMemoryStore.PdaProcessors, pdaId)
}

func (replicaInMemoryStore *ReplicaInMemoryStore) SaveReplica(gid string, processor core.PdaProcessor, group_members []string) {
	replicaInMemoryStore.ReplicaMembers[gid] = group_members
	replicaInMemoryStore.PdaProcessors[gid] = processor
}

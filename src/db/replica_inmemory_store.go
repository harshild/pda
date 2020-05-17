package db

import (
	"core"
	"encoding/json"
)

type ReplicaInMemoryStore struct {
	PdaProcessors  map[string]core.PdaProcessor
	ReplicaMembers map[int]map[string]core.PdaProcessor
}

func (replicaInMemoryStore *ReplicaInMemoryStore) InitStore() {
	replicaInMemoryStore.PdaProcessors = make(map[string]core.PdaProcessor, 0)
	replicaInMemoryStore.ReplicaMembers = make(map[int]map[string]core.PdaProcessor, 0)
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

func (replicaInMemoryStore *ReplicaInMemoryStore) SaveReplica(gid int, processor core.PdaProcessor, group_members []string) {
	//replicaInMemoryStore.ReplicaMembers[gid] = group_members
	// TODO correct usage of id for storing pda processor
	//gidStr := strconv.Itoa(gid)
	members := make(map[string]core.PdaProcessor, 0)

	for pdaid := range group_members {
		members[group_members[pdaid]] = processor
	}

	replicaInMemoryStore.ReplicaMembers[gid] = members
}

func (replicaInMemoryStore *ReplicaInMemoryStore) GetAllReplicaIds() []int {
	var keys []int

	if replicaInMemoryStore.ReplicaMembers != nil {
		for key, _ := range replicaInMemoryStore.ReplicaMembers {
			keys = append(keys, key)
		}
	}
	return keys
}

func (replicaInMemoryStore *ReplicaInMemoryStore) GetAllMembers(id int) []string {
	abc := make([]string, 0)

	for k, _ := range replicaInMemoryStore.ReplicaMembers[id] {
		abc = append(abc, k)
	}

	return abc
}

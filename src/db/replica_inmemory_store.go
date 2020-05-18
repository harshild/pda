package db

import (
	"core"
	"encoding/json"
)

type ReplicaInMemoryStore struct {
	PdaProcessors  map[string]core.PdaProcessor
	ReplicaMembers map[int][]string
}

func (replicaInMemoryStore *ReplicaInMemoryStore) InitStore() {
	replicaInMemoryStore.PdaProcessors = make(map[string]core.PdaProcessor, 0)
	replicaInMemoryStore.ReplicaMembers = make(map[int][]string, 0)
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

func (replicaInMemoryStore *ReplicaInMemoryStore) GetAllPDANames() []string {
	pdaStr := make([]string, 0)
	for _, value := range replicaInMemoryStore.PdaProcessors {
		pdaStr = append(pdaStr, value.GetPDAName())
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

	for index := range group_members {
		replicaInMemoryStore.PdaProcessors[group_members[index]] = processor
	}

	replicaInMemoryStore.ReplicaMembers[gid] = group_members
}

func (replicaInMemoryStore *ReplicaInMemoryStore) JoinReplicaGroup(gid int, pdaid string, processor core.PdaProcessor) {
	replicaInMemoryStore.ReplicaMembers[gid] = append(replicaInMemoryStore.ReplicaMembers[gid], pdaid)
	replicaInMemoryStore.PdaProcessors[pdaid] = processor
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
	return replicaInMemoryStore.ReplicaMembers[id]
}

func (replicaInMemoryStore *ReplicaInMemoryStore) GetPDA(gid int, pdaId string) core.PdaProcessor {
	if contains(replicaInMemoryStore.ReplicaMembers[gid], pdaId) {
		return replicaInMemoryStore.PdaProcessors[pdaId]
	}
	return core.PdaProcessor{}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (replicaInMemoryStore *ReplicaInMemoryStore) DeleteReplicaGrp() {

}

func (replicaInMemoryStore *ReplicaInMemoryStore) GetReplicaConf(gid int) core.PdaProcessor {
	return replicaInMemoryStore.PdaProcessors[replicaInMemoryStore.ReplicaMembers[gid][0]]
}

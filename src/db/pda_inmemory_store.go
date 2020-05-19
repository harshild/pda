package db

import (
	"core"
	"encoding/json"
	"entity"
)

type Replica struct {
	GroupMembers []string
	PdaConf      entity.PDAConf
}

type InMemoryStore struct {
	PdaProcessors  map[string]core.PdaProcessor
	ReplicaMembers map[int]Replica
}

func (inMemoryStore *InMemoryStore) InitStore() {
	inMemoryStore.PdaProcessors = make(map[string]core.PdaProcessor, 0)
	inMemoryStore.ReplicaMembers = make(map[int]Replica, 0)
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

func (inMemoryStore *InMemoryStore) SaveReplica(gid int, processor core.PdaProcessor, group_members []string) {
	for index := range group_members {
		inMemoryStore.PdaProcessors[group_members[index]] = processor
	}

	inMemoryStore.ReplicaMembers[gid] = Replica{
		GroupMembers: group_members,
		PdaConf:      processor.PdaConf,
	}
}

func (inMemoryStore *InMemoryStore) JoinReplicaGroup(gid int, pdaid string, pdaConf entity.PDAConf) {
	members := append(inMemoryStore.ReplicaMembers[gid].GroupMembers, pdaid)
	inMemoryStore.ReplicaMembers[gid] = Replica{
		GroupMembers: members,
		PdaConf:      pdaConf,
	}

}

func (inMemoryStore *InMemoryStore) GetAllReplicaIds() []int {
	var keys []int

	if inMemoryStore.ReplicaMembers != nil {
		for key, _ := range inMemoryStore.ReplicaMembers {
			keys = append(keys, key)
		}
	}
	return keys
}

func (inMemoryStore *InMemoryStore) GetAllMembers(id int) []string {
	return inMemoryStore.ReplicaMembers[id].GroupMembers
}

func (inMemoryStore *InMemoryStore) GetPDA(gid int, pdaId string) core.PdaProcessor {
	if contains(inMemoryStore.ReplicaMembers[gid].GroupMembers, pdaId) {
		return inMemoryStore.PdaProcessors[pdaId]
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

func (inMemoryStore *InMemoryStore) DeleteReplicaGrp(gid int) {
	delete(inMemoryStore.ReplicaMembers, gid)

}

func (inMemoryStore *InMemoryStore) GetReplicaConf(gid int) entity.PDAConf {
	return inMemoryStore.ReplicaMembers[gid].PdaConf
}

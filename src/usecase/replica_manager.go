package usecase

import (
	"core"
	"db"
	"encoding/json"
	"entity"

	"utility"
)

type ReplicaManager struct {
	ReplicaStore db.InMemoryStore
}

func (replicamanager *ReplicaManager) CreateNewReplicagrp(gid string, conf string) error {
	var Replica entity.ReplicaConf
	err := json.Unmarshal([]byte(conf), &Replica)
	if err != nil {
		print("Jason unmarshalling of replica failed!")
	}
	Replica.Gid = gid
	//fmt.Printf("%+v\n", Replica.Group_members)

	pdaProcessor := core.PdaProcessor{}
	marshal, err := json.Marshal(Replica.Pda_code)
	if pdaProcessor.Open(marshal) {
		replicamanager.ReplicaStore.SaveReplica(Replica.Gid, pdaProcessor, Replica.Group_members)
	}
	return nil
}

type Cookie struct {
	Stack             utility.Stack
	State             string
	Clock             int
	InputQueue        map[int]string
	LastConsumedIndex int
}

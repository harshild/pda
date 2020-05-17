package usecase

import (
	"core"
	"db"
	"encoding/json"
	"entity"
)

type ReplicaManager struct {
	ReplicaStore db.ReplicaInMemoryStore
}

func (replicamanager *ReplicaManager) CreateNewReplicaGroup(gid int, conf string) error {
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

func (replicamanager *ReplicaManager) GetAllReplicaIds() []int {
	return replicamanager.ReplicaStore.GetAllReplicaIds()
}

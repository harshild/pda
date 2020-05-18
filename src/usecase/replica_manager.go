package usecase

import (
	"core"
	"db"
	"encoding/json"
	"entity"
	"math/rand"
)

type ReplicaManager struct {
	ReplicaStore db.ReplicaInMemoryStore
}

func (replicamanager *ReplicaManager) CreateNewReplicaGroup(gid int, conf string) error {
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
		replicamanager.ReplicaStore.SaveReplica(Replica.Gid, pdaProcessor, Replica.Group_members)
	}
	return nil
}

func (replicamanager *ReplicaManager) GetAllReplicaIds() []int {
	return replicamanager.ReplicaStore.GetAllReplicaIds()
}

func (replicamanager *ReplicaManager) ResetReplicaMembers(gid string) error {
	get, err := replicamanager.ReplicaStore.Get(gid)

	if err != nil {
		return err
	}
	pdaProcessor := parsePdaProcessor(get)
	//fmt.Printf(" Name:%s  Token:%s Position: N/A  \n", pdaProcessor.PdaConf.Name, "START")
	pdaProcessor.Reset()
	replicamanager.ReplicaStore.Update(gid, pdaProcessor)
	return err
}

func (replicamanager *ReplicaManager) GetMemberAddress(id int) []string {
	return replicamanager.ReplicaStore.GetAllMembers(id)
}

func (replicamanager *ReplicaManager) GetRandomMemberAddress(id int) string {
	members := replicamanager.ReplicaStore.GetAllMembers(id)
	i := len(members)
	if i < 1 {
		return ""
	}
	return members[rand.Intn(i)]
}

func (replicamanager *ReplicaManager) GetCookieFor(gid int, memberId string) entity.PDAStatus {
	pda := replicamanager.ReplicaStore.GetPDA(gid, memberId)
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

func (replicamanager *ReplicaManager) CloseReplicaGrpAndMembers(gid int) {
	all_members := replicamanager.ReplicaStore.GetAllMembers(gid)

	for _, pdaid := range all_members {
		pda := replicamanager.ReplicaStore.GetPDA(gid, pdaid)
		pda.Close()
		replicamanager.ReplicaStore.SavePDA(gid, pdaid, pda)
	}
}

func (replicamanager *ReplicaManager) JoinAReplicaGrp(pdaId string, replicaId int) {
	processor := replicamanager.ReplicaStore.GetReplicaConf(replicaId)
	replicamanager.ReplicaStore.SavePDA(replicaId, pdaId, processor)
}

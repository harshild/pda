package usecase

import (
	"core"
	"db"
)

type PDAManager struct {
}

func (pdaManager *PDAManager) NewPDA(id int, json string) {
	pdaProcessor := core.PdaProcessor{}
	pdaStore := db.PDAStore{}
	if pdaProcessor.Open([]byte(json)) {
		pdaStore.Save(id, pdaProcessor.PdaConf)
	}
}

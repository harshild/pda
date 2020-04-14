package db

import (
	"core"
	_ "github.com/mattn/go-sqlite3"
)

type PDAStore interface {
	InitStore()
	Save(pdaId string, processor core.PdaProcessor)
	Update(pdaId string, processor core.PdaProcessor)
	Get(pdaId string) (string, error)
	idExists(pdaId string) bool
	GetAllPDA() []string
}

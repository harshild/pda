package db

import (
	"core"
)

type FileStore struct {
	FilePath string
}

func (fileStore FileStore) InitStore() {

}

func (fileStore FileStore) Save(pdaId string, processor core.PdaProcessor) {
}

func (fileStore FileStore) Update(pdaId string, processor core.PdaProcessor) {
}

func (fileStore FileStore) Get(pdaId string) (string, error) {
	return "", nil
}

func (fileStore FileStore) idExists(pdaId string) bool {
	return false
}

func (fileStore FileStore) GetAllPDA() []string {
	return nil
}

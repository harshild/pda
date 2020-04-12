package db

import (
	"context"
	"core"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type PDAStore struct {
	dbType string
	dsn    string
	db     *sql.DB
	ctx    context.Context
}

func (pdaStore *PDAStore) initDB() {
	pdaStore.db, _ = sql.Open(pdaStore.dbType, pdaStore.dsn)
	statement, _ := pdaStore.db.Prepare("CREATE TABLE IF NOT EXISTS PDA (id TEXT PRIMARY KEY, pda TEXT)")
	statement.Exec()

}

func (pdaStore *PDAStore) Save(pdaId string, processor core.PdaProcessor) {
	pdaStore.db, _ = sql.Open(pdaStore.dbType, pdaStore.dsn)
	_, err := pdaStore.Get(pdaId)
	if err != nil {
		statement, _ := pdaStore.db.Prepare("update PDA pda=? where id=?")
		statement.Exec("some json file content", pdaId)
	} else {
		statement, _ := pdaStore.db.Prepare("insert into PDA(id,pda) values(?,?)")
		statement.Exec(pdaId, "some json file content")
	}
}

func (pdaStore *PDAStore) Get(pdaId string) (string, error) {
	pdaStore.db, _ = sql.Open(pdaStore.dbType, pdaStore.dsn)
	rows, err := pdaStore.db.QueryContext(pdaStore.ctx, "SELECT pda FROM PDA WHERE id=?", pdaId)
	var jsonconfig string

	if err != nil {
		for rows.Next() {
			if err := rows.Scan(&jsonconfig); err != nil {
				log.Fatal(err)
			}
		}
	}

	return jsonconfig, err
}

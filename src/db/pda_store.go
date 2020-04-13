package db

import (
	"context"
	"core"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type PDAStore struct {
	DbType string
	Dsn    string
	Db     *sql.DB
	Ctx    context.Context
}

func (pdaStore *PDAStore) InitDB() {
	pdaStore.Db, _ = sql.Open(pdaStore.DbType, pdaStore.Dsn)
	statement, _ := pdaStore.Db.Prepare("CREATE TABLE IF NOT EXISTS PDA (id TEXT PRIMARY KEY, pda TEXT)")
	statement.Exec()

}

func (pdaStore *PDAStore) Save(pdaId string, processor core.PdaProcessor) {
	pdaStore.Db, _ = sql.Open(pdaStore.DbType, pdaStore.Dsn)
	_, err := pdaStore.Get(pdaId)
	if err != nil {
		statement, _ := pdaStore.Db.Prepare("update PDA pda=? where id=?")
		statement.Exec("some json file content", pdaId)
	} else {
		statement, _ := pdaStore.Db.Prepare("insert into PDA(id,pda) values(?,?)")
		statement.Exec(pdaId, "some json file content")
	}
}

func (pdaStore *PDAStore) Get(pdaId string) (string, error) {
	pdaStore.Db, _ = sql.Open(pdaStore.DbType, pdaStore.Dsn)
	rows, err := pdaStore.Db.QueryContext(pdaStore.Ctx, "SELECT pda FROM PDA WHERE id=?", pdaId)
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

func (pdaStore *PDAStore) GetAllPDA() []string {
	pdaStore.Db, _ = sql.Open(pdaStore.DbType, pdaStore.Dsn)
	rows, err := pdaStore.Db.QueryContext(pdaStore.Ctx, "SELECT pda FROM PDA")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	pdas := make([]string, 0)

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		pdas = append(pdas, name)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return pdas
}

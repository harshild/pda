package db

import (
	"context"
	"core"
	"database/sql"
	"encoding/json"
	"log"
)

type SqliteStore struct {
	DbType string
	Dsn    string
	Db     *sql.DB
	Ctx    context.Context
}

func (sqliteStore SqliteStore) InitStore() {
	sqliteStore.Db, _ = sql.Open(sqliteStore.DbType, sqliteStore.Dsn)
	statement, _ := sqliteStore.Db.Prepare("CREATE TABLE IF NOT EXISTS PDA (id TEXT PRIMARY KEY, pda TEXT)")
	statement.Exec()

}

func (sqliteStore SqliteStore) Save(pdaId string, processor core.PdaProcessor) {
	sqliteStore.Db, _ = sql.Open(sqliteStore.DbType, sqliteStore.Dsn)
	statement, _ := sqliteStore.Db.Prepare("insert into PDA(id,pda) values(?,?)")
	marshal, _ := json.Marshal(processor)
	statement.Exec(pdaId, marshal)
}

func (sqliteStore SqliteStore) Update(pdaId string, processor core.PdaProcessor) {
	sqliteStore.Db, _ = sql.Open(sqliteStore.DbType, sqliteStore.Dsn)
	statement, _ := sqliteStore.Db.Prepare("update PDA pda=? where id=?")
	marshal, _ := json.Marshal(processor)
	statement.Exec(pdaId, marshal)
}

func (sqliteStore SqliteStore) Get(pdaId string) (string, error) {
	sqliteStore.Db, _ = sql.Open(sqliteStore.DbType, sqliteStore.Dsn)
	rows, err := sqliteStore.Db.QueryContext(sqliteStore.Ctx, "SELECT pda FROM PDA WHERE id=?", pdaId)
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

func (sqliteStore SqliteStore) idExists(pdaId string) bool {
	sqliteStore.Db, _ = sql.Open(sqliteStore.DbType, sqliteStore.Dsn)
	rows, _ := sqliteStore.Db.QueryContext(sqliteStore.Ctx, "SELECT count(*) FROM PDA WHERE id=?", pdaId)
	var count int
	for rows.Next() {
		if rows.Scan(&count); count > 0 {
			return true
		}
	}

	return false
}

func (sqliteStore SqliteStore) GetAllPDA() []string {
	sqliteStore.Db, _ = sql.Open(sqliteStore.DbType, sqliteStore.Dsn)
	rows, err := sqliteStore.Db.Query("SELECT pda FROM PDA")

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

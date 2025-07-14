package utils

import (
  "database/sql"
  "os"
  "github.com/charmbracelet/log"
  _ "github.com/tursodatabase/go-libsql"
)

func IntializeDB() {
  db := TDB()
  _, err := db.Exec("CREATE TABLE IF NOT EXISTS example (example_id TEXT PRIMARY KEY, example_name TEXT, example_value TEXT)")
  if err != nil {
    log.Fatal(err)
  }
  _,err = db.Exec("CREATE TABLE IF NOT EXISTS temp (temp_id TEXT PRIMARY KEY, temp_name TEXT, temp_value TEXT)")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
}

func TDB() *sql.DB {
  db, err := sql.Open("libsql", "file:./tdata.db")
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  return db
}
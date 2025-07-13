package utils

import (
  "database/sql"
  "fmt"
  "os"

  _ "github.com/tursodatabase/go-libsql"
)

func IntializeDB() {
  db := TDB()
  db.Exec("CREATE TABLE IF NOT EXISTS example (example_id TEXT PRIMARY KEY, example_name TEXT, example_value TEXT)")
  db.Exec("CREATE TABLE IF NOT EXISTS temp (temp_id TEXT PRIMARY KEY, temp_name TEXT, temp_value TEXT)")
  defer db.Close()
}

func TDB() *sql.DB {
  db, err := sql.Open("libsql", "file:./tdata.db")
  if err != nil {
    fmt.Fprintf(os.Stderr, "failed to open db %s", err)
    os.Exit(1)
  }
  return db
}
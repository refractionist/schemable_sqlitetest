package schemable_sqlitetest

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/refractionist/schemable"
	"github.com/refractionist/schemable/schemabletest"
)

func TestSqlite(t *testing.T) {
	c, err := schemable.New("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Fatal(err)
	}

	c.SetLogger(schemabletest.QueryLogger(t))
	ctx := schemable.WithClient(context.Background(), c)
	_, err = c.Exec(ctx, `
		CREATE TABLE test_structs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			id_two INTEGER,
			name STRING,
			num INTEGER
		)
`)
	if err != nil {
		t.Fatal(err)
	}

	schemabletest.ClientTests(t, c)
	schemabletest.RecorderTests(t, ctx)
	schemabletest.SchemerTests(t, ctx)
}

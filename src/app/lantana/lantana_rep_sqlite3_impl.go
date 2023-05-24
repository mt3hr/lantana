package lantana

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"strings"
	"sync"

	"github.com/mt3hr/rykv/kyou"
)

var (
	//go:embed lantana/embed
	EmbedDir                    embed.FS
	sqlAddLantana               string
	sqlCreateTables             string
	sqlDeleteLantana            string
	sqlGetAllLantanas           string
	sqlGetLantana               string
	sqlSearchLantanaAll         string
	sqlSearchLantanaGreaterThan string
	sqlSearchLantanaLessThan    string
	sqlSearchLantanaMatch       string
)

func init() {
	sqlAddLantanab, err := EmbedDir.ReadFile("lantana/embed/sql/AddLantana.sql")
	if err != nil {
		panic(err)
	}
	sqlAddLantana = string(sqlAddLantanab)

	sqlCreateTablesB, err := EmbedDir.ReadFile("lantana/embed/sql/CreateTables.sql")
	if err != nil {
		panic(err)
	}
	sqlCreateTables = string(sqlCreateTablesB)

	sqlDeleteLantanab, err := EmbedDir.ReadFile("lantana/embed/sql/DeleteLantana.sql")
	if err != nil {
		panic(err)
	}
	sqlDeleteLantana = string(sqlDeleteLantanab)

	sqlGetAllLantanasb, err := EmbedDir.ReadFile("lantana/embed/sql/GetAllLantanas.sql")
	if err != nil {
		panic(err)
	}
	sqlGetAllLantanas = string(sqlGetAllLantanasb)

	sqlGetLantanab, err := EmbedDir.ReadFile("lantana/embed/sql/GetLantana.sql")
	if err != nil {
		panic(err)
	}
	sqlGetLantana = string(sqlGetLantanab)

	sqlSearchLantanaAllb, err := EmbedDir.ReadFile("lantana/embed/sql/SearchLantanaAll.sql")
	if err != nil {
		panic(err)
	}
	sqlSearchLantanaAll = string(sqlSearchLantanaAllb)

	sqlSearchLantanaGreaterThanb, err := EmbedDir.ReadFile("lantana/embed/sql/SearchLantanaGreaterThan.sql")
	if err != nil {
		panic(err)
	}
	sqlSearchLantanaGreaterThan = string(sqlSearchLantanaGreaterThanb)

	sqlSearchLantanaLessThanb, err := EmbedDir.ReadFile("lantana/embed/sql/SearchLantanaLessThan.sql")
	if err != nil {
		panic(err)
	}
	sqlSearchLantanaLessThan = string(sqlSearchLantanaLessThanb)

	sqlSearchLantanaMatchb, err := EmbedDir.ReadFile("lantana/embed/sql/SearchLantanaMatch.sql")
	if err != nil {
		panic(err)
	}
	sqlSearchLantanaMatch = string(sqlSearchLantanaMatchb)
}

func NewLantanaRepSQLite(dbFileName string) (LantanaRep, error) {
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		err = fmt.Errorf("error at open database %s: %w", dbFileName, err)
		return nil, err
	}
	_, err = db.Exec(sqlCreateTables)
	if err != nil {
		err = fmt.Errorf("error at create table to database at %s: %w", dbFileName, err)
		return nil, err
	}
	return &lantanaRepSQLite3Impl{
		filename: dbFileName,
		db:       db,
		m:        &sync.Mutex{},
	}, nil
}

type lantanaRepSQLite3Impl struct {
	filename string
	db       *sql.DB
	m        *sync.Mutex
}

func (l *lantanaRepSQLite3Impl) GetAllLantanas(ctx context.Context) ([]*Lantana, error) {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) GetLantana(ctx context.Context, lantanaID string) (*Lantana, error) {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) AddLantana(ctx context.Context, lantana *Lantana) error {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) SearchLantana(ctx context.Context, query *LantanaSearchQuery) ([]*Lantana, error) {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) GetAllKyous(ctx context.Context) ([]*kyou.Kyou, error) {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) GetContentHTML(ctx context.Context, id string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) GetPath(ctx context.Context, id string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) Delete(id string) error {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) Close() error {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) Path() string {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) RepName() string {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) Search(ctx context.Context, word string) ([]*kyou.Kyou, error) {
	panic("not implemented") // TODO: Implement
}

func (l *lantanaRepSQLite3Impl) UpdateCache(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func escapeSQLite(str string) string {
	return strings.ReplaceAll(str, "'", "''")
}

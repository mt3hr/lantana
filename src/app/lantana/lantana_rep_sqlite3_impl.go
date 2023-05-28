package lantana

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mt3hr/rykv/kyou"
)

const TimeLayout = kyou.TimeLayout

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
	lantanas := []*Lantana{}
	statement := sqlGetAllLantanas
	rows, err := l.db.QueryContext(ctx, statement)
	if err != nil {
		err = fmt.Errorf("error at get all lantanas: %w", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			lantana := &Lantana{}
			createdTimeStr := ""
			err := rows.Scan(&lantana.LantanaID, &createdTimeStr, &lantana.Mood)
			if err != nil {
				return nil, err
			}

			lantana.Time, err = time.Parse(TimeLayout, createdTimeStr)
			if err != nil {
				err = fmt.Errorf("error at parse time: %w", err)
				return nil, err
			}
			lantanas = append(lantanas, lantana)
		}
	}
	return lantanas, nil
}

func (l *lantanaRepSQLite3Impl) GetLantana(ctx context.Context, lantanaID string) (*Lantana, error) {
	statement := sqlGetLantana
	row := l.db.QueryRowContext(ctx, statement, lantanaID)

	lantana := &Lantana{}
	createdTimeStr := ""
	err := row.Scan(&lantana.LantanaID, &createdTimeStr, &lantana.Mood)
	if err != nil {
		return nil, err
	}

	lantana.Time, err = time.Parse(TimeLayout, createdTimeStr)
	if err != nil {
		err = fmt.Errorf("error at parse time: %w", err)
		return nil, err
	}
	return lantana, nil
}

func (l *lantanaRepSQLite3Impl) AddLantana(ctx context.Context, lantana *Lantana) error {
	l.m.Lock()
	defer l.m.Unlock()
	statement := sqlAddLantana
	_, err := l.db.Exec(statement, lantana.LantanaID, lantana.Time.Format(TimeLayout), lantana.Mood)
	if err != nil {
		err = fmt.Errorf("error at add lantana to to database %s: %w", l.filename, err)
		return err
	}
	return nil
}

func (l *lantanaRepSQLite3Impl) SearchLantana(ctx context.Context, query *LantanaSearchQuery) ([]*Lantana, error) {
	lantanas := []*Lantana{}
	var q LantanaSearchQuery
	q = *query

	statement := ""
	switch *query.LantanaSearchType {
	case All:
		statement = sqlSearchLantanaGreaterThan
		q.Mood = 0
	case Match:
		statement = sqlSearchLantanaMatch
	case GreaterThan:
		statement = sqlSearchLantanaGreaterThan
	case LessThan:
		statement = sqlSearchLantanaLessThan
	}

	rows, err := l.db.QueryContext(ctx, statement, q.Mood, q.Words)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			lantana := &Lantana{}
			createdTimeStr := ""
			err := rows.Scan(&lantana.LantanaID, &createdTimeStr, &lantana.Mood)
			if err != nil {
				return nil, err
			}

			lantana.Time, err = time.Parse(TimeLayout, createdTimeStr)
			if err != nil {
				err = fmt.Errorf("error at parse time: %w", err)
				return nil, err
			}
			lantanas = append(lantanas, lantana)
		}
	}

	return lantanas, nil
}

func (l *lantanaRepSQLite3Impl) GetAllKyous(ctx context.Context) ([]*kyou.Kyou, error) {
	kyous := []*kyou.Kyou{}

	lantanas, err := l.GetAllLantanas(ctx)
	if err != nil {
		return nil, err
	}
	for _, lantana := range lantanas {
		kyous = append(kyous, &kyou.Kyou{
			ID:          lantana.LantanaID,
			Time:        lantana.Time,
			RepName:     l.RepName(),
			ImageSource: "",
		})
	}
	return kyous, nil
}

func (l *lantanaRepSQLite3Impl) GetContentHTML(ctx context.Context, id string) (string, error) {
	//TODO 画像にして
	lantana, err := l.GetLantana(ctx, id)
	if err != nil {
		return "", err
	}
	contentHTML := fmt.Sprintf("<p>気分評価:%d</p>", lantana.Mood)
	return contentHTML, nil
}

func (l *lantanaRepSQLite3Impl) GetPath(ctx context.Context, id string) (string, error) {
	return l.filename, nil
}

func (l *lantanaRepSQLite3Impl) Delete(id string) error {
	l.m.Lock()
	defer l.m.Unlock()
	statement := sqlDeleteLantana
	_, err := l.db.Exec(statement, id)
	if err != nil {
		err = fmt.Errorf("error at delete lantana from database %s: %w", l.filename, err)
		return err
	}
	return nil
}

func (l *lantanaRepSQLite3Impl) Close() error {
	return l.db.Close()
}

func (l *lantanaRepSQLite3Impl) Path() string {
	return l.filename
}

func (l *lantanaRepSQLite3Impl) RepName() string {
	base := filepath.Base(l.Path())
	ext := filepath.Ext(base)
	withoutExt := base[:len(base)-len(ext)]
	return withoutExt
}

func (l *lantanaRepSQLite3Impl) Search(ctx context.Context, word string) ([]*kyou.Kyou, error) {
	kyous := []*kyou.Kyou{}

	query := &LantanaSearchQuery{}
	lantanaSearchType := All
	query.LantanaSearchType = &lantanaSearchType

	*query.LantanaSearchType = All
	if strings.HasPrefix(word, "lantana=") {
		*query.LantanaSearchType = Match
		mood, err := strconv.ParseInt(strings.TrimPrefix(word, "lantana="), 10, 64)
		if err != nil {
			return nil, err
		}
		query.Mood = int(mood)
	} else if strings.HasPrefix(word, "lantana<=") {
		*query.LantanaSearchType = LessThan
		mood, err := strconv.ParseInt(strings.TrimPrefix(word, "lantana<="), 10, 64)
		if err != nil {
			return nil, err
		}
		query.Mood = int(mood)
	} else if strings.HasPrefix(word, "lantana>=") {
		*query.LantanaSearchType = GreaterThan
		mood, err := strconv.ParseInt(strings.TrimPrefix(word, "lantana>="), 10, 64)
		if err != nil {
			return nil, err
		}
		query.Mood = int(mood)
	}

	lantanas, err := l.SearchLantana(ctx, query)
	if err != nil {
		return nil, err
	}
	for _, lantana := range lantanas {
		kyous = append(kyous, &kyou.Kyou{
			ID:          lantana.LantanaID,
			Time:        lantana.Time,
			RepName:     l.RepName(),
			ImageSource: "",
		})
	}
	return kyous, nil
}

func (l *lantanaRepSQLite3Impl) UpdateCache(ctx context.Context) error {
	return nil
}

func escapeSQLite(str string) string {
	return strings.ReplaceAll(str, "'", "''")
}

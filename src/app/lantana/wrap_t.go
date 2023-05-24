package lantana

import (
	"context"

	"github.com/mt3hr/rykv/kyou"
	"github.com/mt3hr/rykv/tag"
)

func WrapLantanaRepT(rep LantanaRep, deleteTagReps tag.DeleteTagReps) LantanaRep {
	return &lantanaRepT{
		lantanaRep:    rep,
		deleteTagReps: deleteTagReps,
	}
}

type lantanaRepT struct {
	lantanaRep    LantanaRep
	deleteTagReps tag.DeleteTagReps
}

func (l *lantanaRepT) GetAllLantanas(ctx context.Context) ([]*Lantana, error) {
	return l.lantanaRep.GetAllLantanas(ctx)
}

func (l *lantanaRepT) GetLantana(ctx context.Context, lantanaID string) (*Lantana, error) {
	return l.lantanaRep.GetLantana(ctx, lantanaID)
}

func (l *lantanaRepT) AddLantana(ctx context.Context, lantana *Lantana) error {
	return l.lantanaRep.AddLantana(ctx, lantana)
}

func (l *lantanaRepT) SearchLantana(ctx context.Context, query *LantanaSearchQuery) ([]*Lantana, error) {
	return l.lantanaRep.SearchLantana(ctx, query)
}

func (l *lantanaRepT) GetAllKyous(ctx context.Context) ([]*kyou.Kyou, error) {
	return l.lantanaRep.GetAllKyous(ctx)
}

func (l *lantanaRepT) GetContentHTML(ctx context.Context, id string) (string, error) {
	return l.lantanaRep.GetContentHTML(ctx, id)
}

func (l *lantanaRepT) GetPath(ctx context.Context, id string) (string, error) {
	return l.lantanaRep.GetPath(ctx, id)
}

func (l *lantanaRepT) Delete(id string) error {
	return l.deleteTagReps.Delete(id)
}

func (l *lantanaRepT) Close() error {
	return l.lantanaRep.Close()
}

func (l *lantanaRepT) Path() string {
	return l.lantanaRep.Path()
}

func (l *lantanaRepT) RepName() string {
	return l.lantanaRep.RepName()
}

func (l *lantanaRepT) Search(ctx context.Context, word string) ([]*kyou.Kyou, error) {
	return l.lantanaRep.Search(ctx, word)
}

func (l *lantanaRepT) UpdateCache(ctx context.Context) error {
	return l.lantanaRep.UpdateCache(ctx)
}

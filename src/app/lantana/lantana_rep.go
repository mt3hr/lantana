// ˅
package main

// ˄

type LantanaRep interface {
	GetAllLantanas() []*Lantana

	GetLantana(ctx context.Context, lantanaID string) *Lantana

	AddLantana(ctx context.Context, lantana Lantana)

	Delete(id string)

	SearchLantana(ctx context.Context, query *LantanaSearchQuery) []*Lantana

	GetAllKyous(ctx context.Context) []*kyou.Kyou

	GetContentHTML(ctx context.Context, id string) string

	GetAllKyous(ctx context.Context) []*kyou.Kyou

	GetContentHTML(ctx context.Context, id string) string

	GetPath(ctx context.Context, id string) string

	Delete(id string)

	Close()

	Path() string

	RepName() string

	Search(ctx context.Context, word string) []*kyou.Kyou

	UpdateCache(ctx context.Context)

	// ˅
	
	// ˄
}

// ˅

// ˄

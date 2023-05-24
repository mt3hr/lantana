package lantana

import (
	"sort"

	"github.com/mt3hr/lantana/src/app/lantana"
	"github.com/mt3hr/rykv/tag"
	"github.com/mt3hr/rykv/text"
)

func sortLantanasByTime(lantanas []*lantana.Lantana) {
	sort.Slice(lantanas, func(i, j int) bool {
		return lantanas[i].Time.After(lantanas[j].Time)
	})
}

func sortTagsByTime(tags []*tag.Tag) {
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Time.Before(tags[j].Time)
	})
}

func sortTextsByTime(texts []*text.Text) {
	sort.Slice(texts, func(i, j int) bool {
		return texts[i].Time.Before(texts[j].Time)
	})
}

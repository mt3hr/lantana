// ˅
package lantana

import (
	"github.com/mt3hr/kmemo"
	"github.com/mt3hr/lantana/src/app/lantana"
	"github.com/mt3hr/rykv/tag"
	"github.com/mt3hr/rykv/text"
)

// ˄

type Repositories struct {
	// ˅

	// ˄

	LantanaRep lantana.LantanaRep

	KmemoRep kmemo.KmemoRep

	TagRep tag.TagRep

	TextRep text.TextRep

	LantanaReps []lantana.LantanaRep

	KmemoReps []kmemo.KmemoRep

	TagReps []tag.TagRep

	TextReps []text.TextRep

	DeleteTagReps tag.DeleteTagReps

	// ˅

	// ˄
}

func (r *Repositories) Close() error {
	// ˅

	panic("notImplements")
	// ˄
}

// ˅

// ˄

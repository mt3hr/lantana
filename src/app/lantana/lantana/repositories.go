// ˅
package lantana

import (
	"fmt"

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
	for _, rep := range r.LantanaReps {
		err := rep.Close()
		if err != nil {
			err = fmt.Errorf("error at close %s: %w", rep.Path(), err)
			return err
		}
	}
	for _, rep := range r.TagReps {
		err := rep.Close()
		if err != nil {
			err = fmt.Errorf("error at close %s: %w", rep.Path(), err)
			return err
		}
	}
	for _, rep := range r.TextReps {
		err := rep.Close()
		if err != nil {
			err = fmt.Errorf("error at close %s: %w", rep.Path(), err)
			return err
		}
	}
	err := r.LantanaRep.Close()
	if err != nil {
		err = fmt.Errorf("error at close %s: %w", r.LantanaRep.Path(), err)
		return err
	}
	err = r.KmemoRep.Close()
	if err != nil {
		err = fmt.Errorf("error at close %s: %w", r.KmemoRep.Path(), err)
		return err
	}
	err = r.TagRep.Close()
	if err != nil {
		err = fmt.Errorf("error at close %s: %w", r.TagRep.Path(), err)
		return err
	}
	err = r.TextRep.Close()
	if err != nil {
		err = fmt.Errorf("error at close %s: %w", r.TextRep.Path(), err)
		return err
	}
	return nil
	// ˄
}

// ˅

// ˄

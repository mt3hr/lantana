package lantana

import (
	"fmt"
	"os"
	"sort"

	"github.com/mattn/go-zglob"
	"github.com/mt3hr/kmemo"
	"github.com/mt3hr/rykv/tag"
	"github.com/mt3hr/rykv/text"
)

func init() {
	setEnv()

	LantanaRepFactories["lantana_db"] = func(contentFile string) ([]LantanaRep, error) {
		reps := []LantanaRep{}

		contentFile = os.ExpandEnv(contentFile)
		matches, _ := zglob.Glob(contentFile)
		for _, match := range matches {
			rep, err := NewLantanaRepSQLite(match)
			if err != nil {
				err = fmt.Errorf("failed to NewLantanaRepSQLite %s: %w", match, err)
				return nil, err
			}
			reps = append(reps, rep)
		}
		return reps, nil
	}

	kmemo.KmemoRepFactories["kmemo_db"] = func(contentFile string) ([]kmemo.KmemoRep, error) {
		reps := []kmemo.KmemoRep{}

		contentFile = os.ExpandEnv(contentFile)
		matches, _ := zglob.Glob(contentFile)
		sort.Strings(matches)
		for _, match := range matches {
			rep, err := kmemo.NewKmemoRepSQLite(match)
			if err != nil {
				err = fmt.Errorf("error at new kmemo rep sqlite %s: %w", match, err)
				return nil, err
			}
			reps = append(reps, rep)
		}
		return reps, nil
	}

	tag.TagRepFactories["db"] = func(contentFile string) ([]tag.TagRep, error) {
		reps := []tag.TagRep{}

		contentFile = os.ExpandEnv(contentFile)
		matches, _ := zglob.Glob(contentFile)
		sort.Strings(matches)
		for _, match := range matches {
			rep, err := tag.NewTagRepSQLite(match)
			if err != nil {
				err = fmt.Errorf("error at new tag rep sqlite %s: %w", match, err)
				return nil, err
			}
			reps = append(reps, rep)
		}
		return reps, nil
	}

	text.TextRepFactories["db"] = func(contentFile string) ([]text.TextRep, error) {
		reps := []text.TextRep{}

		contentFile = os.ExpandEnv(contentFile)
		matches, _ := zglob.Glob(contentFile)
		sort.Strings(matches)
		for _, match := range matches {
			rep, err := text.NewTextRepSQLite(match)
			if err != nil {
				err = fmt.Errorf("error at new text rep sqlite %s: %w", match, err)
				return nil, err
			}
			reps = append(reps, rep)
		}
		return reps, nil
	}
}

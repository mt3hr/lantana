package lantana

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/mt3hr/kmemo"
	"github.com/mt3hr/lantana/src/app/lantana"
	"github.com/mt3hr/rykv"
	"github.com/mt3hr/rykv/tag"
	"github.com/mt3hr/rykv/text"
)

// NoTag . tagが一つもついていないkyouに自動的につけられるタグ名
const NoTag = `no tag`

// FilterMode .
// タグの検索モード。And, Or, Onlyのいずれか
type FilterMode string

// FilterModeの一覧
const (
	And  FilterMode = "and"
	Or   FilterMode = "or"
	Only FilterMode = "only"
)

func filterReps(ctx context.Context, reps []lantana.LantanaRep, repNames []string) ([]lantana.LantanaRep, error) {
	matchReps := []lantana.LantanaRep{}
	for _, rep := range reps {
	loop:
		for _, repname := range repNames {
			if repname == rep.RepName() {
				matchReps = append(matchReps, rep)
				break loop
			}
		}
	}
	return matchReps, nil
}

// kyou := map[kyou.id]
func filterWords(ctx context.Context, reps []lantana.LantanaRep, textReps []text.TextRep, kmemoReps []kmemo.KmemoRep, words []string, notWords []string, and bool, query *lantana.LantanaSearchQuery) (map[string]*lantana.Lantana, error) {
	matchKyous := map[string]*lantana.Lantana{}
	// wordsがないときにはRep内のすべてのLantanaID
	if len(words) == 0 {
		allKyous, err := lantana.LantanaReps(reps).SearchLantana(ctx, query)
		if err != nil {
			err = fmt.Errorf("error at get all kyous: %w", err)
			return nil, err
		}

		// 重複がないようにMapに詰める
		for _, kyou := range allKyous {
			if _, exist := matchKyous[kyou.LantanaID]; !exist {
				matchKyous[kyou.LantanaID] = kyou
			}
		}

		// notWordsにhitしたものを外す
		if len(notWords) != 0 {
			notMatchKyous, err := orSearch(ctx, reps, textReps, kmemoReps, notWords, query)
			if err != nil {
				err := fmt.Errorf("error at orSearch: %w", err)
				return nil, err
			}
			for _, notMatchKyou := range notMatchKyous {
				if _, exist := matchKyous[notMatchKyou.LantanaID]; exist {
					delete(matchKyous, notMatchKyou.LantanaID)
				}
			}
		}
		return matchKyous, nil
	}
	// wordsの長さが1のときはor検索を使う（速いので）
	if len(words) == 1 {
		and = false
	}

	kyous := []*lantana.Lantana{}
	var err error
	if and {
		kyous, err = andSearch(ctx, reps, textReps, words, query)
		if err != nil {
			err = fmt.Errorf("failed to and search: %w", err)
			return nil, err
		}
	} else {
		kyous, err = orSearch(ctx, reps, textReps, kmemoReps, words, query)
		if err != nil {
			err = fmt.Errorf("failed to or search: %w", err)
			return nil, err
		}
	}

	// 重複がないようにMapに詰める
	for _, kyou := range kyous {
		if _, exist := matchKyous[kyou.LantanaID]; !exist {
			matchKyous[kyou.LantanaID] = kyou
		}
	}

	// notWordsにhitしたものを外す
	notLantanaIDs, err := orSearch(ctx, reps, textReps, kmemoReps, notWords, query)
	if err != nil {
		err := fmt.Errorf("error at orSearch: %w", err)
		return nil, err
	}
	for _, notLantanaID := range notLantanaIDs {
		if _, exist := matchKyous[notLantanaID.LantanaID]; exist {
			delete(matchKyous, notLantanaID.LantanaID)
		}
	}
	return matchKyous, nil
}

func orSearch(ctx context.Context, reps []lantana.LantanaRep, textReps []text.TextRep, kmemoReps []kmemo.KmemoRep, words []string, query *lantana.LantanaSearchQuery) ([]*lantana.Lantana, error) {
	// repにSearchしてヒットしたもの
	lantanas := []*lantana.Lantana{}

	allLantanas, err := lantana.LantanaReps(reps).GetAllLantanas(ctx)
	if err != nil {
		return nil, err
	}
	allLantanasTimeMap := map[int64]*lantana.Lantana{}
	for _, lantana := range allLantanas {
		allLantanasTimeMap[lantana.Time.Unix()] = lantana
	}

	matchLantanas, err := lantana.LantanaReps(reps).SearchLantana(ctx, query)
	if err != nil {
		err = fmt.Errorf("error at search: %w", err)
		return nil, err
	}

	// kmemoに検索してヒットしたもの
	rykvRepsWrap := rykv.Reps{}
	for _, kmemoRep := range kmemoReps {
		rykvRepsWrap = append(rykvRepsWrap, kmemoRep)
	}
	for _, word := range words {
		matchKmemoKyous, err := rykv.Reps(rykvRepsWrap).Search(ctx, word)
		if err != nil {
			return nil, err
		}
		for _, kmemoKyou := range matchKmemoKyous {
			if lantana, exist := allLantanasTimeMap[kmemoKyou.Time.Unix()]; exist {
				lantanas = append(lantanas, lantana)
			}
		}
	}

	//textRepにSearchしてヒットしたもの
	for _, word := range words {
		matchTexts, err := text.TextReps(textReps).Search(ctx, word)
		if err != nil {
			err = fmt.Errorf("error at search %s: %w", word, err)
			return nil, err
		}
		for _, text := range matchTexts {
			for _, kyou := range matchLantanas {
				if kyou.LantanaID == text.Target {
					lantanas = append(lantanas, kyou)
				}
			}
		}
	}
	// idが完全に一致するものも
	for _, kyou := range matchLantanas {
		for _, word := range words {
			if kyou.LantanaID == word {
				lantanas = append(lantanas, kyou)
			}
		}
	}
	return lantanas, nil
}

func andSearch(ctx context.Context, reps []lantana.LantanaRep, textReps []text.TextRep, words []string, query *lantana.LantanaSearchQuery) ([]*lantana.Lantana, error) {
	// searchで見つかったかどうか := map[id]map[word]
	m := map[string]map[string]bool{}
	hitKyous := map[string]*lantana.Lantana{}
	allKyous := []*lantana.Lantana{}

	allKyousMap := map[string]*lantana.Lantana{}
	allKyous, err := lantana.LantanaReps(reps).GetAllLantanas(ctx)
	if err != nil {
		err = fmt.Errorf("error at get all kyou: %w", err)
		return nil, err
	}
	for _, kyou := range allKyous {
		if _, exist := allKyousMap[kyou.LantanaID]; !exist {
			allKyousMap[kyou.LantanaID] = kyou
		}
	}
	for _, kyou := range allKyousMap {
		allKyous = append(allKyous, kyou)
	}

	for _, word := range words {
		var q lantana.LantanaSearchQuery
		q = *query
		q.Words = word
		kyous, err := lantana.LantanaReps(reps).SearchLantana(ctx, &q)
		if err != nil {
			err = fmt.Errorf("error at search %s: %w", word, err)
			return nil, err
		}
		for _, kyou := range kyous {
			if _, exist := m[kyou.LantanaID]; !exist {
				m[kyou.LantanaID] = map[string]bool{}
			}
			m[kyou.LantanaID][word] = true
			hitKyous[kyou.LantanaID] = kyou
		}
		texts, err := text.TextReps(textReps).Search(ctx, word)
		if err != nil {
			err = fmt.Errorf("error at search %s: %w", word, err)
			return nil, err
		}
		for _, textObj := range texts {
			texts, err := text.TextReps(textReps).GetTextsByTarget(ctx, textObj.ID)
			if err != nil {
				err = fmt.Errorf("error at get texts by target %s: %w", textObj.ID, err)
				return nil, err
			}

			for _, text := range texts {
				found := false
				for _, kyou := range allKyous {
					if kyou.LantanaID == text.Target {
						found = true
						break
					}
				}
				if !found {
					// repが分散しているとtargetの存在しないtextが出現し得るのでその場合はcontinue
					continue
				}

				if _, exist := m[text.Target]; !exist {
					m[text.Target] = map[string]bool{}
				}
				m[text.Target][word] = true
			}
		}
	}

	for _, word := range words {
		for _, wordMap := range m {
			if _, exist := wordMap[word]; !exist {
				wordMap[word] = false
			}
		}
	}

	kyous := []*lantana.Lantana{}
	ids := []string{}
	for id, wordMap := range m {
		allMatch := true
		for _, exist := range wordMap {
			if !exist && allMatch {
				allMatch = false
				break
			}
		}
		if allMatch {
			ids = append(ids, id)
		}
	}

	for _, id := range ids {
		kyou, exist := hitKyous[id]
		if !exist {
			found := false
			for _, k := range allKyous {
				if k.LantanaID == kyou.LantanaID {
					found = true
					kyou = k
					break
				}
			}
			if !found {
				err := fmt.Errorf("not found %s from all reps", id)
				return nil, err
			}
		}
		kyous = append(kyous, kyou)
	}
	return kyous, nil
}

func filterTags(ctx context.Context, matchKyous map[string]*lantana.Lantana, tagReps []tag.TagRep, tags []string, mode FilterMode, config *Config) (map[string]*lantana.Lantana, error) {
	// タグを持っていないidを取得する
	noHaveTagLantanaIDs := map[string]*lantana.Lantana{}
	haveTagLantanaIDs := map[string]struct{}{}
	allTags, err := tag.TagReps(tagReps).GetAllTags(ctx)
	if err != nil {
		err = fmt.Errorf("error at get all tags from tagrep: %w", err)
		return nil, err
	}
	for _, tag := range allTags {
		haveTagLantanaIDs[tag.Target] = struct{}{}
	}
	for _, id := range matchKyous {
		if _, exist := haveTagLantanaIDs[id.LantanaID]; !exist {
			noHaveTagLantanaIDs[id.LantanaID] = id
		}
	}

	if mode == Or {
		// tagがあり、or検索の場合は、タグにヒットしたやつすべて
		temp := map[string]*lantana.Lantana{}
		for _, tagname := range tags {
			tags, err := tag.TagReps(tagReps).GetTagsByName(ctx, tagname)
			if err != nil {
				err = fmt.Errorf("error at get tag by name %s from tagrep: %w", tagname, err)
				return nil, err
			}
			for _, tag := range tags {
				if id, exist := matchKyous[tag.Target]; exist {
					temp[id.LantanaID] = id
				}
			}
		}
		// notagが含まれたらタグを持っていないkyouを追加する
		for _, tag := range tags {
			if tag == NoTag {
				for _, id := range noHaveTagLantanaIDs {
					temp[id.LantanaID] = id
				}
			}
		}
		matchKyous = map[string]*lantana.Lantana{}
		for _, id := range temp {
			_, exist := matchKyous[id.LantanaID]
			if !exist {
				matchKyous[id.LantanaID] = id
			}
		}
		return filterHiddenTags(ctx, matchKyous, tagReps, tags, config)
	}

	temp := []*lantana.Lantana{}
	for _, tag := range tags {
		if tag == NoTag {
			for _, id := range noHaveTagLantanaIDs {
				temp = append(temp, id)
			}
		}
	}
	for i, tagname := range tags {
		switch i {
		case 0:
			tags, err := tag.TagReps(tagReps).GetTagsByName(ctx, tagname)
			if err != nil {
				err = fmt.Errorf("error at get tags by name %s from tagrep: %w", tagname, err)
				return nil, err
			}
			for _, tag := range tags {
				if id, exist := matchKyous[tag.Target]; exist {
					temp = append(temp, id)
				}
			}
		default:
			temppp := []*lantana.Lantana{}
			tags, err := tag.TagReps(tagReps).GetTagsByName(ctx, tagname)
			if err != nil {
				err = fmt.Errorf("failed to get tag by name %s from tagrep: %w", tagname, err)
				return nil, err
			}

			ids := []*lantana.Lantana{}
			for _, tag := range tags {
				if id, exist := matchKyous[tag.Target]; exist {
					ids = append(ids, id)
				}
			}

			for _, existLantanaID := range temp {
				exist := false
				for _, id := range ids {
					if existLantanaID.LantanaID == id.LantanaID {
						exist = true
					}
				}
				if exist {
					temppp = append(temppp, existLantanaID)
				}
			}
			temp = temppp
		}
	}
	matchKyous = map[string]*lantana.Lantana{}
	for _, id := range temp {
		_, exist := matchKyous[id.LantanaID]
		if !exist {
			matchKyous[id.LantanaID] = id
		}
	}

	// OnlyModeでNoTagが含まれたらAnd検索結果と同義なので
	if mode == And || (mode == Only && equal([]string{NoTag}, tags)) {
		return filterHiddenTags(ctx, matchKyous, tagReps, tags, config)
	} else if mode == Only {
		allTags, err := tag.TagReps(tagReps).GetAllTags(ctx)
		if err != nil {
			err = fmt.Errorf("error at get all tags: %w", err)
			return nil, err
		}

		// requestされたtagじゃないものがあったら除去する
		sortedTags := sort.StringSlice(tags)
		unMatchKyouLantanaIDs := map[string]struct{}{}
		for target := range matchKyous {
			attachedTagsMap := map[string]struct{}{}
			for _, tag := range allTags {
				if tag.Target == target {
					attachedTagsMap[tag.Tag] = struct{}{}
				}
			}
			attachedTags := []string{}
			for attachedTag := range attachedTagsMap {
				attachedTags = append(attachedTags, attachedTag)
			}
			sort.Strings(attachedTags)
			if !equal(sortedTags, attachedTags) {
				unMatchKyouLantanaIDs[target] = struct{}{}
			}
		}
		for unMatchKyouLantanaID := range unMatchKyouLantanaIDs {
			delete(matchKyous, unMatchKyouLantanaID)
		}
		return filterHiddenTags(ctx, matchKyous, tagReps, tags, config)
	}
	err = fmt.Errorf("invalid 'mode' value: %s", mode)
	return nil, err
}

func filterHiddenTags(ctx context.Context, matchKyous map[string]*lantana.Lantana, tagReps []tag.TagRep, tags []string, config *Config) (map[string]*lantana.Lantana, error) {
loop:
	for _, hiddenTag := range config.ApplicationConfig.HiddenTags {
		for _, tag := range tags {
			if hiddenTag == tag {
				continue loop
			}
		}
		tags, err := tag.TagReps(tagReps).GetTagsByName(ctx, hiddenTag)
		if err != nil {
			err = fmt.Errorf("error at get tags by name: %w", err)
			return nil, err
		}
		for _, tag := range tags {
			if _, exist := matchKyous[tag.Target]; exist {
				delete(matchKyous, tag.Target)
			}
		}
	}
	return matchKyous, nil
}

func equal(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func parseWords(word string) (words, notWords []string) {
	nextIsNotWord := false
	for _, word := range strings.Split(word, " ") {
		for _, word := range strings.Split(word, "　") {
			if strings.HasPrefix(word, "-") {
				nextIsNotWord = true
				word = strings.TrimPrefix(word, "-")
			}
			switch word {
			case "":
				continue
			case "-":
				nextIsNotWord = true
			default:
				if nextIsNotWord {
					notWords = append(notWords, word)
				} else {
					words = append(words, word)
				}
				nextIsNotWord = false
			}
		}
	}
	return words, notWords
}

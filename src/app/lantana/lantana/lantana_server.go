// ˅
package lantana

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/mitchellh/go-homedir"
	"github.com/mt3hr/kmemo"
	"github.com/mt3hr/lantana/src/app/lantana"
	"github.com/mt3hr/lantana/src/app/lantana/lantana/api_request_response"
	"github.com/mt3hr/rykv/tag"
	"github.com/mt3hr/rykv/text"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// ˄

type LantanaServer struct {
	// ˅

	// ˄

	ConfigFileName string

	Repositories *Repositories

	Config *Config

	// ˅

	// ˄
}

func (l *LantanaServer) HandleAddLantanaPage(w http.ResponseWriter, r *http.Request) {
	// ˅
	http.RedirectHandler(AddLantanaPageAddress, http.StatusFound).ServeHTTP(w, r)
	// ˄
}

func (l *LantanaServer) HandleLantanaViewerPage(w http.ResponseWriter, r *http.Request) {
	// ˅
	http.RedirectHandler(LantanaLogViewerPageAddress, http.StatusFound).ServeHTTP(w, r)
	// ˄
}

func (l *LantanaServer) HandleSearchLantana(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.SearchLantanaRequest{}
	response := &api_request_response.SearchLantanaResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	words, notWords := parseWords(request.Query.Words)

	lantanas, err := filterWords(r.Context(), l.Repositories.LantanaReps, l.Repositories.TextReps, words, notWords, false, request.Query)
	if err != nil {
		err = fmt.Errorf("lantana検索に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}
	lantanas, err = filterTags(r.Context(), lantanas, l.Repositories.TagReps, request.Query.Tags, Or, l.Config)
	if err != nil {
		err = fmt.Errorf("lantana検索に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}

	hitLantanas := []*lantana.Lantana{}
	for _, lantana := range lantanas {
		hitLantanas = append(hitLantanas, lantana)
	}
	sortLantanasByTime(hitLantanas)

	response.Lantanas = hitLantanas
	// ˄
}

func (l *LantanaServer) HandleAddLantana(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.AddLantanaRequest{}
	response := &api_request_response.AddLantanaResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	err = l.Repositories.LantanaRep.AddLantana(r.Context(), request.Lantana)
	if err != nil {
		err = fmt.Errorf("lantana追加に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}
	// ˄
}

func (l *LantanaServer) HandleDeleteLantana(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.DeleteLantanaRequest{}
	response := &api_request_response.DeleteLantanaResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	err = l.Repositories.LantanaRep.Delete(request.LantanaID)
	if err != nil {
		err = fmt.Errorf("lantana削除に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}
	// ˄
}

func (l *LantanaServer) HandleAddKmemo(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.AddKmemoRequest{}
	response := &api_request_response.AddKmemoResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	err = l.Repositories.KmemoRep.AddKmemo(request.Kmemo)
	if err != nil {
		err = fmt.Errorf("Kmemo追加に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}
	// ˄
}

func (l *LantanaServer) HandleDeleteKmemo(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.DeleteKmemoRequest{}
	response := &api_request_response.DeleteKmemoResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	err = l.Repositories.KmemoRep.Delete(request.KmemoID)
	if err != nil {
		err = fmt.Errorf("Kmemo削除に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}
	// ˄
}

func (l *LantanaServer) HandleAddTag(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.AddTagRequest{}
	response := &api_request_response.AddTagResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	err = l.Repositories.TagRep.AddTag(request.Tag)
	if err != nil {
		err = fmt.Errorf("タグ追加に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}
	// ˄
}

func (l *LantanaServer) HandleDeleteTag(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.DeleteTagRequest{}
	response := &api_request_response.DeleteTagResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}
	err = l.Repositories.TagRep.Delete(request.TagID)
	if err != nil {
		err = fmt.Errorf("タグ削除に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}
	// ˄
}

func (l *LantanaServer) HandleAddText(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.AddTextRequest{}
	response := &api_request_response.AddTextResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	err = l.Repositories.TextRep.AddText(request.Text)
	if err != nil {
		err = fmt.Errorf("テキスト追加に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}
	// ˄
}

func (l *LantanaServer) HandleDeleteText(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.DeleteTextRequest{}
	response := &api_request_response.DeleteTextResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	err = l.Repositories.TextRep.Delete(request.TextID)
	if err != nil {
		err = fmt.Errorf("テキスト削除に失敗しました")
		response.Errors = append(response.Errors, err.Error())
		return
	}
	// ˄
}

func (l *LantanaServer) HandleGetKmemosRelatedLantana(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.GetKmemosRelatedLantanaRequest{}
	response := &api_request_response.GetKmemosRelatedLantanaResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	matchLantana, err := lantana.LantanaReps(l.Repositories.LantanaReps).GetLantana(r.Context(), request.LantanaID)
	if err != nil {
		response.Errors = append(response.Errors, "Lantana取得に失敗しました")
		return
	}

	relatedKmemos, err := kmemo.KmemoReps(l.Repositories.KmemoReps).FindKmemoMatchTime(r.Context(), matchLantana.Time)
	if err != nil {
		response.Errors = append(response.Errors, "Kmemo取得に失敗しました")
		return
	}

	response.Kmemos = relatedKmemos
	// ˄
}

func (l *LantanaServer) HandleGetTagsRelatedLantana(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.GetTagsRelatedLantanaRequest{}
	response := &api_request_response.GetTagsRelatedLantanaResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	tags, err := tag.TagReps(l.Repositories.TagReps).GetTagsByName(r.Context(), request.LantanaID)
	if err != nil {
		response.Errors = append(response.Errors, "タグ取得に失敗しました")
		return
	}

	response.Tags = tags
	// ˄
}

func (l *LantanaServer) HandleGetTextsRelatedLantana(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.GetTextsRelatedLantanaRequest{}
	response := &api_request_response.GetTextsRelatedLantanaResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	texts := map[string]*text.Text{}
	wg := &sync.WaitGroup{}
	ch := make(chan []*text.Text, len(l.Repositories.TextReps))
	for _, textRep := range l.Repositories.TextReps {
		textRep := textRep
		wg.Add(1)
		go func(textRep text.TextRep) {
			defer wg.Done()
			matchTexts, err := textRep.GetTextsByTarget(r.Context(), request.LantanaID)
			if err != nil {
				response.Errors = append(response.Errors, "テキストの取得に失敗しました")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			ch <- matchTexts
		}(textRep)
	}
	wg.Wait()
loop:
	for {
		select {
		case t := <-ch:
			if t == nil {
				continue loop
			}
			for _, text := range t {
				texts[text.ID] = text
			}
		default:
			break loop
		}
	}
	textList := []*text.Text{}
	for _, matchText := range texts {
		textList = append(textList, matchText)
	}
	response.Texts = textList
	// ˄
}

func (l *LantanaServer) HandleGetTagsRelatedKmemo(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.GetTagsRelatedKmemoRequest{}
	response := &api_request_response.GetTagsRelatedKmemoResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	tags := map[string]*tag.Tag{}
	wg := &sync.WaitGroup{}
	ch := make(chan []*tag.Tag, len(l.Repositories.TagReps))
	for _, tagRep := range l.Repositories.TagReps {
		tagRep := tagRep
		wg.Add(1)
		go func(tagRep tag.TagRep) {
			defer wg.Done()
			matchTags, err := tagRep.GetTagsByTarget(r.Context(), request.KmemoID)
			if err != nil {
				response.Errors = append(response.Errors, "タグの取得に失敗しました")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			ch <- matchTags
		}(tagRep)
	}
	wg.Wait()
loop:
	for {
		select {
		case t := <-ch:
			if t == nil {
				continue loop
			}
			for _, tag := range t {
				tags[tag.ID] = tag
			}
		default:
			break loop
		}
	}
	tagList := []*tag.Tag{}
	for _, matchTag := range tags {
		tagList = append(tagList, matchTag)
	}
	response.Tags = tagList
	// ˄
}

func (l *LantanaServer) HandleGetTextsRelatedKmemo(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.GetTextsRelatedKmemoRequest{}
	response := &api_request_response.GetTextsRelatedKmemoResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	texts := map[string]*text.Text{}
	wg := &sync.WaitGroup{}
	ch := make(chan []*text.Text, len(l.Repositories.TextReps))
	for _, textRep := range l.Repositories.TextReps {
		textRep := textRep
		wg.Add(1)
		go func(textRep text.TextRep) {
			defer wg.Done()
			matchTexts, err := textRep.GetTextsByTarget(r.Context(), request.KmemoID)
			if err != nil {
				response.Errors = append(response.Errors, "テキストの取得に失敗しました")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			ch <- matchTexts
		}(textRep)
	}
	wg.Wait()
loop:
	for {
		select {
		case t := <-ch:
			if t == nil {
				continue loop
			}
			for _, text := range t {
				texts[text.ID] = text
			}
		default:
			break loop
		}
	}
	textList := []*text.Text{}
	for _, matchText := range texts {
		textList = append(textList, matchText)
	}
	response.Texts = textList
	// ˄
}

func (l *LantanaServer) HandleGetApplicationConfig(w http.ResponseWriter, r *http.Request) {
	// ˅
	w.Header().Set("Content-Type", "application/json")
	request := &api_request_response.GetApplicationConfigRequest{}
	response := &api_request_response.GetApplicationConfigResponse{}

	defer r.Body.Close()
	defer func() {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}()

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		panic(err)
	}

	response.ApplicationConfig = l.Config.ApplicationConfig
	// ˄
}

func (l *LantanaServer) LoadConfig() error {
	// ˅
	configOpt := l.getConfigFile()
	config := l.getConfig()
	configName := l.getConfigName()
	configExt := l.getConfigExt()

	v := viper.New()
	configPaths := []string{}
	if configOpt != "" {
		// コンフィグファイルが明示的に指定された場合はそれを
		v.SetConfigFile(configOpt)
		configPaths = append(configPaths, configOpt)
	} else {
		// 実行ファイルの親ディレクトリ、カレントディレクトリ、ホームディレクトリの順に
		v.SetConfigName(configName)
		exe, err := os.Executable()
		if err != nil {
			err = fmt.Errorf("error at get executable file path: %w", err)
			log.Printf(err.Error())
		} else {
			v.AddConfigPath(filepath.Dir(exe))
			configPaths = append(configPaths, filepath.Join(filepath.Dir(exe), configName+configExt))
		}

		v.AddConfigPath(".")
		configPaths = append(configPaths, filepath.Join(".", configName+configExt))

		home, err := homedir.Dir()
		if err != nil {
			err = fmt.Errorf("error at get user home directory: %w", err)
			log.Printf(err.Error())
		} else {
			v.AddConfigPath(home)
			configPaths = append(configPaths, filepath.Join(home, configName+configExt))
		}
	}

	// 読み込んでcfgを作成する
	existConfigPath := false
	for _, configPath := range configPaths {
		if _, err := os.Stat(configPath); err == nil {
			existConfigPath = true
			break
		}
	}
	if !existConfigPath {
		// コンフィグファイルが指定されていなくてコンフィグファイルが見つからなかった場合、
		// ホームディレクトリにデフォルトコンフィグファイルを作成する。
		// できなければカレントディレクトリにコンフィグファイルを作成する。
		if configOpt == "" {
			configDir := ""
			home, err := homedir.Dir()
			if err != nil {
				err = fmt.Errorf("error at get user home directory: %w", err)
				log.Printf(err.Error())
				configDir = "."
			} else {
				configDir = home
			}

			configFileName := filepath.Join(configDir, configName+configExt)
			err = os.WriteFile(configFileName, []byte(l.CreateDefaultConfigYAML()), os.ModePerm)
			if err != nil {
				err = fmt.Errorf("error at write file to %s: %w", configFileName, err)
				return err
			}
			v.SetConfigFile(configFileName)
		} else {
			err := fmt.Errorf("コンフィグファイルが見つかりませんでした。")
			return err
		}
	}

	err := v.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("error at read in config: %w", err)
		return err
	}

	err = v.Unmarshal(config)
	if err != nil {
		err = fmt.Errorf("error at unmarshal config file: %w", err)
		return err
	}

	// 各DBファイルの作成
	if l.Repositories.LantanaRep == nil {
		err := fmt.Errorf("configファイルのRepositories.LantanaRepの項目が設定されていないかあるいは不正です")
		return err
	}
	if l.Repositories.TagRep == nil {
		err := fmt.Errorf("configファイルのRepositories.TagRepの項目が設定されていないかあるいは不正です")
		return err
	}
	if l.Repositories.TextRep == nil {
		err := fmt.Errorf("configファイルのRepositories.TextRepの項目が設定されていないかあるいは不正です")
		return err
	}
	files := []string{
		os.ExpandEnv(l.Config.Reps.LantanaRep.File),
		os.ExpandEnv(l.Config.Reps.TagRep.File),
		os.ExpandEnv(l.Config.Reps.TextRep.File),
	}

	for _, filename := range files {
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, os.ModePerm)
		if err != nil {
			err = fmt.Errorf("error at create file %s: %w", filename, err)
			return err
		}
		defer f.Close()
	}

	return nil
	// ˄
}

func (l *LantanaServer) LoadRepositories() error {
	// ˅
	r := &Repositories{}

	if l.Config.Reps.LantanaRep == nil {
		err := fmt.Errorf("configファイルのRepositories.LantanaRepの項目が設定されていないかあるいは不正です")
		return err
	}
	reps, err := lantana.LoadLantanaReps(l.Config.Reps.LantanaRep)
	if err != nil {
		err = fmt.Errorf("error at load rep: %w", err)
		return err
	}
	r.LantanaRep = reps[0]

	if l.Config.Reps.LantanaReps == nil {
		err := fmt.Errorf("configファイルのRepositories.LantanaRepsの項目が設定されていないかあるいは不正です")
		return err
	}
	for _, repInfo := range l.Config.Reps.LantanaReps {
		reps, err := lantana.LoadLantanaReps(repInfo)
		if err != nil {
			err = fmt.Errorf("error at load reps: %w", err)
			return err
		}
		r.LantanaReps = append(r.LantanaReps, reps...)
	}

	if l.Config.Reps.KmemoRep == nil {
		err := fmt.Errorf("configファイルのRepositories.KmemoRepの項目が設定されていないかあるいは不正です")
		return err
	}
	kmemoRep, err := kmemo.LoadKmemoReps(l.Config.Reps.KmemoRep)
	if err != nil {
		err = fmt.Errorf("error at load rep: %w", err)
		return err
	}
	r.KmemoRep = kmemoRep[0]

	if l.Config.Reps.KmemoReps == nil {
		err := fmt.Errorf("configファイルのRepositories.KmemoRepsの項目が設定されていないかあるいは不正です")
		return err
	}
	for _, kmemoRepInfo := range l.Config.Reps.KmemoReps {
		kmemoReps, err := kmemo.LoadKmemoReps(kmemoRepInfo)
		if err != nil {
			err = fmt.Errorf("error at load kmemo reps type=%s file=%s: %w", kmemoRepInfo.Type, kmemoRepInfo.File, err)
			return err
		}
		r.KmemoReps = append(r.KmemoReps, kmemoReps...)
	}

	if l.Config.Reps.TagReps == nil {
		err := fmt.Errorf("configファイルのRepositories.TagRepsの項目が設定されていないかあるいは不正です")
		return err
	}
	for _, tagRepInfo := range l.Config.Reps.TagReps {
		tagReps, err := tag.LoadTagReps(tagRepInfo)
		if err != nil {
			err = fmt.Errorf("error at load tag reps type=%s file=%s: %w", tagRepInfo.Type, tagRepInfo.File, err)
			return err
		}
		r.TagReps = append(r.TagReps, tagReps...)
	}

	if l.Config.Reps.TextReps == nil {
		err := fmt.Errorf("configファイルのRepositories.TextRepsの項目が設定されていないかあるいは不正です")
		return err
	}
	for _, textRepInfo := range l.Config.Reps.TextReps {
		textReps, err := text.LoadTextReps(textRepInfo)
		if err != nil {
			err = fmt.Errorf("error at load text reps type=%s file=%s: %w", textRepInfo.Type, textRepInfo.File, err)
			return err
		}
		r.TextReps = append(r.TextReps, textReps...)
	}

	if l.Config.Reps.TagRep == nil {
		err := fmt.Errorf("configファイルのRepositories.TagRepの項目が設定されていないかあるいは不正です")
		return err
	}
	writetoTagRep, err := tag.LoadTagReps(l.Config.Reps.TagRep)
	if err != nil {
		err = fmt.Errorf("error at load write to tag rep: %w", err)
		return err
	}
	if len(writetoTagRep) != 1 {
		err = fmt.Errorf("見つかったtag repの数が1つではありませんでした。")
		return err
	}
	r.TagRep = writetoTagRep[0]

	if l.Config.Reps.TextRep == nil {
		err := fmt.Errorf("configファイルのRepositories.TextRepの項目が設定されていないかあるいは不正です")
		return err
	}
	writetoTextRep, err := text.LoadTextReps(l.Config.Reps.TextRep)
	if err != nil {
		err = fmt.Errorf("error at to load write to text rep: %w", err)
		return err
	}
	if len(writetoTextRep) != 1 {
		err = fmt.Errorf("見つかったtext repの数が1つではありませんでした。")
		return err
	}
	r.TextRep = writetoTextRep[0]

	r.DeleteTagReps = tag.NewDeleteTagReps(r.TagRep, r.TagReps, time.Hour*24*365)

	l.Repositories = r
	return nil
	// ˄
}

func (l *LantanaServer) LoadTagStruct() error {
	// ˅
	configOpt := l.getConfigFile()
	configName := l.getConfigName()
	configExt := l.getConfigExt()
	configPaths := []string{}
	configFileName := ""
	var b []byte

	if configOpt != "" {
		// コンフィグファイルが明示的に指定された場合はそれを
		configPaths = append(configPaths, configOpt)
	} else {
		// 実行ファイルの親ディレクトリ、カレントディレクトリ、ホームディレクトリの順に
		exe, err := os.Executable()
		if err != nil {
			err = fmt.Errorf("error at get executable file path: %w", err)
			log.Printf(err.Error())
		} else {
			configPaths = append(configPaths, filepath.Join(filepath.Dir(exe), configName+configExt))
		}

		configPaths = append(configPaths, filepath.Join(".", configName+configExt))

		home, err := homedir.Dir()
		if err != nil {
			err = fmt.Errorf("error at get user home directory: %w", err)
			log.Printf(err.Error())
		} else {
			configPaths = append(configPaths, filepath.Join(home, configName+configExt))
		}
	}

	for _, configPath := range configPaths {
		if _, err := os.Stat(configPath); err == nil {
			configFileName = configPath
			break
		}
	}

	b, err := os.ReadFile(configFileName)
	if err != nil {
		err = fmt.Errorf("error at read file %s: %w", configFileName, err)
		return err
	}

	m := yaml.MapSlice{}
	tagStructMap := yaml.MapSlice{}
	err = yaml.Unmarshal(b, &m)
	if err != nil {
		err = fmt.Errorf("error at yaml unmarshall: %w", err)
		return err
	}
	for _, v := range m {
		if v.Key == "ApplicationConfig" {
			i, ok := v.Value.(yaml.MapSlice)
			if !ok {
				err = fmt.Errorf("configファイルが変です。多分ApplicationConfigの項目がありません")
				return err
			}
			for _, v := range i {
				if v.Key == "TagStruct" {
					tagStructMap, ok = v.Value.(yaml.MapSlice)
					if !ok {
						err = fmt.Errorf("configファイルが変です。多分ApplicationConfigの項目、TagStructがありません")
						return err
					}
				}
			}
		}
	}
	l.Config.ApplicationConfig.TagStruct = MapSlice(tagStructMap)
	return nil
	// ˄
}

func (l *LantanaServer) WrapT() error {
	// ˅
	l.Repositories.LantanaReps = l.wrapLantanaRepsT(l.Repositories.LantanaReps, l.Repositories.DeleteTagReps)
	l.Repositories.TagReps = l.wrapTagRepsT(l.Repositories.TagReps, l.Repositories.DeleteTagReps)
	l.Repositories.TextReps = l.wrapTextRepsT(l.Repositories.TextReps, l.Repositories.DeleteTagReps)
	return nil
	// ˄
}

func (l *LantanaServer) wrapLantanaRepsT(reps []lantana.LantanaRep, deleteTagReps tag.DeleteTagReps) []lantana.LantanaRep {
	wrapedReps := []lantana.LantanaRep{}
	for _, rep := range reps {
		wrapedReps = append(wrapedReps, lantana.WrapLantanaRepT(rep, deleteTagReps))
	}
	return wrapedReps
}

func (l *LantanaServer) wrapTagRepsT(reps []tag.TagRep, deleteTagReps tag.DeleteTagReps) []tag.TagRep {
	wrapedReps := []tag.TagRep{}
	for _, rep := range reps {
		wrapedReps = append(wrapedReps, tag.WrapTagRepT(rep, deleteTagReps))
	}
	return wrapedReps
}

func (l *LantanaServer) wrapTextRepsT(reps []text.TextRep, deleteTagReps tag.DeleteTagReps) []text.TextRep {
	wrapedReps := []text.TextRep{}
	for _, rep := range reps {
		wrapedReps = append(wrapedReps, text.WrapTextRepT(rep, deleteTagReps))
	}
	return wrapedReps
}

func (l *LantanaServer) CreateDefaultConfigYAML() string {
	// ˅
	return `ServerConfig:
  # trueにするとlocalhost以外からのリクエストをブロックします。
  LocalOnly: true

  # rykv server でサーバーをたてるときに使うアドレス
  Address: ":7777"

  # 設定するとhttpsで接続するようになります
  TLS:
    Enable: false
    CertFile: ""
    KeyFile: ""

ApplicationConfig:
  # 読み込み時にチェックを入れないTag
  UnCheckTags: []

  # ここに記されたタグのついた情報は、チェックが入っていない限り検索結果に反映されません。
  # UncheckTagsと組み合わせて使います。
  # 削除機能もこの機能で実現されています。
  HiddenTags: []

  # Tagの階層構造の設定。
  TagStruct: 
    no tag: tag

Reps:
  # Lantana記録時の保存先データベースファイル
  LantanaRep:
    type: lantana_db
	file: $HOME/Lantana.db

  # Lantana情報源データベースファイル郡
  LantanaReps:
  - type: lantana_db
	file: $HOME/Lantana.db

  # Kmemo記録時の保存先データベースファイル
  KmemoRep:
    type: db
    file: $HOME/Kmemo.db

  # Kmemo情報源先データベースファイル郡
  KmemoReps:
  - type: db
    file: $HOME/Kmemo.db

  # タグ記録時の保存先データベースファイル
  TagRep:
    type: db
    file: $HOME/Tag.db

  # タグ情報源データベースファイル郡
  TagReps:
  - type: db
    file: $HOME/Tag.db

  # テキスト記録時の保存先データベースファイル
  TextRep:
    type: db
    file: $HOME/Text.db

  # テキスト情報源データベースファイル郡
  TextReps:
  - type: db
    file: $HOME/Text.db
`
	// ˄
}

func (l *LantanaServer) LaunchServer() error {
	// ˅
	router := mux.NewRouter()

	router.PathPrefix(SearchLantanaAddress).HandlerFunc(l.HandleSearchLantana)
	router.PathPrefix(AddLantanaAddress).HandlerFunc(l.HandleAddLantana)
	router.PathPrefix(DeleteLantanaAddress).HandlerFunc(l.HandleDeleteLantana)
	router.PathPrefix(AddKmemoAddress).HandlerFunc(l.HandleAddKmemo)
	router.PathPrefix(DeleteKmemoAddress).HandlerFunc(l.HandleDeleteKmemo)
	router.PathPrefix(AddTagAddress).HandlerFunc(l.HandleAddTag)
	router.PathPrefix(DeleteTagAddress).HandlerFunc(l.HandleDeleteTag)
	router.PathPrefix(AddTextAddress).HandlerFunc(l.HandleAddText)
	router.PathPrefix(DeleteTextAddress).HandlerFunc(l.HandleDeleteText)
	router.PathPrefix(GetKmemosRelatedLantanaAddress).HandlerFunc(l.HandleGetKmemosRelatedLantana)
	router.PathPrefix(GetTagsRelatedLantanaAddress).HandlerFunc(l.HandleGetTagsRelatedLantana)
	router.PathPrefix(GetTextsRelatedLantanaAddress).HandlerFunc(l.HandleGetTextsRelatedLantana)
	router.PathPrefix(GetTagsRelatedKmemoAddress).HandlerFunc(l.HandleGetTagsRelatedKmemo)
	router.PathPrefix(GetTextsRelatedKmemoAddress).HandlerFunc(l.HandleGetTextsRelatedKmemo)
	router.PathPrefix(GetApplicationConfigAddress).HandlerFunc(l.HandleGetApplicationConfig)

	addLantanaPage, err := fs.Sub(lantana.EmbedDir, "lantana/embed/html")
	if err != nil {
		return err
	}
	router.PathPrefix(LantanaLogViewerPageAddress).Handler(http.FileServer(http.FS(addLantanaPage)))
	router.PathPrefix(AddLantanaPageAddress).Handler(http.FileServer(http.FS(addLantanaPage)))

	var handler http.Handler = router
	if l.Config.ServerConfig.LocalOnly {
		h := handler
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			spl := strings.Split(r.RemoteAddr, ":")
			remoteHost := strings.Join(spl[:len(spl)-1], ":")
			switch remoteHost {
			case "localhost":
				fallthrough
			case "127.0.0.1":
				fallthrough
			case "[::1]":
				fallthrough
			case "::1":
				h.ServeHTTP(w, r)
				return
			}
			w.WriteHeader(http.StatusForbidden)
		})
	}

	if l.Config.ServerConfig.TLS.Enable {
		err = http.ListenAndServeTLS(
			l.Config.ServerConfig.Address,
			os.ExpandEnv(l.Config.ServerConfig.TLS.CertFile),
			os.ExpandEnv(l.Config.ServerConfig.TLS.KeyFile),
			handler)
		if err != nil {
			err = fmt.Errorf("failed to launch server: %w", err)
			return err
		}
	} else {
		err = http.ListenAndServe(l.Config.ServerConfig.Address, handler)
		if err != nil {
			err = fmt.Errorf("failed to launch server: %w", err)
			return err
		}
	}

	panic("notImplements")
	// ˄
}

func (l *LantanaServer) getConfigFile() string {
	// ˅
	return l.ConfigFileName
	// ˄
}

func (l *LantanaServer) getConfig() *Config {
	// ˅
	return l.Config
	// ˄
}

func (l *LantanaServer) getConfigName() string {
	// ˅
	return "lantana_config"
	// ˄
}

func (l *LantanaServer) getConfigExt() string {
	// ˅
	return ".yaml"
	// ˄
}

// ˅

// MapSlice . yaml.MapSliceをJSONにするために用意されたものです
type MapSlice yaml.MapSlice

// MapItem . yaml.MapItemをJSONにするために用意されたものです
type MapItem yaml.MapItem

// MarshalJSON . JSONにMarshalします。
func (m MapSlice) MarshalJSON() ([]byte, error) {
	jsonStr := "{"
	for i, item := range m {
		if i != 0 {
			jsonStr += ","
		}
		switch value := interface{}(item.Value).(type) {
		case yaml.MapSlice:
			itemJSON, err := json.Marshal(MapSlice(value))
			if err != nil {
				err = fmt.Errorf("error at marshal json: %w", err)
				return nil, err
			}
			jsonStr += fmt.Sprintf(`"%s": %s`, item.Key, string(itemJSON))
		case yaml.MapItem:
			ValueJSON, err := json.Marshal(MapItem(value))
			if err != nil {
				err = fmt.Errorf("error at marshal json: %w", err)
				return nil, err
			}
			jsonStr += fmt.Sprintf(`"%s": "%s"`, item.Key, string(ValueJSON))
		case string:
			jsonStr += fmt.Sprintf(`"%s": "%s"`, item.Key, value)
		default:
			err := fmt.Errorf("変な型が渡されました %s", reflect.TypeOf(item.Value))
			return nil, err
		}
	}
	jsonStr += "}"
	return []byte(jsonStr), nil
}

// MarshalJSON . JSONにMarshalします。
func (m MapItem) MarshalJSON() ([]byte, error) {
	jsonStr := "{"
	switch value := interface{}(m.Value).(type) {
	case yaml.MapSlice:
		itemJSON, err := json.Marshal(MapSlice(value))
		if err != nil {
			err = fmt.Errorf("error at marshal json: %w", err)
			return nil, err
		}
		jsonStr += fmt.Sprintf(`"%s": %s`, m.Key, string(itemJSON))
	case string:
		jsonStr += fmt.Sprintf(`"%s": "%s" `, m.Key, value)
	default:
		err := fmt.Errorf("変な型が渡されました %s", reflect.TypeOf(m.Value))
		return nil, err
	}
	jsonStr += "}"
	return []byte(jsonStr), nil
}

// ˄

package lantana

import (
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-zglob"
	"github.com/mitchellh/go-homedir"
)

// LantanaRepFactories . 詳しくはLoadLantanaRepsを参照してください。
var LantanaRepFactories = map[string]LantanaRepFactory{}

// LantanaRepInfo . 詳しくはLoadLantanaRepsを参照してください。
type LantanaRepInfo struct {
	Type string
	File string
}

// LantanaRepFactory . 詳しくはLoadLantanaRepsを参照してください。
type LantanaRepFactory func(contentFile string) ([]LantanaRep, error)

// LoadLantanaReps .
// 使い方を示します。
//
// まず使いたいRepInfoを想定します。
// 例えば次のようにします。
//
//	repInfo := RepInfo {
//		type: "db",
//		file: "hoge.db",
//	}
//
// 次に、そのrepInfoが読み込めるように、RepFactoryを作成し、RepFactoriesにを登録します。
// contentFileにはrepInfo.fileが渡されます。
// repFactory := func(contentFile string) ([]Rep, error) { //hogefuga }
// RepFactories[repInfo.type] = repFactory
//
// 最後に、LoadRepsにrepInfoをわたしてRepオブジェクトを読み込みます。
// reps, err := LoadReps(repInfo)
func LoadLantanaReps(repInfo *LantanaRepInfo) ([]LantanaRep, error) {
	reps := []LantanaRep{}
	factory, exist := LantanaRepFactories[repInfo.Type]
	if !exist {
		err := fmt.Errorf("unknown rep type %s", repInfo.Type)
		return nil, err
	}
	rep, err := factory(repInfo.File)
	if err != nil {
		err = fmt.Errorf("failed to load rep %s: %w", repInfo.File, err)
		return nil, err
	}
	reps = append(reps, rep...)
	return reps, nil
}

func init() {
	setEnv()
	registLantanaDirectoryToFactories()
}

func registLantanaDirectoryToFactories() {
	LantanaRepFactories["lantana_db"] = func(contentFile string) ([]LantanaRep, error) {
		reps := []LantanaRep{}

		contentFile = os.ExpandEnv(contentFile)
		matches, _ := zglob.Glob(contentFile)
		for _, match := range matches {
			rep, err := NewLantanaRepSQLite(match)
			if err != nil {
				err = fmt.Errorf("failed to load lantana rep dir %s: %w", match, err)
				return nil, err
			}
			reps = append(reps, rep)
		}
		return reps, nil
	}
}

// 環境変数が設定されていなければ設定します
// ExpandEnvで使います。
func setEnv() {
	// HOME
	home := os.Getenv("HOME")
	if home == "" {
		home, err := homedir.Dir()
		if err != nil {
			err = fmt.Errorf("error at get user home directory: %w", err)
			log.Printf(err.Error())
		} else {
			os.Setenv("HOME", home)
		}
	}

	// EXE
	exe := os.Getenv("EXE")
	if exe == "" {
		exe, err := os.Executable()
		if err != nil {
			err = fmt.Errorf("error at get executable file path: %w", err)
			log.Printf(err.Error())
		} else {
			os.Setenv("EXE", exe)
		}
	}
}

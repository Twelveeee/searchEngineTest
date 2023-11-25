package script

import (
	"encoding/json"
	"os"
	"searchEngineTest/model"
)

const MeilisearchKey = "aSampleMasterKey"

func getJsonDataList(file string) []model.Article {
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	var articleList []model.Article

	decoder := json.NewDecoder(jsonFile)
	if err = decoder.Decode(&articleList); err != nil {
		panic(err)
	}

	return articleList
}

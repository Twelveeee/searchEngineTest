package script

import (
	"fmt"
	"searchEngineTest/model"

	algolia "github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/urfave/cli/v2"
)

type algoliaRecord struct {
	ObjectID string `json:"objectID"`
	model.Article
}

func algoliaGetClient(config *model.Config) *algolia.Client {
	client := algolia.NewClient(config.Algolia.ApplicationId, config.Algolia.AdminApiKey)
	return client
}

// 创建索引
func algoliaCreateIndexAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	client := algoliaGetClient(config)
	indexName := config.Algolia.IndexName
	_ = client.InitIndex(indexName)

	fmt.Printf("algolia: create index success \n")
	return nil
}

// 导入数据
func algoliaImportAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	client := algoliaGetClient(config)
	indexName := config.Algolia.IndexName
	indexClient := client.InitIndex(indexName)

	articleList := getJsonDataList(config.DataFile)

	var interfaceList []interface{}
	for _, article := range articleList {
		obj := &algoliaRecord{}
		obj.ObjectID = article.Rid
		obj.Article = article
		interfaceList = append(interfaceList, obj)
	}

	resList, err := indexClient.SaveObjects(interfaceList)
	if err != nil {
		return err
	}
	for _, res := range resList.Responses {
		fmt.Printf("algolia: add data status:%s TaskUID: %v\n", "success", res.TaskID)
	}

	return nil
}

// 搜索
func algoliaSearchAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	client := algoliaGetClient(config)
	indexName := config.Algolia.IndexName
	indexClient := client.InitIndex(indexName)

	query := ctx.String("query")
	ret, err := indexClient.Search(query)

	fmt.Printf("algolia: TotalHits: %d \n", ret.NbHits)
	for _, hit := range ret.Hits {
		document := hit
		fmt.Printf("algolia: rid: %s name: %s \n", document["Rid"], document["Name"])
	}
	if err != nil {
		return err
	}
	return nil
}

// 删除索引
func algoliaDeleteIndexAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	client := algoliaGetClient(config)
	indexName := config.Algolia.IndexName
	indexClient := client.InitIndex(indexName)

	_, err = indexClient.Delete()
	if err != nil {
		return err
	}
	fmt.Printf("algolia: delete index success \n")
	return nil
}

// 压力测试
func algoliaTestAction(ctx *cli.Context) error {
	fmt.Printf("algolia: do not test \n")
	return nil
}

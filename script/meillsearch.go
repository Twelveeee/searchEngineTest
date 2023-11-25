package script

import (
	"fmt"
	"searchEngineTest/model"

	"github.com/meilisearch/meilisearch-go"
	"github.com/urfave/cli/v2"
)

func meillSearchGetClient(config *model.Config) *meilisearch.Client {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   config.MeillSearch.Host,
		APIKey: config.MeillSearch.APIKey,
	})
	return client
}

// 导入数据
func MeillSearchImportAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	index := config.MeillSearch.IndexName
	client := meillSearchGetClient(config)

	articleList := getJsonDataList(config.DataFile)

	resList, err := client.Index(index).AddDocumentsInBatches(articleList, 500, "Rid")
	if err != nil {
		return err
	}
	for _, res := range resList {
		fmt.Printf("add MeillSearch status:%s TaskUID: %v\n", res.Status, res.TaskUID)

		res2, err := client.GetTask(res.TaskUID)

		if err != nil {
			return err
		}
		fmt.Printf("add MeillSearch result status:%s \n", res2.Status)
	}

	return nil
}

// 搜索一次
func MeillSearchOnceAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	index := config.MeillSearch.IndexName
	client := meillSearchGetClient(config)

	query := ctx.String("query")

	req := &meilisearch.SearchRequest{
		Limit:            10,
		MatchingStrategy: "all",
	}
	ret, err := client.Index(index).Search(query, req)
	if err != nil {
		return err
	}

	fmt.Printf("TotalHits: %d \n", ret.EstimatedTotalHits)
	for _, hit := range ret.Hits {
		document := hit.(map[string]interface{})
		fmt.Printf("rid: %s name: %s \n", document["Rid"], document["Name"])
	}
	return nil
}

// 删除索引
func MeillSearchDeleteIndexAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	index := config.MeillSearch.IndexName
	client := meillSearchGetClient(config)

	res, err := client.DeleteIndex(index)
	if err != nil {
		return err
	}
	fmt.Printf("add MeillSearch status:%s TaskUID: %v\n", res.Status, res.TaskUID)

	return nil
}

// 创建索引
func MeillSearchCreateIndexAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	index := config.MeillSearch.IndexName
	client := meillSearchGetClient(config)

	indexConfig := &meilisearch.IndexConfig{
		Uid:        index,
		PrimaryKey: "Rid",
	}
	res, err := client.CreateIndex(indexConfig)
	if err != nil {
		return err
	}
	fmt.Printf("create MeillSearch index  status:%s TaskUID: %v\n", res.Status, res.TaskUID)

	searchableAttributes := []string{
		"Tags",
		"Author",
		"Name",
		"Full_name",
		"Title",
		"Description",
		"Summary",
	}
	res, err = client.Index(index).UpdateSearchableAttributes(&searchableAttributes)
	if err != nil {
		return err
	}
	fmt.Printf("update MeillSearch index  status:%s TaskUID: %v\n", res.Status, res.TaskUID)

	return nil
}

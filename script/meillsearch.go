package script

import (
	"fmt"
	"net/http"
	"searchEngineTest/model"
	"time"

	"github.com/meilisearch/meilisearch-go"
	vegeta "github.com/tsenart/vegeta/lib"
	"github.com/urfave/cli/v2"
)

func meillSearchGetClient(config *model.Config) *meilisearch.Client {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   config.MeillSearch.Host,
		APIKey: config.MeillSearch.APIKey,
	})
	return client
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

// 搜索
func MeillSearchSearchAction(ctx *cli.Context) error {
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

	_, err = client.DeleteIndex(index)
	if err != nil {
		return err
	}
	fmt.Printf("delete index success \n")

	return nil
}

// 压力测试
func MeillSearchTestAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}

	url := config.MeillSearch.Host + "/indexes/article/search"
	Auth := []string{"Bearer " + config.MeillSearch.APIKey}

	body := `{
		"q": "压力测试",
		"attributesToHighlight": [
			"*"
		],
		"limit": 10,
		"offset": 0
	}`

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    url,
		Body:   []byte(body),
		Header: http.Header{
			"Content-Type":         []string{"application/json"},
			"Cache-Control":        []string{"no-cache"},
			"Authorization":        Auth,
			"Pragma":               []string{"no-cache"},
			"X-Meilisearch-Client": []string{"Meilisearch mini-dashboard (v0.2.11) ; Meilisearch instant-meilisearch (v0.11.1) ; Meilisearch JavaScript (v0.31.1)"},
		},
	})

	rate := vegeta.Rate{Freq: config.TestRate.PerSecond, Per: time.Second}
	duration := time.Duration(config.TestRate.Duration) * time.Second

	if err := pressureTest(targeter, rate, duration); err != nil {
		return err
	}

	return nil
}

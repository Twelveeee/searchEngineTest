package script

import (
	"fmt"
	"net/http"
	"searchEngineTest/model"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"github.com/urfave/cli/v2"
)

func typeSenseGetClient(config *model.Config) *typesense.Client {
	client := typesense.NewClient(
		typesense.WithServer(config.Typesense.Host),
		typesense.WithAPIKey(config.Typesense.APIKey))
	return client
}

// 创建索引
func typeSenseCreateIndexAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	client := typeSenseGetClient(config)
	schema := typeSenseGetIndexSchema(config)
	_, err = client.Collections().Create(schema)
	if err != nil {
		return err
	}
	fmt.Printf("create index success \n")
	return nil
}

// 导入数据
func typeSenseImportAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	index := config.Typesense.IndexName
	client := typeSenseGetClient(config)

	articleList := getJsonDataList("data/data.json")

	var interfaceList []interface{}
	for _, article := range articleList {
		interfaceList = append(interfaceList, article)
	}

	importParams := &api.ImportDocumentsParams{
		Action:    pointer.String("create"),
		BatchSize: pointer.Int(500),
	}

	importResList, err := client.Collection(index).Documents().Import(interfaceList, importParams)
	if err != nil {
		return err
	}
	for _, importRes := range importResList {
		fmt.Printf("import stauts %v document %s\n", importRes.Success, importRes.Document)
	}

	return nil
}

// 搜索
func typeSenseSearchAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	index := config.Typesense.IndexName
	query := ctx.String("query")
	client := typeSenseGetClient(config)

	searchParameters := &api.SearchCollectionParams{
		Q:       query,
		QueryBy: "Name,Full_name,Title,Description,Summary",
		PerPage: pointer.Int(10),
	}
	ret, err := client.Collection(index).Documents().Search(searchParameters)
	if err != nil {
		return err
	}

	fmt.Printf("TotalHits: %d \n", *ret.Found)
	for _, hit := range *ret.Hits {
		document := *hit.Document
		fmt.Printf("rid: %s name: %s \n", document["Rid"], document["Name"])
	}
	return nil
}

// 删除索引
func typeSenseDeleteIndexAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	index := config.Typesense.IndexName
	client := typeSenseGetClient(config)

	_, err = client.Collection(index).Delete()
	if err != nil {
		return err
	}

	fmt.Printf("delete index success \n")
	return nil
}

func typeSenseGetIndexSchema(config *model.Config) *api.CollectionSchema {
	schema := &api.CollectionSchema{
		Name: config.Typesense.IndexName,
		Fields: []api.Field{
			{
				Name:  "Rid",
				Type:  "string",
				Index: pointer.True(),
			},
			{
				Name:  "Tags",
				Type:  "string",
				Index: pointer.True(),
				Facet: pointer.True(),
			},
			{
				Name:  "Author",
				Type:  "string",
				Index: pointer.True(),
			},
			{
				Name:     "Author_avatar",
				Type:     "string",
				Index:    pointer.False(),
				Optional: pointer.True(),
			},
			{
				Name:   "Name",
				Type:   "string",
				Index:  pointer.True(),
				Locale: pointer.String("zh"),
			},
			{
				Name:   "Full_name",
				Type:   "string",
				Index:  pointer.True(),
				Locale: pointer.String("zh"),
			},
			{
				Name:   "Title",
				Type:   "string",
				Index:  pointer.True(),
				Locale: pointer.String("zh"),
			},
			{
				Name:   "Description",
				Type:   "string",
				Index:  pointer.True(),
				Locale: pointer.String("zh"),
			},
			{
				Name:   "Summary",
				Type:   "string",
				Index:  pointer.True(),
				Locale: pointer.String("zh"),
			},
			{
				Name:  "Primary_lang",
				Type:  "string",
				Index: pointer.True(),
				Facet: pointer.True(),
			},
			{
				Name:     "Publish_at",
				Type:     "string",
				Index:    pointer.False(),
				Facet:    pointer.False(),
				Optional: pointer.True(),
			},
			{
				Name:  "Has_chinese",
				Type:  "bool",
				Index: pointer.True(),
				Facet: pointer.True(),
			},
			{
				Name:  "Stars",
				Type:  "int32",
				Index: pointer.True(),
				Facet: pointer.False(),
			},
		},
		DefaultSortingField: pointer.String("Stars"),
	}
	return schema
}

// 压力测试
func typeSenseTestAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	index := config.Typesense.IndexName

	url := config.Typesense.Host + "/collections/" + index + "/documents/search?q=压力测试&query_by=Name,Full_name,Title,Description,Summary&per_page=10"
	Auth := []string{config.Typesense.APIKey}

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    url,
		Header: http.Header{
			"Content-Type":        []string{"application/json"},
			"Cache-Control":       []string{"no-cache"},
			"X-TYPESENSE-API-KEY": Auth,
			"Pragma":              []string{"no-cache"},
		},
	})

	rate := vegeta.Rate{Freq: config.TestRate.PerSecond, Per: time.Second}
	duration := time.Duration(config.TestRate.Duration) * time.Second

	if err := pressureTest(targeter, rate, duration); err != nil {
		return err
	}

	return nil
}

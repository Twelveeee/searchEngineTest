package script

import (
	"fmt"
	"searchEngineTest/model"

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

func TypeSenseImportAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	index := config.Typesense.IndexName
	client := typeSenseGetClient(config)
	// create index
	schema := typeSenseGetIndexSchema(config)
	res, err := client.Collections().Create(schema)
	if err != nil {
		return err
	}
	fmt.Println(res)

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

// 搜索一次
func TypeSenseOnceAction(ctx *cli.Context) error {
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

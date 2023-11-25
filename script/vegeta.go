package script

import (
	"fmt"
	"net/http"
	"os"
	"searchEngineTest/model"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
	"github.com/urfave/cli/v2"
)

func MeillSearchTestAction(ctx *cli.Context) error {
	config, err := model.InitConfig(ctx)
	if err != nil {
		return err
	}
	if err = config.CheckTestRate(); err != nil {
		return err
	}

	rate := vegeta.Rate{Freq: config.TestRate.PerSecond, Per: time.Second}
	duration := time.Duration(config.TestRate.Duration) * time.Second

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
	attacker := vegeta.NewAttacker()

	fmt.Printf("start pressure test, duration:%s , rate: %d/s\n\n", duration, rate.Freq)
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	report := vegeta.NewTextReporter(&metrics)
	report.Report(os.Stdout)

	return nil
}

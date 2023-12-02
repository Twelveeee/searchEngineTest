package model

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DataFile    string       `yaml:"DataFile"`
	MeillSearch *MeillSearch `yaml:"MeillSearch"`
	Typesense   *Typesense   `yaml:"Typesense"`
	Algolia     *Algolia     `yaml:"Algolia"`
	TestRate    *TestRate    `yaml:"TestRate"`
}
type MeillSearch struct {
	Host      string `yaml:"Host"`
	APIKey    string `yaml:"APIKey"`
	IndexName string `yaml:"IndexName"`
}

type Typesense struct {
	Host      string `yaml:"Host"`
	APIKey    string `yaml:"APIKey"`
	IndexName string `yaml:"IndexName"`
}

type Algolia struct {
	ApplicationId string `yaml:"ApplicationId"`
	AdminApiKey   string `yaml:"AdminApiKey"`
	IndexName     string `yaml:"IndexName"`
}

type TestRate struct {
	PerSecond int `yaml:"PerSecond"`
	Duration  int `yaml:"Duration"`
}

func InitConfig(ctx *cli.Context) (*Config, error) {
	configFilePath := ctx.String("config")
	c := &Config{}

	if fileExists(configFilePath) {
		yamlConfig, err := os.ReadFile(configFilePath)
		if err != nil {
			return c, err
		}
		err = yaml.Unmarshal(yamlConfig, c)
		if err != nil {
			return c, err
		}

		return c, nil
	}

	return c, errors.New("config file not find")
}

func (c *Config) CheckTestRate() error {
	fmt.Println(c)
	if c.TestRate == nil {
		return errors.New("config TestRate is not set")
	}
	if c.TestRate.Duration == 0 {
		return errors.New("config TestRate.Duration is not set")
	}
	if c.TestRate.PerSecond == 0 {
		return errors.New("config TestRate.PerSecond is not set")
	}

	return nil
}

func fileExists(fileName string) bool {
	if fileName == "" {
		return false
	}

	info, err := os.Stat(fileName)

	return err == nil && !info.IsDir()
}

package yaml

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Configuration ConfigConfiguration `yaml:"configuration"`
	Rules         []ConfigRule        `yaml:"rules"`
	View          ConfigView          `yaml:"view"`
}

type ConfigConfiguration struct {
	Packages []string `yaml:"pkgs"`
}

type ConfigRule struct {
	PackageRegexps []string            `yaml:"pkg_regexps"`
	NameRegexp     string              `yaml:"name_regexp"`
	Component      ConfigRuleComponent `yaml:"component"`
}

type ConfigRuleComponent struct {
	Description string   `yaml:"description"`
	Technology  string   `yaml:"technology"`
	Tags        []string `yaml:"tags"`
}

type ConfigView struct {
	Title     string            `yaml:"title"`
	LineColor string            `yaml:"line_color"`
	Styles    []ConfigViewStyle `yaml:"styles"`
}

type ConfigViewStyle struct {
	ID              string `yaml:"id"`
	BackgroundColor string `yaml:"background_color"`
	FontColor       string `yaml:"font_color"`
	BorderColor     string `yaml:"border_color"`
}

func LoadFromFile(fileName string) (Config, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return Config{}, err
	}
	defer func() {
		_ = f.Close()
	}()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
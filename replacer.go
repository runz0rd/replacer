package replacer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"gopkg.in/yaml.v1"
)

func Replace(inputFile, outputFile, from, to string, rules []Rule) error {
	if !strings.Contains(inputFile, ".json") {
		return fmt.Errorf("only .json files supported at the moment")
	}
	inBs, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}
	var content map[string]interface{}
	json.Unmarshal(inBs, &content)
	replaced, err := traverse(content, from, to, rules)
	if err != nil {
		return err
	}
	outBs, err := json.Marshal(replaced)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outputFile, outBs, 0644)
}

func traverse(value interface{}, from, to string, rules []Rule) (interface{}, error) {
	switch v := value.(type) {
	default:
		// do replacement
		return replace(v, from, to, rules), nil
	case []interface{}:
		var replaced []interface{}
		for _, v := range v {
			r, err := traverse(v, from, to, rules)
			if err != nil {
				return nil, err
			}
			replaced = append(replaced, r)
		}
		return replaced, nil
	case map[string]interface{}:
		replaced := make(map[string]interface{})
		for k, v := range v {
			r, err := traverse(v, from, to, rules)
			if err != nil {
				return nil, err
			}
			replaced[k] = r
		}
		return replaced, nil
	}
}

func replace(value interface{}, from, to string, rules []Rule) interface{} {
	//do replacement with rules
	for _, rule := range rules {
		fromVal, fromOk := rule[from]
		toVal, toOk := rule[to]
		if !fromOk || !toOk {
			continue
		}
		var re = regexp.MustCompile(fromVal)
		value = re.ReplaceAllString(fmt.Sprint(value), strings.Replace(toVal, "(.*)", "$1", -1))
	}
	return value
}

// map[env]find(replace)
type Rule map[string]string

type Config struct {
	Rules []Rule
}

func LoadConfig(file string) (*Config, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var c Config
	if err = yaml.Unmarshal(bs, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

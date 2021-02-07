package t2y

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

func Toml2Yaml(source *os.File) (string, error) {
	sourceName := source.Name()
	sourcePrefix := sourceName[0:strings.LastIndex(sourceName, ".")]
	in, err := ioutil.ReadAll(source)
	if err != nil {
		return "", err
	}
	tomlTree, err := toml.Load(string(in))
	if err != nil {
		return "", nil
	}
	outMap := Tree2Map(tomlTree)
	outYByte, err := yaml.Marshal(outMap)
	if err != nil {
		return "", nil
	}
	os.Mkdir("./toml_yaml/", 0766)
	destName := fmt.Sprintf("%s%s%s", "./toml_yaml/", sourcePrefix, ".yaml")
	outFile, err := os.Create(destName)
	if err != nil {
		return "", nil
	}
	defer outFile.Close()
	outFile.Write(outYByte)
	return destName, nil
}

func Tree2Map(tree *toml.Tree) map[string]interface{} {
	result := make(map[string]interface{})
	keys := tree.Keys()
	for _, key := range keys {
		value := tree.Get(key)
		switch value.(type) {
		case *toml.Tree:
			result[key] = Tree2Map(value.(*toml.Tree))
		case []*toml.Tree:
			result[key] = make(map[string]interface{}, len(value.([]*toml.Tree)))
			for index, item := range value.([]*toml.Tree) {
				result[key].([]map[string]interface{})[index] = Tree2Map(item)
			}
		default:
			result[key] = &value
		}
	}
	return result
}

func Toml2Yaml2(tomlData []byte) (string, error) {
	tree, err := toml.LoadBytes(tomlData)
	if err != nil {
		return "", err
	}
	tomlMap := tree.ToMap()
	yamlData, err := yaml.Marshal(tomlMap)
	if err != nil {
		return "", err
	}
	fmt.Println(string(yamlData))
	return "", nil

}

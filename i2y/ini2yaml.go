package i2y

import (
	"file2yaml/t2y"
	"fmt"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

func Ini2Yaml(source *os.File) (string, error) {
	sourceName := source.Name()
	sourcePrefix := sourceName[0:strings.LastIndex(sourceName, ".")]
	in, err := ioutil.ReadAll(source)
	if err != nil {
		return "", err
	}
	tr, _ := toml.Load(string(in))
	outMap := t2y.Tree2Map(tr)
	outYaml, err := yaml.Marshal(outMap)
	if err != nil {
		return "", nil
	}
	os.Mkdir("./ini_yaml/", 0766)
	destName := fmt.Sprintf("%s%s%s", "./toml_ini/", sourcePrefix, ".yaml")
	outFile, err := os.Create(destName)
	if err != nil {
		return "", nil
	}
	defer outFile.Close()
	outFile.Write(outYaml)
	return destName, nil
	return "", nil
}

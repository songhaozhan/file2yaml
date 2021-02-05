package j2y

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"strings"
)

func Json2Yaml(source *os.File) (string, error) {
	sourceName := source.Name()
	sourcePrefix := sourceName[0:strings.LastIndex(sourceName, ".")]
	in, err := ioutil.ReadAll(source)
	if err != nil {
		return "", err
	}
	out, err := yaml.JSONToYAML(in)
	if err != nil {
		return "", err
	}
	os.Mkdir("./json_yaml/", 0766)
	destName := fmt.Sprintf("%s%s%s", "./json_yaml/", sourcePrefix, ".yaml")
	outFile, err := os.OpenFile(destName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	outFile.Write(out)
	return destName, nil
}

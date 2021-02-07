package conf2yaml

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type FileConvert struct {
	DataToMap func([]byte) (interface{}, error)
}

func ReadToByte(src string) ([]byte, error) {
	Data, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, err
	}
	return Data, nil
}
func JsonDataToMap(data []byte) (interface{}, error) {
	outData, err := yaml.JSONToYAML(data)
	if err != nil {
		return nil, err
	}
	return outData, nil
}
func TomlDataToMap(data []byte) (interface{}, error) {
	tomlTree, err := toml.Load(string(data))
	if err != nil {
		return nil, err
	}
	resultMap := tomlTree.ToMap()
	return resultMap, nil
}
func IniDataToMap(data []byte) (interface{}, error) {
	iniTree, err := toml.Load(string(data))
	if err != nil {
		return nil, err
	}
	resultMap := iniTree.ToMap()
	return resultMap, nil
}

func WriteYaml(ret interface{}, prefixYaml, dest string) (string, error) {
	os.Mkdir(prefixYaml, 0766)
	fileName := fmt.Sprintf("%s%s%s", prefixYaml, dest, ".yaml")
	yamlData, err := yaml.Marshal(ret)
	if err != nil {
		return "", err
	}
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	file.Write(yamlData)
	return fileName, nil
}
func File2Yaml(srcFile string) {
	if srcFile == "" {
		log.Println("error: name is nil")
		return
	}
	f, err := os.Open(srcFile)
	if err != nil {
		return
	}
	fInfo, err := f.Stat()
	if err != nil {
		return
	}
	srcName := fInfo.Name()
	f.Close()
	var fileCon = FileConvert{}
	var prefixYaml string
	fileType := srcFile[strings.LastIndex(srcName, "."):]
	switch fileType {
	case ".json":
		fileCon.DataToMap = JsonDataToMap
		prefixYaml = "./json_yaml/"
	case ".toml":
		fileCon.DataToMap = TomlDataToMap
		prefixYaml = "./toml_yaml/"
	case ".ini":
		fileCon.DataToMap = IniDataToMap
		prefixYaml = "./ini_yaml/"
	default:
		log.Println("file type not fit,only support json , toml , ini type")
		return
	}
	fileData, err := ReadToByte(srcFile)
	if err != nil {
		log.Println("read file error")
		return
	}
	ret, err := fileCon.DataToMap(fileData)
	if err != nil {
		log.Println("data to map is error")
		return
	}
	destFile, err := WriteYaml(ret, prefixYaml, srcName[:strings.LastIndex(srcName, ".")])
	if err != nil {
		return
	}
	log.Println("create file :", destFile)
}

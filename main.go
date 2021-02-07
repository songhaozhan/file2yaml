package main

import (
	"file2yaml/conf2yaml"
	"flag"
	"log"
	"os"
	"strings"
)

var (
	filepath string
)

const (
	JSON_TYPE = ".json"
	TOML_TYPE = ".toml"
	INI_TYPE  = ".ini"
)

func main() {
	flag.StringVar(&filepath, "filepath", "conf.toml", "profile name")
	flag.Parse()
	if filepath == "" {
		log.Println("error: name is nil")
		return
	}
	f, err := os.Open(filepath)
	if err != nil {
		return
	}
	fInfo, err := f.Stat()
	if err != nil {
		return
	}
	name := fInfo.Name()
	f.Close()
	var fileCon = conf2yaml.FileConvert{}
	var prefixYaml string
	fileType := name[strings.LastIndex(name, "."):]
	switch fileType {
	case JSON_TYPE:
		fileCon.DataToMap = conf2yaml.JsonDataToMap
		prefixYaml = "./json_yaml/"
	case TOML_TYPE:
		fileCon.DataToMap = conf2yaml.TomlDataToMap
		prefixYaml = "./toml_yaml/"
	case INI_TYPE:
		fileCon.DataToMap = conf2yaml.IniDataToMap
		prefixYaml = "./ini_yaml/"
	default:
		log.Println("file type not fit,only support json , toml , ini type")
		return
	}
	fileData, err := conf2yaml.ReadToByte(filepath)
	if err != nil {
		log.Println("read file error")
		return
	}
	ret, err := fileCon.DataToMap(fileData)
	if err != nil {
		log.Println("data to map is error")
		return
	}
	destFile, err := conf2yaml.WriteYaml(ret, prefixYaml, name[:strings.LastIndex(name, ".")])
	if err != nil {
		return
	}
	log.Println("create file  :", destFile)
}

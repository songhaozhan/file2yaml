package test

import (
	"file2yaml/i2y"
	"file2yaml/j2y"
	"file2yaml/t2y"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"testing"
)

func Test_json(t *testing.T) {
	file, err := os.Open("a.json")
	if err != nil {
		fmt.Println("open file is error :", err.Error())
		return
	}
	defer file.Close()
	in, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read io is error", err.Error())
		return
	}
	out, err := yaml.JSONToYAML(in)
	if err != nil {
		fmt.Println("j2y is error :", err.Error())
		return
	}
	fmt.Println(string(out))
	outFile, err := os.Create("b.yaml")
	if err != nil {
		fmt.Println("open out file is error :", err.Error())
		return
	}
	defer outFile.Close()
	outFile.Write(out)

}

func Test_Json2Yaml(t *testing.T) {
	file, _ := os.Open("")
	dest, err := j2y.Json2Yaml(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(dest)
}
func Test_Toml2Yaml(t *testing.T) {
	file, _ := os.Open("")
	dest, _ := t2y.Toml2Yaml(file)
	fmt.Println(dest)
}
func Test_Ini2Yaml(t *testing.T) {
	file, _ := os.Open("c.ini")
	defer file.Close()
	i2y.Ini2Yaml(file)
}

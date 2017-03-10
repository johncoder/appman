package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

var (
	archive  string
	manifest string
	taropts  string
)

type (
	AppManifest struct {
		Name  string
		Files []string
	}
)

func init() {
	flag.StringVar(&manifest, "manifest", "./appman.yml", "The appmain.yml file")
	flag.StringVar(&archive, "archive", "./app.tar", "The path for the archive")
	flag.StringVar(&taropts, "taropts", "cf", "The options to pass to tar")
	flag.Parse()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dump(appmanifest AppManifest) {
	fmt.Println(appmanifest.Name)

	for _, file := range appmanifest.Files {
		fmt.Println(file)
	}
}

func main() {
	data, readFileErr := ioutil.ReadFile(manifest)
	check(readFileErr)

	appmanifest := AppManifest{}

	yamlUnmarshalErr := yaml.Unmarshal(data, &appmanifest)
	check(yamlUnmarshalErr)

	files := strings.Join(appmanifest.Files, " ")
	fmt.Println("tar " + taropts + " " + archive + " " + files)
}

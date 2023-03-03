package main

import (
	_ "embed"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed operation-grouping.tmpl
var operationGroupingTemplate string

func main() {
	var inputFiles string
	flag.StringVar(&inputFiles, "input", "./tests/pb/*.proto", "请指定 proto 文件匹配目录")
	flag.Parse()

	if inputFiles == "" {
		log.Fatal("请指定 proto 文件匹配目录")
	}
	tpl, _ := template.New("result").Parse(operationGroupingTemplate)

	matches, err := filepath.Glob(inputFiles)
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range matches {
		f, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}
		if f.IsDir() {
			continue
		}
		dir := filepath.Dir(path)
		pathSplit := strings.Split(dir, "/")
		packageName := pathSplit[len(pathSplit)-1]
		if !strings.HasSuffix(strings.ToLower(f.Name()), ".proto") {
			continue
		}
		protoParseItem := ProtoParseItem{
			PackageName: packageName,
			GroupFunc:   make(map[string][]FuncParseItem, 0),
		}
		protoParseItem.Parse(path)
		outFile := dir + "/" + strings.Split(f.Name(), ".proto")[0] + ".operation.pb.go"
		file, _ := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0755)
		tpl.Execute(file, protoParseItem)
	}
}

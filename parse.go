package main

import (
	"github.com/emicklei/proto"
	"os"
	"strings"
)

const (
	tagPrefix = "@tags:"
)

type ProtoParseItem struct {
	ProtoName            string
	PackageName          string
	OperationPackageName string
	OperationServiceName string
	GroupFunc            map[string][]FuncParseItem
}

type FuncParseItem struct {
	Operation string
	GroupName string
}

func (o *ProtoParseItem) Parse(path string) {
	reader, _ := os.Open(path)
	defer reader.Close()
	parser := proto.NewParser(reader)
	definition, _ := parser.Parse()

	proto.Walk(definition,
		proto.WithPackage(o.parsePackage),
		proto.WithService(o.parseService),
		proto.WithRPC(o.parseRPC))
}

func (o *ProtoParseItem) parsePackage(p *proto.Package) {
	o.OperationPackageName = p.Name
}

func (o *ProtoParseItem) parseService(s *proto.Service) {
	o.OperationServiceName = s.Name
}

func (o *ProtoParseItem) parseRPC(r *proto.RPC) {
	funcName := "/" + o.OperationPackageName + "." + o.OperationServiceName + "/" + r.Name
	if r.Comment == nil {
		return
	}
	for i := range r.Comment.Lines {
		comment := strings.TrimSpace(r.Comment.Lines[i])
		if strings.HasPrefix(comment, tagPrefix) {
			comment = strings.TrimPrefix(comment, tagPrefix)
			groupNames := strings.Split(comment, ",")
			for _, groupName := range groupNames {
				if groupName = strings.TrimSpace(groupName); groupName == "" {
					continue
				}
				item := FuncParseItem{
					Operation: funcName,
					GroupName: groupName,
				}
				_, ok := o.GroupFunc[groupName]
				if !ok {
					o.GroupFunc[groupName] = make([]FuncParseItem, 0)
				}
				o.GroupFunc[groupName] = append(o.GroupFunc[groupName], item)
			}
		}
	}
}

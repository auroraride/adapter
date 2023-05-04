// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-04
// Based on adapter by liasica, magicrolan@qq.com.

package main

import (
	"flag"
	"log"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	var (
		schemaPath = flag.String("path", "", "path to schema directory")
	)
	flag.Parse()
	if *schemaPath == "" {
		log.Fatal("entproto: must specify schema path. use entproto -path ./ent/schema")
	}
	graph, err := entc.LoadGraph(*schemaPath, &gen.Config{})
	if err != nil {
		log.Fatalf("entproto: failed loading ent graph: %v", err)
	}
	graph.Package = "github.com/auroraride/adapter/rpc"
	if err := entproto.Generate(graph); err != nil {
		log.Fatalf("entproto: failed generating protos: %s", err)
	}
}

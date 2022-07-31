//go:build tools
// +build tools

package tools

import (
	// Tools we use during development.
	_ "github.com/golang/mock/mockgen"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "k8s.io/code-generator/cmd/client-gen"
	_ "k8s.io/code-generator/cmd/deepcopy-gen"
	_ "k8s.io/code-generator/cmd/defaulter-gen"
	_ "k8s.io/code-generator/cmd/informer-gen"
	_ "k8s.io/code-generator/cmd/lister-gen"
)

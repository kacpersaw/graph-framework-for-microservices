package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/go-openapi/spec"

	"gitlab.eng.vmware.com/nexus/compiler/pkg/openapi"
	generator "gitlab.eng.vmware.com/nexus/compiler/pkg/openapi_generator"
)

func main() {
	var yamlsPath string
	flag.StringVar(&yamlsPath, "yamls-path", "", "Path to directory containing CRD YAML definitions")
	flag.Parse()
	if yamlsPath == "" {
		panic("yamls-path is empty. Run with -h for help")
	}

	ref := func(pkg string) spec.Ref {
		r, err := spec.NewRef(strings.ToLower(pkg))
		if err != nil {
			panic(err)
		}
		return r
	}
	g, err := generator.NewGenerator(openapi.GetOpenAPIDefinitions(ref))
	if err != nil {
		panic(fmt.Sprintf("Failed creating Generator: %v", err))
	}
	err = g.ResolveRefs()
	if err != nil {
		panic(err)
	}
	if len(g.MissingDefinitions()) > 0 {
		for pkg := range g.MissingDefinitions() {
			fmt.Printf("\n***\nMissing schema for %q\n***\n", pkg)
		}
		readmePath := "https://gitlab.eng.vmware.com/nexus/compiler/blob/master/" +
			"cmd/generate-openapischema/README.md" +
			"#possible-missing-schema-error-messages-and-how-to-solve-them"
		fmt.Printf("\"openapi-gen\" did not generate all the needed schemas.\n"+
			"Refer to %q for possible causes and solutions\n", readmePath)
		panic("Missing schemas!")
	}
	err = g.UpdateYAMLs(yamlsPath)
	if err != nil {
		panic(err)
	}
}

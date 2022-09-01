package declarative

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"sigs.k8s.io/yaml"
	"strings"
	"sync"

	"github.com/getkin/kin-openapi/openapi3"
)

var supportedOperations = []string{"GET", "DELETE", "PUT"}

const NexusKindName = "x-nexus-kind-name"
const NexusGroupName = "x-nexus-group-name"
const NexusListEndpoint = "x-nexus-list-endpoint"

var (
	Paths              = make(map[string]*openapi3.PathItem)
	ApisList           = make(map[string]map[string]interface{})
	apisListMutex      = sync.Mutex{}
	Schemas            openapi3.Schemas
	parsedSchemas      = make(map[string]interface{})
	parsedSchemasMutex = sync.Mutex{}
)

func Setup() error {
	_, err := os.Stat("/openapi/openapi.yaml")
	if err == nil {
		f, err := ioutil.ReadFile("/openapi/openapi.yaml")
		if err != nil {
			return err
		}

		return Load(f)
	}
	return nil
}

func Load(data []byte) error {
	doc, err := openapi3.NewLoader().LoadFromData(data)
	if err != nil {
		return err
	}

	Schemas = doc.Components.Schemas

	for uri, pathInfo := range doc.Paths {
		if !ValidateNexusAnnotations(pathInfo) {
			continue
		}
		Paths[uri] = pathInfo
	}

	ParseSchemas()

	return nil
}

func ValidateNexusAnnotations(item *openapi3.PathItem) bool {
	for _, supportedOperation := range supportedOperations {
		op := item.GetOperation(supportedOperation)
		if op != nil {
			if GetExtensionVal(op, NexusKindName) == "" {
				return false
			}

			if GetExtensionVal(op, NexusGroupName) == "" {
				return false
			}
		}
	}

	return true
}

func GetExtensionVal(operation *openapi3.Operation, key string) string {
	val, ok := operation.ExtensionProps.Extensions[key]
	if val == nil || !ok {
		return ""
	}

	out, _ := val.(json.RawMessage).MarshalJSON()
	outStr := string(out)

	if strings.HasPrefix(outStr, `"`) && strings.HasSuffix(outStr, `"`) && len(outStr) > 2 {
		return outStr[1 : len(outStr)-1]
	}

	return outStr
}

func AddApisEndpoint(ec *EndpointContext) {
	apisListMutex.Lock()
	defer apisListMutex.Unlock()

	if ApisList[ec.Uri] == nil {
		ApisList[ec.Uri] = make(map[string]interface{})
	}

	var params []string
	for _, param := range ec.Params {
		params = append(params, param[1])
	}

	ApisList[ec.Uri][ec.Method] = map[string]interface{}{
		"group":  ec.GroupName,
		"kind":   ec.KindName,
		"params": params,
		"uri":    ec.SpecUri,
	}

	if ec.SchemaName != "" {
		ApisList[ec.Uri]["yaml"] = ConvertSchemaToYaml(ec.SchemaName, ec.GroupName, ec.KindName, params)
	}
}

func ConvertSchemaToYaml(schemaName string, group string, kind string, params []string) string {
	labels := map[string]interface{}{}
	for _, param := range params {
		labels[param] = "string"
	}

	obj := map[string]interface{}{
		"apiVersion": group + "/v1",
		"kind":       kind,
		"metadata": map[string]interface{}{
			"name":   "string",
			"labels": labels,
		},
	}
	obj["spec"] = parsedSchemas[schemaName]
	yamlObj, err := yaml.Marshal(obj)
	if err != nil {
		log.Warn(err)
	}
	return string(yamlObj)
}

func parseSchema(schemaName string, wg *sync.WaitGroup) {
	parsedSchemasMutex.Lock()
	defer func() {
		parsedSchemasMutex.Unlock()
		wg.Done()
	}()

	spec := make(map[string]interface{})

	for field, val := range Schemas[schemaName].Value.Properties {
		switch val.Value.Type {
		case "string":
			spec[field] = "string"
			if len(val.Value.Enum) > 0 {
				spec[field] = val.Value.Enum[0]
			}
		case "boolean":
			spec[field] = true
		case "number":
			spec[field] = 1.2
		case "integer":
			spec[field] = 1
		case "array":
			if val.Value.Items.Ref != "" {
				ref := openapi3.DefaultRefNameResolver(val.Value.Items.Ref)
				spec[field] = map[string]interface{}{
					"ref":  ref,
					"type": "array",
				}
			}
		case "object":
			spec[field] = "object"
		}

		if val.Ref != "" {
			ref := openapi3.DefaultRefNameResolver(val.Ref)
			spec[field] = map[string]interface{}{
				"ref": ref,
			}
		}
	}

	parsedSchemas[schemaName] = spec
}

func parseSchemaRefs(schemaName string, wg *sync.WaitGroup) {
	parsedSchemasMutex.Lock()
	defer func() {
		parsedSchemasMutex.Unlock()
		wg.Done()
	}()

	for fieldName, fieldVal := range parsedSchemas[schemaName].(map[string]interface{}) {
		if _, ok := fieldVal.(map[string]interface{}); !ok {
			continue
		}

		fv := fieldVal.(map[string]interface{})
		ref := fv["ref"]
		refType := fv["type"]

		if ref == nil {
			continue
		}

		refStr := ref.(string)

		if refType == "array" {
			parsedSchemas[schemaName].(map[string]interface{})[fieldName] = []map[string]interface{}{
				parsedSchemas[refStr].(map[string]interface{}),
			}
			continue
		}

		parsedSchemas[schemaName].(map[string]interface{})[fieldName] = parsedSchemas[refStr]
	}
}

func ParseSchemas() {
	wg := &sync.WaitGroup{}
	for schemaName := range Schemas {
		wg.Add(1)
		log.Debugf("Parsing %s schema", schemaName)
		go parseSchema(schemaName, wg)
	}
	wg.Wait()

	for schemaName := range parsedSchemas {
		wg.Add(1)
		log.Debugf("Parsing %s schema refs", schemaName)
		go parseSchemaRefs(schemaName, wg)
	}
	wg.Wait()
	log.Debugf("Finished parsing schemas")
}

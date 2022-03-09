package crd_generator_test

import (
	"go/format"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	crdgenerator "gitlab.eng.vmware.com/nexus/compiler/pkg/crd-generator"
	"gitlab.eng.vmware.com/nexus/compiler/pkg/parser"
)

const (
	baseGroupName        = "tsm.tanzu.vmware.com"
	crdModulePath        = "gitlab.eng.vmware.com/nexus/compiler/example/output/crd/"
	examplePath          = "../../example/"
	exampleDSLPath       = examplePath + "datamodel"
	exampleCRDOutputPath = examplePath + "output/crd"
	gnsExamplePath       = exampleCRDOutputPath + "/gns.tsm.tanzu.vmware.com/"
	gnsDocPath           = gnsExamplePath + "v1/doc.go"
	gnsRegisterGroupPath = gnsExamplePath + "register.go"
	gnsRegisterCRDPath   = gnsExamplePath + "v1/register.go"
	gnsTypesPath         = gnsExamplePath + "v1/types.go"
	gnsCrdBasePath       = gnsExamplePath + "gns_gns.yaml"
)

var _ = Describe("Template renderers tests", func() {
	var (
		//err error
		pkg parser.Package
		ok  bool
	)

	BeforeEach(func() {
		pkgs := parser.ParseDSLPkg(exampleDSLPath)
		pkg, ok = pkgs["gitlab.eng.vmware.com/nexus/compiler/example/datamodel//config/gns"]
		Expect(ok).To(BeTrue())
	})

	It("should parse doc template", func() {
		docBytes, err := crdgenerator.RenderDocTemplate(baseGroupName, pkg)
		Expect(err).NotTo(HaveOccurred())
		formatted, err := format.Source(docBytes.Bytes())
		Expect(err).NotTo(HaveOccurred())

		expectedDoc, err := ioutil.ReadFile(gnsDocPath)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(formatted)).To(Equal(string(expectedDoc)))
	})

	It("should parse register group template", func() {
		regBytes, err := crdgenerator.RenderRegisterGroupTemplate(baseGroupName, pkg)
		Expect(err).NotTo(HaveOccurred())
		formatted, err := format.Source(regBytes.Bytes())
		Expect(err).NotTo(HaveOccurred())

		expectedRegisterGroup, err := ioutil.ReadFile(gnsRegisterGroupPath)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(formatted)).To(Equal(string(expectedRegisterGroup)))
	})

	It("should parse register CRD template", func() {
		regBytes, err := crdgenerator.RenderRegisterCRDTemplate(crdModulePath, baseGroupName, pkg)
		Expect(err).NotTo(HaveOccurred())
		formatted, err := format.Source(regBytes.Bytes())
		Expect(err).NotTo(HaveOccurred())

		expectedRegisterCRD, err := ioutil.ReadFile(gnsRegisterCRDPath)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(formatted)).To(Equal(string(expectedRegisterCRD)))
	})

	It("should parse base crd template", func() {
		files, err := crdgenerator.RenderCRDBaseTemplate(baseGroupName, pkg)
		Expect(err).NotTo(HaveOccurred())
		Expect(files).To(HaveLen(2))

		expectedSdk, err := ioutil.ReadFile(gnsCrdBasePath)
		Expect(err).NotTo(HaveOccurred())

		Expect("gns_gns.yaml").To(Or(Equal(files[0].Name)), Equal(files[1].Name))
		Expect(string(expectedSdk)).To(Or(Equal(files[0].File.String()), Equal(files[1].File.String())))
	})

	It("should parse types template", func() {
		typesBytes, err := crdgenerator.RenderTypesTemplate(pkg)
		Expect(err).NotTo(HaveOccurred())

		formatted, err := format.Source(typesBytes.Bytes())
		Expect(err).NotTo(HaveOccurred())

		expectedTypes, err := ioutil.ReadFile(gnsTypesPath)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(formatted)).To(Equal(string(expectedTypes)))
	})
})

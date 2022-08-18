// Code generated by gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/gqlgen.git, DO NOT EDIT.

package generated

type Hello struct {
	Name      string `json:"name"`
	Secondary string `json:"secondary"`
}

func (Hello) IsEntity() {}

type HelloMultiSingleKeys struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func (HelloMultiSingleKeys) IsEntity() {}

type HelloWithErrors struct {
	Name string `json:"name"`
}

func (HelloWithErrors) IsEntity() {}

type MultiHello struct {
	Name string `json:"name"`
}

func (MultiHello) IsEntity() {}

type MultiHelloByNamesInput struct {
	Name string `json:"Name"`
}

type MultiHelloMultipleRequires struct {
	Name string `json:"name"`
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	Key3 string `json:"key3"`
}

func (MultiHelloMultipleRequires) IsEntity() {}

type MultiHelloMultipleRequiresByNamesInput struct {
	Name string `json:"Name"`
}

type MultiHelloRequires struct {
	Name string `json:"name"`
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func (MultiHelloRequires) IsEntity() {}

type MultiHelloRequiresByNamesInput struct {
	Name string `json:"Name"`
}

type MultiHelloWithError struct {
	Name string `json:"name"`
}

func (MultiHelloWithError) IsEntity() {}

type MultiHelloWithErrorByNamesInput struct {
	Name string `json:"Name"`
}

type MultiPlanetRequiresNested struct {
	Name  string `json:"name"`
	World *World `json:"world"`
	Size  int    `json:"size"`
}

func (MultiPlanetRequiresNested) IsEntity() {}

type MultiPlanetRequiresNestedByNamesInput struct {
	Name string `json:"Name"`
}

type PlanetMultipleRequires struct {
	Name     string `json:"name"`
	Diameter int    `json:"diameter"`
	Density  int    `json:"density"`
	Weight   int    `json:"weight"`
}

func (PlanetMultipleRequires) IsEntity() {}

type PlanetRequires struct {
	Name     string `json:"name"`
	Size     int    `json:"size"`
	Diameter int    `json:"diameter"`
}

func (PlanetRequires) IsEntity() {}

type PlanetRequiresNested struct {
	Name  string `json:"name"`
	World *World `json:"world"`
	Size  int    `json:"size"`
}

func (PlanetRequiresNested) IsEntity() {}

type World struct {
	Foo   string `json:"foo"`
	Bar   int    `json:"bar"`
	Hello *Hello `json:"hello"`
}

func (World) IsEntity() {}

type WorldName struct {
	Name string `json:"name"`
}

func (WorldName) IsEntity() {}

type WorldWithMultipleKeys struct {
	Foo   string `json:"foo"`
	Bar   int    `json:"bar"`
	Hello *Hello `json:"hello"`
}

func (WorldWithMultipleKeys) IsEntity() {}

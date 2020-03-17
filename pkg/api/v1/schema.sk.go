// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"sort"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewSchema(namespace, name string) *Schema {
	schema := &Schema{}
	schema.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return schema
}

func (r *Schema) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Schema) SetStatus(status core.Status) {
	r.Status = status
}

func (r *Schema) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.InlineSchema,
	)
}

type SchemaList []*Schema

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list SchemaList) Find(namespace, name string) (*Schema, error) {
	for _, schema := range list {
		if schema.GetMetadata().Name == name {
			if namespace == "" || schema.GetMetadata().Namespace == namespace {
				return schema, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find schema %v.%v", namespace, name)
}

func (list SchemaList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, schema := range list {
		ress = append(ress, schema)
	}
	return ress
}

func (list SchemaList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, schema := range list {
		ress = append(ress, schema)
	}
	return ress
}

func (list SchemaList) Names() []string {
	var names []string
	for _, schema := range list {
		names = append(names, schema.GetMetadata().Name)
	}
	return names
}

func (list SchemaList) NamespacesDotNames() []string {
	var names []string
	for _, schema := range list {
		names = append(names, schema.GetMetadata().Namespace+"."+schema.GetMetadata().Name)
	}
	return names
}

func (list SchemaList) Sort() SchemaList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list SchemaList) Clone() SchemaList {
	var schemaList SchemaList
	for _, schema := range list {
		schemaList = append(schemaList, resources.Clone(schema).(*Schema))
	}
	return schemaList
}

func (list SchemaList) Each(f func(element *Schema)) {
	for _, schema := range list {
		f(schema)
	}
}

func (list SchemaList) EachResource(f func(element resources.Resource)) {
	for _, schema := range list {
		f(schema)
	}
}

func (list SchemaList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Schema) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var _ resources.Resource = &Schema{}

// Kubernetes Adapter for Schema

func (o *Schema) GetObjectKind() schema.ObjectKind {
	t := SchemaCrd.TypeMeta()
	return &t
}

func (o *Schema) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Schema)
}

var (
	SchemaGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "sqoop.solo.io",
		Kind:    "Schema",
	}

	SchemaCrd = crd.NewCrd(
	"schemas",
	"sqoop.solo.io",
	"v1",
	"Schema",
	"sc",
	false,
	&Schema{})
)


func init() {
	if err := crd.AddCrd(SchemaCrd); err != nil {
		log.Fatalf("could not add crd to global registry")
	}
}

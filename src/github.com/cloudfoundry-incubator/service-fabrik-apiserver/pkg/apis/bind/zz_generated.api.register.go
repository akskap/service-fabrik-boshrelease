/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package bind

import (
	"fmt"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalDirectorBind = builders.NewInternalResource(
		"directorbinds",
		"DirectorBind",
		func() runtime.Object { return &DirectorBind{} },
		func() runtime.Object { return &DirectorBindList{} },
	)
	InternalDirectorBindStatus = builders.NewInternalResourceStatus(
		"directorbinds",
		"DirectorBindStatus",
		func() runtime.Object { return &DirectorBind{} },
		func() runtime.Object { return &DirectorBindList{} },
	)
	InternalDockerBind = builders.NewInternalResource(
		"dockerbinds",
		"DockerBind",
		func() runtime.Object { return &DockerBind{} },
		func() runtime.Object { return &DockerBindList{} },
	)
	InternalDockerBindStatus = builders.NewInternalResourceStatus(
		"dockerbinds",
		"DockerBindStatus",
		func() runtime.Object { return &DockerBind{} },
		func() runtime.Object { return &DockerBindList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("bind.servicefabrik.io").WithKinds(
		InternalDirectorBind,
		InternalDirectorBindStatus,
		InternalDockerBind,
		InternalDockerBindStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DirectorBind struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   DirectorBindSpec
	Status DirectorBindStatus
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DockerBind struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   DockerBindSpec
	Status DockerBindStatus
}

type DirectorBindStatus struct {
	State    string
	Response string
}

type DockerBindStatus struct {
	State    string
	Response string
}

type DockerBindSpec struct {
	Instance string
	Options  string
}

type DirectorBindSpec struct {
	Instance string
	Options  string
}

//
// DirectorBind Functions and Structs
//
// +k8s:deepcopy-gen=false
type DirectorBindStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type DirectorBindStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DirectorBindList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []DirectorBind
}

func (DirectorBind) NewStatus() interface{} {
	return DirectorBindStatus{}
}

func (pc *DirectorBind) GetStatus() interface{} {
	return pc.Status
}

func (pc *DirectorBind) SetStatus(s interface{}) {
	pc.Status = s.(DirectorBindStatus)
}

func (pc *DirectorBind) GetSpec() interface{} {
	return pc.Spec
}

func (pc *DirectorBind) SetSpec(s interface{}) {
	pc.Spec = s.(DirectorBindSpec)
}

func (pc *DirectorBind) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *DirectorBind) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc DirectorBind) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store DirectorBind.
// +k8s:deepcopy-gen=false
type DirectorBindRegistry interface {
	ListDirectorBinds(ctx request.Context, options *internalversion.ListOptions) (*DirectorBindList, error)
	GetDirectorBind(ctx request.Context, id string, options *metav1.GetOptions) (*DirectorBind, error)
	CreateDirectorBind(ctx request.Context, id *DirectorBind) (*DirectorBind, error)
	UpdateDirectorBind(ctx request.Context, id *DirectorBind) (*DirectorBind, error)
	DeleteDirectorBind(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewDirectorBindRegistry(sp builders.StandardStorageProvider) DirectorBindRegistry {
	return &storageDirectorBind{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageDirectorBind struct {
	builders.StandardStorageProvider
}

func (s *storageDirectorBind) ListDirectorBinds(ctx request.Context, options *internalversion.ListOptions) (*DirectorBindList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*DirectorBindList), err
}

func (s *storageDirectorBind) GetDirectorBind(ctx request.Context, id string, options *metav1.GetOptions) (*DirectorBind, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*DirectorBind), nil
}

func (s *storageDirectorBind) CreateDirectorBind(ctx request.Context, object *DirectorBind) (*DirectorBind, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*DirectorBind), nil
}

func (s *storageDirectorBind) UpdateDirectorBind(ctx request.Context, object *DirectorBind) (*DirectorBind, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*DirectorBind), nil
}

func (s *storageDirectorBind) DeleteDirectorBind(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}

//
// DockerBind Functions and Structs
//
// +k8s:deepcopy-gen=false
type DockerBindStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type DockerBindStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DockerBindList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []DockerBind
}

func (DockerBind) NewStatus() interface{} {
	return DockerBindStatus{}
}

func (pc *DockerBind) GetStatus() interface{} {
	return pc.Status
}

func (pc *DockerBind) SetStatus(s interface{}) {
	pc.Status = s.(DockerBindStatus)
}

func (pc *DockerBind) GetSpec() interface{} {
	return pc.Spec
}

func (pc *DockerBind) SetSpec(s interface{}) {
	pc.Spec = s.(DockerBindSpec)
}

func (pc *DockerBind) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *DockerBind) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc DockerBind) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store DockerBind.
// +k8s:deepcopy-gen=false
type DockerBindRegistry interface {
	ListDockerBinds(ctx request.Context, options *internalversion.ListOptions) (*DockerBindList, error)
	GetDockerBind(ctx request.Context, id string, options *metav1.GetOptions) (*DockerBind, error)
	CreateDockerBind(ctx request.Context, id *DockerBind) (*DockerBind, error)
	UpdateDockerBind(ctx request.Context, id *DockerBind) (*DockerBind, error)
	DeleteDockerBind(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewDockerBindRegistry(sp builders.StandardStorageProvider) DockerBindRegistry {
	return &storageDockerBind{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageDockerBind struct {
	builders.StandardStorageProvider
}

func (s *storageDockerBind) ListDockerBinds(ctx request.Context, options *internalversion.ListOptions) (*DockerBindList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*DockerBindList), err
}

func (s *storageDockerBind) GetDockerBind(ctx request.Context, id string, options *metav1.GetOptions) (*DockerBind, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*DockerBind), nil
}

func (s *storageDockerBind) CreateDockerBind(ctx request.Context, object *DockerBind) (*DockerBind, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*DockerBind), nil
}

func (s *storageDockerBind) UpdateDockerBind(ctx request.Context, object *DockerBind) (*DockerBind, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*DockerBind), nil
}

func (s *storageDockerBind) DeleteDockerBind(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}

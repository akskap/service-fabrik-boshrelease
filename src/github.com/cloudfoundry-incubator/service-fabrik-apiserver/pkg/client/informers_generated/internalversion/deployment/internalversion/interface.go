//TODO copyright header

// This file was automatically generated by informer-gen

package internalversion

import (
	internalinterfaces "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/informers_generated/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Directors returns a DirectorInformer.
	Directors() DirectorInformer
	// Dockers returns a DockerInformer.
	Dockers() DockerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Directors returns a DirectorInformer.
func (v *version) Directors() DirectorInformer {
	return &directorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Dockers returns a DockerInformer.
func (v *version) Dockers() DockerInformer {
	return &dockerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

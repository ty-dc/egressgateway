// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	egressgatewayspidernetiov1 "github.com/spidernet-io/egressgateway/pkg/k8s/apis/egressgateway.spidernet.io/v1"
	versioned "github.com/spidernet-io/egressgateway/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/spidernet-io/egressgateway/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v1 "github.com/spidernet-io/egressgateway/pkg/k8s/client/listers/egressgateway.spidernet.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EgressGatewayInformer provides access to a shared informer and lister for
// EgressGateways.
type EgressGatewayInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.EgressGatewayLister
}

type egressGatewayInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewEgressGatewayInformer constructs a new informer for EgressGateway type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEgressGatewayInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEgressGatewayInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredEgressGatewayInformer constructs a new informer for EgressGateway type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEgressGatewayInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EgressgatewayV1().EgressGateways().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EgressgatewayV1().EgressGateways().Watch(context.TODO(), options)
			},
		},
		&egressgatewayspidernetiov1.EgressGateway{},
		resyncPeriod,
		indexers,
	)
}

func (f *egressGatewayInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEgressGatewayInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *egressGatewayInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&egressgatewayspidernetiov1.EgressGateway{}, f.defaultInformer)
}

func (f *egressGatewayInformer) Lister() v1.EgressGatewayLister {
	return v1.NewEgressGatewayLister(f.Informer().GetIndexer())
}
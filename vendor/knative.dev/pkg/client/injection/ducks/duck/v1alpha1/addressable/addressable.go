/*
Copyright 2020 The Knative Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package addressable

import (
	"context"

	duck "knative.dev/pkg/apis/duck"
	v1alpha1 "knative.dev/pkg/apis/duck/v1alpha1"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	dynamicclient "knative.dev/pkg/injection/clients/dynamicclient"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterDuck(WithDuck)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func WithDuck(ctx context.Context) context.Context {
	dc := dynamicclient.Get(ctx)
	dif := &duck.CachedInformerFactory{
		Delegate: &duck.TypedInformerFactory{
			Client:       dc,
			Type:         (&v1alpha1.Addressable{}).GetFullType(),
			ResyncPeriod: controller.GetResyncPeriod(ctx),
			StopChannel:  ctx.Done(),
		},
	}
	return context.WithValue(ctx, Key{}, dif)
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) duck.InformerFactory {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch knative.dev/pkg/apis/duck.InformerFactory from context.")
	}
	return untyped.(duck.InformerFactory)
}

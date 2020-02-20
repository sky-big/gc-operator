/*
Copyright 2019 JD Cloud

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

package fake

import (
	"context"

	job "github.com/sky-big/gc-operator/pkg/client/kube/injection/informers/batch/v1/job"
	fake "github.com/sky-big/gc-operator/pkg/client/kube/injection/informers/factory/fake"
	controller "github.com/sky-big/gc-operator/pkg/common/controller"
	injection "github.com/sky-big/gc-operator/pkg/common/injection"
)

var Get = job.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Batch().V1().Jobs()
	return context.WithValue(ctx, job.Key{}, inf), inf.Informer()
}

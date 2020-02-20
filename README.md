# Kubernetes CRD Controller

## Overview

Kubernetes CRD Controller is to manage custom resources on the Kubernetes cluster.

## Quick Start

### Deploy Kubernetes CRD Controller

1. Clone the project on your Kubernetes cluster master node:
```
$ git clone https://github.com/sky-big/gc-operator
$ cd gc-operator
```

2. To deploy the Kubernetes CRD Controller on your Kubernetes cluster, please run the following script:
```
$ make install
```

3. Use command ```kubectl get pods``` to check Pulsar Operator deploy status like:
```
$ kubectl get pods
NAME                                         READY   STATUS    RESTARTS   AGE
gc-operator-86d5f5cbd7-6kq2x   1/1     Running   0          40m
```

Now you can use the CRDs provide by Kubernetes CRD Controller to deploy your Pulsar Cluster.

### Define Custom Resource

1. example [define crd1 custom resource](./pkg/apis/example/v1alpha1/crd1_types.go)

```
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRD1 struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec CRD1Spec `json:"spec,omitempty"`
	// +optional
	Status CRD1Status `json:"status,omitempty"`
}

type CRD1Spec struct {
	// +optional
	Generation int64 `json:"generation,omitempty"`

	// +optional
	V1alpha1 int64 `json:"v1alpha1,omitempty"`
}

type CRD1Status struct {
	// +optional
	Phase int64 `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRD1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CRD1 `json:"items"`
}
```

2. example [define crd2 custom resource](./pkg/apis/example/v1alpha1/crd2_types.go)

```
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRD2 struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec CRD2Spec `json:"spec,omitempty"`
	// +optional
	Status CRD2Status `json:"status,omitempty"`
}

type CRD2Spec struct {
	// +optional
	Generation int64 `json:"generation,omitempty"`

	// +optional
	V1alpha1 int64 `json:"v1alpha1,omitempty"`
}

type CRD2Status struct {
	// +optional
	Phase int64 `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CRD2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CRD2 `json:"items"`
}
```

### Create Custom Resource

1. example [create custom resource](./deploy/crds/crs.yaml):

```
apiVersion: example.dev/v1alpha1
kind: CRD1
metadata:
  name: test-crd1
spec:
  v1alpha1: 1

---

apiVersion: example.dev/v1alpha1
kind: CRD2
metadata:
  name: test-crd2
spec:
  v1alpha1: 2
```

2. create in kubernetes

```
$ kubectl apply -f deploy/crds/crs.yaml
```

3. kubernetes crd controller logs

```
{"level":"info","ts":1578450836.597406,"logger":"fallback","caller":"crd1/controller.go:25","msg":"CRD1 Controller Started"}
{"level":"info","ts":1578450836.5974965,"logger":"fallback","caller":"crd2/controller.go:27","msg":"CRD2 Controller Started"}
{"level":"info","ts":1578450836.5975108,"logger":"fallback","caller":"sharemain/main.go:50","msg":"Starting informers."}
{"level":"info","ts":1578450836.7009065,"logger":"fallback","caller":"sharemain/main.go:56","msg":"Starting controllers."}
{"level":"info","ts":1578450836.7011466,"logger":"fallback","caller":"controller/controller.go:282","msg":"Starting controller and workers"}
{"level":"info","ts":1578450836.7011998,"logger":"fallback","caller":"controller/controller.go:292","msg":"Started workers"}
{"level":"info","ts":1578450836.7011847,"logger":"fallback","caller":"controller/controller.go:282","msg":"Starting controller and workers"}
{"level":"info","ts":1578450836.7013545,"logger":"fallback","caller":"controller/controller.go:292","msg":"Started workers"}
{"level":"info","ts":1578450872.271796,"logger":"fallback","caller":"crd1/crd1.go:36","msg":"Reconcile CRD1 Resource &{TypeMeta:{Kind: APIVersion:} ObjectMeta:{Name:test-crd1 GenerateName: Namespace:default SelfLink:/apis/example.dev/v1alpha1/namespaces/default/crd1s/test-crd1 UID:7f085e03-d75f-4821-a967-775188e23e0d ResourceVersion:26823628 Generation:1 CreationTimestamp:2020-01-08 02:34:32 +0000 UTC DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[kubectl.kubernetes.io/last-applied-configuration:{\"apiVersion\":\"example.dev/v1alpha1\",\"kind\":\"CRD1\",\"metadata\":{\"annotations\":{},\"name\":\"test-crd1\",\"namespace\":\"default\"},\"spec\":{\"v1alpha1\":1}}\n] OwnerReferences:[] Finalizers:[] ClusterName: ManagedFields:[]} Spec:{Generation:0 V1alpha1:1} Status:{Phase:0}}","traceid":"b721879e-2518-465a-aedc-8efb57fe21d8","key":"default/test-crd1"}
{"level":"info","ts":1578450872.2721462,"logger":"fallback","caller":"controller/controller.go:338","msg":"Reconcile succeeded. Time taken: 652.085µs.","traceid":"b721879e-2518-465a-aedc-8efb57fe21d8","key":"default/test-crd1"}
{"level":"info","ts":1578450872.2769606,"logger":"fallback","caller":"crd2/crd2.go:36","msg":"Reconcile CRD2 Resource &{TypeMeta:{Kind: APIVersion:} ObjectMeta:{Name:test-crd2 GenerateName: Namespace:default SelfLink:/apis/example.dev/v1alpha1/namespaces/default/crd2s/test-crd2 UID:2982e467-14d8-4147-9922-e6509f427aea ResourceVersion:26823629 Generation:1 CreationTimestamp:2020-01-08 02:34:32 +0000 UTC DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[kubectl.kubernetes.io/last-applied-configuration:{\"apiVersion\":\"example.dev/v1alpha1\",\"kind\":\"CRD2\",\"metadata\":{\"annotations\":{},\"name\":\"test-crd2\",\"namespace\":\"default\"},\"spec\":{\"v1alpha1\":2}}\n] OwnerReferences:[] Finalizers:[] ClusterName: ManagedFields:[]} Spec:{Generation:0 V1alpha1:2} Status:{Phase:0}}","traceid":"f4dbd50f-bd97-45b8-80c1-c36d4ef6f890","key":"default/test-crd2"}
{"level":"info","ts":1578450872.2770426,"logger":"fallback","caller":"controller/controller.go:338","msg":"Reconcile succeeded. Time taken: 246.213µs.","traceid":"f4dbd50f-bd97-45b8-80c1-c36d4ef6f890","key":"default/test-crd2"}
```

## Generate Custom Resource SDK

```
$ make generate
```

## Generate Image

```
$ make image
```

## Push Image

```
$ make push
```

## Generate Project Vendor

```
$ make vendor
```

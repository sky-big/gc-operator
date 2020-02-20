# Kubernetes CRD Controller

## Overview

GC Operator test pod gc

## Quick Start

### Deploy GC Operator

1. Clone the project on your Kubernetes cluster master node:

```
$ git clone https://github.com/sky-big/gc-operator
$ cd gc-operator
```

2. To deploy the GC Operator on your Kubernetes cluster, please run the following script:

```
$ make install
```

3. Use command ```kubectl get pods``` to check GC Operator deploy status like:

```
$ kubectl get pods
NAME                                         READY   STATUS    RESTARTS   AGE
gc-operator-86d5f5cbd7-6kq2x   1/1     Running   0          40m
```

4. uninstall

```
$ make uninstall
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

package pod

import (
	"context"
	"fmt"

	"github.com/sky-big/gc-operator/pkg/common/logging"

	"go.uber.org/zap"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/informers/core/v1"
	listerv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	// ReconcilerName is the name of the reconciler
	ReconcilerName = "PodReconciler"
)

const (
	NodeFinalizerName = "finalizers.jdcloud.com/node"
	NetworkFinalizerName = "finalizers.jdcloud.com/network"
)

type Reconciler struct {
	podInformer v1.PodInformer
	podLister   listerv1.PodLister
	client kubernetes.Interface
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	logger := logging.FromContext(ctx)

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		logger.Errorw("Invalid resource key", zap.Error(err))
		return nil
	}

	if name != "my-test-0" {
		return nil
	}

	// Get pod resource with this namespace/name
	original, err := c.podLister.Pods(namespace).Get(name)
	if err != nil {
		return nil
	}
	p := original.DeepCopy()

	logger.Infof("Reconcile Pod(%s) Resource Finalizers : %v", p.Name, p.Finalizers)

	if original.GetDeletionTimestamp().IsZero() {
		// 1. 如果Pod没有被调度且该Pod对应的Node节点没有被创建则创建该节点
		// 同时创建该节点的时候给该node打上该pod才能容忍的污点信息

		// 2. 给pod添加容忍该node的容忍信息

		// 3. 给pod添加上要垃圾回收node的信息，用来在pod要删除的时候能够有钩子来删除node节点
		if !containsString(p.Finalizers, NodeFinalizerName) {
			p.Finalizers = append(p.Finalizers, NodeFinalizerName)
			if _, err := c.client.CoreV1().Pods(namespace).Update(p); err != nil {
				return fmt.Errorf("pod(%s) add node finalizer error %v", p.Name, err)
			} else {
				logger.Infof("pod(%s) add node finalizer success", p.Name)
				return nil
			}
		}

		// 4. 如果需要则给该node插上弹性网卡

		// 5. 给pod添加上要垃圾回收弹性网卡的信息，用来在pod要删除的时候能够有钩子来删除弹性网卡
		if !containsString(p.Finalizers, NetworkFinalizerName) {
			p.Finalizers = append(p.Finalizers, NetworkFinalizerName)
			if _, err := c.client.CoreV1().Pods(namespace).Update(p); err != nil {
				return fmt.Errorf("pod(%s) add network finalizer error %v", p.Name, err)
			} else {
				logger.Infof("pod(%s) add network finalizer success", p.Name)
				return nil
			}
		}
	} else {
		// 1. 删除该pod对应的node

		// 2. pod将要被删除，在此处根据node垃圾回收信息，将node资源回收
		if containsString(p.Finalizers, NodeFinalizerName) {
			p.Finalizers = removeString(p.Finalizers, NodeFinalizerName)
			if _, err := c.client.CoreV1().Pods(namespace).Update(p); err != nil {
				return fmt.Errorf("pod(%s) delete node finalizer error %v", p.Name, err)
			} else {
				logger.Infof("pod(%s) delete node finalizer success", p.Name)
				return nil
			}
		}

		// 3. 删除该pod对应的弹性网卡

		// 4. pod将要被删除，在此处根据弹性网卡回收信息，将弹性网卡回收
		if containsString(p.Finalizers, NetworkFinalizerName) {
			p.Finalizers = removeString(p.Finalizers, NetworkFinalizerName)
			if _, err := c.client.CoreV1().Pods(namespace).Update(p); err != nil {
				return fmt.Errorf("pod(%s) delete network finalizer error %v", p.Name, err)
			} else {
				logger.Infof("pod(%s) delete network finalizer success", p.Name)
				return nil
			}
		}
	}
	return nil
}

func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
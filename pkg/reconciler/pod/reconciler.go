package pod

import (
	"context"

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

	// Get pod resource with this namespace/name
	original, err := c.podLister.Pods(namespace).Get(name)

	logger.Infof("Reconcile Pod Resource %+v", original)

	if original.GetDeletionTimestamp() == nil {
		// 1. 如果Pod没有被调度且该Pod对应的Node节点没有被创建则创建该节点
		// 同时创建该节点的时候给该node打上该pod才能容忍的污点信息

		// 2. 给pod添加容忍该node的容忍信息

		// 3. 给pod添加上要垃圾回收node的信息，用来在pod要删除的时候能够有钩子来删除node节点

		// 4. 如果需要则给该node插上弹性网卡

		// 5. 给pod添加上要垃圾回收弹性网卡的信息，用来在pod要删除的时候能够有钩子来删除弹性网卡
	} else {
		// 1. pod将要被删除，在此处根据node垃圾回收信息，将node资源回收

		// 2. pod将要被删除，在此处根据弹性网卡回收信息，将弹性网卡回收
	}
	return nil
}

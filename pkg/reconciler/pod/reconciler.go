package pod

import (
	"context"

	"github.com/sky-big/gc-operator/pkg/common/logging"

	"go.uber.org/zap"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/informers/core/v1"
	listerv1 "k8s.io/client-go/listers/core/v1"
)

const (
	// ReconcilerName is the name of the reconciler
	ReconcilerName = "PodReconciler"
)

type Reconciler struct {
	podInformer v1.PodInformer
	podLister   listerv1.PodLister
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
	return nil
}

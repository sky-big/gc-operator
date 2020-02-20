package pod

import (
	"context"

	"github.com/sky-big/gc-operator/pkg/client/kube/injection/informers/core/v1/pod"
	"github.com/sky-big/gc-operator/pkg/client/kube/injection/client"
	"github.com/sky-big/gc-operator/pkg/common/controller"
	"github.com/sky-big/gc-operator/pkg/common/logging"
)

func NewController(
	ctx context.Context,
) *controller.Impl {
	logger := logging.FromContext(ctx)
	podInformer := pod.Get(ctx)

	c := &Reconciler{
		podInformer: podInformer,
		podLister:   podInformer.Lister(),
		client: client.Get(ctx),
	}
	impl := controller.NewImpl(c, logger, ReconcilerName)

	podInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	logger.Info("Pod Controller Started")
	return impl
}

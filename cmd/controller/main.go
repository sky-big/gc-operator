package main

import (
	"github.com/sky-big/gc-operator/pkg/common/sharemain"
	"github.com/sky-big/gc-operator/pkg/reconciler/pod"
)

func main() {
	sharemain.Main(
		pod.NewController,
	)
}

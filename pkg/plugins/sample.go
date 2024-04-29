package plugins

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/klog"

	// "k8s.io/kubernetes/pkg/scheduler/framework"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"k8s.io/kubernetes/pkg/scheduler/nodeinfo"

	"k8s.io/apimachinery/pkg/runtime"
)

const Name string = "sample"

type SampleScheduler struct{}

func New(configuration *runtime.Unknown, f framework.FrameworkHandle) (framework.Plugin, error) {
	return &SampleScheduler{}, nil
}

func (s *SampleScheduler) Name() string {
	return Name
}

func (s *SampleScheduler) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *nodeinfo.NodeInfo) *framework.Status {
	klog.V(3).Info("<ddd...ddd...>")

	if v, ok := nodeInfo.Node().Labels["sample"]; ok && v == "true" {
		return framework.NewStatus(framework.Success, "<ok filter label>")
	}

	return framework.NewStatus(framework.Unschedulable, "<no filter label>")

	// return framework.NewStatus(framework.Unschedulable, "<no filter label>")
}

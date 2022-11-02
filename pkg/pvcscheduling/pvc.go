package pvcscheduling

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "pvc-plugin"

type PvcScheduler struct {
	handle framework.Handle
}

func (p *PvcScheduler) Name() string {
	return Name
}

func (p *PvcScheduler) Filter(state *framework.CycleState, pod *v1.Pod, node *v1.Node) *framework.Status {
	klog.V(3).Infof("filter pod: %v,node: %v", pod.Name, node.Name)
	return framework.NewStatus(framework.Success, "")
}

func (p *PvcScheduler) Bind(state *framework.CycleState, pod *v1.Pod, node *v1.Node) *framework.Status {
	klog.V(3).Infof("filter pod: %v,node: %v", pod.Name, node.Name)
	return framework.NewStatus(framework.Success, "")
}

// New initializes a new plugin and returns it.
func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &PvcScheduler{handle: h}, nil
}

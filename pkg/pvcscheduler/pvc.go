package pvcscheduler

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"sigs.k8s.io/scheduler-plugins/apis/config"
)

type Args struct {
	EnvType string `json:"env_type,omitempty"`
	EnvName string `json:"env_name,omitempty"`
}
type PvcScheduler struct {
	args   *Args
	handle framework.Handle
}

var _ = framework.PostBindPlugin(&PvcScheduler{})
var _ = framework.PreBindPlugin(&PvcScheduler{})

const Name = "GinoPvc"

func (p *PvcScheduler) Name() string {
	return Name
}

func (ps *PvcScheduler) PreBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).InfoS("PostBind", "nodename", nodeName)
	if nodeInfo, err := ps.handle.SnapshotSharedLister().NodeInfos().Get(nodeName); err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("prebind get node info error: %+v", nodeName))
	} else {
		klog.V(3).Infof("prebind node info: %+v", nodeInfo.Node())
		return framework.NewStatus(framework.Success, "")
	}
}

func (ps *PvcScheduler) PostBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) {
	klog.V(3).InfoS("PostBind", "pod", klog.KObj(pod))
}

func (ps *PvcScheduler) Filter(state *framework.CycleState, pod *v1.Pod, node *v1.Node) *framework.Status {
	klog.V(3).Infof("filter pod: %v,node: %v", pod.Name, node.Name)
	return framework.NewStatus(framework.Success, "")
}

func (ps *PvcScheduler) Bind(state *framework.CycleState, pod *v1.Pod, node *v1.Node) *framework.Status {
	klog.V(3).Infof("filter pod: %v,node: %v", pod.Name, node.Name)
	return framework.NewStatus(framework.Success, "")
}

// New initializes a new plugin and returns it.
func New(configuration runtime.Object, h framework.Handle) (framework.Plugin, error) {
	a := configuration.(*config.GinoPvcArgs)
	args := &Args{}
	if configuration == nil {
		return nil, errors.Errorf("No config is found")
	}
	klog.V(3).Infof("get plugin config args: %+v", a)
	return &PvcScheduler{args: args, handle: h}, nil
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Cluster Alert Group resource. This can be used to create Cluster Alert Group for Rancher v2 environments and retrieve their information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/cluster_alert_group.html.markdown.
type ClusterAlterGroup struct {
	s *pulumi.ResourceState
}

// NewClusterAlterGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterAlterGroup(ctx *pulumi.Context,
	name string, args *ClusterAlterGroupArgs, opts ...pulumi.ResourceOpt) (*ClusterAlterGroup, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["annotations"] = nil
		inputs["clusterId"] = nil
		inputs["description"] = nil
		inputs["groupIntervalSeconds"] = nil
		inputs["groupWaitSeconds"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["recipients"] = nil
		inputs["repeatIntervalSeconds"] = nil
	} else {
		inputs["annotations"] = args.Annotations
		inputs["clusterId"] = args.ClusterId
		inputs["description"] = args.Description
		inputs["groupIntervalSeconds"] = args.GroupIntervalSeconds
		inputs["groupWaitSeconds"] = args.GroupWaitSeconds
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["recipients"] = args.Recipients
		inputs["repeatIntervalSeconds"] = args.RepeatIntervalSeconds
	}
	s, err := ctx.RegisterResource("rancher2:index/clusterAlterGroup:ClusterAlterGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterAlterGroup{s: s}, nil
}

// GetClusterAlterGroup gets an existing ClusterAlterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterAlterGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterAlterGroupState, opts ...pulumi.ResourceOpt) (*ClusterAlterGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["annotations"] = state.Annotations
		inputs["clusterId"] = state.ClusterId
		inputs["description"] = state.Description
		inputs["groupIntervalSeconds"] = state.GroupIntervalSeconds
		inputs["groupWaitSeconds"] = state.GroupWaitSeconds
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["recipients"] = state.Recipients
		inputs["repeatIntervalSeconds"] = state.RepeatIntervalSeconds
	}
	s, err := ctx.ReadResource("rancher2:index/clusterAlterGroup:ClusterAlterGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterAlterGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ClusterAlterGroup) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ClusterAlterGroup) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The cluster alert group annotations (map)
func (r *ClusterAlterGroup) Annotations() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["annotations"])
}

// The cluster id where create cluster alert group (string)
func (r *ClusterAlterGroup) ClusterId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clusterId"])
}

// The cluster alert group description (string)
func (r *ClusterAlterGroup) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The cluster alert group interval seconds. Default: `180` (int)
func (r *ClusterAlterGroup) GroupIntervalSeconds() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["groupIntervalSeconds"])
}

// The cluster alert group wait seconds. Default: `180` (int)
func (r *ClusterAlterGroup) GroupWaitSeconds() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["groupWaitSeconds"])
}

// The cluster alert group labels (map)
func (r *ClusterAlterGroup) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The cluster alert group name (string)
func (r *ClusterAlterGroup) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The cluster alert group recipients (list)
func (r *ClusterAlterGroup) Recipients() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["recipients"])
}

// The cluster alert group wait seconds. Default: `3600` (int)
func (r *ClusterAlterGroup) RepeatIntervalSeconds() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["repeatIntervalSeconds"])
}

// Input properties used for looking up and filtering ClusterAlterGroup resources.
type ClusterAlterGroupState struct {
	// The cluster alert group annotations (map)
	Annotations interface{}
	// The cluster id where create cluster alert group (string)
	ClusterId interface{}
	// The cluster alert group description (string)
	Description interface{}
	// The cluster alert group interval seconds. Default: `180` (int)
	GroupIntervalSeconds interface{}
	// The cluster alert group wait seconds. Default: `180` (int)
	GroupWaitSeconds interface{}
	// The cluster alert group labels (map)
	Labels interface{}
	// The cluster alert group name (string)
	Name interface{}
	// The cluster alert group recipients (list)
	Recipients interface{}
	// The cluster alert group wait seconds. Default: `3600` (int)
	RepeatIntervalSeconds interface{}
}

// The set of arguments for constructing a ClusterAlterGroup resource.
type ClusterAlterGroupArgs struct {
	// The cluster alert group annotations (map)
	Annotations interface{}
	// The cluster id where create cluster alert group (string)
	ClusterId interface{}
	// The cluster alert group description (string)
	Description interface{}
	// The cluster alert group interval seconds. Default: `180` (int)
	GroupIntervalSeconds interface{}
	// The cluster alert group wait seconds. Default: `180` (int)
	GroupWaitSeconds interface{}
	// The cluster alert group labels (map)
	Labels interface{}
	// The cluster alert group name (string)
	Name interface{}
	// The cluster alert group recipients (list)
	Recipients interface{}
	// The cluster alert group wait seconds. Default: `3600` (int)
	RepeatIntervalSeconds interface{}
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Cluster Alert Rule resource. This can be used to create Cluster Alert Rule for Rancher v2 environments and retrieve their information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/cluster_alert_rule.html.markdown.
type ClusterAlterRule struct {
	s *pulumi.ResourceState
}

// NewClusterAlterRule registers a new resource with the given unique name, arguments, and options.
func NewClusterAlterRule(ctx *pulumi.Context,
	name string, args *ClusterAlterRuleArgs, opts ...pulumi.ResourceOpt) (*ClusterAlterRule, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["annotations"] = nil
		inputs["clusterId"] = nil
		inputs["eventRule"] = nil
		inputs["groupId"] = nil
		inputs["groupIntervalSeconds"] = nil
		inputs["groupWaitSeconds"] = nil
		inputs["inherited"] = nil
		inputs["labels"] = nil
		inputs["metricRule"] = nil
		inputs["name"] = nil
		inputs["nodeRule"] = nil
		inputs["repeatIntervalSeconds"] = nil
		inputs["severity"] = nil
		inputs["systemServiceRule"] = nil
	} else {
		inputs["annotations"] = args.Annotations
		inputs["clusterId"] = args.ClusterId
		inputs["eventRule"] = args.EventRule
		inputs["groupId"] = args.GroupId
		inputs["groupIntervalSeconds"] = args.GroupIntervalSeconds
		inputs["groupWaitSeconds"] = args.GroupWaitSeconds
		inputs["inherited"] = args.Inherited
		inputs["labels"] = args.Labels
		inputs["metricRule"] = args.MetricRule
		inputs["name"] = args.Name
		inputs["nodeRule"] = args.NodeRule
		inputs["repeatIntervalSeconds"] = args.RepeatIntervalSeconds
		inputs["severity"] = args.Severity
		inputs["systemServiceRule"] = args.SystemServiceRule
	}
	s, err := ctx.RegisterResource("rancher2:index/clusterAlterRule:ClusterAlterRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterAlterRule{s: s}, nil
}

// GetClusterAlterRule gets an existing ClusterAlterRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterAlterRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterAlterRuleState, opts ...pulumi.ResourceOpt) (*ClusterAlterRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["annotations"] = state.Annotations
		inputs["clusterId"] = state.ClusterId
		inputs["eventRule"] = state.EventRule
		inputs["groupId"] = state.GroupId
		inputs["groupIntervalSeconds"] = state.GroupIntervalSeconds
		inputs["groupWaitSeconds"] = state.GroupWaitSeconds
		inputs["inherited"] = state.Inherited
		inputs["labels"] = state.Labels
		inputs["metricRule"] = state.MetricRule
		inputs["name"] = state.Name
		inputs["nodeRule"] = state.NodeRule
		inputs["repeatIntervalSeconds"] = state.RepeatIntervalSeconds
		inputs["severity"] = state.Severity
		inputs["systemServiceRule"] = state.SystemServiceRule
	}
	s, err := ctx.ReadResource("rancher2:index/clusterAlterRule:ClusterAlterRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterAlterRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ClusterAlterRule) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ClusterAlterRule) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The cluster alert rule annotations (map)
func (r *ClusterAlterRule) Annotations() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["annotations"])
}

// The cluster id where create cluster alert rule (string)
func (r *ClusterAlterRule) ClusterId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clusterId"])
}

// The cluster alert rule event rule. ConflictsWith: `"metricRule", "nodeRule", "systemServiceRule"`` (list Maxitems:1)
func (r *ClusterAlterRule) EventRule() pulumi.Output {
	return r.s.State["eventRule"]
}

// The cluster alert rule alert group ID (string)
func (r *ClusterAlterRule) GroupId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["groupId"])
}

// The cluster alert rule group interval seconds. Default: `180` (int)
func (r *ClusterAlterRule) GroupIntervalSeconds() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["groupIntervalSeconds"])
}

// The cluster alert rule group wait seconds. Default: `180` (int)
func (r *ClusterAlterRule) GroupWaitSeconds() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["groupWaitSeconds"])
}

// The cluster alert rule inherited. Default: `true` (bool)
func (r *ClusterAlterRule) Inherited() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["inherited"])
}

// The cluster alert rule labels (map)
func (r *ClusterAlterRule) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The cluster alert rule metric rule. ConflictsWith: `"eventRule", "nodeRule", "systemServiceRule"`` (list Maxitems:1)
func (r *ClusterAlterRule) MetricRule() pulumi.Output {
	return r.s.State["metricRule"]
}

// The cluster alert rule name (string)
func (r *ClusterAlterRule) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The cluster alert rule node rule. ConflictsWith: `"eventRule", "metricRule", "systemServiceRule"`` (list Maxitems:1)
func (r *ClusterAlterRule) NodeRule() pulumi.Output {
	return r.s.State["nodeRule"]
}

// The cluster alert rule wait seconds. Default: `3600` (int)
func (r *ClusterAlterRule) RepeatIntervalSeconds() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["repeatIntervalSeconds"])
}

// The cluster alert rule severity. Supported values : `"critical" | "info" | "warning"`. Default: `critical` (string)
func (r *ClusterAlterRule) Severity() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["severity"])
}

// The cluster alert rule system service rule. ConflictsWith: `"eventRule", "metricRule", "nodeRule"`` (list Maxitems:1)
func (r *ClusterAlterRule) SystemServiceRule() pulumi.Output {
	return r.s.State["systemServiceRule"]
}

// Input properties used for looking up and filtering ClusterAlterRule resources.
type ClusterAlterRuleState struct {
	// The cluster alert rule annotations (map)
	Annotations interface{}
	// The cluster id where create cluster alert rule (string)
	ClusterId interface{}
	// The cluster alert rule event rule. ConflictsWith: `"metricRule", "nodeRule", "systemServiceRule"`` (list Maxitems:1)
	EventRule interface{}
	// The cluster alert rule alert group ID (string)
	GroupId interface{}
	// The cluster alert rule group interval seconds. Default: `180` (int)
	GroupIntervalSeconds interface{}
	// The cluster alert rule group wait seconds. Default: `180` (int)
	GroupWaitSeconds interface{}
	// The cluster alert rule inherited. Default: `true` (bool)
	Inherited interface{}
	// The cluster alert rule labels (map)
	Labels interface{}
	// The cluster alert rule metric rule. ConflictsWith: `"eventRule", "nodeRule", "systemServiceRule"`` (list Maxitems:1)
	MetricRule interface{}
	// The cluster alert rule name (string)
	Name interface{}
	// The cluster alert rule node rule. ConflictsWith: `"eventRule", "metricRule", "systemServiceRule"`` (list Maxitems:1)
	NodeRule interface{}
	// The cluster alert rule wait seconds. Default: `3600` (int)
	RepeatIntervalSeconds interface{}
	// The cluster alert rule severity. Supported values : `"critical" | "info" | "warning"`. Default: `critical` (string)
	Severity interface{}
	// The cluster alert rule system service rule. ConflictsWith: `"eventRule", "metricRule", "nodeRule"`` (list Maxitems:1)
	SystemServiceRule interface{}
}

// The set of arguments for constructing a ClusterAlterRule resource.
type ClusterAlterRuleArgs struct {
	// The cluster alert rule annotations (map)
	Annotations interface{}
	// The cluster id where create cluster alert rule (string)
	ClusterId interface{}
	// The cluster alert rule event rule. ConflictsWith: `"metricRule", "nodeRule", "systemServiceRule"`` (list Maxitems:1)
	EventRule interface{}
	// The cluster alert rule alert group ID (string)
	GroupId interface{}
	// The cluster alert rule group interval seconds. Default: `180` (int)
	GroupIntervalSeconds interface{}
	// The cluster alert rule group wait seconds. Default: `180` (int)
	GroupWaitSeconds interface{}
	// The cluster alert rule inherited. Default: `true` (bool)
	Inherited interface{}
	// The cluster alert rule labels (map)
	Labels interface{}
	// The cluster alert rule metric rule. ConflictsWith: `"eventRule", "nodeRule", "systemServiceRule"`` (list Maxitems:1)
	MetricRule interface{}
	// The cluster alert rule name (string)
	Name interface{}
	// The cluster alert rule node rule. ConflictsWith: `"eventRule", "metricRule", "systemServiceRule"`` (list Maxitems:1)
	NodeRule interface{}
	// The cluster alert rule wait seconds. Default: `3600` (int)
	RepeatIntervalSeconds interface{}
	// The cluster alert rule severity. Supported values : `"critical" | "info" | "warning"`. Default: `critical` (string)
	Severity interface{}
	// The cluster alert rule system service rule. ConflictsWith: `"eventRule", "metricRule", "nodeRule"`` (list Maxitems:1)
	SystemServiceRule interface{}
}
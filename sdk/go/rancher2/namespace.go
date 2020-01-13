// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Namespace resource. This can be used to create namespaces for Rancher v2 environments and retrieve their information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/namespace.html.markdown.
type Namespace struct {
	s *pulumi.ResourceState
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOpt) (*Namespace, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["annotations"] = nil
		inputs["containerResourceLimit"] = nil
		inputs["description"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["projectId"] = nil
		inputs["resourceQuota"] = nil
		inputs["waitForCluster"] = nil
	} else {
		inputs["annotations"] = args.Annotations
		inputs["containerResourceLimit"] = args.ContainerResourceLimit
		inputs["description"] = args.Description
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["projectId"] = args.ProjectId
		inputs["resourceQuota"] = args.ResourceQuota
		inputs["waitForCluster"] = args.WaitForCluster
	}
	s, err := ctx.RegisterResource("rancher2:index/namespace:Namespace", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Namespace{s: s}, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NamespaceState, opts ...pulumi.ResourceOpt) (*Namespace, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["annotations"] = state.Annotations
		inputs["containerResourceLimit"] = state.ContainerResourceLimit
		inputs["description"] = state.Description
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["projectId"] = state.ProjectId
		inputs["resourceQuota"] = state.ResourceQuota
		inputs["waitForCluster"] = state.WaitForCluster
	}
	s, err := ctx.ReadResource("rancher2:index/namespace:Namespace", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Namespace{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Namespace) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Namespace) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Annotations for Node Pool object (map)
func (r *Namespace) Annotations() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["annotations"])
}

// Default containers resource limits on namespace (List maxitem:1)
func (r *Namespace) ContainerResourceLimit() pulumi.Output {
	return r.s.State["containerResourceLimit"]
}

// A namespace description (string)
func (r *Namespace) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Labels for Node Pool object (map)
func (r *Namespace) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The name of the namespace (string)
func (r *Namespace) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The project id where assign namespace. It's on the form `project_id=<cluster_id>:<id>`. Updating `<id>` part on same `<cluster_id>` namespace will be moved between projects (string)
func (r *Namespace) ProjectId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["projectId"])
}

// Resource quota for namespace. Rancher v2.1.x or higher (list maxitems:1)
func (r *Namespace) ResourceQuota() pulumi.Output {
	return r.s.State["resourceQuota"]
}

// Wait for cluster becomes active. Default `false` (bool)
func (r *Namespace) WaitForCluster() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["waitForCluster"])
}

// Input properties used for looking up and filtering Namespace resources.
type NamespaceState struct {
	// Annotations for Node Pool object (map)
	Annotations interface{}
	// Default containers resource limits on namespace (List maxitem:1)
	ContainerResourceLimit interface{}
	// A namespace description (string)
	Description interface{}
	// Labels for Node Pool object (map)
	Labels interface{}
	// The name of the namespace (string)
	Name interface{}
	// The project id where assign namespace. It's on the form `project_id=<cluster_id>:<id>`. Updating `<id>` part on same `<cluster_id>` namespace will be moved between projects (string)
	ProjectId interface{}
	// Resource quota for namespace. Rancher v2.1.x or higher (list maxitems:1)
	ResourceQuota interface{}
	// Wait for cluster becomes active. Default `false` (bool)
	WaitForCluster interface{}
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Annotations for Node Pool object (map)
	Annotations interface{}
	// Default containers resource limits on namespace (List maxitem:1)
	ContainerResourceLimit interface{}
	// A namespace description (string)
	Description interface{}
	// Labels for Node Pool object (map)
	Labels interface{}
	// The name of the namespace (string)
	Name interface{}
	// The project id where assign namespace. It's on the form `project_id=<cluster_id>:<id>`. Updating `<id>` part on same `<cluster_id>` namespace will be moved between projects (string)
	ProjectId interface{}
	// Resource quota for namespace. Rancher v2.1.x or higher (list maxitems:1)
	ResourceQuota interface{}
	// Wait for cluster becomes active. Default `false` (bool)
	WaitForCluster interface{}
}

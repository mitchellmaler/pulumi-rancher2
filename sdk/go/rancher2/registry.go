// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Registry resource. This can be used to create docker registries for Rancher v2 environments and retrieve their information.
// 
// Depending of the availability, there are 2 types of Rancher v2 docker registries:
// - Project registry: Available to all namespaces in the `projectId`
// - Namespaced regitry: Available to just `namespaceId` in the `projectId`
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/registry.html.markdown.
type Registry struct {
	s *pulumi.ResourceState
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOpt) (*Registry, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil || args.Registries == nil {
		return nil, errors.New("missing required argument 'Registries'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["annotations"] = nil
		inputs["description"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["namespaceId"] = nil
		inputs["projectId"] = nil
		inputs["registries"] = nil
	} else {
		inputs["annotations"] = args.Annotations
		inputs["description"] = args.Description
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["namespaceId"] = args.NamespaceId
		inputs["projectId"] = args.ProjectId
		inputs["registries"] = args.Registries
	}
	s, err := ctx.RegisterResource("rancher2:index/registry:Registry", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Registry{s: s}, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegistryState, opts ...pulumi.ResourceOpt) (*Registry, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["annotations"] = state.Annotations
		inputs["description"] = state.Description
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["namespaceId"] = state.NamespaceId
		inputs["projectId"] = state.ProjectId
		inputs["registries"] = state.Registries
	}
	s, err := ctx.ReadResource("rancher2:index/registry:Registry", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Registry{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Registry) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Registry) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Annotations for Registry object (map)
func (r *Registry) Annotations() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["annotations"])
}

// A registry description (string)
func (r *Registry) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Labels for Registry object (map)
func (r *Registry) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The name of the registry (string)
func (r *Registry) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The namespace id where to assign the namespaced registry (string)
func (r *Registry) NamespaceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namespaceId"])
}

// The project id where to assign the registry (string)
func (r *Registry) ProjectId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["projectId"])
}

// Registries data for registry (list)
func (r *Registry) Registries() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["registries"])
}

// Input properties used for looking up and filtering Registry resources.
type RegistryState struct {
	// Annotations for Registry object (map)
	Annotations interface{}
	// A registry description (string)
	Description interface{}
	// Labels for Registry object (map)
	Labels interface{}
	// The name of the registry (string)
	Name interface{}
	// The namespace id where to assign the namespaced registry (string)
	NamespaceId interface{}
	// The project id where to assign the registry (string)
	ProjectId interface{}
	// Registries data for registry (list)
	Registries interface{}
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// Annotations for Registry object (map)
	Annotations interface{}
	// A registry description (string)
	Description interface{}
	// Labels for Registry object (map)
	Labels interface{}
	// The name of the registry (string)
	Name interface{}
	// The namespace id where to assign the namespaced registry (string)
	NamespaceId interface{}
	// The project id where to assign the registry (string)
	ProjectId interface{}
	// Registries data for registry (list)
	Registries interface{}
}
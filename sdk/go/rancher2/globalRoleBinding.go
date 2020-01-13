// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Global Role Binding resource. This can be used to create Global Role Bindings for Rancher v2 environments and retrieve their information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/global_role_binding.html.markdown.
type GlobalRoleBinding struct {
	s *pulumi.ResourceState
}

// NewGlobalRoleBinding registers a new resource with the given unique name, arguments, and options.
func NewGlobalRoleBinding(ctx *pulumi.Context,
	name string, args *GlobalRoleBindingArgs, opts ...pulumi.ResourceOpt) (*GlobalRoleBinding, error) {
	if args == nil || args.GlobalRoleId == nil {
		return nil, errors.New("missing required argument 'GlobalRoleId'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["annotations"] = nil
		inputs["globalRoleId"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["userId"] = nil
	} else {
		inputs["annotations"] = args.Annotations
		inputs["globalRoleId"] = args.GlobalRoleId
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["userId"] = args.UserId
	}
	s, err := ctx.RegisterResource("rancher2:index/globalRoleBinding:GlobalRoleBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GlobalRoleBinding{s: s}, nil
}

// GetGlobalRoleBinding gets an existing GlobalRoleBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalRoleBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GlobalRoleBindingState, opts ...pulumi.ResourceOpt) (*GlobalRoleBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["annotations"] = state.Annotations
		inputs["globalRoleId"] = state.GlobalRoleId
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["userId"] = state.UserId
	}
	s, err := ctx.ReadResource("rancher2:index/globalRoleBinding:GlobalRoleBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GlobalRoleBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *GlobalRoleBinding) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *GlobalRoleBinding) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Annotations for global role binding (map)
func (r *GlobalRoleBinding) Annotations() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["annotations"])
}

// The role id from create global role binding (string)
func (r *GlobalRoleBinding) GlobalRoleId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["globalRoleId"])
}

// Labels for global role binding (map)
func (r *GlobalRoleBinding) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The name of the global role binding (string)
func (r *GlobalRoleBinding) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The user ID to assign global role binding (string)
func (r *GlobalRoleBinding) UserId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userId"])
}

// Input properties used for looking up and filtering GlobalRoleBinding resources.
type GlobalRoleBindingState struct {
	// Annotations for global role binding (map)
	Annotations interface{}
	// The role id from create global role binding (string)
	GlobalRoleId interface{}
	// Labels for global role binding (map)
	Labels interface{}
	// The name of the global role binding (string)
	Name interface{}
	// The user ID to assign global role binding (string)
	UserId interface{}
}

// The set of arguments for constructing a GlobalRoleBinding resource.
type GlobalRoleBindingArgs struct {
	// Annotations for global role binding (map)
	Annotations interface{}
	// The role id from create global role binding (string)
	GlobalRoleId interface{}
	// Labels for global role binding (map)
	Labels interface{}
	// The name of the global role binding (string)
	Name interface{}
	// The user ID to assign global role binding (string)
	UserId interface{}
}
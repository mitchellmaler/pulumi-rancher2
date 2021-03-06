// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"reflect"

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
	pulumi.CustomResourceState

	// Annotations for Registry object (map)
	Annotations pulumi.MapOutput `pulumi:"annotations"`
	// A registry description (string)
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Labels for Registry object (map)
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The name of the registry (string)
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace id where to assign the namespaced registry (string)
	NamespaceId pulumi.StringPtrOutput `pulumi:"namespaceId"`
	// The project id where to assign the registry (string)
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Registries data for registry (list)
	Registries RegistryRegistryArrayOutput `pulumi:"registries"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil || args.Registries == nil {
		return nil, errors.New("missing required argument 'Registries'")
	}
	if args == nil {
		args = &RegistryArgs{}
	}
	var resource Registry
	err := ctx.RegisterResource("rancher2:index/registry:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("rancher2:index/registry:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registry resources.
type registryState struct {
	// Annotations for Registry object (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// A registry description (string)
	Description *string `pulumi:"description"`
	// Labels for Registry object (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The name of the registry (string)
	Name *string `pulumi:"name"`
	// The namespace id where to assign the namespaced registry (string)
	NamespaceId *string `pulumi:"namespaceId"`
	// The project id where to assign the registry (string)
	ProjectId *string `pulumi:"projectId"`
	// Registries data for registry (list)
	Registries []RegistryRegistry `pulumi:"registries"`
}

type RegistryState struct {
	// Annotations for Registry object (map)
	Annotations pulumi.MapInput
	// A registry description (string)
	Description pulumi.StringPtrInput
	// Labels for Registry object (map)
	Labels pulumi.MapInput
	// The name of the registry (string)
	Name pulumi.StringPtrInput
	// The namespace id where to assign the namespaced registry (string)
	NamespaceId pulumi.StringPtrInput
	// The project id where to assign the registry (string)
	ProjectId pulumi.StringPtrInput
	// Registries data for registry (list)
	Registries RegistryRegistryArrayInput
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	// Annotations for Registry object (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// A registry description (string)
	Description *string `pulumi:"description"`
	// Labels for Registry object (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The name of the registry (string)
	Name *string `pulumi:"name"`
	// The namespace id where to assign the namespaced registry (string)
	NamespaceId *string `pulumi:"namespaceId"`
	// The project id where to assign the registry (string)
	ProjectId string `pulumi:"projectId"`
	// Registries data for registry (list)
	Registries []RegistryRegistry `pulumi:"registries"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// Annotations for Registry object (map)
	Annotations pulumi.MapInput
	// A registry description (string)
	Description pulumi.StringPtrInput
	// Labels for Registry object (map)
	Labels pulumi.MapInput
	// The name of the registry (string)
	Name pulumi.StringPtrInput
	// The namespace id where to assign the namespaced registry (string)
	NamespaceId pulumi.StringPtrInput
	// The project id where to assign the registry (string)
	ProjectId pulumi.StringInput
	// Registries data for registry (list)
	Registries RegistryRegistryArrayInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}


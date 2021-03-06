// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Project Alert Group resource. This can be used to create Project Alert Group for Rancher v2 environments and retrieve their information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/projectAlertGroup.html.markdown.
type ProjectAlertGroup struct {
	pulumi.CustomResourceState

	// The project alert group annotations (map)
	Annotations pulumi.MapOutput `pulumi:"annotations"`
	// The project alert group description (string)
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The project alert group interval seconds. Default: `180` (int)
	GroupIntervalSeconds pulumi.IntPtrOutput `pulumi:"groupIntervalSeconds"`
	// The project alert group wait seconds. Default: `180` (int)
	GroupWaitSeconds pulumi.IntPtrOutput `pulumi:"groupWaitSeconds"`
	// The project alert group labels (map)
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The project alert group name (string)
	Name pulumi.StringOutput `pulumi:"name"`
	// The project id where create project alert group (string)
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The project alert group recipients (list)
	Recipients ProjectAlertGroupRecipientArrayOutput `pulumi:"recipients"`
	// The project alert group wait seconds. Default: `3600` (int)
	RepeatIntervalSeconds pulumi.IntPtrOutput `pulumi:"repeatIntervalSeconds"`
}

// NewProjectAlertGroup registers a new resource with the given unique name, arguments, and options.
func NewProjectAlertGroup(ctx *pulumi.Context,
	name string, args *ProjectAlertGroupArgs, opts ...pulumi.ResourceOption) (*ProjectAlertGroup, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil {
		args = &ProjectAlertGroupArgs{}
	}
	var resource ProjectAlertGroup
	err := ctx.RegisterResource("rancher2:index/projectAlertGroup:ProjectAlertGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectAlertGroup gets an existing ProjectAlertGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectAlertGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectAlertGroupState, opts ...pulumi.ResourceOption) (*ProjectAlertGroup, error) {
	var resource ProjectAlertGroup
	err := ctx.ReadResource("rancher2:index/projectAlertGroup:ProjectAlertGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectAlertGroup resources.
type projectAlertGroupState struct {
	// The project alert group annotations (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// The project alert group description (string)
	Description *string `pulumi:"description"`
	// The project alert group interval seconds. Default: `180` (int)
	GroupIntervalSeconds *int `pulumi:"groupIntervalSeconds"`
	// The project alert group wait seconds. Default: `180` (int)
	GroupWaitSeconds *int `pulumi:"groupWaitSeconds"`
	// The project alert group labels (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The project alert group name (string)
	Name *string `pulumi:"name"`
	// The project id where create project alert group (string)
	ProjectId *string `pulumi:"projectId"`
	// The project alert group recipients (list)
	Recipients []ProjectAlertGroupRecipient `pulumi:"recipients"`
	// The project alert group wait seconds. Default: `3600` (int)
	RepeatIntervalSeconds *int `pulumi:"repeatIntervalSeconds"`
}

type ProjectAlertGroupState struct {
	// The project alert group annotations (map)
	Annotations pulumi.MapInput
	// The project alert group description (string)
	Description pulumi.StringPtrInput
	// The project alert group interval seconds. Default: `180` (int)
	GroupIntervalSeconds pulumi.IntPtrInput
	// The project alert group wait seconds. Default: `180` (int)
	GroupWaitSeconds pulumi.IntPtrInput
	// The project alert group labels (map)
	Labels pulumi.MapInput
	// The project alert group name (string)
	Name pulumi.StringPtrInput
	// The project id where create project alert group (string)
	ProjectId pulumi.StringPtrInput
	// The project alert group recipients (list)
	Recipients ProjectAlertGroupRecipientArrayInput
	// The project alert group wait seconds. Default: `3600` (int)
	RepeatIntervalSeconds pulumi.IntPtrInput
}

func (ProjectAlertGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectAlertGroupState)(nil)).Elem()
}

type projectAlertGroupArgs struct {
	// The project alert group annotations (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// The project alert group description (string)
	Description *string `pulumi:"description"`
	// The project alert group interval seconds. Default: `180` (int)
	GroupIntervalSeconds *int `pulumi:"groupIntervalSeconds"`
	// The project alert group wait seconds. Default: `180` (int)
	GroupWaitSeconds *int `pulumi:"groupWaitSeconds"`
	// The project alert group labels (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The project alert group name (string)
	Name *string `pulumi:"name"`
	// The project id where create project alert group (string)
	ProjectId string `pulumi:"projectId"`
	// The project alert group recipients (list)
	Recipients []ProjectAlertGroupRecipient `pulumi:"recipients"`
	// The project alert group wait seconds. Default: `3600` (int)
	RepeatIntervalSeconds *int `pulumi:"repeatIntervalSeconds"`
}

// The set of arguments for constructing a ProjectAlertGroup resource.
type ProjectAlertGroupArgs struct {
	// The project alert group annotations (map)
	Annotations pulumi.MapInput
	// The project alert group description (string)
	Description pulumi.StringPtrInput
	// The project alert group interval seconds. Default: `180` (int)
	GroupIntervalSeconds pulumi.IntPtrInput
	// The project alert group wait seconds. Default: `180` (int)
	GroupWaitSeconds pulumi.IntPtrInput
	// The project alert group labels (map)
	Labels pulumi.MapInput
	// The project alert group name (string)
	Name pulumi.StringPtrInput
	// The project id where create project alert group (string)
	ProjectId pulumi.StringInput
	// The project alert group recipients (list)
	Recipients ProjectAlertGroupRecipientArrayInput
	// The project alert group wait seconds. Default: `3600` (int)
	RepeatIntervalSeconds pulumi.IntPtrInput
}

func (ProjectAlertGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectAlertGroupArgs)(nil)).Elem()
}


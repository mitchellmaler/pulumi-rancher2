// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a Rancher v2 global role binding.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/global_role_binding.html.markdown.
func LookupGlobalRoleBinding(ctx *pulumi.Context, args *GetGlobalRoleBindingArgs) (*GetGlobalRoleBindingResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["globalRoleId"] = args.GlobalRoleId
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("rancher2:index/getGlobalRoleBinding:getGlobalRoleBinding", inputs)
	if err != nil {
		return nil, err
	}
	return &GetGlobalRoleBindingResult{
		Annotations: outputs["annotations"],
		GlobalRoleId: outputs["globalRoleId"],
		Labels: outputs["labels"],
		Name: outputs["name"],
		UserId: outputs["userId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getGlobalRoleBinding.
type GetGlobalRoleBindingArgs struct {
	// The global role id (string)
	GlobalRoleId interface{}
	// The name of the global role binding (string)
	Name interface{}
}

// A collection of values returned by getGlobalRoleBinding.
type GetGlobalRoleBindingResult struct {
	// (Computed) Annotations of the resource (map)
	Annotations interface{}
	GlobalRoleId interface{}
	// (Computed) Labels of the resource (map)
	Labels interface{}
	Name interface{}
	// (Computed) The user ID to assign global role binding (string)
	UserId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
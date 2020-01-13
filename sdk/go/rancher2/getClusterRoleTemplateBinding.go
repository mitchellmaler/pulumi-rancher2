// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a Rancher v2 cluster role template binding.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/cluster_role_template_binding.html.markdown.
func LookupClusterRoleTemplateBinding(ctx *pulumi.Context, args *GetClusterRoleTemplateBindingArgs) (*GetClusterRoleTemplateBindingResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["clusterId"] = args.ClusterId
		inputs["name"] = args.Name
		inputs["roleTemplateId"] = args.RoleTemplateId
	}
	outputs, err := ctx.Invoke("rancher2:index/getClusterRoleTemplateBinding:getClusterRoleTemplateBinding", inputs)
	if err != nil {
		return nil, err
	}
	return &GetClusterRoleTemplateBindingResult{
		Annotations: outputs["annotations"],
		ClusterId: outputs["clusterId"],
		GroupId: outputs["groupId"],
		GroupPrincipalId: outputs["groupPrincipalId"],
		Labels: outputs["labels"],
		Name: outputs["name"],
		RoleTemplateId: outputs["roleTemplateId"],
		UserId: outputs["userId"],
		UserPrincipalId: outputs["userPrincipalId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getClusterRoleTemplateBinding.
type GetClusterRoleTemplateBindingArgs struct {
	// The cluster id where bind cluster role template (string)
	ClusterId interface{}
	// The name of the cluster role template binding (string)
	Name interface{}
	// The role template id from create cluster role template binding (string)
	RoleTemplateId interface{}
}

// A collection of values returned by getClusterRoleTemplateBinding.
type GetClusterRoleTemplateBindingResult struct {
	// (Computed) Annotations of the resource (map)
	Annotations interface{}
	ClusterId interface{}
	// (Computed) The group ID to assign cluster role template binding (string)
	GroupId interface{}
	// (Computed) The groupPrincipal ID to assign cluster role template binding (string)
	GroupPrincipalId interface{}
	// (Computed) Labels of the resource (map)
	Labels interface{}
	Name interface{}
	RoleTemplateId interface{}
	// (Computed) The user ID to assign cluster role template binding (string)
	UserId interface{}
	// (Computed) The userPrincipal ID to assign cluster role template binding (string)
	UserPrincipalId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
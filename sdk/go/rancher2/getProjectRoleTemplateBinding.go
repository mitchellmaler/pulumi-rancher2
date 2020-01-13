// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a Rancher v2 project role template binding.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/project_role_template_binding.html.markdown.
func LookupProjectRoleTemplateBinding(ctx *pulumi.Context, args *GetProjectRoleTemplateBindingArgs) (*GetProjectRoleTemplateBindingResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["projectId"] = args.ProjectId
		inputs["roleTemplateId"] = args.RoleTemplateId
	}
	outputs, err := ctx.Invoke("rancher2:index/getProjectRoleTemplateBinding:getProjectRoleTemplateBinding", inputs)
	if err != nil {
		return nil, err
	}
	return &GetProjectRoleTemplateBindingResult{
		Annotations: outputs["annotations"],
		GroupId: outputs["groupId"],
		GroupPrincipalId: outputs["groupPrincipalId"],
		Labels: outputs["labels"],
		Name: outputs["name"],
		ProjectId: outputs["projectId"],
		RoleTemplateId: outputs["roleTemplateId"],
		UserId: outputs["userId"],
		UserPrincipalId: outputs["userPrincipalId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getProjectRoleTemplateBinding.
type GetProjectRoleTemplateBindingArgs struct {
	// The name of the project role template binding (string)
	Name interface{}
	// The project id where bind project role template (string)
	ProjectId interface{}
	// The role template id from create project role template binding (string)
	RoleTemplateId interface{}
}

// A collection of values returned by getProjectRoleTemplateBinding.
type GetProjectRoleTemplateBindingResult struct {
	// (Computed) Annotations of the resource (map)
	Annotations interface{}
	// (Computed) The group ID to assign project role template binding (string)
	GroupId interface{}
	// (Computed) The groupPrincipal ID to assign project role template binding (string)
	GroupPrincipalId interface{}
	// (Computed) Labels of the resource (map)
	Labels interface{}
	Name interface{}
	ProjectId interface{}
	RoleTemplateId interface{}
	// (Computed) The user ID to assign project role template binding (string)
	UserId interface{}
	// (Computed) The userPrincipal ID to assign project role template binding (string)
	UserPrincipalId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
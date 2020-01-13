// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a Rancher v2 app.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/app.html.markdown.
func LookupApp(ctx *pulumi.Context, args *GetAppArgs) (*GetAppResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["annotations"] = args.Annotations
		inputs["name"] = args.Name
		inputs["projectId"] = args.ProjectId
		inputs["targetNamespace"] = args.TargetNamespace
	}
	outputs, err := ctx.Invoke("rancher2:index/getApp:getApp", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAppResult{
		Annotations: outputs["annotations"],
		Answers: outputs["answers"],
		CatalogName: outputs["catalogName"],
		Description: outputs["description"],
		ExternalId: outputs["externalId"],
		Labels: outputs["labels"],
		Name: outputs["name"],
		ProjectId: outputs["projectId"],
		RevisionId: outputs["revisionId"],
		TargetNamespace: outputs["targetNamespace"],
		TemplateName: outputs["templateName"],
		TemplateVersion: outputs["templateVersion"],
		ValuesYaml: outputs["valuesYaml"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getApp.
type GetAppArgs struct {
	Annotations interface{}
	// The app name (string)
	Name interface{}
	// The id of the project where the app is deployed (string)
	ProjectId interface{}
	// The namespace name where the app is deployed (string)
	TargetNamespace interface{}
}

// A collection of values returned by getApp.
type GetAppResult struct {
	// (Computed) Annotations for the catalog (map)
	Annotations interface{}
	// (Computed) Answers for the app (map)
	Answers interface{}
	// (Computed) Catalog name of the app (string)
	CatalogName interface{}
	// (Computed) Description for the app (string)
	Description interface{}
	// (Computed) The URL of the helm catalog app (string)
	ExternalId interface{}
	// (Computed) Labels for the catalog (map)
	Labels interface{}
	Name interface{}
	ProjectId interface{}
	// (Computed) Current revision id for the app (string)
	RevisionId interface{}
	TargetNamespace interface{}
	// (Computed) Template name of the app (string)
	TemplateName interface{}
	// (Computed) Template version of the app (string)
	TemplateVersion interface{}
	// (Computed) values.yaml base64 encoded file content for the app (string)
	ValuesYaml interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}

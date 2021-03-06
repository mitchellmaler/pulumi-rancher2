// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a Rancher v2 app.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/app.html.markdown.
func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	var rv LookupAppResult
	err := ctx.Invoke("rancher2:index/getApp:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApp.
type LookupAppArgs struct {
	Annotations map[string]interface{} `pulumi:"annotations"`
	// The app name (string)
	Name string `pulumi:"name"`
	// The id of the project where the app is deployed (string)
	ProjectId string `pulumi:"projectId"`
	// The namespace name where the app is deployed (string)
	TargetNamespace *string `pulumi:"targetNamespace"`
}


// A collection of values returned by getApp.
type LookupAppResult struct {
	// (Computed) Annotations for the catalog (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// (Computed) Answers for the app (map)
	Answers map[string]interface{} `pulumi:"answers"`
	// (Computed) Catalog name of the app (string)
	CatalogName string `pulumi:"catalogName"`
	// (Computed) Description for the app (string)
	Description string `pulumi:"description"`
	// (Computed) The URL of the helm catalog app (string)
	ExternalId string `pulumi:"externalId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) Labels for the catalog (map)
	Labels map[string]interface{} `pulumi:"labels"`
	Name string `pulumi:"name"`
	ProjectId string `pulumi:"projectId"`
	// (Computed) Current revision id for the app (string)
	RevisionId string `pulumi:"revisionId"`
	TargetNamespace string `pulumi:"targetNamespace"`
	// (Computed) Template name of the app (string)
	TemplateName string `pulumi:"templateName"`
	// (Computed) Template version of the app (string)
	TemplateVersion string `pulumi:"templateVersion"`
	// (Computed) values.yaml base64 encoded file content for the app (string)
	ValuesYaml string `pulumi:"valuesYaml"`
}


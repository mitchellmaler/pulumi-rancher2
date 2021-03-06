// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a Rancher v2 multi cluster app.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/multiClusterApp.html.markdown.
func LookupMultiClusterApp(ctx *pulumi.Context, args *LookupMultiClusterAppArgs, opts ...pulumi.InvokeOption) (*LookupMultiClusterAppResult, error) {
	var rv LookupMultiClusterAppResult
	err := ctx.Invoke("rancher2:index/getMultiClusterApp:getMultiClusterApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMultiClusterApp.
type LookupMultiClusterAppArgs struct {
	// The multi cluster app name (string)
	Name string `pulumi:"name"`
}


// A collection of values returned by getMultiClusterApp.
type LookupMultiClusterAppResult struct {
	// (Computed) Annotations for multi cluster app object (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// (Computed) The multi cluster app answers (list)
	Answers []GetMultiClusterAppAnswer `pulumi:"answers"`
	// (Computed) The multi cluster app catalog name (string)
	CatalogName string `pulumi:"catalogName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) Labels for multi cluster app object (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// (Computed) The multi cluster app members (list)
	Members []GetMultiClusterAppMember `pulumi:"members"`
	Name string `pulumi:"name"`
	// (Computed) The multi cluster app revision history limit (int)
	RevisionHistoryLimit int `pulumi:"revisionHistoryLimit"`
	// (Computed) Current revision id for the multi cluster app (string)
	RevisionId string `pulumi:"revisionId"`
	// (Computed) The multi cluster app roles (list)
	Roles []string `pulumi:"roles"`
	// (Computed) The multi cluster app target projects (list)
	Targets []GetMultiClusterAppTarget `pulumi:"targets"`
	// (Computed) The multi cluster app template name (string)
	TemplateName string `pulumi:"templateName"`
	// (Computed) The multi cluster app template version (string)
	TemplateVersion string `pulumi:"templateVersion"`
	// (Computed) The multi cluster app template version ID (string)
	TemplateVersionId string `pulumi:"templateVersionId"`
	// (Computed) The multi cluster app upgrade strategy (list)
	UpgradeStrategies []GetMultiClusterAppUpgradeStrategy `pulumi:"upgradeStrategies"`
}


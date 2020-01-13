// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a Rancher v2 setting.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/setting.html.markdown.
func LookupSetting(ctx *pulumi.Context, args *GetSettingArgs) (*GetSettingResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("rancher2:index/getSetting:getSetting", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSettingResult{
		Name: outputs["name"],
		Value: outputs["value"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSetting.
type GetSettingArgs struct {
	// The setting name.
	Name interface{}
}

// A collection of values returned by getSetting.
type GetSettingResult struct {
	Name interface{}
	// the settting's value.
	Value interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
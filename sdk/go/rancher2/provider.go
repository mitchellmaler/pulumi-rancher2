// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The provider type for the rancher2 package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/index.html.markdown.
type Provider struct {
	pulumi.ProviderResourceState

}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.AccessKey == nil {
		args.AccessKey = pulumi.StringPtr(getEnvOrDefault("", nil, "RANCHER_ACCESS_KEY").(string))
	}
	if args.ApiUrl == nil {
		args.ApiUrl = pulumi.StringPtr(getEnvOrDefault("", nil, "RANCHER_URL").(string))
	}
	if args.Bootstrap == nil {
		args.Bootstrap = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "RANCHER_BOOTSTRAP").(bool))
	}
	if args.CaCerts == nil {
		args.CaCerts = pulumi.StringPtr(getEnvOrDefault("", nil, "RANCHER_CA_CERTS").(string))
	}
	if args.Insecure == nil {
		args.Insecure = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "RANCHER_INSECURE").(bool))
	}
	if args.SecretKey == nil {
		args.SecretKey = pulumi.StringPtr(getEnvOrDefault("", nil, "RANCHER_SECRET_KEY").(string))
	}
	if args.TokenKey == nil {
		args.TokenKey = pulumi.StringPtr(getEnvOrDefault("", nil, "RANCHER_TOKEN_KEY").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:rancher2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// API Key used to authenticate with the rancher server
	AccessKey *string `pulumi:"accessKey"`
	// The URL to the rancher API
	ApiUrl *string `pulumi:"apiUrl"`
	// Bootstrap rancher server
	Bootstrap *bool `pulumi:"bootstrap"`
	// CA certificates used to sign rancher server tls certificates. Mandatory if self signed tls and insecure option false
	CaCerts *string `pulumi:"caCerts"`
	// Allow insecure connections to Rancher. Mandatory if self signed tls and not ca_certs provided
	Insecure *bool `pulumi:"insecure"`
	// API secret used to authenticate with the rancher server
	SecretKey *string `pulumi:"secretKey"`
	// API token used to authenticate with the rancher server
	TokenKey *string `pulumi:"tokenKey"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// API Key used to authenticate with the rancher server
	AccessKey pulumi.StringPtrInput
	// The URL to the rancher API
	ApiUrl pulumi.StringPtrInput
	// Bootstrap rancher server
	Bootstrap pulumi.BoolPtrInput
	// CA certificates used to sign rancher server tls certificates. Mandatory if self signed tls and insecure option false
	CaCerts pulumi.StringPtrInput
	// Allow insecure connections to Rancher. Mandatory if self signed tls and not ca_certs provided
	Insecure pulumi.BoolPtrInput
	// API secret used to authenticate with the rancher server
	SecretKey pulumi.StringPtrInput
	// API token used to authenticate with the rancher server
	TokenKey pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}


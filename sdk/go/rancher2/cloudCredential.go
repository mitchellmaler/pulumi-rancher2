// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Cloud Credential resource. This can be used to create Cloud Credential for Rancher v2.2.x and retrieve their information.
//
// amazonec2, azure, digitalocean, openstack and vsphere credentials config are supported for Cloud Credential.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/cloudCredential.html.markdown.
type CloudCredential struct {
	pulumi.CustomResourceState

	// AWS config for the Cloud Credential (list maxitems:1)
	Amazonec2CredentialConfig CloudCredentialAmazonec2CredentialConfigPtrOutput `pulumi:"amazonec2CredentialConfig"`
	// Annotations for Cloud Credential object (map)
	Annotations pulumi.MapOutput `pulumi:"annotations"`
	// Azure config for the Cloud Credential (list maxitems:1)
	AzureCredentialConfig CloudCredentialAzureCredentialConfigPtrOutput `pulumi:"azureCredentialConfig"`
	// Description for the Cloud Credential (string)
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// DigitalOcean config for the Cloud Credential (list maxitems:1)
	DigitaloceanCredentialConfig CloudCredentialDigitaloceanCredentialConfigPtrOutput `pulumi:"digitaloceanCredentialConfig"`
	// (Computed) The driver of the Cloud Credential (string)
	Driver pulumi.StringOutput `pulumi:"driver"`
	// Labels for Cloud Credential object (map)
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The name of the Cloud Credential (string)
	Name pulumi.StringOutput `pulumi:"name"`
	// OpenStack config for the Cloud Credential (list maxitems:1)
	OpenstackCredentialConfig CloudCredentialOpenstackCredentialConfigPtrOutput `pulumi:"openstackCredentialConfig"`
	// vSphere config for the Cloud Credential (list maxitems:1)
	VsphereCredentialConfig CloudCredentialVsphereCredentialConfigPtrOutput `pulumi:"vsphereCredentialConfig"`
}

// NewCloudCredential registers a new resource with the given unique name, arguments, and options.
func NewCloudCredential(ctx *pulumi.Context,
	name string, args *CloudCredentialArgs, opts ...pulumi.ResourceOption) (*CloudCredential, error) {
	if args == nil {
		args = &CloudCredentialArgs{}
	}
	var resource CloudCredential
	err := ctx.RegisterResource("rancher2:index/cloudCredential:CloudCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudCredential gets an existing CloudCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudCredentialState, opts ...pulumi.ResourceOption) (*CloudCredential, error) {
	var resource CloudCredential
	err := ctx.ReadResource("rancher2:index/cloudCredential:CloudCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudCredential resources.
type cloudCredentialState struct {
	// AWS config for the Cloud Credential (list maxitems:1)
	Amazonec2CredentialConfig *CloudCredentialAmazonec2CredentialConfig `pulumi:"amazonec2CredentialConfig"`
	// Annotations for Cloud Credential object (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// Azure config for the Cloud Credential (list maxitems:1)
	AzureCredentialConfig *CloudCredentialAzureCredentialConfig `pulumi:"azureCredentialConfig"`
	// Description for the Cloud Credential (string)
	Description *string `pulumi:"description"`
	// DigitalOcean config for the Cloud Credential (list maxitems:1)
	DigitaloceanCredentialConfig *CloudCredentialDigitaloceanCredentialConfig `pulumi:"digitaloceanCredentialConfig"`
	// (Computed) The driver of the Cloud Credential (string)
	Driver *string `pulumi:"driver"`
	// Labels for Cloud Credential object (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The name of the Cloud Credential (string)
	Name *string `pulumi:"name"`
	// OpenStack config for the Cloud Credential (list maxitems:1)
	OpenstackCredentialConfig *CloudCredentialOpenstackCredentialConfig `pulumi:"openstackCredentialConfig"`
	// vSphere config for the Cloud Credential (list maxitems:1)
	VsphereCredentialConfig *CloudCredentialVsphereCredentialConfig `pulumi:"vsphereCredentialConfig"`
}

type CloudCredentialState struct {
	// AWS config for the Cloud Credential (list maxitems:1)
	Amazonec2CredentialConfig CloudCredentialAmazonec2CredentialConfigPtrInput
	// Annotations for Cloud Credential object (map)
	Annotations pulumi.MapInput
	// Azure config for the Cloud Credential (list maxitems:1)
	AzureCredentialConfig CloudCredentialAzureCredentialConfigPtrInput
	// Description for the Cloud Credential (string)
	Description pulumi.StringPtrInput
	// DigitalOcean config for the Cloud Credential (list maxitems:1)
	DigitaloceanCredentialConfig CloudCredentialDigitaloceanCredentialConfigPtrInput
	// (Computed) The driver of the Cloud Credential (string)
	Driver pulumi.StringPtrInput
	// Labels for Cloud Credential object (map)
	Labels pulumi.MapInput
	// The name of the Cloud Credential (string)
	Name pulumi.StringPtrInput
	// OpenStack config for the Cloud Credential (list maxitems:1)
	OpenstackCredentialConfig CloudCredentialOpenstackCredentialConfigPtrInput
	// vSphere config for the Cloud Credential (list maxitems:1)
	VsphereCredentialConfig CloudCredentialVsphereCredentialConfigPtrInput
}

func (CloudCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudCredentialState)(nil)).Elem()
}

type cloudCredentialArgs struct {
	// AWS config for the Cloud Credential (list maxitems:1)
	Amazonec2CredentialConfig *CloudCredentialAmazonec2CredentialConfig `pulumi:"amazonec2CredentialConfig"`
	// Annotations for Cloud Credential object (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// Azure config for the Cloud Credential (list maxitems:1)
	AzureCredentialConfig *CloudCredentialAzureCredentialConfig `pulumi:"azureCredentialConfig"`
	// Description for the Cloud Credential (string)
	Description *string `pulumi:"description"`
	// DigitalOcean config for the Cloud Credential (list maxitems:1)
	DigitaloceanCredentialConfig *CloudCredentialDigitaloceanCredentialConfig `pulumi:"digitaloceanCredentialConfig"`
	// Labels for Cloud Credential object (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The name of the Cloud Credential (string)
	Name *string `pulumi:"name"`
	// OpenStack config for the Cloud Credential (list maxitems:1)
	OpenstackCredentialConfig *CloudCredentialOpenstackCredentialConfig `pulumi:"openstackCredentialConfig"`
	// vSphere config for the Cloud Credential (list maxitems:1)
	VsphereCredentialConfig *CloudCredentialVsphereCredentialConfig `pulumi:"vsphereCredentialConfig"`
}

// The set of arguments for constructing a CloudCredential resource.
type CloudCredentialArgs struct {
	// AWS config for the Cloud Credential (list maxitems:1)
	Amazonec2CredentialConfig CloudCredentialAmazonec2CredentialConfigPtrInput
	// Annotations for Cloud Credential object (map)
	Annotations pulumi.MapInput
	// Azure config for the Cloud Credential (list maxitems:1)
	AzureCredentialConfig CloudCredentialAzureCredentialConfigPtrInput
	// Description for the Cloud Credential (string)
	Description pulumi.StringPtrInput
	// DigitalOcean config for the Cloud Credential (list maxitems:1)
	DigitaloceanCredentialConfig CloudCredentialDigitaloceanCredentialConfigPtrInput
	// Labels for Cloud Credential object (map)
	Labels pulumi.MapInput
	// The name of the Cloud Credential (string)
	Name pulumi.StringPtrInput
	// OpenStack config for the Cloud Credential (list maxitems:1)
	OpenstackCredentialConfig CloudCredentialOpenstackCredentialConfigPtrInput
	// vSphere config for the Cloud Credential (list maxitems:1)
	VsphereCredentialConfig CloudCredentialVsphereCredentialConfigPtrInput
}

func (CloudCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudCredentialArgs)(nil)).Elem()
}


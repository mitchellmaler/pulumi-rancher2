// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rancher2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Node Template resource. This can be used to create Node Template for Rancher v2 and retrieve their information. 
// 
// amazonec2, azure, digitalocean, openstack and vsphere drivers are supported for node templates.
// 
// **Note** If you are upgrading to Rancher v2.3.3, please take a look to final section
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/node_template.html.markdown.
type NodeTemplate struct {
	s *pulumi.ResourceState
}

// NewNodeTemplate registers a new resource with the given unique name, arguments, and options.
func NewNodeTemplate(ctx *pulumi.Context,
	name string, args *NodeTemplateArgs, opts ...pulumi.ResourceOpt) (*NodeTemplate, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["amazonec2Config"] = nil
		inputs["annotations"] = nil
		inputs["authCertificateAuthority"] = nil
		inputs["authKey"] = nil
		inputs["azureConfig"] = nil
		inputs["cloudCredentialId"] = nil
		inputs["description"] = nil
		inputs["digitaloceanConfig"] = nil
		inputs["engineEnv"] = nil
		inputs["engineInsecureRegistries"] = nil
		inputs["engineInstallUrl"] = nil
		inputs["engineLabel"] = nil
		inputs["engineOpt"] = nil
		inputs["engineRegistryMirrors"] = nil
		inputs["engineStorageDriver"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["openstackConfig"] = nil
		inputs["useInternalIpAddress"] = nil
		inputs["vsphereConfig"] = nil
	} else {
		inputs["amazonec2Config"] = args.Amazonec2Config
		inputs["annotations"] = args.Annotations
		inputs["authCertificateAuthority"] = args.AuthCertificateAuthority
		inputs["authKey"] = args.AuthKey
		inputs["azureConfig"] = args.AzureConfig
		inputs["cloudCredentialId"] = args.CloudCredentialId
		inputs["description"] = args.Description
		inputs["digitaloceanConfig"] = args.DigitaloceanConfig
		inputs["engineEnv"] = args.EngineEnv
		inputs["engineInsecureRegistries"] = args.EngineInsecureRegistries
		inputs["engineInstallUrl"] = args.EngineInstallUrl
		inputs["engineLabel"] = args.EngineLabel
		inputs["engineOpt"] = args.EngineOpt
		inputs["engineRegistryMirrors"] = args.EngineRegistryMirrors
		inputs["engineStorageDriver"] = args.EngineStorageDriver
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["openstackConfig"] = args.OpenstackConfig
		inputs["useInternalIpAddress"] = args.UseInternalIpAddress
		inputs["vsphereConfig"] = args.VsphereConfig
	}
	inputs["driver"] = nil
	s, err := ctx.RegisterResource("rancher2:index/nodeTemplate:NodeTemplate", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NodeTemplate{s: s}, nil
}

// GetNodeTemplate gets an existing NodeTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeTemplate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NodeTemplateState, opts ...pulumi.ResourceOpt) (*NodeTemplate, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["amazonec2Config"] = state.Amazonec2Config
		inputs["annotations"] = state.Annotations
		inputs["authCertificateAuthority"] = state.AuthCertificateAuthority
		inputs["authKey"] = state.AuthKey
		inputs["azureConfig"] = state.AzureConfig
		inputs["cloudCredentialId"] = state.CloudCredentialId
		inputs["description"] = state.Description
		inputs["digitaloceanConfig"] = state.DigitaloceanConfig
		inputs["driver"] = state.Driver
		inputs["engineEnv"] = state.EngineEnv
		inputs["engineInsecureRegistries"] = state.EngineInsecureRegistries
		inputs["engineInstallUrl"] = state.EngineInstallUrl
		inputs["engineLabel"] = state.EngineLabel
		inputs["engineOpt"] = state.EngineOpt
		inputs["engineRegistryMirrors"] = state.EngineRegistryMirrors
		inputs["engineStorageDriver"] = state.EngineStorageDriver
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["openstackConfig"] = state.OpenstackConfig
		inputs["useInternalIpAddress"] = state.UseInternalIpAddress
		inputs["vsphereConfig"] = state.VsphereConfig
	}
	s, err := ctx.ReadResource("rancher2:index/nodeTemplate:NodeTemplate", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NodeTemplate{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NodeTemplate) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NodeTemplate) ID() pulumi.IDOutput {
	return r.s.ID()
}

// AWS config for the Node Template (list maxitems:1)
func (r *NodeTemplate) Amazonec2Config() pulumi.Output {
	return r.s.State["amazonec2Config"]
}

// Annotations for Node Template object (map)
func (r *NodeTemplate) Annotations() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["annotations"])
}

// Auth certificate authority for the Node Template (string)
func (r *NodeTemplate) AuthCertificateAuthority() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["authCertificateAuthority"])
}

// Auth key for the Node Template (string)
func (r *NodeTemplate) AuthKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["authKey"])
}

// Azure config for the Node Template (list maxitems:1)
func (r *NodeTemplate) AzureConfig() pulumi.Output {
	return r.s.State["azureConfig"]
}

// Cloud credential ID for the Node Template. Required from Rancher v2.2.x (string)
func (r *NodeTemplate) CloudCredentialId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cloudCredentialId"])
}

// Description for the Node Template (string)
func (r *NodeTemplate) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Digitalocean config for the Node Template (list maxitems:1)
func (r *NodeTemplate) DigitaloceanConfig() pulumi.Output {
	return r.s.State["digitaloceanConfig"]
}

// (Computed) The driver of the node template (string)
func (r *NodeTemplate) Driver() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["driver"])
}

// Engine environment for the node template (string)
func (r *NodeTemplate) EngineEnv() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["engineEnv"])
}

// Insecure registry for the node template (list)
func (r *NodeTemplate) EngineInsecureRegistries() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["engineInsecureRegistries"])
}

// Docker engine install URL for the node template. Default `https://releases.rancher.com/install-docker/18.09.sh`. Available install docker versions at `https://github.com/rancher/install-docker` (string)
func (r *NodeTemplate) EngineInstallUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engineInstallUrl"])
}

// Engine label for the node template (string)
func (r *NodeTemplate) EngineLabel() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["engineLabel"])
}

// Engine options for the node template (map)
func (r *NodeTemplate) EngineOpt() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["engineOpt"])
}

// Engine registry mirror for the node template (list)
func (r *NodeTemplate) EngineRegistryMirrors() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["engineRegistryMirrors"])
}

// Engine storage driver for the node template (string)
func (r *NodeTemplate) EngineStorageDriver() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engineStorageDriver"])
}

// Labels for Node Template object (map)
func (r *NodeTemplate) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The name of the Node Template (string)
func (r *NodeTemplate) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Openstack config for the Node Template (list maxitems:1)
func (r *NodeTemplate) OpenstackConfig() pulumi.Output {
	return r.s.State["openstackConfig"]
}

// Engine storage driver for the node template (bool)
func (r *NodeTemplate) UseInternalIpAddress() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["useInternalIpAddress"])
}

// vSphere config for the Node Template (list maxitems:1)
func (r *NodeTemplate) VsphereConfig() pulumi.Output {
	return r.s.State["vsphereConfig"]
}

// Input properties used for looking up and filtering NodeTemplate resources.
type NodeTemplateState struct {
	// AWS config for the Node Template (list maxitems:1)
	Amazonec2Config interface{}
	// Annotations for Node Template object (map)
	Annotations interface{}
	// Auth certificate authority for the Node Template (string)
	AuthCertificateAuthority interface{}
	// Auth key for the Node Template (string)
	AuthKey interface{}
	// Azure config for the Node Template (list maxitems:1)
	AzureConfig interface{}
	// Cloud credential ID for the Node Template. Required from Rancher v2.2.x (string)
	CloudCredentialId interface{}
	// Description for the Node Template (string)
	Description interface{}
	// Digitalocean config for the Node Template (list maxitems:1)
	DigitaloceanConfig interface{}
	// (Computed) The driver of the node template (string)
	Driver interface{}
	// Engine environment for the node template (string)
	EngineEnv interface{}
	// Insecure registry for the node template (list)
	EngineInsecureRegistries interface{}
	// Docker engine install URL for the node template. Default `https://releases.rancher.com/install-docker/18.09.sh`. Available install docker versions at `https://github.com/rancher/install-docker` (string)
	EngineInstallUrl interface{}
	// Engine label for the node template (string)
	EngineLabel interface{}
	// Engine options for the node template (map)
	EngineOpt interface{}
	// Engine registry mirror for the node template (list)
	EngineRegistryMirrors interface{}
	// Engine storage driver for the node template (string)
	EngineStorageDriver interface{}
	// Labels for Node Template object (map)
	Labels interface{}
	// The name of the Node Template (string)
	Name interface{}
	// Openstack config for the Node Template (list maxitems:1)
	OpenstackConfig interface{}
	// Engine storage driver for the node template (bool)
	UseInternalIpAddress interface{}
	// vSphere config for the Node Template (list maxitems:1)
	VsphereConfig interface{}
}

// The set of arguments for constructing a NodeTemplate resource.
type NodeTemplateArgs struct {
	// AWS config for the Node Template (list maxitems:1)
	Amazonec2Config interface{}
	// Annotations for Node Template object (map)
	Annotations interface{}
	// Auth certificate authority for the Node Template (string)
	AuthCertificateAuthority interface{}
	// Auth key for the Node Template (string)
	AuthKey interface{}
	// Azure config for the Node Template (list maxitems:1)
	AzureConfig interface{}
	// Cloud credential ID for the Node Template. Required from Rancher v2.2.x (string)
	CloudCredentialId interface{}
	// Description for the Node Template (string)
	Description interface{}
	// Digitalocean config for the Node Template (list maxitems:1)
	DigitaloceanConfig interface{}
	// Engine environment for the node template (string)
	EngineEnv interface{}
	// Insecure registry for the node template (list)
	EngineInsecureRegistries interface{}
	// Docker engine install URL for the node template. Default `https://releases.rancher.com/install-docker/18.09.sh`. Available install docker versions at `https://github.com/rancher/install-docker` (string)
	EngineInstallUrl interface{}
	// Engine label for the node template (string)
	EngineLabel interface{}
	// Engine options for the node template (map)
	EngineOpt interface{}
	// Engine registry mirror for the node template (list)
	EngineRegistryMirrors interface{}
	// Engine storage driver for the node template (string)
	EngineStorageDriver interface{}
	// Labels for Node Template object (map)
	Labels interface{}
	// The name of the Node Template (string)
	Name interface{}
	// Openstack config for the Node Template (list maxitems:1)
	OpenstackConfig interface{}
	// Engine storage driver for the node template (bool)
	UseInternalIpAddress interface{}
	// vSphere config for the Node Template (list maxitems:1)
	VsphereConfig interface{}
}
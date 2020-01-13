# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class CloudCredential(pulumi.CustomResource):
    amazonec2_credential_config: pulumi.Output[dict]
    """
    AWS config for the Cloud Credential (list maxitems:1)
    
      * `access_key` (`str`) - AWS access key (string)
      * `secret_key` (`str`) - AWS secret key (string)
    """
    annotations: pulumi.Output[dict]
    """
    Annotations for Cloud Credential object (map)
    """
    azure_credential_config: pulumi.Output[dict]
    """
    Azure config for the Cloud Credential (list maxitems:1)
    
      * `client_id` (`str`) - Azure Service Principal Account ID (string)
      * `client_secret` (`str`) - Azure Service Principal Account password (string)
      * `subscriptionId` (`str`) - Azure Subscription ID (string)
    """
    description: pulumi.Output[str]
    """
    Description for the Cloud Credential (string)
    """
    digitalocean_credential_config: pulumi.Output[dict]
    """
    DigitalOcean config for the Cloud Credential (list maxitems:1)
    
      * `accessToken` (`str`) - DigitalOcean access token (string)
    """
    driver: pulumi.Output[str]
    """
    (Computed) The driver of the Cloud Credential (string)
    """
    labels: pulumi.Output[dict]
    """
    Labels for Cloud Credential object (map)
    """
    name: pulumi.Output[str]
    """
    The name of the Cloud Credential (string)
    """
    openstack_credential_config: pulumi.Output[dict]
    """
    OpenStack config for the Cloud Credential (list maxitems:1)
    
      * `password` (`str`) - vSphere password (string)
    """
    vsphere_credential_config: pulumi.Output[dict]
    """
    vSphere config for the Cloud Credential (list maxitems:1)
    
      * `password` (`str`) - vSphere password (string)
      * `username` (`str`) - vSphere username (string)
      * `vcenter` (`str`) - vSphere IP/hostname for vCenter (string)
      * `vcenterPort` (`str`) - vSphere Port for vCenter. Default `443` (string)
    """
    def __init__(__self__, resource_name, opts=None, amazonec2_credential_config=None, annotations=None, azure_credential_config=None, description=None, digitalocean_credential_config=None, labels=None, name=None, openstack_credential_config=None, vsphere_credential_config=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Rancher v2 Cloud Credential resource. This can be used to create Cloud Credential for Rancher v2.2.x and retrieve their information.
        
        amazonec2, azure, digitalocean, openstack and vsphere credentials config are supported for Cloud Credential.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] amazonec2_credential_config: AWS config for the Cloud Credential (list maxitems:1)
        :param pulumi.Input[dict] annotations: Annotations for Cloud Credential object (map)
        :param pulumi.Input[dict] azure_credential_config: Azure config for the Cloud Credential (list maxitems:1)
        :param pulumi.Input[str] description: Description for the Cloud Credential (string)
        :param pulumi.Input[dict] digitalocean_credential_config: DigitalOcean config for the Cloud Credential (list maxitems:1)
        :param pulumi.Input[dict] labels: Labels for Cloud Credential object (map)
        :param pulumi.Input[str] name: The name of the Cloud Credential (string)
        :param pulumi.Input[dict] openstack_credential_config: OpenStack config for the Cloud Credential (list maxitems:1)
        :param pulumi.Input[dict] vsphere_credential_config: vSphere config for the Cloud Credential (list maxitems:1)
        
        The **amazonec2_credential_config** object supports the following:
        
          * `access_key` (`pulumi.Input[str]`) - AWS access key (string)
          * `secret_key` (`pulumi.Input[str]`) - AWS secret key (string)
        
        The **azure_credential_config** object supports the following:
        
          * `client_id` (`pulumi.Input[str]`) - Azure Service Principal Account ID (string)
          * `client_secret` (`pulumi.Input[str]`) - Azure Service Principal Account password (string)
          * `subscriptionId` (`pulumi.Input[str]`) - Azure Subscription ID (string)
        
        The **digitalocean_credential_config** object supports the following:
        
          * `accessToken` (`pulumi.Input[str]`) - DigitalOcean access token (string)
        
        The **openstack_credential_config** object supports the following:
        
          * `password` (`pulumi.Input[str]`) - vSphere password (string)
        
        The **vsphere_credential_config** object supports the following:
        
          * `password` (`pulumi.Input[str]`) - vSphere password (string)
          * `username` (`pulumi.Input[str]`) - vSphere username (string)
          * `vcenter` (`pulumi.Input[str]`) - vSphere IP/hostname for vCenter (string)
          * `vcenterPort` (`pulumi.Input[str]`) - vSphere Port for vCenter. Default `443` (string)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/cloud_credential.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['amazonec2_credential_config'] = amazonec2_credential_config
            __props__['annotations'] = annotations
            __props__['azure_credential_config'] = azure_credential_config
            __props__['description'] = description
            __props__['digitalocean_credential_config'] = digitalocean_credential_config
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['openstack_credential_config'] = openstack_credential_config
            __props__['vsphere_credential_config'] = vsphere_credential_config
            __props__['driver'] = None
        super(CloudCredential, __self__).__init__(
            'rancher2:index/cloudCredential:CloudCredential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, amazonec2_credential_config=None, annotations=None, azure_credential_config=None, description=None, digitalocean_credential_config=None, driver=None, labels=None, name=None, openstack_credential_config=None, vsphere_credential_config=None):
        """
        Get an existing CloudCredential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] amazonec2_credential_config: AWS config for the Cloud Credential (list maxitems:1)
        :param pulumi.Input[dict] annotations: Annotations for Cloud Credential object (map)
        :param pulumi.Input[dict] azure_credential_config: Azure config for the Cloud Credential (list maxitems:1)
        :param pulumi.Input[str] description: Description for the Cloud Credential (string)
        :param pulumi.Input[dict] digitalocean_credential_config: DigitalOcean config for the Cloud Credential (list maxitems:1)
        :param pulumi.Input[str] driver: (Computed) The driver of the Cloud Credential (string)
        :param pulumi.Input[dict] labels: Labels for Cloud Credential object (map)
        :param pulumi.Input[str] name: The name of the Cloud Credential (string)
        :param pulumi.Input[dict] openstack_credential_config: OpenStack config for the Cloud Credential (list maxitems:1)
        :param pulumi.Input[dict] vsphere_credential_config: vSphere config for the Cloud Credential (list maxitems:1)
        
        The **amazonec2_credential_config** object supports the following:
        
          * `access_key` (`pulumi.Input[str]`) - AWS access key (string)
          * `secret_key` (`pulumi.Input[str]`) - AWS secret key (string)
        
        The **azure_credential_config** object supports the following:
        
          * `client_id` (`pulumi.Input[str]`) - Azure Service Principal Account ID (string)
          * `client_secret` (`pulumi.Input[str]`) - Azure Service Principal Account password (string)
          * `subscriptionId` (`pulumi.Input[str]`) - Azure Subscription ID (string)
        
        The **digitalocean_credential_config** object supports the following:
        
          * `accessToken` (`pulumi.Input[str]`) - DigitalOcean access token (string)
        
        The **openstack_credential_config** object supports the following:
        
          * `password` (`pulumi.Input[str]`) - vSphere password (string)
        
        The **vsphere_credential_config** object supports the following:
        
          * `password` (`pulumi.Input[str]`) - vSphere password (string)
          * `username` (`pulumi.Input[str]`) - vSphere username (string)
          * `vcenter` (`pulumi.Input[str]`) - vSphere IP/hostname for vCenter (string)
          * `vcenterPort` (`pulumi.Input[str]`) - vSphere Port for vCenter. Default `443` (string)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/cloud_credential.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["amazonec2_credential_config"] = amazonec2_credential_config
        __props__["annotations"] = annotations
        __props__["azure_credential_config"] = azure_credential_config
        __props__["description"] = description
        __props__["digitalocean_credential_config"] = digitalocean_credential_config
        __props__["driver"] = driver
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["openstack_credential_config"] = openstack_credential_config
        __props__["vsphere_credential_config"] = vsphere_credential_config
        return CloudCredential(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

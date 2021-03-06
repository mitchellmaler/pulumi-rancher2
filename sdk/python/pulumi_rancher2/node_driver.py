# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class NodeDriver(pulumi.CustomResource):
    active: pulumi.Output[bool]
    """
    Specify if the node driver state (bool)
    """
    annotations: pulumi.Output[dict]
    """
    Annotations of the resource (map)
    """
    builtin: pulumi.Output[bool]
    """
    Specify wheter the node driver is an internal node driver or not (bool)
    """
    checksum: pulumi.Output[str]
    """
    Verify that the downloaded driver matches the expected checksum (string)
    """
    description: pulumi.Output[str]
    """
    Description of the node driver (string)
    """
    external_id: pulumi.Output[str]
    """
    External ID (string)
    """
    labels: pulumi.Output[dict]
    """
    Labels of the resource (map)
    """
    name: pulumi.Output[str]
    """
    Name of the node driver (string)
    """
    ui_url: pulumi.Output[str]
    """
    The URL to load for customized Add Nodes screen for this driver (string)
    """
    url: pulumi.Output[str]
    """
    The URL to download the machine driver binary for 64-bit Linux (string)
    """
    whitelist_domains: pulumi.Output[list]
    """
    Domains to whitelist for the ui (list)
    """
    def __init__(__self__, resource_name, opts=None, active=None, annotations=None, builtin=None, checksum=None, description=None, external_id=None, labels=None, name=None, ui_url=None, url=None, whitelist_domains=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Rancher v2 Node Driver resource. This can be used to create Node Driver for Rancher v2 RKE clusters and retrieve their information.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/nodeDriver.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Specify if the node driver state (bool)
        :param pulumi.Input[dict] annotations: Annotations of the resource (map)
        :param pulumi.Input[bool] builtin: Specify wheter the node driver is an internal node driver or not (bool)
        :param pulumi.Input[str] checksum: Verify that the downloaded driver matches the expected checksum (string)
        :param pulumi.Input[str] description: Description of the node driver (string)
        :param pulumi.Input[str] external_id: External ID (string)
        :param pulumi.Input[dict] labels: Labels of the resource (map)
        :param pulumi.Input[str] name: Name of the node driver (string)
        :param pulumi.Input[str] ui_url: The URL to load for customized Add Nodes screen for this driver (string)
        :param pulumi.Input[str] url: The URL to download the machine driver binary for 64-bit Linux (string)
        :param pulumi.Input[list] whitelist_domains: Domains to whitelist for the ui (list)
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

            if active is None:
                raise TypeError("Missing required property 'active'")
            __props__['active'] = active
            __props__['annotations'] = annotations
            if builtin is None:
                raise TypeError("Missing required property 'builtin'")
            __props__['builtin'] = builtin
            __props__['checksum'] = checksum
            __props__['description'] = description
            __props__['external_id'] = external_id
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['ui_url'] = ui_url
            if url is None:
                raise TypeError("Missing required property 'url'")
            __props__['url'] = url
            __props__['whitelist_domains'] = whitelist_domains
        super(NodeDriver, __self__).__init__(
            'rancher2:index/nodeDriver:NodeDriver',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, active=None, annotations=None, builtin=None, checksum=None, description=None, external_id=None, labels=None, name=None, ui_url=None, url=None, whitelist_domains=None):
        """
        Get an existing NodeDriver resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Specify if the node driver state (bool)
        :param pulumi.Input[dict] annotations: Annotations of the resource (map)
        :param pulumi.Input[bool] builtin: Specify wheter the node driver is an internal node driver or not (bool)
        :param pulumi.Input[str] checksum: Verify that the downloaded driver matches the expected checksum (string)
        :param pulumi.Input[str] description: Description of the node driver (string)
        :param pulumi.Input[str] external_id: External ID (string)
        :param pulumi.Input[dict] labels: Labels of the resource (map)
        :param pulumi.Input[str] name: Name of the node driver (string)
        :param pulumi.Input[str] ui_url: The URL to load for customized Add Nodes screen for this driver (string)
        :param pulumi.Input[str] url: The URL to download the machine driver binary for 64-bit Linux (string)
        :param pulumi.Input[list] whitelist_domains: Domains to whitelist for the ui (list)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["active"] = active
        __props__["annotations"] = annotations
        __props__["builtin"] = builtin
        __props__["checksum"] = checksum
        __props__["description"] = description
        __props__["external_id"] = external_id
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["ui_url"] = ui_url
        __props__["url"] = url
        __props__["whitelist_domains"] = whitelist_domains
        return NodeDriver(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


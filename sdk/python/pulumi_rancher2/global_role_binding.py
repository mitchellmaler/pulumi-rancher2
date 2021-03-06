# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GlobalRoleBinding(pulumi.CustomResource):
    annotations: pulumi.Output[dict]
    """
    Annotations for global role binding (map)
    """
    global_role_id: pulumi.Output[str]
    """
    The role id from create global role binding (string)
    """
    labels: pulumi.Output[dict]
    """
    Labels for global role binding (map)
    """
    name: pulumi.Output[str]
    """
    The name of the global role binding (string)
    """
    user_id: pulumi.Output[str]
    """
    The user ID to assign global role binding (string)
    """
    def __init__(__self__, resource_name, opts=None, annotations=None, global_role_id=None, labels=None, name=None, user_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Rancher v2 Global Role Binding resource. This can be used to create Global Role Bindings for Rancher v2 environments and retrieve their information.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/globalRole.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] annotations: Annotations for global role binding (map)
        :param pulumi.Input[str] global_role_id: The role id from create global role binding (string)
        :param pulumi.Input[dict] labels: Labels for global role binding (map)
        :param pulumi.Input[str] name: The name of the global role binding (string)
        :param pulumi.Input[str] user_id: The user ID to assign global role binding (string)
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

            __props__['annotations'] = annotations
            if global_role_id is None:
                raise TypeError("Missing required property 'global_role_id'")
            __props__['global_role_id'] = global_role_id
            __props__['labels'] = labels
            __props__['name'] = name
            if user_id is None:
                raise TypeError("Missing required property 'user_id'")
            __props__['user_id'] = user_id
        super(GlobalRoleBinding, __self__).__init__(
            'rancher2:index/globalRoleBinding:GlobalRoleBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, annotations=None, global_role_id=None, labels=None, name=None, user_id=None):
        """
        Get an existing GlobalRoleBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] annotations: Annotations for global role binding (map)
        :param pulumi.Input[str] global_role_id: The role id from create global role binding (string)
        :param pulumi.Input[dict] labels: Labels for global role binding (map)
        :param pulumi.Input[str] name: The name of the global role binding (string)
        :param pulumi.Input[str] user_id: The user ID to assign global role binding (string)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["annotations"] = annotations
        __props__["global_role_id"] = global_role_id
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["user_id"] = user_id
        return GlobalRoleBinding(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


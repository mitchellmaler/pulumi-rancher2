# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Token(pulumi.CustomResource):
    access_key: pulumi.Output[str]
    """
    (Computed) Token access key part (string)
    """
    annotations: pulumi.Output[dict]
    """
    Annotations of the token (map)
    """
    cluster_id: pulumi.Output[str]
    """
    Cluster ID for scoped token (string)
    """
    description: pulumi.Output[str]
    """
    Token description (string)
    """
    enabled: pulumi.Output[bool]
    """
    (Computed) Token is enabled (bool)
    """
    expired: pulumi.Output[bool]
    """
    (Computed) Token is expired (bool)
    """
    labels: pulumi.Output[dict]
    """
    Labels of the token (map)
    """
    name: pulumi.Output[str]
    """
    (Computed) Token name (string)
    """
    renew: pulumi.Output[bool]
    secret_key: pulumi.Output[str]
    """
    (Computed/Sensitive) Token secret key part (string)
    """
    token: pulumi.Output[str]
    """
    (Computed/Sensitive) Token value (string)
    """
    ttl: pulumi.Output[float]
    """
    Token time to live in seconds. Default `0` (int)
    """
    user_id: pulumi.Output[str]
    """
    (Computed) Token user ID (string)
    """
    def __init__(__self__, resource_name, opts=None, annotations=None, cluster_id=None, description=None, labels=None, renew=None, ttl=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Rancher v2 Token resource. This can be used to create Tokens for Rancher v2 provider user and retrieve their information.
        
        There are 2 kind of tokens:
        - no scoped: valid for global system.
        - scoped: valid for just a specific cluster (`cluster_id` should be provided).
        
        Tokens can't be updated once created. Any diff in token data will recreate the token. If any token expire, Rancher2 provider will generate a diff to regenerate it.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] annotations: Annotations of the token (map)
        :param pulumi.Input[str] cluster_id: Cluster ID for scoped token (string)
        :param pulumi.Input[str] description: Token description (string)
        :param pulumi.Input[dict] labels: Labels of the token (map)
        :param pulumi.Input[float] ttl: Token time to live in seconds. Default `0` (int)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/token.html.markdown.
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
            __props__['cluster_id'] = cluster_id
            __props__['description'] = description
            __props__['labels'] = labels
            __props__['renew'] = renew
            __props__['ttl'] = ttl
            __props__['access_key'] = None
            __props__['enabled'] = None
            __props__['expired'] = None
            __props__['name'] = None
            __props__['secret_key'] = None
            __props__['token'] = None
            __props__['user_id'] = None
        super(Token, __self__).__init__(
            'rancher2:index/token:Token',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_key=None, annotations=None, cluster_id=None, description=None, enabled=None, expired=None, labels=None, name=None, renew=None, secret_key=None, token=None, ttl=None, user_id=None):
        """
        Get an existing Token resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: (Computed) Token access key part (string)
        :param pulumi.Input[dict] annotations: Annotations of the token (map)
        :param pulumi.Input[str] cluster_id: Cluster ID for scoped token (string)
        :param pulumi.Input[str] description: Token description (string)
        :param pulumi.Input[bool] enabled: (Computed) Token is enabled (bool)
        :param pulumi.Input[bool] expired: (Computed) Token is expired (bool)
        :param pulumi.Input[dict] labels: Labels of the token (map)
        :param pulumi.Input[str] name: (Computed) Token name (string)
        :param pulumi.Input[str] secret_key: (Computed/Sensitive) Token secret key part (string)
        :param pulumi.Input[str] token: (Computed/Sensitive) Token value (string)
        :param pulumi.Input[float] ttl: Token time to live in seconds. Default `0` (int)
        :param pulumi.Input[str] user_id: (Computed) Token user ID (string)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/token.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["access_key"] = access_key
        __props__["annotations"] = annotations
        __props__["cluster_id"] = cluster_id
        __props__["description"] = description
        __props__["enabled"] = enabled
        __props__["expired"] = expired
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["renew"] = renew
        __props__["secret_key"] = secret_key
        __props__["token"] = token
        __props__["ttl"] = ttl
        __props__["user_id"] = user_id
        return Token(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

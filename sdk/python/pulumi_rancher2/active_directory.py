# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class ActiveDirectory(pulumi.CustomResource):
    access_mode: pulumi.Output[str]
    """
    Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
    """
    allowed_principal_ids: pulumi.Output[list]
    """
    Allowed principal ids for auth. Required if `access_mode` is `required` or `restricted`. Ex: `activedirectory_user://<DN>`  `activedirectory_group://<DN>` (list)
    """
    annotations: pulumi.Output[dict]
    """
    Annotations of the resource (map)
    """
    certificate: pulumi.Output[str]
    """
    CA certificate for TLS if selfsigned (string)
    """
    connection_timeout: pulumi.Output[float]
    """
    ActiveDirectory connection timeout. Default `5000` (int)
    """
    default_login_domain: pulumi.Output[str]
    """
    ActiveDirectory defult lgoin domain (string)
    """
    enabled: pulumi.Output[bool]
    """
    Enable auth config provider. Default `true` (bool)
    """
    group_dn_attribute: pulumi.Output[str]
    """
    Group DN attribute. Default `distinguishedName` (string)
    """
    group_member_mapping_attribute: pulumi.Output[str]
    """
    Group member mapping attribute. Default `member` (string)
    """
    group_member_user_attribute: pulumi.Output[str]
    """
    Group member user attribute. Default `distinguishedName` (string)
    """
    group_name_attribute: pulumi.Output[str]
    """
    Group name attribute. Default `name` (string)
    """
    group_object_class: pulumi.Output[str]
    """
    Group object class. Default `group` (string)
    """
    group_search_attribute: pulumi.Output[str]
    """
    Group search attribute. Default `sAMAccountName` (string)
    """
    group_search_base: pulumi.Output[str]
    """
    Group search base (string)
    """
    group_search_filter: pulumi.Output[str]
    """
    Group search filter (string)
    """
    labels: pulumi.Output[dict]
    """
    Labels of the resource (map)
    """
    name: pulumi.Output[str]
    """
    (Computed) The name of the resource (string)
    """
    nested_group_membership_enabled: pulumi.Output[bool]
    """
    Nested group membership enable. Default `false` (bool)
    """
    port: pulumi.Output[float]
    """
    ActiveDirectory port. Default `389` (int)
    """
    servers: pulumi.Output[list]
    """
    ActiveDirectory servers list (list)
    """
    service_account_password: pulumi.Output[str]
    """
    Service account password for access ActiveDirectory service (string)
    """
    service_account_username: pulumi.Output[str]
    """
    Service account DN for access ActiveDirectory service (string)
    """
    tls: pulumi.Output[bool]
    """
    Enable TLS connection (bool)
    """
    type: pulumi.Output[str]
    """
    (Computed) The type of the resource (string)
    """
    user_disabled_bit_mask: pulumi.Output[float]
    """
    User disabled bit mask. Default `2` (int)
    """
    user_enabled_attribute: pulumi.Output[str]
    """
    User enable attribute (string)
    """
    user_login_attribute: pulumi.Output[str]
    """
    User login attribute. Default `sAMAccountName` (string)
    """
    user_name_attribute: pulumi.Output[str]
    """
    User name attribute. Default `name` (string)
    """
    user_object_class: pulumi.Output[str]
    """
    User object class. Default `person` (string)
    """
    user_search_attribute: pulumi.Output[str]
    """
    User search attribute. Default `sAMAccountName|sn|givenName` (string)
    """
    user_search_base: pulumi.Output[str]
    """
    User search base DN (string)
    """
    user_search_filter: pulumi.Output[str]
    """
    User search filter (string)
    """
    def __init__(__self__, resource_name, opts=None, access_mode=None, allowed_principal_ids=None, annotations=None, certificate=None, connection_timeout=None, default_login_domain=None, enabled=None, group_dn_attribute=None, group_member_mapping_attribute=None, group_member_user_attribute=None, group_name_attribute=None, group_object_class=None, group_search_attribute=None, group_search_base=None, group_search_filter=None, labels=None, nested_group_membership_enabled=None, port=None, servers=None, service_account_password=None, service_account_username=None, tls=None, user_disabled_bit_mask=None, user_enabled_attribute=None, user_login_attribute=None, user_name_attribute=None, user_object_class=None, user_search_attribute=None, user_search_base=None, user_search_filter=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Rancher v2 Auth Config ActiveDirectory resource. This can be used to configure and enable Auth Config ActiveDirectory for Rancher v2 RKE clusters and retrieve their information.
        
        In addition to the built-in local auth, only one external auth config provider can be enabled at a time.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_mode: Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
        :param pulumi.Input[list] allowed_principal_ids: Allowed principal ids for auth. Required if `access_mode` is `required` or `restricted`. Ex: `activedirectory_user://<DN>`  `activedirectory_group://<DN>` (list)
        :param pulumi.Input[dict] annotations: Annotations of the resource (map)
        :param pulumi.Input[str] certificate: CA certificate for TLS if selfsigned (string)
        :param pulumi.Input[float] connection_timeout: ActiveDirectory connection timeout. Default `5000` (int)
        :param pulumi.Input[str] default_login_domain: ActiveDirectory defult lgoin domain (string)
        :param pulumi.Input[bool] enabled: Enable auth config provider. Default `true` (bool)
        :param pulumi.Input[str] group_dn_attribute: Group DN attribute. Default `distinguishedName` (string)
        :param pulumi.Input[str] group_member_mapping_attribute: Group member mapping attribute. Default `member` (string)
        :param pulumi.Input[str] group_member_user_attribute: Group member user attribute. Default `distinguishedName` (string)
        :param pulumi.Input[str] group_name_attribute: Group name attribute. Default `name` (string)
        :param pulumi.Input[str] group_object_class: Group object class. Default `group` (string)
        :param pulumi.Input[str] group_search_attribute: Group search attribute. Default `sAMAccountName` (string)
        :param pulumi.Input[str] group_search_base: Group search base (string)
        :param pulumi.Input[str] group_search_filter: Group search filter (string)
        :param pulumi.Input[dict] labels: Labels of the resource (map)
        :param pulumi.Input[bool] nested_group_membership_enabled: Nested group membership enable. Default `false` (bool)
        :param pulumi.Input[float] port: ActiveDirectory port. Default `389` (int)
        :param pulumi.Input[list] servers: ActiveDirectory servers list (list)
        :param pulumi.Input[str] service_account_password: Service account password for access ActiveDirectory service (string)
        :param pulumi.Input[str] service_account_username: Service account DN for access ActiveDirectory service (string)
        :param pulumi.Input[bool] tls: Enable TLS connection (bool)
        :param pulumi.Input[float] user_disabled_bit_mask: User disabled bit mask. Default `2` (int)
        :param pulumi.Input[str] user_enabled_attribute: User enable attribute (string)
        :param pulumi.Input[str] user_login_attribute: User login attribute. Default `sAMAccountName` (string)
        :param pulumi.Input[str] user_name_attribute: User name attribute. Default `name` (string)
        :param pulumi.Input[str] user_object_class: User object class. Default `person` (string)
        :param pulumi.Input[str] user_search_attribute: User search attribute. Default `sAMAccountName|sn|givenName` (string)
        :param pulumi.Input[str] user_search_base: User search base DN (string)
        :param pulumi.Input[str] user_search_filter: User search filter (string)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/auth_config_activedirectory.html.markdown.
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

            __props__['access_mode'] = access_mode
            __props__['allowed_principal_ids'] = allowed_principal_ids
            __props__['annotations'] = annotations
            __props__['certificate'] = certificate
            __props__['connection_timeout'] = connection_timeout
            __props__['default_login_domain'] = default_login_domain
            __props__['enabled'] = enabled
            __props__['group_dn_attribute'] = group_dn_attribute
            __props__['group_member_mapping_attribute'] = group_member_mapping_attribute
            __props__['group_member_user_attribute'] = group_member_user_attribute
            __props__['group_name_attribute'] = group_name_attribute
            __props__['group_object_class'] = group_object_class
            __props__['group_search_attribute'] = group_search_attribute
            __props__['group_search_base'] = group_search_base
            __props__['group_search_filter'] = group_search_filter
            __props__['labels'] = labels
            __props__['nested_group_membership_enabled'] = nested_group_membership_enabled
            __props__['port'] = port
            if servers is None:
                raise TypeError("Missing required property 'servers'")
            __props__['servers'] = servers
            if service_account_password is None:
                raise TypeError("Missing required property 'service_account_password'")
            __props__['service_account_password'] = service_account_password
            if service_account_username is None:
                raise TypeError("Missing required property 'service_account_username'")
            __props__['service_account_username'] = service_account_username
            __props__['tls'] = tls
            __props__['user_disabled_bit_mask'] = user_disabled_bit_mask
            __props__['user_enabled_attribute'] = user_enabled_attribute
            __props__['user_login_attribute'] = user_login_attribute
            __props__['user_name_attribute'] = user_name_attribute
            __props__['user_object_class'] = user_object_class
            __props__['user_search_attribute'] = user_search_attribute
            if user_search_base is None:
                raise TypeError("Missing required property 'user_search_base'")
            __props__['user_search_base'] = user_search_base
            __props__['user_search_filter'] = user_search_filter
            __props__['name'] = None
            __props__['type'] = None
        super(ActiveDirectory, __self__).__init__(
            'rancher2:index/activeDirectory:ActiveDirectory',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_mode=None, allowed_principal_ids=None, annotations=None, certificate=None, connection_timeout=None, default_login_domain=None, enabled=None, group_dn_attribute=None, group_member_mapping_attribute=None, group_member_user_attribute=None, group_name_attribute=None, group_object_class=None, group_search_attribute=None, group_search_base=None, group_search_filter=None, labels=None, name=None, nested_group_membership_enabled=None, port=None, servers=None, service_account_password=None, service_account_username=None, tls=None, type=None, user_disabled_bit_mask=None, user_enabled_attribute=None, user_login_attribute=None, user_name_attribute=None, user_object_class=None, user_search_attribute=None, user_search_base=None, user_search_filter=None):
        """
        Get an existing ActiveDirectory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_mode: Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
        :param pulumi.Input[list] allowed_principal_ids: Allowed principal ids for auth. Required if `access_mode` is `required` or `restricted`. Ex: `activedirectory_user://<DN>`  `activedirectory_group://<DN>` (list)
        :param pulumi.Input[dict] annotations: Annotations of the resource (map)
        :param pulumi.Input[str] certificate: CA certificate for TLS if selfsigned (string)
        :param pulumi.Input[float] connection_timeout: ActiveDirectory connection timeout. Default `5000` (int)
        :param pulumi.Input[str] default_login_domain: ActiveDirectory defult lgoin domain (string)
        :param pulumi.Input[bool] enabled: Enable auth config provider. Default `true` (bool)
        :param pulumi.Input[str] group_dn_attribute: Group DN attribute. Default `distinguishedName` (string)
        :param pulumi.Input[str] group_member_mapping_attribute: Group member mapping attribute. Default `member` (string)
        :param pulumi.Input[str] group_member_user_attribute: Group member user attribute. Default `distinguishedName` (string)
        :param pulumi.Input[str] group_name_attribute: Group name attribute. Default `name` (string)
        :param pulumi.Input[str] group_object_class: Group object class. Default `group` (string)
        :param pulumi.Input[str] group_search_attribute: Group search attribute. Default `sAMAccountName` (string)
        :param pulumi.Input[str] group_search_base: Group search base (string)
        :param pulumi.Input[str] group_search_filter: Group search filter (string)
        :param pulumi.Input[dict] labels: Labels of the resource (map)
        :param pulumi.Input[str] name: (Computed) The name of the resource (string)
        :param pulumi.Input[bool] nested_group_membership_enabled: Nested group membership enable. Default `false` (bool)
        :param pulumi.Input[float] port: ActiveDirectory port. Default `389` (int)
        :param pulumi.Input[list] servers: ActiveDirectory servers list (list)
        :param pulumi.Input[str] service_account_password: Service account password for access ActiveDirectory service (string)
        :param pulumi.Input[str] service_account_username: Service account DN for access ActiveDirectory service (string)
        :param pulumi.Input[bool] tls: Enable TLS connection (bool)
        :param pulumi.Input[str] type: (Computed) The type of the resource (string)
        :param pulumi.Input[float] user_disabled_bit_mask: User disabled bit mask. Default `2` (int)
        :param pulumi.Input[str] user_enabled_attribute: User enable attribute (string)
        :param pulumi.Input[str] user_login_attribute: User login attribute. Default `sAMAccountName` (string)
        :param pulumi.Input[str] user_name_attribute: User name attribute. Default `name` (string)
        :param pulumi.Input[str] user_object_class: User object class. Default `person` (string)
        :param pulumi.Input[str] user_search_attribute: User search attribute. Default `sAMAccountName|sn|givenName` (string)
        :param pulumi.Input[str] user_search_base: User search base DN (string)
        :param pulumi.Input[str] user_search_filter: User search filter (string)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/auth_config_activedirectory.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["access_mode"] = access_mode
        __props__["allowed_principal_ids"] = allowed_principal_ids
        __props__["annotations"] = annotations
        __props__["certificate"] = certificate
        __props__["connection_timeout"] = connection_timeout
        __props__["default_login_domain"] = default_login_domain
        __props__["enabled"] = enabled
        __props__["group_dn_attribute"] = group_dn_attribute
        __props__["group_member_mapping_attribute"] = group_member_mapping_attribute
        __props__["group_member_user_attribute"] = group_member_user_attribute
        __props__["group_name_attribute"] = group_name_attribute
        __props__["group_object_class"] = group_object_class
        __props__["group_search_attribute"] = group_search_attribute
        __props__["group_search_base"] = group_search_base
        __props__["group_search_filter"] = group_search_filter
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["nested_group_membership_enabled"] = nested_group_membership_enabled
        __props__["port"] = port
        __props__["servers"] = servers
        __props__["service_account_password"] = service_account_password
        __props__["service_account_username"] = service_account_username
        __props__["tls"] = tls
        __props__["type"] = type
        __props__["user_disabled_bit_mask"] = user_disabled_bit_mask
        __props__["user_enabled_attribute"] = user_enabled_attribute
        __props__["user_login_attribute"] = user_login_attribute
        __props__["user_name_attribute"] = user_name_attribute
        __props__["user_object_class"] = user_object_class
        __props__["user_search_attribute"] = user_search_attribute
        __props__["user_search_base"] = user_search_base
        __props__["user_search_filter"] = user_search_filter
        return ActiveDirectory(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

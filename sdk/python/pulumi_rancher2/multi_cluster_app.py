# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class MultiClusterApp(pulumi.CustomResource):
    annotations: pulumi.Output[dict]
    """
    Annotations for multi cluster app object (map)
    """
    answers: pulumi.Output[list]
    """
    The multi cluster app answers (list)

      * `cluster_id` (`str`) - Cluster ID for answer (string)
      * `project_id` (`str`) - Project ID for target (string)
      * `values` (`dict`) - Key/values for answer (map)
    """
    catalog_name: pulumi.Output[str]
    """
    The multi cluster app catalog name (string)
    """
    labels: pulumi.Output[dict]
    """
    Labels for multi cluster app object (map)
    """
    members: pulumi.Output[list]
    """
    The multi cluster app answers (list)

      * `accessType` (`str`) - Member access type. Valid values: `["member" | "owner" | "read-only"]` (string)
      * `group_principal_id` (`str`) - Member group principal id (string)
      * `user_principal_id` (`str`) - Member user principal id (string)
    """
    name: pulumi.Output[str]
    """
    The multi cluster app name (string)
    """
    revision_history_limit: pulumi.Output[float]
    """
    The multi cluster app revision history limit. Default `10` (int)
    """
    revision_id: pulumi.Output[str]
    """
    Current revision id for the multi cluster app (string)
    """
    roles: pulumi.Output[list]
    """
    The multi cluster app roles (list)
    """
    targets: pulumi.Output[list]
    """
    The multi cluster app target projects (list)

      * `appId` (`str`) - App ID for target (string)
      * `healthState` (`str`) - App health state for target (string)
      * `project_id` (`str`) - Project ID for target (string)
      * `state` (`str`) - App state for target (string)
    """
    template_name: pulumi.Output[str]
    """
    The multi cluster app template name (string)
    """
    template_version: pulumi.Output[str]
    """
    The multi cluster app template version. Default: `latest` (string)
    """
    template_version_id: pulumi.Output[str]
    """
    (Computed) The multi cluster app template version ID (string)
    """
    upgrade_strategy: pulumi.Output[dict]
    """
    The multi cluster app upgrade strategy (list MaxItems:1)

      * `rollingUpdate` (`dict`) - Upgrade strategy rolling update (list MaxItems:1)
        * `batchSize` (`float`) - Rolling update batch size. Default `1` (int)
        * `interval` (`float`) - Rolling update interval. Default `1` (int)
    """
    wait: pulumi.Output[bool]
    """
    Wait until the multi cluster app is active. Default `true` (bool)
    """
    def __init__(__self__, resource_name, opts=None, annotations=None, answers=None, catalog_name=None, labels=None, members=None, name=None, revision_history_limit=None, revision_id=None, roles=None, targets=None, template_name=None, template_version=None, upgrade_strategy=None, wait=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a MultiClusterApp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] annotations: Annotations for multi cluster app object (map)
        :param pulumi.Input[list] answers: The multi cluster app answers (list)
        :param pulumi.Input[str] catalog_name: The multi cluster app catalog name (string)
        :param pulumi.Input[dict] labels: Labels for multi cluster app object (map)
        :param pulumi.Input[list] members: The multi cluster app answers (list)
        :param pulumi.Input[str] name: The multi cluster app name (string)
        :param pulumi.Input[float] revision_history_limit: The multi cluster app revision history limit. Default `10` (int)
        :param pulumi.Input[str] revision_id: Current revision id for the multi cluster app (string)
        :param pulumi.Input[list] roles: The multi cluster app roles (list)
        :param pulumi.Input[list] targets: The multi cluster app target projects (list)
        :param pulumi.Input[str] template_name: The multi cluster app template name (string)
        :param pulumi.Input[str] template_version: The multi cluster app template version. Default: `latest` (string)
        :param pulumi.Input[dict] upgrade_strategy: The multi cluster app upgrade strategy (list MaxItems:1)
        :param pulumi.Input[bool] wait: Wait until the multi cluster app is active. Default `true` (bool)

        The **answers** object supports the following:

          * `cluster_id` (`pulumi.Input[str]`) - Cluster ID for answer (string)
          * `project_id` (`pulumi.Input[str]`) - Project ID for target (string)
          * `values` (`pulumi.Input[dict]`) - Key/values for answer (map)

        The **members** object supports the following:

          * `accessType` (`pulumi.Input[str]`) - Member access type. Valid values: `["member" | "owner" | "read-only"]` (string)
          * `group_principal_id` (`pulumi.Input[str]`) - Member group principal id (string)
          * `user_principal_id` (`pulumi.Input[str]`) - Member user principal id (string)

        The **targets** object supports the following:

          * `appId` (`pulumi.Input[str]`) - App ID for target (string)
          * `healthState` (`pulumi.Input[str]`) - App health state for target (string)
          * `project_id` (`pulumi.Input[str]`) - Project ID for target (string)
          * `state` (`pulumi.Input[str]`) - App state for target (string)

        The **upgrade_strategy** object supports the following:

          * `rollingUpdate` (`pulumi.Input[dict]`) - Upgrade strategy rolling update (list MaxItems:1)
            * `batchSize` (`pulumi.Input[float]`) - Rolling update batch size. Default `1` (int)
            * `interval` (`pulumi.Input[float]`) - Rolling update interval. Default `1` (int)
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
            __props__['answers'] = answers
            if catalog_name is None:
                raise TypeError("Missing required property 'catalog_name'")
            __props__['catalog_name'] = catalog_name
            __props__['labels'] = labels
            __props__['members'] = members
            __props__['name'] = name
            __props__['revision_history_limit'] = revision_history_limit
            __props__['revision_id'] = revision_id
            if roles is None:
                raise TypeError("Missing required property 'roles'")
            __props__['roles'] = roles
            if targets is None:
                raise TypeError("Missing required property 'targets'")
            __props__['targets'] = targets
            if template_name is None:
                raise TypeError("Missing required property 'template_name'")
            __props__['template_name'] = template_name
            __props__['template_version'] = template_version
            __props__['upgrade_strategy'] = upgrade_strategy
            __props__['wait'] = wait
            __props__['template_version_id'] = None
        super(MultiClusterApp, __self__).__init__(
            'rancher2:index/multiClusterApp:MultiClusterApp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, annotations=None, answers=None, catalog_name=None, labels=None, members=None, name=None, revision_history_limit=None, revision_id=None, roles=None, targets=None, template_name=None, template_version=None, template_version_id=None, upgrade_strategy=None, wait=None):
        """
        Get an existing MultiClusterApp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] annotations: Annotations for multi cluster app object (map)
        :param pulumi.Input[list] answers: The multi cluster app answers (list)
        :param pulumi.Input[str] catalog_name: The multi cluster app catalog name (string)
        :param pulumi.Input[dict] labels: Labels for multi cluster app object (map)
        :param pulumi.Input[list] members: The multi cluster app answers (list)
        :param pulumi.Input[str] name: The multi cluster app name (string)
        :param pulumi.Input[float] revision_history_limit: The multi cluster app revision history limit. Default `10` (int)
        :param pulumi.Input[str] revision_id: Current revision id for the multi cluster app (string)
        :param pulumi.Input[list] roles: The multi cluster app roles (list)
        :param pulumi.Input[list] targets: The multi cluster app target projects (list)
        :param pulumi.Input[str] template_name: The multi cluster app template name (string)
        :param pulumi.Input[str] template_version: The multi cluster app template version. Default: `latest` (string)
        :param pulumi.Input[str] template_version_id: (Computed) The multi cluster app template version ID (string)
        :param pulumi.Input[dict] upgrade_strategy: The multi cluster app upgrade strategy (list MaxItems:1)
        :param pulumi.Input[bool] wait: Wait until the multi cluster app is active. Default `true` (bool)

        The **answers** object supports the following:

          * `cluster_id` (`pulumi.Input[str]`) - Cluster ID for answer (string)
          * `project_id` (`pulumi.Input[str]`) - Project ID for target (string)
          * `values` (`pulumi.Input[dict]`) - Key/values for answer (map)

        The **members** object supports the following:

          * `accessType` (`pulumi.Input[str]`) - Member access type. Valid values: `["member" | "owner" | "read-only"]` (string)
          * `group_principal_id` (`pulumi.Input[str]`) - Member group principal id (string)
          * `user_principal_id` (`pulumi.Input[str]`) - Member user principal id (string)

        The **targets** object supports the following:

          * `appId` (`pulumi.Input[str]`) - App ID for target (string)
          * `healthState` (`pulumi.Input[str]`) - App health state for target (string)
          * `project_id` (`pulumi.Input[str]`) - Project ID for target (string)
          * `state` (`pulumi.Input[str]`) - App state for target (string)

        The **upgrade_strategy** object supports the following:

          * `rollingUpdate` (`pulumi.Input[dict]`) - Upgrade strategy rolling update (list MaxItems:1)
            * `batchSize` (`pulumi.Input[float]`) - Rolling update batch size. Default `1` (int)
            * `interval` (`pulumi.Input[float]`) - Rolling update interval. Default `1` (int)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["annotations"] = annotations
        __props__["answers"] = answers
        __props__["catalog_name"] = catalog_name
        __props__["labels"] = labels
        __props__["members"] = members
        __props__["name"] = name
        __props__["revision_history_limit"] = revision_history_limit
        __props__["revision_id"] = revision_id
        __props__["roles"] = roles
        __props__["targets"] = targets
        __props__["template_name"] = template_name
        __props__["template_version"] = template_version
        __props__["template_version_id"] = template_version_id
        __props__["upgrade_strategy"] = upgrade_strategy
        __props__["wait"] = wait
        return MultiClusterApp(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


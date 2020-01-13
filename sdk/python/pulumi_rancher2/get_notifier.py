# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetNotifierResult:
    """
    A collection of values returned by getNotifier.
    """
    def __init__(__self__, annotations=None, cluster_id=None, description=None, labels=None, name=None, pagerduty_config=None, slack_config=None, smtp_config=None, webhook_config=None, wechat_config=None, id=None):
        if annotations and not isinstance(annotations, dict):
            raise TypeError("Expected argument 'annotations' to be a dict")
        __self__.annotations = annotations
        """
        (Computed) Annotations for notifier object (map)
        """
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        __self__.cluster_id = cluster_id
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        (Computed) The notifier description (string)
        """
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        """
        (Computed) Labels for notifier object (map)
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if pagerduty_config and not isinstance(pagerduty_config, dict):
            raise TypeError("Expected argument 'pagerduty_config' to be a dict")
        __self__.pagerduty_config = pagerduty_config
        """
        (Computed) Pagerduty config for notifier (list maxitems:1)
        """
        if slack_config and not isinstance(slack_config, dict):
            raise TypeError("Expected argument 'slack_config' to be a dict")
        __self__.slack_config = slack_config
        """
        (Computed) Slack config for notifier (list maxitems:1)
        """
        if smtp_config and not isinstance(smtp_config, dict):
            raise TypeError("Expected argument 'smtp_config' to be a dict")
        __self__.smtp_config = smtp_config
        """
        (Computed) SMTP config for notifier (list maxitems:1)
        """
        if webhook_config and not isinstance(webhook_config, dict):
            raise TypeError("Expected argument 'webhook_config' to be a dict")
        __self__.webhook_config = webhook_config
        """
        (Computed) Webhook config for notifier (list maxitems:1)
        """
        if wechat_config and not isinstance(wechat_config, dict):
            raise TypeError("Expected argument 'wechat_config' to be a dict")
        __self__.wechat_config = wechat_config
        """
        (Computed) Wechat config for notifier (list maxitems:1)
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetNotifierResult(GetNotifierResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNotifierResult(
            annotations=self.annotations,
            cluster_id=self.cluster_id,
            description=self.description,
            labels=self.labels,
            name=self.name,
            pagerduty_config=self.pagerduty_config,
            slack_config=self.slack_config,
            smtp_config=self.smtp_config,
            webhook_config=self.webhook_config,
            wechat_config=self.wechat_config,
            id=self.id)

def get_notifier(cluster_id=None,name=None,opts=None):
    """
    Use this data source to retrieve information about a Rancher v2 notifier.
    
    :param str cluster_id: The cluster id where create notifier (string)
    :param str name: The name of the notifier (string)

    > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/notifier.html.markdown.
    """
    __args__ = dict()

    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('rancher2:index/getNotifier:getNotifier', __args__, opts=opts).value

    return AwaitableGetNotifierResult(
        annotations=__ret__.get('annotations'),
        cluster_id=__ret__.get('clusterId'),
        description=__ret__.get('description'),
        labels=__ret__.get('labels'),
        name=__ret__.get('name'),
        pagerduty_config=__ret__.get('pagerdutyConfig'),
        slack_config=__ret__.get('slackConfig'),
        smtp_config=__ret__.get('smtpConfig'),
        webhook_config=__ret__.get('webhookConfig'),
        wechat_config=__ret__.get('wechatConfig'),
        id=__ret__.get('id'))

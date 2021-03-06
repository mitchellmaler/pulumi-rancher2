# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetProjectAlertGroupResult:
    """
    A collection of values returned by getProjectAlertGroup.
    """
    def __init__(__self__, annotations=None, description=None, group_interval_seconds=None, group_wait_seconds=None, id=None, labels=None, name=None, project_id=None, recipients=None, repeat_interval_seconds=None):
        if annotations and not isinstance(annotations, dict):
            raise TypeError("Expected argument 'annotations' to be a dict")
        __self__.annotations = annotations
        """
        (Computed) The project alert group annotations (map)
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        (Computed) The project alert group description (string)
        """
        if group_interval_seconds and not isinstance(group_interval_seconds, float):
            raise TypeError("Expected argument 'group_interval_seconds' to be a float")
        __self__.group_interval_seconds = group_interval_seconds
        """
        (Computed) The project alert group interval seconds. Default: `180` (int)
        """
        if group_wait_seconds and not isinstance(group_wait_seconds, float):
            raise TypeError("Expected argument 'group_wait_seconds' to be a float")
        __self__.group_wait_seconds = group_wait_seconds
        """
        (Computed) The project alert group wait seconds. Default: `180` (int)
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        """
        (Computed) The project alert group labels (map)
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        if recipients and not isinstance(recipients, list):
            raise TypeError("Expected argument 'recipients' to be a list")
        __self__.recipients = recipients
        """
        (Computed) The project alert group recipients (list)
        """
        if repeat_interval_seconds and not isinstance(repeat_interval_seconds, float):
            raise TypeError("Expected argument 'repeat_interval_seconds' to be a float")
        __self__.repeat_interval_seconds = repeat_interval_seconds
        """
        (Computed) The project alert group wait seconds. Default: `3600` (int)
        """
class AwaitableGetProjectAlertGroupResult(GetProjectAlertGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectAlertGroupResult(
            annotations=self.annotations,
            description=self.description,
            group_interval_seconds=self.group_interval_seconds,
            group_wait_seconds=self.group_wait_seconds,
            id=self.id,
            labels=self.labels,
            name=self.name,
            project_id=self.project_id,
            recipients=self.recipients,
            repeat_interval_seconds=self.repeat_interval_seconds)

def get_project_alert_group(name=None,project_id=None,opts=None):
    """
    Use this data source to retrieve information about a Rancher v2 project alert group.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/projectAlertGroup.html.markdown.


    :param str name: The project alert group name (string)
    :param str project_id: The project id where create project alert group (string)
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('rancher2:index/getProjectAlertGroup:getProjectAlertGroup', __args__, opts=opts).value

    return AwaitableGetProjectAlertGroupResult(
        annotations=__ret__.get('annotations'),
        description=__ret__.get('description'),
        group_interval_seconds=__ret__.get('groupIntervalSeconds'),
        group_wait_seconds=__ret__.get('groupWaitSeconds'),
        id=__ret__.get('id'),
        labels=__ret__.get('labels'),
        name=__ret__.get('name'),
        project_id=__ret__.get('projectId'),
        recipients=__ret__.get('recipients'),
        repeat_interval_seconds=__ret__.get('repeatIntervalSeconds'))

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, annotations=None, enabled=None, id=None, is_external=None, labels=None, name=None, principal_ids=None, username=None):
        if annotations and not isinstance(annotations, dict):
            raise TypeError("Expected argument 'annotations' to be a dict")
        __self__.annotations = annotations
        """
        (Computed) Annotations of the resource (map)
        """
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        """
        (Computed) The user is enabled (bool)
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if is_external and not isinstance(is_external, bool):
            raise TypeError("Expected argument 'is_external' to be a bool")
        __self__.is_external = is_external
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        """
        (Computed) Labels of the resource (map)
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        (Computed) The user common name (string)
        """
        if principal_ids and not isinstance(principal_ids, list):
            raise TypeError("Expected argument 'principal_ids' to be a list")
        __self__.principal_ids = principal_ids
        """
        (Computed) The user principal IDs (list)
        """
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        __self__.username = username
class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            annotations=self.annotations,
            enabled=self.enabled,
            id=self.id,
            is_external=self.is_external,
            labels=self.labels,
            name=self.name,
            principal_ids=self.principal_ids,
            username=self.username)

def get_user(is_external=None,name=None,username=None,opts=None):
    """
    Use this data source to retrieve information about a Rancher v2 user

    > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/user.html.markdown.


    :param bool is_external: Set is the user if the user is external. Default: `false` (bool)
    :param str name: The name of the user (string)
    :param str username: The username of the user (string)
    """
    __args__ = dict()


    __args__['isExternal'] = is_external
    __args__['name'] = name
    __args__['username'] = username
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('rancher2:index/getUser:getUser', __args__, opts=opts).value

    return AwaitableGetUserResult(
        annotations=__ret__.get('annotations'),
        enabled=__ret__.get('enabled'),
        id=__ret__.get('id'),
        is_external=__ret__.get('isExternal'),
        labels=__ret__.get('labels'),
        name=__ret__.get('name'),
        principal_ids=__ret__.get('principalIds'),
        username=__ret__.get('username'))

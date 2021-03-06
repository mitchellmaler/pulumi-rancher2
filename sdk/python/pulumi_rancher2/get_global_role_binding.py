# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetGlobalRoleBindingResult:
    """
    A collection of values returned by getGlobalRoleBinding.
    """
    def __init__(__self__, annotations=None, global_role_id=None, id=None, labels=None, name=None, user_id=None):
        if annotations and not isinstance(annotations, dict):
            raise TypeError("Expected argument 'annotations' to be a dict")
        __self__.annotations = annotations
        """
        (Computed) Annotations of the resource (map)
        """
        if global_role_id and not isinstance(global_role_id, str):
            raise TypeError("Expected argument 'global_role_id' to be a str")
        __self__.global_role_id = global_role_id
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
        (Computed) Labels of the resource (map)
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        __self__.user_id = user_id
        """
        (Computed) The user ID to assign global role binding (string)
        """
class AwaitableGetGlobalRoleBindingResult(GetGlobalRoleBindingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGlobalRoleBindingResult(
            annotations=self.annotations,
            global_role_id=self.global_role_id,
            id=self.id,
            labels=self.labels,
            name=self.name,
            user_id=self.user_id)

def get_global_role_binding(global_role_id=None,name=None,opts=None):
    """
    Use this data source to retrieve information about a Rancher v2 global role binding.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/globalRole.html.markdown.


    :param str global_role_id: The global role id (string)
    :param str name: The name of the global role binding (string)
    """
    __args__ = dict()


    __args__['globalRoleId'] = global_role_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('rancher2:index/getGlobalRoleBinding:getGlobalRoleBinding', __args__, opts=opts).value

    return AwaitableGetGlobalRoleBindingResult(
        annotations=__ret__.get('annotations'),
        global_role_id=__ret__.get('globalRoleId'),
        id=__ret__.get('id'),
        labels=__ret__.get('labels'),
        name=__ret__.get('name'),
        user_id=__ret__.get('userId'))

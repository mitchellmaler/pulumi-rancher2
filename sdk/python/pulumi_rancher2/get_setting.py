# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetSettingResult:
    """
    A collection of values returned by getSetting.
    """
    def __init__(__self__, name=None, value=None, id=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        __self__.value = value
        """
        the settting's value.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetSettingResult(GetSettingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSettingResult(
            name=self.name,
            value=self.value,
            id=self.id)

def get_setting(name=None,opts=None):
    """
    Use this data source to retrieve information about a Rancher v2 setting.
    
    :param str name: The setting name.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/setting.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('rancher2:index/getSetting:getSetting', __args__, opts=opts).value

    return AwaitableGetSettingResult(
        name=__ret__.get('name'),
        value=__ret__.get('value'),
        id=__ret__.get('id'))

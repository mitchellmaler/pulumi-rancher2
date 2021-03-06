# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

__config__ = pulumi.Config('rancher2')

access_key = __config__.get('accessKey') or utilities.get_env('RANCHER_ACCESS_KEY')
"""
API Key used to authenticate with the rancher server
"""

api_url = __config__.get('apiUrl') or utilities.get_env('RANCHER_URL')
"""
The URL to the rancher API
"""

bootstrap = __config__.get('bootstrap') or (utilities.get_env_bool('RANCHER_BOOTSTRAP') or False)
"""
Bootstrap rancher server
"""

ca_certs = __config__.get('caCerts') or utilities.get_env('RANCHER_CA_CERTS')
"""
CA certificates used to sign rancher server tls certificates. Mandatory if self signed tls and insecure option false
"""

insecure = __config__.get('insecure') or (utilities.get_env_bool('RANCHER_INSECURE') or False)
"""
Allow insecure connections to Rancher. Mandatory if self signed tls and not ca_certs provided
"""

secret_key = __config__.get('secretKey') or utilities.get_env('RANCHER_SECRET_KEY')
"""
API secret used to authenticate with the rancher server
"""

token_key = __config__.get('tokenKey') or utilities.get_env('RANCHER_TOKEN_KEY')
"""
API token used to authenticate with the rancher server
"""


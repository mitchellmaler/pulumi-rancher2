[![Build Status](https://travis-ci.com/pulumi/pulumi-rancher2.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-rancher2)

# Rancher2 Resource Provider

The Rancher2 resource provider for Pulumi lets you manage Rancher2 resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/rancher2

or `yarn`:

    $ yarn add @pulumi/rancher2

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_rancher2

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-rancher2/sdk/go/...

## Configuration

The following configuration points are available for the `xyz` provider:

- `rancher2:apiUrl` (Required) - The URL to the rancher API. It must be provided, but it can also be sourced from the 
  `RANCHER_URL` environment variable.
- `rancher2:accessKey` (Optional) - API Key used to authenticate with the rancher server. It can also be sourced from the
  `RANCHER_ACCESS_KEY` environment variable.
- `rancher2:secretKey` (Optional) - API secret used to authenticate with the rancher server. It can also be sourced from
  the `RANCHER_SECRET_KEY` environment variable.
- `rancher2:tokenKey` (Optional) - API token used to authenticate with the rancher server. It can also be sourced from 
  the `RANCHER_TOKEN_KEY` environment variable.
- `rancher2:caCerts` (Optional) - CA certificates used to sign rancher server tls certificates. Mandatory if self signed
  tls and insecure option false. It can also be sourced from the `RANCHER_CA_CERTS` environment variable.
- `rancher2:bootstrap` (Optional) - Bootstrap rancher server. Default value is `false`. It can also be sourced from the
  `RANCHER_BOOTSTRAP` environment variable.
- `rancher2:insecure` (Optional) - Allow insecure connections to Rancher. Mandatory if self signed tls and no caCerts
  provided. Default value is `false`. It can also be sourced from the `RANCHER_INSECURE` environment variable.


## Reference

For detailed reference documentation, please visit [the API docs][1].


[1]: https://pulumi.io/reference/pkg/nodejs/pulumi/rancher2/

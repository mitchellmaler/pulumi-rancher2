// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rancher2
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve information about a Rancher v2 setting.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/setting.html.markdown.
        /// </summary>
        public static Task<GetSettingResult> GetSetting(GetSettingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSettingResult>("rancher2:index/getSetting:getSetting", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSettingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The setting name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetSettingArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSettingResult
    {
        public readonly string Name;
        /// <summary>
        /// the settting's value.
        /// </summary>
        public readonly string Value;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSettingResult(
            string name,
            string value,
            string id)
        {
            Name = name;
            Value = value;
            Id = id;
        }
    }
}

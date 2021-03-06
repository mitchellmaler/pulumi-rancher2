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
        /// Use this data source to retrieve information about a Rancher v2 cluster alert group.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/clusterAlertGroup.html.markdown.
        /// </summary>
        public static Task<GetClusterAlertGroupResult> GetClusterAlertGroup(GetClusterAlertGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterAlertGroupResult>("rancher2:index/getClusterAlertGroup:getClusterAlertGroup", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterAlertGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster id where create cluster alert group (string)
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The cluster alert group name (string)
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetClusterAlertGroupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterAlertGroupResult
    {
        /// <summary>
        /// (Computed) The cluster alert group annotations (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Annotations;
        public readonly string ClusterId;
        /// <summary>
        /// (Computed) The cluster alert group description (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Computed) The cluster alert group interval seconds. Default: `180` (int)
        /// </summary>
        public readonly int GroupIntervalSeconds;
        /// <summary>
        /// (Computed) The cluster alert group wait seconds. Default: `180` (int)
        /// </summary>
        public readonly int GroupWaitSeconds;
        /// <summary>
        /// (Computed) The cluster alert group labels (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        public readonly string Name;
        /// <summary>
        /// (Computed) The cluster alert group recipients (list)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterAlertGroupRecipientsResult> Recipients;
        /// <summary>
        /// (Computed) The cluster alert group wait seconds. Default: `3600` (int)
        /// </summary>
        public readonly int RepeatIntervalSeconds;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterAlertGroupResult(
            ImmutableDictionary<string, object> annotations,
            string clusterId,
            string description,
            int groupIntervalSeconds,
            int groupWaitSeconds,
            ImmutableDictionary<string, object> labels,
            string name,
            ImmutableArray<Outputs.GetClusterAlertGroupRecipientsResult> recipients,
            int repeatIntervalSeconds,
            string id)
        {
            Annotations = annotations;
            ClusterId = clusterId;
            Description = description;
            GroupIntervalSeconds = groupIntervalSeconds;
            GroupWaitSeconds = groupWaitSeconds;
            Labels = labels;
            Name = name;
            Recipients = recipients;
            RepeatIntervalSeconds = repeatIntervalSeconds;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetClusterAlertGroupRecipientsResult
    {
        public readonly string NotifierId;
        public readonly string NotifierType;
        public readonly string Recipient;

        [OutputConstructor]
        private GetClusterAlertGroupRecipientsResult(
            string notifierId,
            string notifierType,
            string recipient)
        {
            NotifierId = notifierId;
            NotifierType = notifierType;
            Recipient = recipient;
        }
    }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rancher2
{
    public static partial class Invokes
    {
        public static Task<GetClusterAlterGroupResult> GetClusterAlterGroup(GetClusterAlterGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterAlterGroupResult>("rancher2:index/getClusterAlterGroup:getClusterAlterGroup", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterAlterGroupArgs : Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetClusterAlterGroupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterAlterGroupResult
    {
        public readonly ImmutableDictionary<string, object> Annotations;
        public readonly string ClusterId;
        public readonly string Description;
        public readonly int GroupIntervalSeconds;
        public readonly int GroupWaitSeconds;
        public readonly ImmutableDictionary<string, object> Labels;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetClusterAlterGroupRecipientsResult> Recipients;
        public readonly int RepeatIntervalSeconds;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterAlterGroupResult(
            ImmutableDictionary<string, object> annotations,
            string clusterId,
            string description,
            int groupIntervalSeconds,
            int groupWaitSeconds,
            ImmutableDictionary<string, object> labels,
            string name,
            ImmutableArray<Outputs.GetClusterAlterGroupRecipientsResult> recipients,
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
    public sealed class GetClusterAlterGroupRecipientsResult
    {
        public readonly string NotifierId;
        public readonly string NotifierType;
        public readonly string Recipient;

        [OutputConstructor]
        private GetClusterAlterGroupRecipientsResult(
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

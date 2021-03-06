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
        /// Use this data source to retrieve information about a Rancher v2 project alert group.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/projectAlertGroup.html.markdown.
        /// </summary>
        public static Task<GetProjectAlertGroupResult> GetProjectAlertGroup(GetProjectAlertGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectAlertGroupResult>("rancher2:index/getProjectAlertGroup:getProjectAlertGroup", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetProjectAlertGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project alert group name (string)
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project id where create project alert group (string)
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetProjectAlertGroupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetProjectAlertGroupResult
    {
        /// <summary>
        /// (Computed) The project alert group annotations (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Annotations;
        /// <summary>
        /// (Computed) The project alert group description (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Computed) The project alert group interval seconds. Default: `180` (int)
        /// </summary>
        public readonly int GroupIntervalSeconds;
        /// <summary>
        /// (Computed) The project alert group wait seconds. Default: `180` (int)
        /// </summary>
        public readonly int GroupWaitSeconds;
        /// <summary>
        /// (Computed) The project alert group labels (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        public readonly string Name;
        public readonly string ProjectId;
        /// <summary>
        /// (Computed) The project alert group recipients (list)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectAlertGroupRecipientsResult> Recipients;
        /// <summary>
        /// (Computed) The project alert group wait seconds. Default: `3600` (int)
        /// </summary>
        public readonly int RepeatIntervalSeconds;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetProjectAlertGroupResult(
            ImmutableDictionary<string, object> annotations,
            string description,
            int groupIntervalSeconds,
            int groupWaitSeconds,
            ImmutableDictionary<string, object> labels,
            string name,
            string projectId,
            ImmutableArray<Outputs.GetProjectAlertGroupRecipientsResult> recipients,
            int repeatIntervalSeconds,
            string id)
        {
            Annotations = annotations;
            Description = description;
            GroupIntervalSeconds = groupIntervalSeconds;
            GroupWaitSeconds = groupWaitSeconds;
            Labels = labels;
            Name = name;
            ProjectId = projectId;
            Recipients = recipients;
            RepeatIntervalSeconds = repeatIntervalSeconds;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetProjectAlertGroupRecipientsResult
    {
        public readonly string NotifierId;
        public readonly string NotifierType;
        public readonly string Recipient;

        [OutputConstructor]
        private GetProjectAlertGroupRecipientsResult(
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

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
        /// Use this data source to retrieve information about a Rancher v2 cluster alert rule.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/clusterAlertRule.html.markdown.
        /// </summary>
        public static Task<GetClusterAlterRuleResult> GetClusterAlterRule(GetClusterAlterRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterAlterRuleResult>("rancher2:index/getClusterAlterRule:getClusterAlterRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterAlterRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster id where create cluster alert rule (string)
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("labels")]
        private Dictionary<string, object>? _labels;
        public Dictionary<string, object> Labels
        {
            get => _labels ?? (_labels = new Dictionary<string, object>());
            set => _labels = value;
        }

        /// <summary>
        /// The cluster alert rule name (string)
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetClusterAlterRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterAlterRuleResult
    {
        /// <summary>
        /// (Computed) The cluster alert rule annotations (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Annotations;
        public readonly string ClusterId;
        /// <summary>
        /// (Computed) The cluster alert rule event rule. ConflictsWith: `"metric_rule", "node_rule", "system_service_rule"`` (list Maxitems:1)
        /// </summary>
        public readonly Outputs.GetClusterAlterRuleEventRuleResult EventRule;
        /// <summary>
        /// (Computed) The cluster alert rule alert group ID (string)
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// (Computed) The cluster alert rule group interval seconds. Default: `180` (int)
        /// </summary>
        public readonly int GroupIntervalSeconds;
        /// <summary>
        /// (Computed) The cluster alert rule group wait seconds. Default: `180` (int)
        /// </summary>
        public readonly int GroupWaitSeconds;
        /// <summary>
        /// (Computed) The cluster alert rule inherited. Default: `true` (bool)
        /// </summary>
        public readonly bool Inherited;
        /// <summary>
        /// (Computed) The cluster alert rule labels (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Labels;
        /// <summary>
        /// (Computed) The cluster alert rule metric rule. ConflictsWith: `"event_rule", "node_rule", "system_service_rule"`` (list Maxitems:1)
        /// </summary>
        public readonly Outputs.GetClusterAlterRuleMetricRuleResult MetricRule;
        public readonly string Name;
        /// <summary>
        /// (Computed) The cluster alert rule node rule. ConflictsWith: `"event_rule", "metric_rule", "system_service_rule"`` (list Maxitems:1)
        /// </summary>
        public readonly Outputs.GetClusterAlterRuleNodeRuleResult NodeRule;
        /// <summary>
        /// (Optional) The cluster alert rule wait seconds. Default: `3600` (int)
        /// </summary>
        public readonly int RepeatIntervalSeconds;
        /// <summary>
        /// (Computed) The cluster alert rule severity. Supported values : `"critical" | "info" | "warning"`. Default: `critical` (string)
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// (Computed) The cluster alert rule system service rule. ConflictsWith: `"event_rule", "metric_rule", "node_rule"`` (list Maxitems:1)
        /// </summary>
        public readonly Outputs.GetClusterAlterRuleSystemServiceRuleResult SystemServiceRule;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterAlterRuleResult(
            ImmutableDictionary<string, object> annotations,
            string clusterId,
            Outputs.GetClusterAlterRuleEventRuleResult eventRule,
            string groupId,
            int groupIntervalSeconds,
            int groupWaitSeconds,
            bool inherited,
            ImmutableDictionary<string, object>? labels,
            Outputs.GetClusterAlterRuleMetricRuleResult metricRule,
            string name,
            Outputs.GetClusterAlterRuleNodeRuleResult nodeRule,
            int repeatIntervalSeconds,
            string severity,
            Outputs.GetClusterAlterRuleSystemServiceRuleResult systemServiceRule,
            string id)
        {
            Annotations = annotations;
            ClusterId = clusterId;
            EventRule = eventRule;
            GroupId = groupId;
            GroupIntervalSeconds = groupIntervalSeconds;
            GroupWaitSeconds = groupWaitSeconds;
            Inherited = inherited;
            Labels = labels;
            MetricRule = metricRule;
            Name = name;
            NodeRule = nodeRule;
            RepeatIntervalSeconds = repeatIntervalSeconds;
            Severity = severity;
            SystemServiceRule = systemServiceRule;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetClusterAlterRuleEventRuleResult
    {
        public readonly string? EventType;
        public readonly string ResourceKind;

        [OutputConstructor]
        private GetClusterAlterRuleEventRuleResult(
            string? eventType,
            string resourceKind)
        {
            EventType = eventType;
            ResourceKind = resourceKind;
        }
    }

    [OutputType]
    public sealed class GetClusterAlterRuleMetricRuleResult
    {
        public readonly string? Comparison;
        public readonly string? Description;
        public readonly string Duration;
        public readonly string Expression;
        public readonly double ThresholdValue;

        [OutputConstructor]
        private GetClusterAlterRuleMetricRuleResult(
            string? comparison,
            string? description,
            string duration,
            string expression,
            double thresholdValue)
        {
            Comparison = comparison;
            Description = description;
            Duration = duration;
            Expression = expression;
            ThresholdValue = thresholdValue;
        }
    }

    [OutputType]
    public sealed class GetClusterAlterRuleNodeRuleResult
    {
        public readonly string? Condition;
        public readonly int? CpuThreshold;
        public readonly int? MemThreshold;
        public readonly string? NodeId;
        public readonly ImmutableDictionary<string, object>? Selector;

        [OutputConstructor]
        private GetClusterAlterRuleNodeRuleResult(
            string? condition,
            int? cpuThreshold,
            int? memThreshold,
            string? nodeId,
            ImmutableDictionary<string, object>? selector)
        {
            Condition = condition;
            CpuThreshold = cpuThreshold;
            MemThreshold = memThreshold;
            NodeId = nodeId;
            Selector = selector;
        }
    }

    [OutputType]
    public sealed class GetClusterAlterRuleSystemServiceRuleResult
    {
        public readonly string? Condition;

        [OutputConstructor]
        private GetClusterAlterRuleSystemServiceRuleResult(string? condition)
        {
            Condition = condition;
        }
    }
    }
}

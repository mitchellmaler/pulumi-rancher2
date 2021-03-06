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
        public static Task<GetProjectResult> GetProject(GetProjectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("rancher2:index/getProject:getProject", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the Rancher 2 cluster (string)
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The project name (string)
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetProjectArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// (Computed) Annotations of the rancher2 project (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Annotations;
        public readonly string ClusterId;
        /// <summary>
        /// (Computed) Default containers resource limits on project (List maxitem:1)
        /// </summary>
        public readonly Outputs.GetProjectContainerResourceLimitResult ContainerResourceLimit;
        /// <summary>
        /// (Computed) The project's description (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (Computed) Enable built-in project monitoring. Default `false` (bool)
        /// </summary>
        public readonly bool EnableProjectMonitoring;
        /// <summary>
        /// (Computed) Labels of the rancher2 project (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        public readonly string Name;
        /// <summary>
        /// (Computed) Default Pod Security Policy ID for the project (string)
        /// </summary>
        public readonly string PodSecurityPolicyTemplateId;
        /// <summary>
        /// (Computed) Resource quota for project. Rancher v2.1.x or higher (list maxitems:1)
        /// </summary>
        public readonly Outputs.GetProjectResourceQuotaResult ResourceQuota;
        /// <summary>
        /// (Computed) UUID of the project as stored by Rancher 2 (string)
        /// </summary>
        public readonly string Uuid;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetProjectResult(
            ImmutableDictionary<string, object> annotations,
            string clusterId,
            Outputs.GetProjectContainerResourceLimitResult containerResourceLimit,
            string description,
            bool enableProjectMonitoring,
            ImmutableDictionary<string, object> labels,
            string name,
            string podSecurityPolicyTemplateId,
            Outputs.GetProjectResourceQuotaResult resourceQuota,
            string uuid,
            string id)
        {
            Annotations = annotations;
            ClusterId = clusterId;
            ContainerResourceLimit = containerResourceLimit;
            Description = description;
            EnableProjectMonitoring = enableProjectMonitoring;
            Labels = labels;
            Name = name;
            PodSecurityPolicyTemplateId = podSecurityPolicyTemplateId;
            ResourceQuota = resourceQuota;
            Uuid = uuid;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetProjectContainerResourceLimitResult
    {
        public readonly string? LimitsCpu;
        public readonly string? LimitsMemory;
        public readonly string? RequestsCpu;
        public readonly string? RequestsMemory;

        [OutputConstructor]
        private GetProjectContainerResourceLimitResult(
            string? limitsCpu,
            string? limitsMemory,
            string? requestsCpu,
            string? requestsMemory)
        {
            LimitsCpu = limitsCpu;
            LimitsMemory = limitsMemory;
            RequestsCpu = requestsCpu;
            RequestsMemory = requestsMemory;
        }
    }

    [OutputType]
    public sealed class GetProjectResourceQuotaNamespaceDefaultLimitResult
    {
        public readonly string? ConfigMaps;
        public readonly string? LimitsCpu;
        public readonly string? LimitsMemory;
        public readonly string? PersistentVolumeClaims;
        public readonly string? Pods;
        public readonly string? ReplicationControllers;
        public readonly string? RequestsCpu;
        public readonly string? RequestsMemory;
        public readonly string? RequestsStorage;
        public readonly string? Secrets;
        public readonly string? Services;
        public readonly string? ServicesLoadBalancers;
        public readonly string? ServicesNodePorts;

        [OutputConstructor]
        private GetProjectResourceQuotaNamespaceDefaultLimitResult(
            string? configMaps,
            string? limitsCpu,
            string? limitsMemory,
            string? persistentVolumeClaims,
            string? pods,
            string? replicationControllers,
            string? requestsCpu,
            string? requestsMemory,
            string? requestsStorage,
            string? secrets,
            string? services,
            string? servicesLoadBalancers,
            string? servicesNodePorts)
        {
            ConfigMaps = configMaps;
            LimitsCpu = limitsCpu;
            LimitsMemory = limitsMemory;
            PersistentVolumeClaims = persistentVolumeClaims;
            Pods = pods;
            ReplicationControllers = replicationControllers;
            RequestsCpu = requestsCpu;
            RequestsMemory = requestsMemory;
            RequestsStorage = requestsStorage;
            Secrets = secrets;
            Services = services;
            ServicesLoadBalancers = servicesLoadBalancers;
            ServicesNodePorts = servicesNodePorts;
        }
    }

    [OutputType]
    public sealed class GetProjectResourceQuotaProjectLimitResult
    {
        public readonly string? ConfigMaps;
        public readonly string? LimitsCpu;
        public readonly string? LimitsMemory;
        public readonly string? PersistentVolumeClaims;
        public readonly string? Pods;
        public readonly string? ReplicationControllers;
        public readonly string? RequestsCpu;
        public readonly string? RequestsMemory;
        public readonly string? RequestsStorage;
        public readonly string? Secrets;
        public readonly string? Services;
        public readonly string? ServicesLoadBalancers;
        public readonly string? ServicesNodePorts;

        [OutputConstructor]
        private GetProjectResourceQuotaProjectLimitResult(
            string? configMaps,
            string? limitsCpu,
            string? limitsMemory,
            string? persistentVolumeClaims,
            string? pods,
            string? replicationControllers,
            string? requestsCpu,
            string? requestsMemory,
            string? requestsStorage,
            string? secrets,
            string? services,
            string? servicesLoadBalancers,
            string? servicesNodePorts)
        {
            ConfigMaps = configMaps;
            LimitsCpu = limitsCpu;
            LimitsMemory = limitsMemory;
            PersistentVolumeClaims = persistentVolumeClaims;
            Pods = pods;
            ReplicationControllers = replicationControllers;
            RequestsCpu = requestsCpu;
            RequestsMemory = requestsMemory;
            RequestsStorage = requestsStorage;
            Secrets = secrets;
            Services = services;
            ServicesLoadBalancers = servicesLoadBalancers;
            ServicesNodePorts = servicesNodePorts;
        }
    }

    [OutputType]
    public sealed class GetProjectResourceQuotaResult
    {
        public readonly GetProjectResourceQuotaNamespaceDefaultLimitResult NamespaceDefaultLimit;
        public readonly GetProjectResourceQuotaProjectLimitResult ProjectLimit;

        [OutputConstructor]
        private GetProjectResourceQuotaResult(
            GetProjectResourceQuotaNamespaceDefaultLimitResult namespaceDefaultLimit,
            GetProjectResourceQuotaProjectLimitResult projectLimit)
        {
            NamespaceDefaultLimit = namespaceDefaultLimit;
            ProjectLimit = projectLimit;
        }
    }
    }
}

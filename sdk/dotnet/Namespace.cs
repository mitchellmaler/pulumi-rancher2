// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rancher2
{
    /// <summary>
    /// Provides a Rancher v2 Namespace resource. This can be used to create namespaces for Rancher v2 environments and retrieve their information.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/namespace.html.markdown.
    /// </summary>
    public partial class Namespace : Pulumi.CustomResource
    {
        /// <summary>
        /// Annotations for Node Pool object (map)
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, object>> Annotations { get; private set; } = null!;

        /// <summary>
        /// Default containers resource limits on namespace (List maxitem:1)
        /// </summary>
        [Output("containerResourceLimit")]
        public Output<Outputs.NamespaceContainerResourceLimit?> ContainerResourceLimit { get; private set; } = null!;

        /// <summary>
        /// A namespace description (string)
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Labels for Node Pool object (map)
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>> Labels { get; private set; } = null!;

        /// <summary>
        /// The name of the namespace (string)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project id where assign namespace. It's on the form `project_id=&lt;cluster_id&gt;:&lt;id&gt;`. Updating `&lt;id&gt;` part on same `&lt;cluster_id&gt;` namespace will be moved between projects (string)
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Resource quota for namespace. Rancher v2.1.x or higher (list maxitems:1)
        /// </summary>
        [Output("resourceQuota")]
        public Output<Outputs.NamespaceResourceQuota?> ResourceQuota { get; private set; } = null!;

        /// <summary>
        /// Wait for cluster becomes active. Default `false` (bool)
        /// </summary>
        [Output("waitForCluster")]
        public Output<bool?> WaitForCluster { get; private set; } = null!;


        /// <summary>
        /// Create a Namespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Namespace(string name, NamespaceArgs args, CustomResourceOptions? options = null)
            : base("rancher2:index/namespace:Namespace", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Namespace(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
            : base("rancher2:index/namespace:Namespace", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Namespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Namespace Get(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Namespace(name, id, state, options);
        }
    }

    public sealed class NamespaceArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<object>? _annotations;

        /// <summary>
        /// Annotations for Node Pool object (map)
        /// </summary>
        public InputMap<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<object>());
            set => _annotations = value;
        }

        /// <summary>
        /// Default containers resource limits on namespace (List maxitem:1)
        /// </summary>
        [Input("containerResourceLimit")]
        public Input<Inputs.NamespaceContainerResourceLimitArgs>? ContainerResourceLimit { get; set; }

        /// <summary>
        /// A namespace description (string)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels for Node Pool object (map)
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the namespace (string)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project id where assign namespace. It's on the form `project_id=&lt;cluster_id&gt;:&lt;id&gt;`. Updating `&lt;id&gt;` part on same `&lt;cluster_id&gt;` namespace will be moved between projects (string)
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Resource quota for namespace. Rancher v2.1.x or higher (list maxitems:1)
        /// </summary>
        [Input("resourceQuota")]
        public Input<Inputs.NamespaceResourceQuotaArgs>? ResourceQuota { get; set; }

        /// <summary>
        /// Wait for cluster becomes active. Default `false` (bool)
        /// </summary>
        [Input("waitForCluster")]
        public Input<bool>? WaitForCluster { get; set; }

        public NamespaceArgs()
        {
        }
    }

    public sealed class NamespaceState : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<object>? _annotations;

        /// <summary>
        /// Annotations for Node Pool object (map)
        /// </summary>
        public InputMap<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<object>());
            set => _annotations = value;
        }

        /// <summary>
        /// Default containers resource limits on namespace (List maxitem:1)
        /// </summary>
        [Input("containerResourceLimit")]
        public Input<Inputs.NamespaceContainerResourceLimitGetArgs>? ContainerResourceLimit { get; set; }

        /// <summary>
        /// A namespace description (string)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels for Node Pool object (map)
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the namespace (string)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project id where assign namespace. It's on the form `project_id=&lt;cluster_id&gt;:&lt;id&gt;`. Updating `&lt;id&gt;` part on same `&lt;cluster_id&gt;` namespace will be moved between projects (string)
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Resource quota for namespace. Rancher v2.1.x or higher (list maxitems:1)
        /// </summary>
        [Input("resourceQuota")]
        public Input<Inputs.NamespaceResourceQuotaGetArgs>? ResourceQuota { get; set; }

        /// <summary>
        /// Wait for cluster becomes active. Default `false` (bool)
        /// </summary>
        [Input("waitForCluster")]
        public Input<bool>? WaitForCluster { get; set; }

        public NamespaceState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class NamespaceContainerResourceLimitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limit for limits cpu in namespace (string)
        /// </summary>
        [Input("limitsCpu")]
        public Input<string>? LimitsCpu { get; set; }

        /// <summary>
        /// Limit for limits memory in namespace (string)
        /// </summary>
        [Input("limitsMemory")]
        public Input<string>? LimitsMemory { get; set; }

        /// <summary>
        /// Limit for requests cpu in namespace (string)
        /// </summary>
        [Input("requestsCpu")]
        public Input<string>? RequestsCpu { get; set; }

        /// <summary>
        /// Limit for requests memory in namespace (string)
        /// </summary>
        [Input("requestsMemory")]
        public Input<string>? RequestsMemory { get; set; }

        public NamespaceContainerResourceLimitArgs()
        {
        }
    }

    public sealed class NamespaceContainerResourceLimitGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limit for limits cpu in namespace (string)
        /// </summary>
        [Input("limitsCpu")]
        public Input<string>? LimitsCpu { get; set; }

        /// <summary>
        /// Limit for limits memory in namespace (string)
        /// </summary>
        [Input("limitsMemory")]
        public Input<string>? LimitsMemory { get; set; }

        /// <summary>
        /// Limit for requests cpu in namespace (string)
        /// </summary>
        [Input("requestsCpu")]
        public Input<string>? RequestsCpu { get; set; }

        /// <summary>
        /// Limit for requests memory in namespace (string)
        /// </summary>
        [Input("requestsMemory")]
        public Input<string>? RequestsMemory { get; set; }

        public NamespaceContainerResourceLimitGetArgs()
        {
        }
    }

    public sealed class NamespaceResourceQuotaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource quota limit for namespace (list maxitems:1)
        /// </summary>
        [Input("limit", required: true)]
        public Input<NamespaceResourceQuotaLimitArgs> Limit { get; set; } = null!;

        public NamespaceResourceQuotaArgs()
        {
        }
    }

    public sealed class NamespaceResourceQuotaGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource quota limit for namespace (list maxitems:1)
        /// </summary>
        [Input("limit", required: true)]
        public Input<NamespaceResourceQuotaLimitGetArgs> Limit { get; set; } = null!;

        public NamespaceResourceQuotaGetArgs()
        {
        }
    }

    public sealed class NamespaceResourceQuotaLimitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limit for config maps in namespace (string)
        /// </summary>
        [Input("configMaps")]
        public Input<string>? ConfigMaps { get; set; }

        /// <summary>
        /// Limit for limits cpu in namespace (string)
        /// </summary>
        [Input("limitsCpu")]
        public Input<string>? LimitsCpu { get; set; }

        /// <summary>
        /// Limit for limits memory in namespace (string)
        /// </summary>
        [Input("limitsMemory")]
        public Input<string>? LimitsMemory { get; set; }

        /// <summary>
        /// Limit for persistent volume claims in namespace (string)
        /// </summary>
        [Input("persistentVolumeClaims")]
        public Input<string>? PersistentVolumeClaims { get; set; }

        /// <summary>
        /// Limit for pods in namespace (string)
        /// </summary>
        [Input("pods")]
        public Input<string>? Pods { get; set; }

        /// <summary>
        /// Limit for replication controllers in namespace (string)
        /// </summary>
        [Input("replicationControllers")]
        public Input<string>? ReplicationControllers { get; set; }

        /// <summary>
        /// Limit for requests cpu in namespace (string)
        /// </summary>
        [Input("requestsCpu")]
        public Input<string>? RequestsCpu { get; set; }

        /// <summary>
        /// Limit for requests memory in namespace (string)
        /// </summary>
        [Input("requestsMemory")]
        public Input<string>? RequestsMemory { get; set; }

        /// <summary>
        /// Limit for requests storage in namespace (string)
        /// </summary>
        [Input("requestsStorage")]
        public Input<string>? RequestsStorage { get; set; }

        /// <summary>
        /// Limit for secrets in namespace (string)
        /// </summary>
        [Input("secrets")]
        public Input<string>? Secrets { get; set; }

        [Input("services")]
        public Input<string>? Services { get; set; }

        /// <summary>
        /// Limit for services load balancers in namespace (string)
        /// </summary>
        [Input("servicesLoadBalancers")]
        public Input<string>? ServicesLoadBalancers { get; set; }

        /// <summary>
        /// Limit for services node ports in namespace (string)
        /// </summary>
        [Input("servicesNodePorts")]
        public Input<string>? ServicesNodePorts { get; set; }

        public NamespaceResourceQuotaLimitArgs()
        {
        }
    }

    public sealed class NamespaceResourceQuotaLimitGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limit for config maps in namespace (string)
        /// </summary>
        [Input("configMaps")]
        public Input<string>? ConfigMaps { get; set; }

        /// <summary>
        /// Limit for limits cpu in namespace (string)
        /// </summary>
        [Input("limitsCpu")]
        public Input<string>? LimitsCpu { get; set; }

        /// <summary>
        /// Limit for limits memory in namespace (string)
        /// </summary>
        [Input("limitsMemory")]
        public Input<string>? LimitsMemory { get; set; }

        /// <summary>
        /// Limit for persistent volume claims in namespace (string)
        /// </summary>
        [Input("persistentVolumeClaims")]
        public Input<string>? PersistentVolumeClaims { get; set; }

        /// <summary>
        /// Limit for pods in namespace (string)
        /// </summary>
        [Input("pods")]
        public Input<string>? Pods { get; set; }

        /// <summary>
        /// Limit for replication controllers in namespace (string)
        /// </summary>
        [Input("replicationControllers")]
        public Input<string>? ReplicationControllers { get; set; }

        /// <summary>
        /// Limit for requests cpu in namespace (string)
        /// </summary>
        [Input("requestsCpu")]
        public Input<string>? RequestsCpu { get; set; }

        /// <summary>
        /// Limit for requests memory in namespace (string)
        /// </summary>
        [Input("requestsMemory")]
        public Input<string>? RequestsMemory { get; set; }

        /// <summary>
        /// Limit for requests storage in namespace (string)
        /// </summary>
        [Input("requestsStorage")]
        public Input<string>? RequestsStorage { get; set; }

        /// <summary>
        /// Limit for secrets in namespace (string)
        /// </summary>
        [Input("secrets")]
        public Input<string>? Secrets { get; set; }

        [Input("services")]
        public Input<string>? Services { get; set; }

        /// <summary>
        /// Limit for services load balancers in namespace (string)
        /// </summary>
        [Input("servicesLoadBalancers")]
        public Input<string>? ServicesLoadBalancers { get; set; }

        /// <summary>
        /// Limit for services node ports in namespace (string)
        /// </summary>
        [Input("servicesNodePorts")]
        public Input<string>? ServicesNodePorts { get; set; }

        public NamespaceResourceQuotaLimitGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class NamespaceContainerResourceLimit
    {
        /// <summary>
        /// Limit for limits cpu in namespace (string)
        /// </summary>
        public readonly string? LimitsCpu;
        /// <summary>
        /// Limit for limits memory in namespace (string)
        /// </summary>
        public readonly string? LimitsMemory;
        /// <summary>
        /// Limit for requests cpu in namespace (string)
        /// </summary>
        public readonly string? RequestsCpu;
        /// <summary>
        /// Limit for requests memory in namespace (string)
        /// </summary>
        public readonly string? RequestsMemory;

        [OutputConstructor]
        private NamespaceContainerResourceLimit(
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
    public sealed class NamespaceResourceQuota
    {
        /// <summary>
        /// Resource quota limit for namespace (list maxitems:1)
        /// </summary>
        public readonly NamespaceResourceQuotaLimit Limit;

        [OutputConstructor]
        private NamespaceResourceQuota(NamespaceResourceQuotaLimit limit)
        {
            Limit = limit;
        }
    }

    [OutputType]
    public sealed class NamespaceResourceQuotaLimit
    {
        /// <summary>
        /// Limit for config maps in namespace (string)
        /// </summary>
        public readonly string? ConfigMaps;
        /// <summary>
        /// Limit for limits cpu in namespace (string)
        /// </summary>
        public readonly string? LimitsCpu;
        /// <summary>
        /// Limit for limits memory in namespace (string)
        /// </summary>
        public readonly string? LimitsMemory;
        /// <summary>
        /// Limit for persistent volume claims in namespace (string)
        /// </summary>
        public readonly string? PersistentVolumeClaims;
        /// <summary>
        /// Limit for pods in namespace (string)
        /// </summary>
        public readonly string? Pods;
        /// <summary>
        /// Limit for replication controllers in namespace (string)
        /// </summary>
        public readonly string? ReplicationControllers;
        /// <summary>
        /// Limit for requests cpu in namespace (string)
        /// </summary>
        public readonly string? RequestsCpu;
        /// <summary>
        /// Limit for requests memory in namespace (string)
        /// </summary>
        public readonly string? RequestsMemory;
        /// <summary>
        /// Limit for requests storage in namespace (string)
        /// </summary>
        public readonly string? RequestsStorage;
        /// <summary>
        /// Limit for secrets in namespace (string)
        /// </summary>
        public readonly string? Secrets;
        public readonly string? Services;
        /// <summary>
        /// Limit for services load balancers in namespace (string)
        /// </summary>
        public readonly string? ServicesLoadBalancers;
        /// <summary>
        /// Limit for services node ports in namespace (string)
        /// </summary>
        public readonly string? ServicesNodePorts;

        [OutputConstructor]
        private NamespaceResourceQuotaLimit(
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
    }
}

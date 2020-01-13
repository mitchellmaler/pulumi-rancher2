// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a Rancher v2 Cluster resource. This can be used to create Clusters for Rancher v2 environments and retrieve their information.
 * 
 * ## Example Usage
 * 
 * Creating Rancher v2 imported cluster
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rancher2 from "@pulumi/rancher2";
 * 
 * // Create a new rancher2 imported Cluster
 * const fooImported = new rancher2.Cluster("foo-imported", {
 *     description: "Foo rancher2 imported cluster",
 * });
 * ```
 * 
 * Creating Rancher v2 RKE cluster
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rancher2 from "@pulumi/rancher2";
 * 
 * // Create a new rancher2 RKE Cluster
 * const fooCustom = new rancher2.Cluster("foo-custom", {
 *     description: "Foo rancher2 custom cluster",
 *     rkeConfig: {
 *         network: {
 *             plugin: "canal",
 *         },
 *     },
 * });
 * ```
 * 
 * Creating Rancher v2 RKE cluster enabling and customizing monitoring
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rancher2 from "@pulumi/rancher2";
 * 
 * // Create a new rancher2 RKE Cluster
 * const fooCustom = new rancher2.Cluster("foo-custom", {
 *     clusterMonitoringInput: {
 *         answers: {
 *             "exporter-kubelets.https": true,
 *             "exporter-node.enabled": true,
 *             "exporter-node.ports.metrics.port": 9796,
 *             "exporter-node.resources.limits.cpu": "200m",
 *             "exporter-node.resources.limits.memory": "200Mi",
 *             "grafana.persistence.enabled": false,
 *             "grafana.persistence.size": "10Gi",
 *             "grafana.persistence.storageClass": "default",
 *             "operator.resources.limits.memory": "500Mi",
 *             "prometheus.persistence.enabled": "false",
 *             "prometheus.persistence.size": "50Gi",
 *             "prometheus.persistence.storageClass": "default",
 *             "prometheus.persistent.useReleaseName": "true",
 *             "prometheus.resources.core.limits.cpu": "1000m",
 *             "prometheus.resources.core.limits.memory": "1500Mi",
 *             "prometheus.resources.core.requests.cpu": "750m",
 *             "prometheus.resources.core.requests.memory": "750Mi",
 *             "prometheus.retention": "12h",
 *         },
 *     },
 *     description: "Foo rancher2 custom cluster",
 *     enableClusterMonitoring: true,
 *     rkeConfig: {
 *         network: {
 *             plugin: "canal",
 *         },
 *     },
 * });
 * ```
 * 
 * 
 * Creating Rancher v2 RKE cluster assigning a node pool (overlapped planes)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rancher2 from "@pulumi/rancher2";
 * 
 * // Create a new rancher2 RKE Cluster
 * const fooCustom = new rancher2.Cluster("foo-custom", {
 *     description: "Foo rancher2 custom cluster",
 *     rkeConfig: {
 *         network: {
 *             plugin: "canal",
 *         },
 *     },
 * });
 * // Create a new rancher2 Node Template
 * const fooNodeTemplate = new rancher2.NodeTemplate("foo", {
 *     amazonec2Config: {
 *         accessKey: "AWS_ACCESS_KEY",
 *         ami: "<AMI_ID>",
 *         region: "<REGION>",
 *         secretKey: "<AWS_SECRET_KEY>",
 *         securityGroups: ["<AWS_SECURITY_GROUP>"],
 *         subnetId: "<SUBNET_ID>",
 *         vpcId: "<VPC_ID>",
 *         zone: "<ZONE>",
 *     },
 *     description: "foo test",
 * });
 * // Create a new rancher2 Node Pool
 * const fooNodePool = new rancher2.NodePool("foo", {
 *     clusterId: foo_custom.id,
 *     controlPlane: true,
 *     etcd: true,
 *     hostnamePrefix: "foo-cluster-0",
 *     nodeTemplateId: fooNodeTemplate.id,
 *     quantity: 3,
 *     worker: true,
 * });
 * ```
 * 
 * Creating Rancher v2 RKE cluster from template. For Rancher v2.3.x or above.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rancher2 from "@pulumi/rancher2";
 * 
 * // Create a new rancher2 cluster template
 * const fooClusterTemplate = new rancher2.ClusterTemplate("foo", {
 *     description: "Test cluster template v2",
 *     members: [{
 *         accessType: "owner",
 *         userPrincipalId: "local://user-XXXXX",
 *     }],
 *     templateRevisions: [{
 *         clusterConfig: {
 *             rkeConfig: {
 *                 network: {
 *                     plugin: "canal",
 *                 },
 *                 services: {
 *                     etcd: {
 *                         creation: "6h",
 *                         retention: "24h",
 *                     },
 *                 },
 *             },
 *         },
 *         default: true,
 *         name: "V1",
 *     }],
 * });
 * // Create a new rancher2 RKE Cluster from template
 * const fooCluster = new rancher2.Cluster("foo", {
 *     clusterTemplateId: fooClusterTemplate.id,
 *     clusterTemplateRevisionId: fooClusterTemplate.defaultRevisionId,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/cluster.html.markdown.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rancher2:index/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The Azure AKS configuration for `aks` Clusters. Conflicts with `eksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
     */
    public readonly aksConfig!: pulumi.Output<outputs.ClusterAksConfig | undefined>;
    /**
     * Annotations for cluster registration token object (map)
     */
    public readonly annotations!: pulumi.Output<{[key: string]: any}>;
    /**
     * Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)
     */
    public readonly clusterAuthEndpoint!: pulumi.Output<outputs.ClusterClusterAuthEndpoint>;
    /**
     * Cluster monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured  (list maxitems:1)
     */
    public readonly clusterMonitoringInput!: pulumi.Output<outputs.ClusterClusterMonitoringInput>;
    /**
     * (Computed) Cluster Registration Token generated for the cluster (list maxitems:1)
     */
    public /*out*/ readonly clusterRegistrationToken!: pulumi.Output<outputs.ClusterClusterRegistrationToken>;
    /**
     * Cluster template answers. Just for Rancher v2.3.x and above (list maxitems:1)
     */
    public readonly clusterTemplateAnswers!: pulumi.Output<outputs.ClusterClusterTemplateAnswers | undefined>;
    /**
     * Cluster template ID. Just for Rancher v2.3.x and above (string)
     */
    public readonly clusterTemplateId!: pulumi.Output<string | undefined>;
    /**
     * Cluster template questions. Just for Rancher v2.3.x and above (list)
     */
    public readonly clusterTemplateQuestions!: pulumi.Output<outputs.ClusterClusterTemplateQuestion[] | undefined>;
    /**
     * Cluster template revision ID. Just for Rancher v2.3.x and above (string)
     */
    public readonly clusterTemplateRevisionId!: pulumi.Output<string | undefined>;
    /**
     * [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support) (string)
     */
    public readonly defaultPodSecurityPolicyTemplateId!: pulumi.Output<string>;
    /**
     * (Computed) Default project ID for the cluster (string)
     */
    public /*out*/ readonly defaultProjectId!: pulumi.Output<string>;
    /**
     * An optional description of this cluster (string)
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Desired agent image. Just for Rancher v2.3.x and above (string)
     */
    public readonly desiredAgentImage!: pulumi.Output<string>;
    /**
     * Desired auth image. Just for Rancher v2.3.x and above (string)
     */
    public readonly desiredAuthImage!: pulumi.Output<string>;
    /**
     * Desired auth image. Just for Rancher v2.3.x and above (string)
     */
    public readonly dockerRootDir!: pulumi.Output<string>;
    /**
     * (Computed) The driver used for the Cluster. `imported`, `azurekubernetesservice`, `amazonelasticcontainerservice`, `googlekubernetesengine` and `rancherKubernetesEngine` are supported (string)
     */
    public readonly driver!: pulumi.Output<string>;
    /**
     * The Amazon EKS configuration for `eks` Clusters. Conflicts with `aksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
     */
    public readonly eksConfig!: pulumi.Output<outputs.ClusterEksConfig | undefined>;
    /**
     * Enable built-in cluster alerting. Default `false` (bool)
     */
    public readonly enableClusterAlerting!: pulumi.Output<boolean | undefined>;
    /**
     * Enable built-in cluster istio. Default `false`. Just for Rancher v2.3.x and above (bool)
     */
    public readonly enableClusterIstio!: pulumi.Output<boolean | undefined>;
    /**
     * Enable built-in cluster monitoring. Default `false` (bool)
     */
    public readonly enableClusterMonitoring!: pulumi.Output<boolean | undefined>;
    /**
     * Enable project network isolation. Default `false` (bool)
     */
    public readonly enableNetworkPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The Google GKE configuration for `gke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `rkeConfig` (list maxitems:1)
     */
    public readonly gkeConfig!: pulumi.Output<outputs.ClusterGkeConfig | undefined>;
    /**
     * (Computed) Kube Config generated for the cluster (string)
     */
    public /*out*/ readonly kubeConfig!: pulumi.Output<string>;
    /**
     * Labels for cluster registration token object (map)
     */
    public readonly labels!: pulumi.Output<{[key: string]: any}>;
    /**
     * Name of cluster registration token (string)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The RKE configuration for `rke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `gkeConfig` (list maxitems:1)
     */
    public readonly rkeConfig!: pulumi.Output<outputs.ClusterRkeConfig>;
    /**
     * (Computed) System project ID for the cluster (string)
     */
    public /*out*/ readonly systemProjectId!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["aksConfig"] = state ? state.aksConfig : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["clusterAuthEndpoint"] = state ? state.clusterAuthEndpoint : undefined;
            inputs["clusterMonitoringInput"] = state ? state.clusterMonitoringInput : undefined;
            inputs["clusterRegistrationToken"] = state ? state.clusterRegistrationToken : undefined;
            inputs["clusterTemplateAnswers"] = state ? state.clusterTemplateAnswers : undefined;
            inputs["clusterTemplateId"] = state ? state.clusterTemplateId : undefined;
            inputs["clusterTemplateQuestions"] = state ? state.clusterTemplateQuestions : undefined;
            inputs["clusterTemplateRevisionId"] = state ? state.clusterTemplateRevisionId : undefined;
            inputs["defaultPodSecurityPolicyTemplateId"] = state ? state.defaultPodSecurityPolicyTemplateId : undefined;
            inputs["defaultProjectId"] = state ? state.defaultProjectId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["desiredAgentImage"] = state ? state.desiredAgentImage : undefined;
            inputs["desiredAuthImage"] = state ? state.desiredAuthImage : undefined;
            inputs["dockerRootDir"] = state ? state.dockerRootDir : undefined;
            inputs["driver"] = state ? state.driver : undefined;
            inputs["eksConfig"] = state ? state.eksConfig : undefined;
            inputs["enableClusterAlerting"] = state ? state.enableClusterAlerting : undefined;
            inputs["enableClusterIstio"] = state ? state.enableClusterIstio : undefined;
            inputs["enableClusterMonitoring"] = state ? state.enableClusterMonitoring : undefined;
            inputs["enableNetworkPolicy"] = state ? state.enableNetworkPolicy : undefined;
            inputs["gkeConfig"] = state ? state.gkeConfig : undefined;
            inputs["kubeConfig"] = state ? state.kubeConfig : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rkeConfig"] = state ? state.rkeConfig : undefined;
            inputs["systemProjectId"] = state ? state.systemProjectId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            inputs["aksConfig"] = args ? args.aksConfig : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["clusterAuthEndpoint"] = args ? args.clusterAuthEndpoint : undefined;
            inputs["clusterMonitoringInput"] = args ? args.clusterMonitoringInput : undefined;
            inputs["clusterTemplateAnswers"] = args ? args.clusterTemplateAnswers : undefined;
            inputs["clusterTemplateId"] = args ? args.clusterTemplateId : undefined;
            inputs["clusterTemplateQuestions"] = args ? args.clusterTemplateQuestions : undefined;
            inputs["clusterTemplateRevisionId"] = args ? args.clusterTemplateRevisionId : undefined;
            inputs["defaultPodSecurityPolicyTemplateId"] = args ? args.defaultPodSecurityPolicyTemplateId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["desiredAgentImage"] = args ? args.desiredAgentImage : undefined;
            inputs["desiredAuthImage"] = args ? args.desiredAuthImage : undefined;
            inputs["dockerRootDir"] = args ? args.dockerRootDir : undefined;
            inputs["driver"] = args ? args.driver : undefined;
            inputs["eksConfig"] = args ? args.eksConfig : undefined;
            inputs["enableClusterAlerting"] = args ? args.enableClusterAlerting : undefined;
            inputs["enableClusterIstio"] = args ? args.enableClusterIstio : undefined;
            inputs["enableClusterMonitoring"] = args ? args.enableClusterMonitoring : undefined;
            inputs["enableNetworkPolicy"] = args ? args.enableNetworkPolicy : undefined;
            inputs["gkeConfig"] = args ? args.gkeConfig : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rkeConfig"] = args ? args.rkeConfig : undefined;
            inputs["clusterRegistrationToken"] = undefined /*out*/;
            inputs["defaultProjectId"] = undefined /*out*/;
            inputs["kubeConfig"] = undefined /*out*/;
            inputs["systemProjectId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The Azure AKS configuration for `aks` Clusters. Conflicts with `eksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
     */
    readonly aksConfig?: pulumi.Input<inputs.ClusterAksConfig>;
    /**
     * Annotations for cluster registration token object (map)
     */
    readonly annotations?: pulumi.Input<{[key: string]: any}>;
    /**
     * Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)
     */
    readonly clusterAuthEndpoint?: pulumi.Input<inputs.ClusterClusterAuthEndpoint>;
    /**
     * Cluster monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured  (list maxitems:1)
     */
    readonly clusterMonitoringInput?: pulumi.Input<inputs.ClusterClusterMonitoringInput>;
    /**
     * (Computed) Cluster Registration Token generated for the cluster (list maxitems:1)
     */
    readonly clusterRegistrationToken?: pulumi.Input<inputs.ClusterClusterRegistrationToken>;
    /**
     * Cluster template answers. Just for Rancher v2.3.x and above (list maxitems:1)
     */
    readonly clusterTemplateAnswers?: pulumi.Input<inputs.ClusterClusterTemplateAnswers>;
    /**
     * Cluster template ID. Just for Rancher v2.3.x and above (string)
     */
    readonly clusterTemplateId?: pulumi.Input<string>;
    /**
     * Cluster template questions. Just for Rancher v2.3.x and above (list)
     */
    readonly clusterTemplateQuestions?: pulumi.Input<pulumi.Input<inputs.ClusterClusterTemplateQuestion>[]>;
    /**
     * Cluster template revision ID. Just for Rancher v2.3.x and above (string)
     */
    readonly clusterTemplateRevisionId?: pulumi.Input<string>;
    /**
     * [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support) (string)
     */
    readonly defaultPodSecurityPolicyTemplateId?: pulumi.Input<string>;
    /**
     * (Computed) Default project ID for the cluster (string)
     */
    readonly defaultProjectId?: pulumi.Input<string>;
    /**
     * An optional description of this cluster (string)
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Desired agent image. Just for Rancher v2.3.x and above (string)
     */
    readonly desiredAgentImage?: pulumi.Input<string>;
    /**
     * Desired auth image. Just for Rancher v2.3.x and above (string)
     */
    readonly desiredAuthImage?: pulumi.Input<string>;
    /**
     * Desired auth image. Just for Rancher v2.3.x and above (string)
     */
    readonly dockerRootDir?: pulumi.Input<string>;
    /**
     * (Computed) The driver used for the Cluster. `imported`, `azurekubernetesservice`, `amazonelasticcontainerservice`, `googlekubernetesengine` and `rancherKubernetesEngine` are supported (string)
     */
    readonly driver?: pulumi.Input<string>;
    /**
     * The Amazon EKS configuration for `eks` Clusters. Conflicts with `aksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
     */
    readonly eksConfig?: pulumi.Input<inputs.ClusterEksConfig>;
    /**
     * Enable built-in cluster alerting. Default `false` (bool)
     */
    readonly enableClusterAlerting?: pulumi.Input<boolean>;
    /**
     * Enable built-in cluster istio. Default `false`. Just for Rancher v2.3.x and above (bool)
     */
    readonly enableClusterIstio?: pulumi.Input<boolean>;
    /**
     * Enable built-in cluster monitoring. Default `false` (bool)
     */
    readonly enableClusterMonitoring?: pulumi.Input<boolean>;
    /**
     * Enable project network isolation. Default `false` (bool)
     */
    readonly enableNetworkPolicy?: pulumi.Input<boolean>;
    /**
     * The Google GKE configuration for `gke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `rkeConfig` (list maxitems:1)
     */
    readonly gkeConfig?: pulumi.Input<inputs.ClusterGkeConfig>;
    /**
     * (Computed) Kube Config generated for the cluster (string)
     */
    readonly kubeConfig?: pulumi.Input<string>;
    /**
     * Labels for cluster registration token object (map)
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Name of cluster registration token (string)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The RKE configuration for `rke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `gkeConfig` (list maxitems:1)
     */
    readonly rkeConfig?: pulumi.Input<inputs.ClusterRkeConfig>;
    /**
     * (Computed) System project ID for the cluster (string)
     */
    readonly systemProjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The Azure AKS configuration for `aks` Clusters. Conflicts with `eksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
     */
    readonly aksConfig?: pulumi.Input<inputs.ClusterAksConfig>;
    /**
     * Annotations for cluster registration token object (map)
     */
    readonly annotations?: pulumi.Input<{[key: string]: any}>;
    /**
     * Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)
     */
    readonly clusterAuthEndpoint?: pulumi.Input<inputs.ClusterClusterAuthEndpoint>;
    /**
     * Cluster monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured  (list maxitems:1)
     */
    readonly clusterMonitoringInput?: pulumi.Input<inputs.ClusterClusterMonitoringInput>;
    /**
     * Cluster template answers. Just for Rancher v2.3.x and above (list maxitems:1)
     */
    readonly clusterTemplateAnswers?: pulumi.Input<inputs.ClusterClusterTemplateAnswers>;
    /**
     * Cluster template ID. Just for Rancher v2.3.x and above (string)
     */
    readonly clusterTemplateId?: pulumi.Input<string>;
    /**
     * Cluster template questions. Just for Rancher v2.3.x and above (list)
     */
    readonly clusterTemplateQuestions?: pulumi.Input<pulumi.Input<inputs.ClusterClusterTemplateQuestion>[]>;
    /**
     * Cluster template revision ID. Just for Rancher v2.3.x and above (string)
     */
    readonly clusterTemplateRevisionId?: pulumi.Input<string>;
    /**
     * [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support) (string)
     */
    readonly defaultPodSecurityPolicyTemplateId?: pulumi.Input<string>;
    /**
     * An optional description of this cluster (string)
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Desired agent image. Just for Rancher v2.3.x and above (string)
     */
    readonly desiredAgentImage?: pulumi.Input<string>;
    /**
     * Desired auth image. Just for Rancher v2.3.x and above (string)
     */
    readonly desiredAuthImage?: pulumi.Input<string>;
    /**
     * Desired auth image. Just for Rancher v2.3.x and above (string)
     */
    readonly dockerRootDir?: pulumi.Input<string>;
    /**
     * (Computed) The driver used for the Cluster. `imported`, `azurekubernetesservice`, `amazonelasticcontainerservice`, `googlekubernetesengine` and `rancherKubernetesEngine` are supported (string)
     */
    readonly driver?: pulumi.Input<string>;
    /**
     * The Amazon EKS configuration for `eks` Clusters. Conflicts with `aksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
     */
    readonly eksConfig?: pulumi.Input<inputs.ClusterEksConfig>;
    /**
     * Enable built-in cluster alerting. Default `false` (bool)
     */
    readonly enableClusterAlerting?: pulumi.Input<boolean>;
    /**
     * Enable built-in cluster istio. Default `false`. Just for Rancher v2.3.x and above (bool)
     */
    readonly enableClusterIstio?: pulumi.Input<boolean>;
    /**
     * Enable built-in cluster monitoring. Default `false` (bool)
     */
    readonly enableClusterMonitoring?: pulumi.Input<boolean>;
    /**
     * Enable project network isolation. Default `false` (bool)
     */
    readonly enableNetworkPolicy?: pulumi.Input<boolean>;
    /**
     * The Google GKE configuration for `gke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `rkeConfig` (list maxitems:1)
     */
    readonly gkeConfig?: pulumi.Input<inputs.ClusterGkeConfig>;
    /**
     * Labels for cluster registration token object (map)
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Name of cluster registration token (string)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The RKE configuration for `rke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `gkeConfig` (list maxitems:1)
     */
    readonly rkeConfig?: pulumi.Input<inputs.ClusterRkeConfig>;
}

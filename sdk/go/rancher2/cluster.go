// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Cluster resource. This can be used to create Clusters for Rancher v2 environments and retrieve their information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/cluster.html.markdown.
type Cluster struct {
	pulumi.CustomResourceState

	// The Azure AKS configuration for `aks` Clusters. Conflicts with `eksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	AksConfig ClusterAksConfigPtrOutput `pulumi:"aksConfig"`
	// Annotations for cluster registration token object (map)
	Annotations pulumi.MapOutput `pulumi:"annotations"`
	// Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)
	ClusterAuthEndpoint ClusterClusterAuthEndpointOutput `pulumi:"clusterAuthEndpoint"`
	// Cluster monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured  (list maxitems:1)
	ClusterMonitoringInput ClusterClusterMonitoringInputOutput `pulumi:"clusterMonitoringInput"`
	// (Computed) Cluster Registration Token generated for the cluster (list maxitems:1)
	ClusterRegistrationToken ClusterClusterRegistrationTokenOutput `pulumi:"clusterRegistrationToken"`
	// Cluster template answers. Just for Rancher v2.3.x and above (list maxitems:1)
	ClusterTemplateAnswers ClusterClusterTemplateAnswersOutput `pulumi:"clusterTemplateAnswers"`
	// Cluster template ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateId pulumi.StringPtrOutput `pulumi:"clusterTemplateId"`
	// Cluster template questions. Just for Rancher v2.3.x and above (list)
	ClusterTemplateQuestions ClusterClusterTemplateQuestionArrayOutput `pulumi:"clusterTemplateQuestions"`
	// Cluster template revision ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateRevisionId pulumi.StringPtrOutput `pulumi:"clusterTemplateRevisionId"`
	// [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support) (string)
	DefaultPodSecurityPolicyTemplateId pulumi.StringOutput `pulumi:"defaultPodSecurityPolicyTemplateId"`
	// (Computed) Default project ID for the cluster (string)
	DefaultProjectId pulumi.StringOutput `pulumi:"defaultProjectId"`
	// An optional description of this cluster (string)
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Desired agent image. Just for Rancher v2.3.x and above (string)
	DesiredAgentImage pulumi.StringOutput `pulumi:"desiredAgentImage"`
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DesiredAuthImage pulumi.StringOutput `pulumi:"desiredAuthImage"`
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DockerRootDir pulumi.StringOutput `pulumi:"dockerRootDir"`
	// (Computed) The driver used for the Cluster. `imported`, `azurekubernetesservice`, `amazonelasticcontainerservice`, `googlekubernetesengine` and `rancherKubernetesEngine` are supported (string)
	Driver pulumi.StringOutput `pulumi:"driver"`
	// The Amazon EKS configuration for `eks` Clusters. Conflicts with `aksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	EksConfig ClusterEksConfigPtrOutput `pulumi:"eksConfig"`
	// Enable built-in cluster alerting. Default `false` (bool)
	EnableClusterAlerting pulumi.BoolPtrOutput `pulumi:"enableClusterAlerting"`
	// Enable built-in cluster istio. Default `false`. Just for Rancher v2.3.x and above (bool)
	EnableClusterIstio pulumi.BoolPtrOutput `pulumi:"enableClusterIstio"`
	// Enable built-in cluster monitoring. Default `false` (bool)
	EnableClusterMonitoring pulumi.BoolPtrOutput `pulumi:"enableClusterMonitoring"`
	// Enable project network isolation. Default `false` (bool)
	EnableNetworkPolicy pulumi.BoolPtrOutput `pulumi:"enableNetworkPolicy"`
	// The Google GKE configuration for `gke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `rkeConfig` (list maxitems:1)
	GkeConfig ClusterGkeConfigPtrOutput `pulumi:"gkeConfig"`
	// (Computed) Kube Config generated for the cluster (string)
	KubeConfig pulumi.StringOutput `pulumi:"kubeConfig"`
	// Labels for cluster registration token object (map)
	Labels pulumi.MapOutput `pulumi:"labels"`
	// Name of cluster registration token (string)
	Name pulumi.StringOutput `pulumi:"name"`
	// The RKE configuration for `rke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `gkeConfig` (list maxitems:1)
	RkeConfig ClusterRkeConfigOutput `pulumi:"rkeConfig"`
	// (Computed) System project ID for the cluster (string)
	SystemProjectId pulumi.StringOutput `pulumi:"systemProjectId"`
	// Windows preferred cluster. Default: `false` (bool)
	WindowsPreferedCluster pulumi.BoolPtrOutput `pulumi:"windowsPreferedCluster"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("rancher2:index/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("rancher2:index/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The Azure AKS configuration for `aks` Clusters. Conflicts with `eksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	AksConfig *ClusterAksConfig `pulumi:"aksConfig"`
	// Annotations for cluster registration token object (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)
	ClusterAuthEndpoint *ClusterClusterAuthEndpoint `pulumi:"clusterAuthEndpoint"`
	// Cluster monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured  (list maxitems:1)
	ClusterMonitoringInput *ClusterClusterMonitoringInput `pulumi:"clusterMonitoringInput"`
	// (Computed) Cluster Registration Token generated for the cluster (list maxitems:1)
	ClusterRegistrationToken *ClusterClusterRegistrationToken `pulumi:"clusterRegistrationToken"`
	// Cluster template answers. Just for Rancher v2.3.x and above (list maxitems:1)
	ClusterTemplateAnswers *ClusterClusterTemplateAnswers `pulumi:"clusterTemplateAnswers"`
	// Cluster template ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateId *string `pulumi:"clusterTemplateId"`
	// Cluster template questions. Just for Rancher v2.3.x and above (list)
	ClusterTemplateQuestions []ClusterClusterTemplateQuestion `pulumi:"clusterTemplateQuestions"`
	// Cluster template revision ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateRevisionId *string `pulumi:"clusterTemplateRevisionId"`
	// [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support) (string)
	DefaultPodSecurityPolicyTemplateId *string `pulumi:"defaultPodSecurityPolicyTemplateId"`
	// (Computed) Default project ID for the cluster (string)
	DefaultProjectId *string `pulumi:"defaultProjectId"`
	// An optional description of this cluster (string)
	Description *string `pulumi:"description"`
	// Desired agent image. Just for Rancher v2.3.x and above (string)
	DesiredAgentImage *string `pulumi:"desiredAgentImage"`
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DesiredAuthImage *string `pulumi:"desiredAuthImage"`
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DockerRootDir *string `pulumi:"dockerRootDir"`
	// (Computed) The driver used for the Cluster. `imported`, `azurekubernetesservice`, `amazonelasticcontainerservice`, `googlekubernetesengine` and `rancherKubernetesEngine` are supported (string)
	Driver *string `pulumi:"driver"`
	// The Amazon EKS configuration for `eks` Clusters. Conflicts with `aksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	EksConfig *ClusterEksConfig `pulumi:"eksConfig"`
	// Enable built-in cluster alerting. Default `false` (bool)
	EnableClusterAlerting *bool `pulumi:"enableClusterAlerting"`
	// Enable built-in cluster istio. Default `false`. Just for Rancher v2.3.x and above (bool)
	EnableClusterIstio *bool `pulumi:"enableClusterIstio"`
	// Enable built-in cluster monitoring. Default `false` (bool)
	EnableClusterMonitoring *bool `pulumi:"enableClusterMonitoring"`
	// Enable project network isolation. Default `false` (bool)
	EnableNetworkPolicy *bool `pulumi:"enableNetworkPolicy"`
	// The Google GKE configuration for `gke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `rkeConfig` (list maxitems:1)
	GkeConfig *ClusterGkeConfig `pulumi:"gkeConfig"`
	// (Computed) Kube Config generated for the cluster (string)
	KubeConfig *string `pulumi:"kubeConfig"`
	// Labels for cluster registration token object (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// Name of cluster registration token (string)
	Name *string `pulumi:"name"`
	// The RKE configuration for `rke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `gkeConfig` (list maxitems:1)
	RkeConfig *ClusterRkeConfig `pulumi:"rkeConfig"`
	// (Computed) System project ID for the cluster (string)
	SystemProjectId *string `pulumi:"systemProjectId"`
	// Windows preferred cluster. Default: `false` (bool)
	WindowsPreferedCluster *bool `pulumi:"windowsPreferedCluster"`
}

type ClusterState struct {
	// The Azure AKS configuration for `aks` Clusters. Conflicts with `eksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	AksConfig ClusterAksConfigPtrInput
	// Annotations for cluster registration token object (map)
	Annotations pulumi.MapInput
	// Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)
	ClusterAuthEndpoint ClusterClusterAuthEndpointPtrInput
	// Cluster monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured  (list maxitems:1)
	ClusterMonitoringInput ClusterClusterMonitoringInputPtrInput
	// (Computed) Cluster Registration Token generated for the cluster (list maxitems:1)
	ClusterRegistrationToken ClusterClusterRegistrationTokenPtrInput
	// Cluster template answers. Just for Rancher v2.3.x and above (list maxitems:1)
	ClusterTemplateAnswers ClusterClusterTemplateAnswersPtrInput
	// Cluster template ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateId pulumi.StringPtrInput
	// Cluster template questions. Just for Rancher v2.3.x and above (list)
	ClusterTemplateQuestions ClusterClusterTemplateQuestionArrayInput
	// Cluster template revision ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateRevisionId pulumi.StringPtrInput
	// [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support) (string)
	DefaultPodSecurityPolicyTemplateId pulumi.StringPtrInput
	// (Computed) Default project ID for the cluster (string)
	DefaultProjectId pulumi.StringPtrInput
	// An optional description of this cluster (string)
	Description pulumi.StringPtrInput
	// Desired agent image. Just for Rancher v2.3.x and above (string)
	DesiredAgentImage pulumi.StringPtrInput
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DesiredAuthImage pulumi.StringPtrInput
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DockerRootDir pulumi.StringPtrInput
	// (Computed) The driver used for the Cluster. `imported`, `azurekubernetesservice`, `amazonelasticcontainerservice`, `googlekubernetesengine` and `rancherKubernetesEngine` are supported (string)
	Driver pulumi.StringPtrInput
	// The Amazon EKS configuration for `eks` Clusters. Conflicts with `aksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	EksConfig ClusterEksConfigPtrInput
	// Enable built-in cluster alerting. Default `false` (bool)
	EnableClusterAlerting pulumi.BoolPtrInput
	// Enable built-in cluster istio. Default `false`. Just for Rancher v2.3.x and above (bool)
	EnableClusterIstio pulumi.BoolPtrInput
	// Enable built-in cluster monitoring. Default `false` (bool)
	EnableClusterMonitoring pulumi.BoolPtrInput
	// Enable project network isolation. Default `false` (bool)
	EnableNetworkPolicy pulumi.BoolPtrInput
	// The Google GKE configuration for `gke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `rkeConfig` (list maxitems:1)
	GkeConfig ClusterGkeConfigPtrInput
	// (Computed) Kube Config generated for the cluster (string)
	KubeConfig pulumi.StringPtrInput
	// Labels for cluster registration token object (map)
	Labels pulumi.MapInput
	// Name of cluster registration token (string)
	Name pulumi.StringPtrInput
	// The RKE configuration for `rke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `gkeConfig` (list maxitems:1)
	RkeConfig ClusterRkeConfigPtrInput
	// (Computed) System project ID for the cluster (string)
	SystemProjectId pulumi.StringPtrInput
	// Windows preferred cluster. Default: `false` (bool)
	WindowsPreferedCluster pulumi.BoolPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The Azure AKS configuration for `aks` Clusters. Conflicts with `eksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	AksConfig *ClusterAksConfig `pulumi:"aksConfig"`
	// Annotations for cluster registration token object (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)
	ClusterAuthEndpoint *ClusterClusterAuthEndpoint `pulumi:"clusterAuthEndpoint"`
	// Cluster monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured  (list maxitems:1)
	ClusterMonitoringInput *ClusterClusterMonitoringInput `pulumi:"clusterMonitoringInput"`
	// Cluster template answers. Just for Rancher v2.3.x and above (list maxitems:1)
	ClusterTemplateAnswers *ClusterClusterTemplateAnswers `pulumi:"clusterTemplateAnswers"`
	// Cluster template ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateId *string `pulumi:"clusterTemplateId"`
	// Cluster template questions. Just for Rancher v2.3.x and above (list)
	ClusterTemplateQuestions []ClusterClusterTemplateQuestion `pulumi:"clusterTemplateQuestions"`
	// Cluster template revision ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateRevisionId *string `pulumi:"clusterTemplateRevisionId"`
	// [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support) (string)
	DefaultPodSecurityPolicyTemplateId *string `pulumi:"defaultPodSecurityPolicyTemplateId"`
	// An optional description of this cluster (string)
	Description *string `pulumi:"description"`
	// Desired agent image. Just for Rancher v2.3.x and above (string)
	DesiredAgentImage *string `pulumi:"desiredAgentImage"`
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DesiredAuthImage *string `pulumi:"desiredAuthImage"`
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DockerRootDir *string `pulumi:"dockerRootDir"`
	// (Computed) The driver used for the Cluster. `imported`, `azurekubernetesservice`, `amazonelasticcontainerservice`, `googlekubernetesengine` and `rancherKubernetesEngine` are supported (string)
	Driver *string `pulumi:"driver"`
	// The Amazon EKS configuration for `eks` Clusters. Conflicts with `aksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	EksConfig *ClusterEksConfig `pulumi:"eksConfig"`
	// Enable built-in cluster alerting. Default `false` (bool)
	EnableClusterAlerting *bool `pulumi:"enableClusterAlerting"`
	// Enable built-in cluster istio. Default `false`. Just for Rancher v2.3.x and above (bool)
	EnableClusterIstio *bool `pulumi:"enableClusterIstio"`
	// Enable built-in cluster monitoring. Default `false` (bool)
	EnableClusterMonitoring *bool `pulumi:"enableClusterMonitoring"`
	// Enable project network isolation. Default `false` (bool)
	EnableNetworkPolicy *bool `pulumi:"enableNetworkPolicy"`
	// The Google GKE configuration for `gke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `rkeConfig` (list maxitems:1)
	GkeConfig *ClusterGkeConfig `pulumi:"gkeConfig"`
	// Labels for cluster registration token object (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// Name of cluster registration token (string)
	Name *string `pulumi:"name"`
	// The RKE configuration for `rke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `gkeConfig` (list maxitems:1)
	RkeConfig *ClusterRkeConfig `pulumi:"rkeConfig"`
	// Windows preferred cluster. Default: `false` (bool)
	WindowsPreferedCluster *bool `pulumi:"windowsPreferedCluster"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The Azure AKS configuration for `aks` Clusters. Conflicts with `eksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	AksConfig ClusterAksConfigPtrInput
	// Annotations for cluster registration token object (map)
	Annotations pulumi.MapInput
	// Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)
	ClusterAuthEndpoint ClusterClusterAuthEndpointPtrInput
	// Cluster monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured  (list maxitems:1)
	ClusterMonitoringInput ClusterClusterMonitoringInputPtrInput
	// Cluster template answers. Just for Rancher v2.3.x and above (list maxitems:1)
	ClusterTemplateAnswers ClusterClusterTemplateAnswersPtrInput
	// Cluster template ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateId pulumi.StringPtrInput
	// Cluster template questions. Just for Rancher v2.3.x and above (list)
	ClusterTemplateQuestions ClusterClusterTemplateQuestionArrayInput
	// Cluster template revision ID. Just for Rancher v2.3.x and above (string)
	ClusterTemplateRevisionId pulumi.StringPtrInput
	// [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support) (string)
	DefaultPodSecurityPolicyTemplateId pulumi.StringPtrInput
	// An optional description of this cluster (string)
	Description pulumi.StringPtrInput
	// Desired agent image. Just for Rancher v2.3.x and above (string)
	DesiredAgentImage pulumi.StringPtrInput
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DesiredAuthImage pulumi.StringPtrInput
	// Desired auth image. Just for Rancher v2.3.x and above (string)
	DockerRootDir pulumi.StringPtrInput
	// (Computed) The driver used for the Cluster. `imported`, `azurekubernetesservice`, `amazonelasticcontainerservice`, `googlekubernetesengine` and `rancherKubernetesEngine` are supported (string)
	Driver pulumi.StringPtrInput
	// The Amazon EKS configuration for `eks` Clusters. Conflicts with `aksConfig`, `gkeConfig` and `rkeConfig` (list maxitems:1)
	EksConfig ClusterEksConfigPtrInput
	// Enable built-in cluster alerting. Default `false` (bool)
	EnableClusterAlerting pulumi.BoolPtrInput
	// Enable built-in cluster istio. Default `false`. Just for Rancher v2.3.x and above (bool)
	EnableClusterIstio pulumi.BoolPtrInput
	// Enable built-in cluster monitoring. Default `false` (bool)
	EnableClusterMonitoring pulumi.BoolPtrInput
	// Enable project network isolation. Default `false` (bool)
	EnableNetworkPolicy pulumi.BoolPtrInput
	// The Google GKE configuration for `gke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `rkeConfig` (list maxitems:1)
	GkeConfig ClusterGkeConfigPtrInput
	// Labels for cluster registration token object (map)
	Labels pulumi.MapInput
	// Name of cluster registration token (string)
	Name pulumi.StringPtrInput
	// The RKE configuration for `rke` Clusters. Conflicts with `aksConfig`, `eksConfig` and `gkeConfig` (list maxitems:1)
	RkeConfig ClusterRkeConfigPtrInput
	// Windows preferred cluster. Default: `false` (bool)
	WindowsPreferedCluster pulumi.BoolPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}


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
        /// Use this data source to retrieve information about a Rancher v2 cluster role template binding.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/clusterRole.html.markdown.
        /// </summary>
        public static Task<GetClusterRoleTemplateBindingResult> GetClusterRoleTemplateBinding(GetClusterRoleTemplateBindingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterRoleTemplateBindingResult>("rancher2:index/getClusterRoleTemplateBinding:getClusterRoleTemplateBinding", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetClusterRoleTemplateBindingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster id where bind cluster role template (string)
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the cluster role template binding (string)
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The role template id from create cluster role template binding (string)
        /// </summary>
        [Input("roleTemplateId")]
        public string? RoleTemplateId { get; set; }

        public GetClusterRoleTemplateBindingArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClusterRoleTemplateBindingResult
    {
        /// <summary>
        /// (Computed) Annotations of the resource (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Annotations;
        public readonly string ClusterId;
        /// <summary>
        /// (Computed) The group ID to assign cluster role template binding (string)
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// (Computed) The group_principal ID to assign cluster role template binding (string)
        /// </summary>
        public readonly string GroupPrincipalId;
        /// <summary>
        /// (Computed) Labels of the resource (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        public readonly string Name;
        public readonly string RoleTemplateId;
        /// <summary>
        /// (Computed) The user ID to assign cluster role template binding (string)
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// (Computed) The user_principal ID to assign cluster role template binding (string)
        /// </summary>
        public readonly string UserPrincipalId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClusterRoleTemplateBindingResult(
            ImmutableDictionary<string, object> annotations,
            string clusterId,
            string groupId,
            string groupPrincipalId,
            ImmutableDictionary<string, object> labels,
            string name,
            string roleTemplateId,
            string userId,
            string userPrincipalId,
            string id)
        {
            Annotations = annotations;
            ClusterId = clusterId;
            GroupId = groupId;
            GroupPrincipalId = groupPrincipalId;
            Labels = labels;
            Name = name;
            RoleTemplateId = roleTemplateId;
            UserId = userId;
            UserPrincipalId = userPrincipalId;
            Id = id;
        }
    }
}

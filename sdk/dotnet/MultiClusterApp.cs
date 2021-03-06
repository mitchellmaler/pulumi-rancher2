// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rancher2
{
    public partial class MultiClusterApp : Pulumi.CustomResource
    {
        /// <summary>
        /// Annotations for multi cluster app object (map)
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, object>> Annotations { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app answers (list)
        /// </summary>
        [Output("answers")]
        public Output<ImmutableArray<Outputs.MultiClusterAppAnswers>> Answers { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app catalog name (string)
        /// </summary>
        [Output("catalogName")]
        public Output<string> CatalogName { get; private set; } = null!;

        /// <summary>
        /// Labels for multi cluster app object (map)
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>> Labels { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app answers (list)
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.MultiClusterAppMembers>> Members { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app name (string)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app revision history limit. Default `10` (int)
        /// </summary>
        [Output("revisionHistoryLimit")]
        public Output<int?> RevisionHistoryLimit { get; private set; } = null!;

        /// <summary>
        /// Current revision id for the multi cluster app (string)
        /// </summary>
        [Output("revisionId")]
        public Output<string> RevisionId { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app roles (list)
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app target projects (list)
        /// </summary>
        [Output("targets")]
        public Output<ImmutableArray<Outputs.MultiClusterAppTargets>> Targets { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app template name (string)
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app template version. Default: `latest` (string)
        /// </summary>
        [Output("templateVersion")]
        public Output<string> TemplateVersion { get; private set; } = null!;

        /// <summary>
        /// (Computed) The multi cluster app template version ID (string)
        /// </summary>
        [Output("templateVersionId")]
        public Output<string> TemplateVersionId { get; private set; } = null!;

        /// <summary>
        /// The multi cluster app upgrade strategy (list MaxItems:1)
        /// </summary>
        [Output("upgradeStrategy")]
        public Output<Outputs.MultiClusterAppUpgradeStrategy> UpgradeStrategy { get; private set; } = null!;

        /// <summary>
        /// Wait until the multi cluster app is active. Default `true` (bool)
        /// </summary>
        [Output("wait")]
        public Output<bool?> Wait { get; private set; } = null!;


        /// <summary>
        /// Create a MultiClusterApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MultiClusterApp(string name, MultiClusterAppArgs args, CustomResourceOptions? options = null)
            : base("rancher2:index/multiClusterApp:MultiClusterApp", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private MultiClusterApp(string name, Input<string> id, MultiClusterAppState? state = null, CustomResourceOptions? options = null)
            : base("rancher2:index/multiClusterApp:MultiClusterApp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MultiClusterApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MultiClusterApp Get(string name, Input<string> id, MultiClusterAppState? state = null, CustomResourceOptions? options = null)
        {
            return new MultiClusterApp(name, id, state, options);
        }
    }

    public sealed class MultiClusterAppArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<object>? _annotations;

        /// <summary>
        /// Annotations for multi cluster app object (map)
        /// </summary>
        public InputMap<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<object>());
            set => _annotations = value;
        }

        [Input("answers")]
        private InputList<Inputs.MultiClusterAppAnswersArgs>? _answers;

        /// <summary>
        /// The multi cluster app answers (list)
        /// </summary>
        public InputList<Inputs.MultiClusterAppAnswersArgs> Answers
        {
            get => _answers ?? (_answers = new InputList<Inputs.MultiClusterAppAnswersArgs>());
            set => _answers = value;
        }

        /// <summary>
        /// The multi cluster app catalog name (string)
        /// </summary>
        [Input("catalogName", required: true)]
        public Input<string> CatalogName { get; set; } = null!;

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels for multi cluster app object (map)
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        [Input("members")]
        private InputList<Inputs.MultiClusterAppMembersArgs>? _members;

        /// <summary>
        /// The multi cluster app answers (list)
        /// </summary>
        public InputList<Inputs.MultiClusterAppMembersArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.MultiClusterAppMembersArgs>());
            set => _members = value;
        }

        /// <summary>
        /// The multi cluster app name (string)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The multi cluster app revision history limit. Default `10` (int)
        /// </summary>
        [Input("revisionHistoryLimit")]
        public Input<int>? RevisionHistoryLimit { get; set; }

        /// <summary>
        /// Current revision id for the multi cluster app (string)
        /// </summary>
        [Input("revisionId")]
        public Input<string>? RevisionId { get; set; }

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// The multi cluster app roles (list)
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("targets", required: true)]
        private InputList<Inputs.MultiClusterAppTargetsArgs>? _targets;

        /// <summary>
        /// The multi cluster app target projects (list)
        /// </summary>
        public InputList<Inputs.MultiClusterAppTargetsArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.MultiClusterAppTargetsArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// The multi cluster app template name (string)
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        /// <summary>
        /// The multi cluster app template version. Default: `latest` (string)
        /// </summary>
        [Input("templateVersion")]
        public Input<string>? TemplateVersion { get; set; }

        /// <summary>
        /// The multi cluster app upgrade strategy (list MaxItems:1)
        /// </summary>
        [Input("upgradeStrategy")]
        public Input<Inputs.MultiClusterAppUpgradeStrategyArgs>? UpgradeStrategy { get; set; }

        /// <summary>
        /// Wait until the multi cluster app is active. Default `true` (bool)
        /// </summary>
        [Input("wait")]
        public Input<bool>? Wait { get; set; }

        public MultiClusterAppArgs()
        {
        }
    }

    public sealed class MultiClusterAppState : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<object>? _annotations;

        /// <summary>
        /// Annotations for multi cluster app object (map)
        /// </summary>
        public InputMap<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<object>());
            set => _annotations = value;
        }

        [Input("answers")]
        private InputList<Inputs.MultiClusterAppAnswersGetArgs>? _answers;

        /// <summary>
        /// The multi cluster app answers (list)
        /// </summary>
        public InputList<Inputs.MultiClusterAppAnswersGetArgs> Answers
        {
            get => _answers ?? (_answers = new InputList<Inputs.MultiClusterAppAnswersGetArgs>());
            set => _answers = value;
        }

        /// <summary>
        /// The multi cluster app catalog name (string)
        /// </summary>
        [Input("catalogName")]
        public Input<string>? CatalogName { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels for multi cluster app object (map)
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        [Input("members")]
        private InputList<Inputs.MultiClusterAppMembersGetArgs>? _members;

        /// <summary>
        /// The multi cluster app answers (list)
        /// </summary>
        public InputList<Inputs.MultiClusterAppMembersGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.MultiClusterAppMembersGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// The multi cluster app name (string)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The multi cluster app revision history limit. Default `10` (int)
        /// </summary>
        [Input("revisionHistoryLimit")]
        public Input<int>? RevisionHistoryLimit { get; set; }

        /// <summary>
        /// Current revision id for the multi cluster app (string)
        /// </summary>
        [Input("revisionId")]
        public Input<string>? RevisionId { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// The multi cluster app roles (list)
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("targets")]
        private InputList<Inputs.MultiClusterAppTargetsGetArgs>? _targets;

        /// <summary>
        /// The multi cluster app target projects (list)
        /// </summary>
        public InputList<Inputs.MultiClusterAppTargetsGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.MultiClusterAppTargetsGetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// The multi cluster app template name (string)
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// The multi cluster app template version. Default: `latest` (string)
        /// </summary>
        [Input("templateVersion")]
        public Input<string>? TemplateVersion { get; set; }

        /// <summary>
        /// (Computed) The multi cluster app template version ID (string)
        /// </summary>
        [Input("templateVersionId")]
        public Input<string>? TemplateVersionId { get; set; }

        /// <summary>
        /// The multi cluster app upgrade strategy (list MaxItems:1)
        /// </summary>
        [Input("upgradeStrategy")]
        public Input<Inputs.MultiClusterAppUpgradeStrategyGetArgs>? UpgradeStrategy { get; set; }

        /// <summary>
        /// Wait until the multi cluster app is active. Default `true` (bool)
        /// </summary>
        [Input("wait")]
        public Input<bool>? Wait { get; set; }

        public MultiClusterAppState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class MultiClusterAppAnswersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID for answer (string)
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Project ID for target (string)
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("values")]
        private InputMap<object>? _values;

        /// <summary>
        /// Key/values for answer (map)
        /// </summary>
        public InputMap<object> Values
        {
            get => _values ?? (_values = new InputMap<object>());
            set => _values = value;
        }

        public MultiClusterAppAnswersArgs()
        {
        }
    }

    public sealed class MultiClusterAppAnswersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID for answer (string)
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Project ID for target (string)
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("values")]
        private InputMap<object>? _values;

        /// <summary>
        /// Key/values for answer (map)
        /// </summary>
        public InputMap<object> Values
        {
            get => _values ?? (_values = new InputMap<object>());
            set => _values = value;
        }

        public MultiClusterAppAnswersGetArgs()
        {
        }
    }

    public sealed class MultiClusterAppMembersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Member access type. Valid values: `["member" | "owner" | "read-only"]` (string)
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// Member group principal id (string)
        /// </summary>
        [Input("groupPrincipalId")]
        public Input<string>? GroupPrincipalId { get; set; }

        /// <summary>
        /// Member user principal id (string)
        /// </summary>
        [Input("userPrincipalId")]
        public Input<string>? UserPrincipalId { get; set; }

        public MultiClusterAppMembersArgs()
        {
        }
    }

    public sealed class MultiClusterAppMembersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Member access type. Valid values: `["member" | "owner" | "read-only"]` (string)
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// Member group principal id (string)
        /// </summary>
        [Input("groupPrincipalId")]
        public Input<string>? GroupPrincipalId { get; set; }

        /// <summary>
        /// Member user principal id (string)
        /// </summary>
        [Input("userPrincipalId")]
        public Input<string>? UserPrincipalId { get; set; }

        public MultiClusterAppMembersGetArgs()
        {
        }
    }

    public sealed class MultiClusterAppTargetsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// App ID for target (string)
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// App health state for target (string)
        /// </summary>
        [Input("healthState")]
        public Input<string>? HealthState { get; set; }

        /// <summary>
        /// Project ID for target (string)
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// App state for target (string)
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public MultiClusterAppTargetsArgs()
        {
        }
    }

    public sealed class MultiClusterAppTargetsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// App ID for target (string)
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// App health state for target (string)
        /// </summary>
        [Input("healthState")]
        public Input<string>? HealthState { get; set; }

        /// <summary>
        /// Project ID for target (string)
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// App state for target (string)
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public MultiClusterAppTargetsGetArgs()
        {
        }
    }

    public sealed class MultiClusterAppUpgradeStrategyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Upgrade strategy rolling update (list MaxItems:1)
        /// </summary>
        [Input("rollingUpdate")]
        public Input<MultiClusterAppUpgradeStrategyRollingUpdateArgs>? RollingUpdate { get; set; }

        public MultiClusterAppUpgradeStrategyArgs()
        {
        }
    }

    public sealed class MultiClusterAppUpgradeStrategyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Upgrade strategy rolling update (list MaxItems:1)
        /// </summary>
        [Input("rollingUpdate")]
        public Input<MultiClusterAppUpgradeStrategyRollingUpdateGetArgs>? RollingUpdate { get; set; }

        public MultiClusterAppUpgradeStrategyGetArgs()
        {
        }
    }

    public sealed class MultiClusterAppUpgradeStrategyRollingUpdateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rolling update batch size. Default `1` (int)
        /// </summary>
        [Input("batchSize")]
        public Input<int>? BatchSize { get; set; }

        /// <summary>
        /// Rolling update interval. Default `1` (int)
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        public MultiClusterAppUpgradeStrategyRollingUpdateArgs()
        {
        }
    }

    public sealed class MultiClusterAppUpgradeStrategyRollingUpdateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rolling update batch size. Default `1` (int)
        /// </summary>
        [Input("batchSize")]
        public Input<int>? BatchSize { get; set; }

        /// <summary>
        /// Rolling update interval. Default `1` (int)
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        public MultiClusterAppUpgradeStrategyRollingUpdateGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class MultiClusterAppAnswers
    {
        /// <summary>
        /// Cluster ID for answer (string)
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Project ID for target (string)
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Key/values for answer (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Values;

        [OutputConstructor]
        private MultiClusterAppAnswers(
            string clusterId,
            string projectId,
            ImmutableDictionary<string, object> values)
        {
            ClusterId = clusterId;
            ProjectId = projectId;
            Values = values;
        }
    }

    [OutputType]
    public sealed class MultiClusterAppMembers
    {
        /// <summary>
        /// Member access type. Valid values: `["member" | "owner" | "read-only"]` (string)
        /// </summary>
        public readonly string? AccessType;
        /// <summary>
        /// Member group principal id (string)
        /// </summary>
        public readonly string? GroupPrincipalId;
        /// <summary>
        /// Member user principal id (string)
        /// </summary>
        public readonly string? UserPrincipalId;

        [OutputConstructor]
        private MultiClusterAppMembers(
            string? accessType,
            string? groupPrincipalId,
            string? userPrincipalId)
        {
            AccessType = accessType;
            GroupPrincipalId = groupPrincipalId;
            UserPrincipalId = userPrincipalId;
        }
    }

    [OutputType]
    public sealed class MultiClusterAppTargets
    {
        /// <summary>
        /// App ID for target (string)
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// App health state for target (string)
        /// </summary>
        public readonly string HealthState;
        /// <summary>
        /// Project ID for target (string)
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// App state for target (string)
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private MultiClusterAppTargets(
            string appId,
            string healthState,
            string projectId,
            string state)
        {
            AppId = appId;
            HealthState = healthState;
            ProjectId = projectId;
            State = state;
        }
    }

    [OutputType]
    public sealed class MultiClusterAppUpgradeStrategy
    {
        /// <summary>
        /// Upgrade strategy rolling update (list MaxItems:1)
        /// </summary>
        public readonly MultiClusterAppUpgradeStrategyRollingUpdate? RollingUpdate;

        [OutputConstructor]
        private MultiClusterAppUpgradeStrategy(MultiClusterAppUpgradeStrategyRollingUpdate? rollingUpdate)
        {
            RollingUpdate = rollingUpdate;
        }
    }

    [OutputType]
    public sealed class MultiClusterAppUpgradeStrategyRollingUpdate
    {
        /// <summary>
        /// Rolling update batch size. Default `1` (int)
        /// </summary>
        public readonly int? BatchSize;
        /// <summary>
        /// Rolling update interval. Default `1` (int)
        /// </summary>
        public readonly int? Interval;

        [OutputConstructor]
        private MultiClusterAppUpgradeStrategyRollingUpdate(
            int? batchSize,
            int? interval)
        {
            BatchSize = batchSize;
            Interval = interval;
        }
    }
    }
}

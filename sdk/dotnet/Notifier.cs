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
    /// Provides a Rancher v2 Notifier resource. This can be used to create notifiers for Rancher v2 environments and retrieve their information.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/notifier.html.markdown.
    /// </summary>
    public partial class Notifier : Pulumi.CustomResource
    {
        /// <summary>
        /// Annotations for notifier object (map)
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, object>> Annotations { get; private set; } = null!;

        /// <summary>
        /// The cluster id where create notifier (string)
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The notifier description (string)
        /// * `send_resolved` = (Optional) Enable the notifier to send resolved notifications. Default `false` (bool)
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Labels for notifier object (map)
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>> Labels { get; private set; } = null!;

        /// <summary>
        /// The name of the notifier (string)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Pagerduty config for notifier (list maxitems:1)
        /// </summary>
        [Output("pagerdutyConfig")]
        public Output<Outputs.NotifierPagerdutyConfig?> PagerdutyConfig { get; private set; } = null!;

        /// <summary>
        /// Notifier send resolved
        /// </summary>
        [Output("sendResolved")]
        public Output<bool?> SendResolved { get; private set; } = null!;

        /// <summary>
        /// Slack config for notifier (list maxitems:1)
        /// </summary>
        [Output("slackConfig")]
        public Output<Outputs.NotifierSlackConfig?> SlackConfig { get; private set; } = null!;

        /// <summary>
        /// SMTP config for notifier (list maxitems:1)
        /// </summary>
        [Output("smtpConfig")]
        public Output<Outputs.NotifierSmtpConfig?> SmtpConfig { get; private set; } = null!;

        /// <summary>
        /// Webhook config for notifier (list maxitems:1)
        /// </summary>
        [Output("webhookConfig")]
        public Output<Outputs.NotifierWebhookConfig?> WebhookConfig { get; private set; } = null!;

        /// <summary>
        /// Wechat config for notifier (list maxitems:1)
        /// </summary>
        [Output("wechatConfig")]
        public Output<Outputs.NotifierWechatConfig?> WechatConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Notifier resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Notifier(string name, NotifierArgs args, CustomResourceOptions? options = null)
            : base("rancher2:index/notifier:Notifier", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Notifier(string name, Input<string> id, NotifierState? state = null, CustomResourceOptions? options = null)
            : base("rancher2:index/notifier:Notifier", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Notifier resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Notifier Get(string name, Input<string> id, NotifierState? state = null, CustomResourceOptions? options = null)
        {
            return new Notifier(name, id, state, options);
        }
    }

    public sealed class NotifierArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<object>? _annotations;

        /// <summary>
        /// Annotations for notifier object (map)
        /// </summary>
        public InputMap<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<object>());
            set => _annotations = value;
        }

        /// <summary>
        /// The cluster id where create notifier (string)
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The notifier description (string)
        /// * `send_resolved` = (Optional) Enable the notifier to send resolved notifications. Default `false` (bool)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels for notifier object (map)
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the notifier (string)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Pagerduty config for notifier (list maxitems:1)
        /// </summary>
        [Input("pagerdutyConfig")]
        public Input<Inputs.NotifierPagerdutyConfigArgs>? PagerdutyConfig { get; set; }

        /// <summary>
        /// Notifier send resolved
        /// </summary>
        [Input("sendResolved")]
        public Input<bool>? SendResolved { get; set; }

        /// <summary>
        /// Slack config for notifier (list maxitems:1)
        /// </summary>
        [Input("slackConfig")]
        public Input<Inputs.NotifierSlackConfigArgs>? SlackConfig { get; set; }

        /// <summary>
        /// SMTP config for notifier (list maxitems:1)
        /// </summary>
        [Input("smtpConfig")]
        public Input<Inputs.NotifierSmtpConfigArgs>? SmtpConfig { get; set; }

        /// <summary>
        /// Webhook config for notifier (list maxitems:1)
        /// </summary>
        [Input("webhookConfig")]
        public Input<Inputs.NotifierWebhookConfigArgs>? WebhookConfig { get; set; }

        /// <summary>
        /// Wechat config for notifier (list maxitems:1)
        /// </summary>
        [Input("wechatConfig")]
        public Input<Inputs.NotifierWechatConfigArgs>? WechatConfig { get; set; }

        public NotifierArgs()
        {
        }
    }

    public sealed class NotifierState : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<object>? _annotations;

        /// <summary>
        /// Annotations for notifier object (map)
        /// </summary>
        public InputMap<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<object>());
            set => _annotations = value;
        }

        /// <summary>
        /// The cluster id where create notifier (string)
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The notifier description (string)
        /// * `send_resolved` = (Optional) Enable the notifier to send resolved notifications. Default `false` (bool)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels for notifier object (map)
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the notifier (string)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Pagerduty config for notifier (list maxitems:1)
        /// </summary>
        [Input("pagerdutyConfig")]
        public Input<Inputs.NotifierPagerdutyConfigGetArgs>? PagerdutyConfig { get; set; }

        /// <summary>
        /// Notifier send resolved
        /// </summary>
        [Input("sendResolved")]
        public Input<bool>? SendResolved { get; set; }

        /// <summary>
        /// Slack config for notifier (list maxitems:1)
        /// </summary>
        [Input("slackConfig")]
        public Input<Inputs.NotifierSlackConfigGetArgs>? SlackConfig { get; set; }

        /// <summary>
        /// SMTP config for notifier (list maxitems:1)
        /// </summary>
        [Input("smtpConfig")]
        public Input<Inputs.NotifierSmtpConfigGetArgs>? SmtpConfig { get; set; }

        /// <summary>
        /// Webhook config for notifier (list maxitems:1)
        /// </summary>
        [Input("webhookConfig")]
        public Input<Inputs.NotifierWebhookConfigGetArgs>? WebhookConfig { get; set; }

        /// <summary>
        /// Wechat config for notifier (list maxitems:1)
        /// </summary>
        [Input("wechatConfig")]
        public Input<Inputs.NotifierWechatConfigGetArgs>? WechatConfig { get; set; }

        public NotifierState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class NotifierPagerdutyConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// Pagerduty service key (string)
        /// </summary>
        [Input("serviceKey", required: true)]
        public Input<string> ServiceKey { get; set; } = null!;

        public NotifierPagerdutyConfigArgs()
        {
        }
    }

    public sealed class NotifierPagerdutyConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// Pagerduty service key (string)
        /// </summary>
        [Input("serviceKey", required: true)]
        public Input<string> ServiceKey { get; set; } = null!;

        public NotifierPagerdutyConfigGetArgs()
        {
        }
    }

    public sealed class NotifierSlackConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat default recipient (string)
        /// </summary>
        [Input("defaultRecipient", required: true)]
        public Input<string> DefaultRecipient { get; set; } = null!;

        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// Webhook url (string)
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public NotifierSlackConfigArgs()
        {
        }
    }

    public sealed class NotifierSlackConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat default recipient (string)
        /// </summary>
        [Input("defaultRecipient", required: true)]
        public Input<string> DefaultRecipient { get; set; } = null!;

        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// Webhook url (string)
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public NotifierSlackConfigGetArgs()
        {
        }
    }

    public sealed class NotifierSmtpConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat default recipient (string)
        /// </summary>
        [Input("defaultRecipient", required: true)]
        public Input<string> DefaultRecipient { get; set; } = null!;

        /// <summary>
        /// SMTP host (string)
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// SMTP password (string)
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// SMTP port (int)
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// SMTP sender (string)
        /// </summary>
        [Input("sender", required: true)]
        public Input<string> Sender { get; set; } = null!;

        /// <summary>
        /// SMTP tls. Default `true` (bool)
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// SMTP username (string)
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public NotifierSmtpConfigArgs()
        {
        }
    }

    public sealed class NotifierSmtpConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat default recipient (string)
        /// </summary>
        [Input("defaultRecipient", required: true)]
        public Input<string> DefaultRecipient { get; set; } = null!;

        /// <summary>
        /// SMTP host (string)
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// SMTP password (string)
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// SMTP port (int)
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// SMTP sender (string)
        /// </summary>
        [Input("sender", required: true)]
        public Input<string> Sender { get; set; } = null!;

        /// <summary>
        /// SMTP tls. Default `true` (bool)
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// SMTP username (string)
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public NotifierSmtpConfigGetArgs()
        {
        }
    }

    public sealed class NotifierWebhookConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// Webhook url (string)
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public NotifierWebhookConfigArgs()
        {
        }
    }

    public sealed class NotifierWebhookConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// Webhook url (string)
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public NotifierWebhookConfigGetArgs()
        {
        }
    }

    public sealed class NotifierWechatConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat agent ID (string)
        /// </summary>
        [Input("agent", required: true)]
        public Input<string> Agent { get; set; } = null!;

        /// <summary>
        /// Wechat corporation ID (string)
        /// </summary>
        [Input("corp", required: true)]
        public Input<string> Corp { get; set; } = null!;

        /// <summary>
        /// Wechat default recipient (string)
        /// </summary>
        [Input("defaultRecipient", required: true)]
        public Input<string> DefaultRecipient { get; set; } = null!;

        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// Wechat recipient type. Allowed values: `party` | `tag` | `user` (string)
        /// </summary>
        [Input("recipientType")]
        public Input<string>? RecipientType { get; set; }

        /// <summary>
        /// Wechat agent ID (string)
        /// </summary>
        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        public NotifierWechatConfigArgs()
        {
        }
    }

    public sealed class NotifierWechatConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wechat agent ID (string)
        /// </summary>
        [Input("agent", required: true)]
        public Input<string> Agent { get; set; } = null!;

        /// <summary>
        /// Wechat corporation ID (string)
        /// </summary>
        [Input("corp", required: true)]
        public Input<string> Corp { get; set; } = null!;

        /// <summary>
        /// Wechat default recipient (string)
        /// </summary>
        [Input("defaultRecipient", required: true)]
        public Input<string> DefaultRecipient { get; set; } = null!;

        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        /// <summary>
        /// Wechat recipient type. Allowed values: `party` | `tag` | `user` (string)
        /// </summary>
        [Input("recipientType")]
        public Input<string>? RecipientType { get; set; }

        /// <summary>
        /// Wechat agent ID (string)
        /// </summary>
        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        public NotifierWechatConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class NotifierPagerdutyConfig
    {
        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        public readonly string? ProxyUrl;
        /// <summary>
        /// Pagerduty service key (string)
        /// </summary>
        public readonly string ServiceKey;

        [OutputConstructor]
        private NotifierPagerdutyConfig(
            string? proxyUrl,
            string serviceKey)
        {
            ProxyUrl = proxyUrl;
            ServiceKey = serviceKey;
        }
    }

    [OutputType]
    public sealed class NotifierSlackConfig
    {
        /// <summary>
        /// Wechat default recipient (string)
        /// </summary>
        public readonly string DefaultRecipient;
        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        public readonly string? ProxyUrl;
        /// <summary>
        /// Webhook url (string)
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private NotifierSlackConfig(
            string defaultRecipient,
            string? proxyUrl,
            string url)
        {
            DefaultRecipient = defaultRecipient;
            ProxyUrl = proxyUrl;
            Url = url;
        }
    }

    [OutputType]
    public sealed class NotifierSmtpConfig
    {
        /// <summary>
        /// Wechat default recipient (string)
        /// </summary>
        public readonly string DefaultRecipient;
        /// <summary>
        /// SMTP host (string)
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// SMTP password (string)
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// SMTP port (int)
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// SMTP sender (string)
        /// </summary>
        public readonly string Sender;
        /// <summary>
        /// SMTP tls. Default `true` (bool)
        /// </summary>
        public readonly bool? Tls;
        /// <summary>
        /// SMTP username (string)
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private NotifierSmtpConfig(
            string defaultRecipient,
            string host,
            string? password,
            int port,
            string sender,
            bool? tls,
            string? username)
        {
            DefaultRecipient = defaultRecipient;
            Host = host;
            Password = password;
            Port = port;
            Sender = sender;
            Tls = tls;
            Username = username;
        }
    }

    [OutputType]
    public sealed class NotifierWebhookConfig
    {
        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        public readonly string? ProxyUrl;
        /// <summary>
        /// Webhook url (string)
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private NotifierWebhookConfig(
            string? proxyUrl,
            string url)
        {
            ProxyUrl = proxyUrl;
            Url = url;
        }
    }

    [OutputType]
    public sealed class NotifierWechatConfig
    {
        /// <summary>
        /// Wechat agent ID (string)
        /// </summary>
        public readonly string Agent;
        /// <summary>
        /// Wechat corporation ID (string)
        /// </summary>
        public readonly string Corp;
        /// <summary>
        /// Wechat default recipient (string)
        /// </summary>
        public readonly string DefaultRecipient;
        /// <summary>
        /// Wechat proxy url (string)
        /// </summary>
        public readonly string? ProxyUrl;
        /// <summary>
        /// Wechat recipient type. Allowed values: `party` | `tag` | `user` (string)
        /// </summary>
        public readonly string? RecipientType;
        /// <summary>
        /// Wechat agent ID (string)
        /// </summary>
        public readonly string Secret;

        [OutputConstructor]
        private NotifierWechatConfig(
            string agent,
            string corp,
            string defaultRecipient,
            string? proxyUrl,
            string? recipientType,
            string secret)
        {
            Agent = agent;
            Corp = corp;
            DefaultRecipient = defaultRecipient;
            ProxyUrl = proxyUrl;
            RecipientType = recipientType;
            Secret = secret;
        }
    }
    }
}

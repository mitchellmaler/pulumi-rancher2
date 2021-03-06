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
        /// Use this data source to retrieve information about a Rancher v2 Project Logging.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/projectLogging.html.markdown.
        /// </summary>
        public static Task<GetProjectLoggingResult> GetProjectLogging(GetProjectLoggingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectLoggingResult>("rancher2:index/getProjectLogging:getProjectLogging", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetProjectLoggingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project id to configure logging (string)
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetProjectLoggingArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetProjectLoggingResult
    {
        /// <summary>
        /// (Computed) Annotations for Cluster Logging object (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Annotations;
        /// <summary>
        /// (Computed) The elasticsearch config for Cluster Logging. For `kind = elasticsearch`  (list maxitems:1)
        /// </summary>
        public readonly Outputs.GetProjectLoggingElasticsearchConfigResult ElasticsearchConfig;
        /// <summary>
        /// (Computed) The fluentd config for Cluster Logging. For `kind = fluentd` (list maxitems:1)
        /// </summary>
        public readonly Outputs.GetProjectLoggingFluentdConfigResult FluentdConfig;
        /// <summary>
        /// (Computed) The kafka config for Cluster Logging. For `kind = kafka` (list maxitems:1)
        /// </summary>
        public readonly Outputs.GetProjectLoggingKafkaConfigResult KafkaConfig;
        /// <summary>
        /// (Computed) The kind of the Cluster Logging. `elasticsearch`, `fluentd`, `kafka`, `splunk` and `syslog` are supported (string)
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// (Computed) Labels for Cluster Logging object (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        /// <summary>
        /// (Computed) The name of the cluster logging config (string)
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (Computed) The namespace id from cluster logging (string)
        /// </summary>
        public readonly string NamespaceId;
        /// <summary>
        /// (Computed) How often buffered logs would be flushed. Default: `3` seconds (int)
        /// </summary>
        public readonly int OutputFlushInterval;
        /// <summary>
        /// (computed) The output tags for Cluster Logging (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> OutputTags;
        public readonly string ProjectId;
        /// <summary>
        /// (Computed) The splunk config for Cluster Logging. For `kind = splunk` (list maxitems:1)
        /// </summary>
        public readonly Outputs.GetProjectLoggingSplunkConfigResult SplunkConfig;
        /// <summary>
        /// (Computed) The syslog config for Cluster Logging. For `kind = syslog` (list maxitems:1)
        /// </summary>
        public readonly Outputs.GetProjectLoggingSyslogConfigResult SyslogConfig;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetProjectLoggingResult(
            ImmutableDictionary<string, object> annotations,
            Outputs.GetProjectLoggingElasticsearchConfigResult elasticsearchConfig,
            Outputs.GetProjectLoggingFluentdConfigResult fluentdConfig,
            Outputs.GetProjectLoggingKafkaConfigResult kafkaConfig,
            string kind,
            ImmutableDictionary<string, object> labels,
            string name,
            string namespaceId,
            int outputFlushInterval,
            ImmutableDictionary<string, object> outputTags,
            string projectId,
            Outputs.GetProjectLoggingSplunkConfigResult splunkConfig,
            Outputs.GetProjectLoggingSyslogConfigResult syslogConfig,
            string id)
        {
            Annotations = annotations;
            ElasticsearchConfig = elasticsearchConfig;
            FluentdConfig = fluentdConfig;
            KafkaConfig = kafkaConfig;
            Kind = kind;
            Labels = labels;
            Name = name;
            NamespaceId = namespaceId;
            OutputFlushInterval = outputFlushInterval;
            OutputTags = outputTags;
            ProjectId = projectId;
            SplunkConfig = splunkConfig;
            SyslogConfig = syslogConfig;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetProjectLoggingElasticsearchConfigResult
    {
        public readonly string? AuthPassword;
        public readonly string? AuthUsername;
        public readonly string? Certificate;
        public readonly string? ClientCert;
        public readonly string? ClientKey;
        public readonly string? ClientKeyPass;
        public readonly string? DateFormat;
        public readonly string Endpoint;
        public readonly string? IndexPrefix;
        public readonly bool SslVerify;
        public readonly string? SslVersion;

        [OutputConstructor]
        private GetProjectLoggingElasticsearchConfigResult(
            string? authPassword,
            string? authUsername,
            string? certificate,
            string? clientCert,
            string? clientKey,
            string? clientKeyPass,
            string? dateFormat,
            string endpoint,
            string? indexPrefix,
            bool sslVerify,
            string? sslVersion)
        {
            AuthPassword = authPassword;
            AuthUsername = authUsername;
            Certificate = certificate;
            ClientCert = clientCert;
            ClientKey = clientKey;
            ClientKeyPass = clientKeyPass;
            DateFormat = dateFormat;
            Endpoint = endpoint;
            IndexPrefix = indexPrefix;
            SslVerify = sslVerify;
            SslVersion = sslVersion;
        }
    }

    [OutputType]
    public sealed class GetProjectLoggingFluentdConfigFluentServersResult
    {
        public readonly string Endpoint;
        public readonly string? Hostname;
        public readonly string? Password;
        public readonly string? SharedKey;
        public readonly bool? Standby;
        public readonly string? Username;
        public readonly int? Weight;

        [OutputConstructor]
        private GetProjectLoggingFluentdConfigFluentServersResult(
            string endpoint,
            string? hostname,
            string? password,
            string? sharedKey,
            bool? standby,
            string? username,
            int? weight)
        {
            Endpoint = endpoint;
            Hostname = hostname;
            Password = password;
            SharedKey = sharedKey;
            Standby = standby;
            Username = username;
            Weight = weight;
        }
    }

    [OutputType]
    public sealed class GetProjectLoggingFluentdConfigResult
    {
        public readonly string? Certificate;
        public readonly bool? Compress;
        public readonly bool? EnableTls;
        public readonly ImmutableArray<GetProjectLoggingFluentdConfigFluentServersResult> FluentServers;

        [OutputConstructor]
        private GetProjectLoggingFluentdConfigResult(
            string? certificate,
            bool? compress,
            bool? enableTls,
            ImmutableArray<GetProjectLoggingFluentdConfigFluentServersResult> fluentServers)
        {
            Certificate = certificate;
            Compress = compress;
            EnableTls = enableTls;
            FluentServers = fluentServers;
        }
    }

    [OutputType]
    public sealed class GetProjectLoggingKafkaConfigResult
    {
        public readonly ImmutableArray<string> BrokerEndpoints;
        public readonly string? Certificate;
        public readonly string? ClientCert;
        public readonly string? ClientKey;
        public readonly string Topic;
        public readonly string? ZookeeperEndpoint;

        [OutputConstructor]
        private GetProjectLoggingKafkaConfigResult(
            ImmutableArray<string> brokerEndpoints,
            string? certificate,
            string? clientCert,
            string? clientKey,
            string topic,
            string? zookeeperEndpoint)
        {
            BrokerEndpoints = brokerEndpoints;
            Certificate = certificate;
            ClientCert = clientCert;
            ClientKey = clientKey;
            Topic = topic;
            ZookeeperEndpoint = zookeeperEndpoint;
        }
    }

    [OutputType]
    public sealed class GetProjectLoggingSplunkConfigResult
    {
        public readonly string? Certificate;
        public readonly string? ClientCert;
        public readonly string? ClientKey;
        public readonly string? ClientKeyPass;
        public readonly string Endpoint;
        public readonly string? Index;
        public readonly string? Source;
        public readonly bool SslVerify;
        public readonly string Token;

        [OutputConstructor]
        private GetProjectLoggingSplunkConfigResult(
            string? certificate,
            string? clientCert,
            string? clientKey,
            string? clientKeyPass,
            string endpoint,
            string? index,
            string? source,
            bool sslVerify,
            string token)
        {
            Certificate = certificate;
            ClientCert = clientCert;
            ClientKey = clientKey;
            ClientKeyPass = clientKeyPass;
            Endpoint = endpoint;
            Index = index;
            Source = source;
            SslVerify = sslVerify;
            Token = token;
        }
    }

    [OutputType]
    public sealed class GetProjectLoggingSyslogConfigResult
    {
        public readonly string? Certificate;
        public readonly string? ClientCert;
        public readonly string? ClientKey;
        public readonly string Endpoint;
        public readonly string? Program;
        public readonly string? Protocol;
        public readonly string? Severity;
        public readonly bool SslVerify;
        public readonly string? Token;

        [OutputConstructor]
        private GetProjectLoggingSyslogConfigResult(
            string? certificate,
            string? clientCert,
            string? clientKey,
            string endpoint,
            string? program,
            string? protocol,
            string? severity,
            bool sslVerify,
            string? token)
        {
            Certificate = certificate;
            ClientCert = clientCert;
            ClientKey = clientKey;
            Endpoint = endpoint;
            Program = program;
            Protocol = protocol;
            Severity = severity;
            SslVerify = sslVerify;
            Token = token;
        }
    }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rancher2
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve information about a Rancher v2 etcd backup.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/etcd_backup.html.markdown.
        /// </summary>
        public static Task<GetEtcdBackupResult> GetEtcdBackup(GetEtcdBackupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEtcdBackupResult>("rancher2:index/getEtcdBackup:getEtcdBackup", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetEtcdBackupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID to config Etcd Backup (string)
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the Etcd Backup (string)
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetEtcdBackupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetEtcdBackupResult
    {
        /// <summary>
        /// (Computed) Annotations for Etcd Backup object (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Annotations;
        /// <summary>
        /// (Computed) Backup config for etcd backup (list maxitems:1)
        /// </summary>
        public readonly Outputs.GetEtcdBackupBackupConfigResult BackupConfig;
        public readonly string ClusterId;
        /// <summary>
        /// (Computed) Filename of the Etcd Backup (string)
        /// </summary>
        public readonly string Filename;
        /// <summary>
        /// (Computed) Labels for Etcd Backup object (map)
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        /// <summary>
        /// (Computed) Manual execution of the Etcd Backup. Default `false` (bool)
        /// </summary>
        public readonly bool Manual;
        public readonly string Name;
        /// <summary>
        /// (Computed) Description for the Etcd Backup (string)
        /// </summary>
        public readonly string NamespaceId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetEtcdBackupResult(
            ImmutableDictionary<string, object> annotations,
            Outputs.GetEtcdBackupBackupConfigResult backupConfig,
            string clusterId,
            string filename,
            ImmutableDictionary<string, object> labels,
            bool manual,
            string name,
            string namespaceId,
            string id)
        {
            Annotations = annotations;
            BackupConfig = backupConfig;
            ClusterId = clusterId;
            Filename = filename;
            Labels = labels;
            Manual = manual;
            Name = name;
            NamespaceId = namespaceId;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetEtcdBackupBackupConfigResult
    {
        public readonly bool? Enabled;
        public readonly int? IntervalHours;
        public readonly int? Retention;
        public readonly GetEtcdBackupBackupConfigS3BackupConfigResult? S3BackupConfig;

        [OutputConstructor]
        private GetEtcdBackupBackupConfigResult(
            bool? enabled,
            int? intervalHours,
            int? retention,
            GetEtcdBackupBackupConfigS3BackupConfigResult? s3BackupConfig)
        {
            Enabled = enabled;
            IntervalHours = intervalHours;
            Retention = retention;
            S3BackupConfig = s3BackupConfig;
        }
    }

    [OutputType]
    public sealed class GetEtcdBackupBackupConfigS3BackupConfigResult
    {
        public readonly string? AccessKey;
        public readonly string BucketName;
        public readonly string? CustomCa;
        public readonly string Endpoint;
        public readonly string? Folder;
        public readonly string? Region;
        public readonly string? SecretKey;

        [OutputConstructor]
        private GetEtcdBackupBackupConfigS3BackupConfigResult(
            string? accessKey,
            string bucketName,
            string? customCa,
            string endpoint,
            string? folder,
            string? region,
            string? secretKey)
        {
            AccessKey = accessKey;
            BucketName = bucketName;
            CustomCa = customCa;
            Endpoint = endpoint;
            Folder = folder;
            Region = region;
            SecretKey = secretKey;
        }
    }
    }
}

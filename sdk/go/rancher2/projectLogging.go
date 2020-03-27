// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Project Logging resource. This can be used to create Project Logging for Rancher v2 environments and retrieve their information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/projectLogging.html.markdown.
type ProjectLogging struct {
	pulumi.CustomResourceState

	// Annotations for Project Logging object (map)
	Annotations pulumi.MapOutput `pulumi:"annotations"`
	// The custom target config for Cluster Logging. For `kind = custom`. Conflicts with `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	CustomTargetConfig ProjectLoggingCustomTargetConfigPtrOutput `pulumi:"customTargetConfig"`
	// The elasticsearch config for Project Logging. For `kind = elasticsearch`. Conflicts with `customTargetConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	ElasticsearchConfig ProjectLoggingElasticsearchConfigPtrOutput `pulumi:"elasticsearchConfig"`
	// The fluentd config for Project Logging. For `kind = fluentd`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	FluentdConfig ProjectLoggingFluentdConfigPtrOutput `pulumi:"fluentdConfig"`
	// The kafka config for Project Logging. For `kind = kafka`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	KafkaConfig ProjectLoggingKafkaConfigPtrOutput `pulumi:"kafkaConfig"`
	// The kind of the Project Logging. `elasticsearch`, `fluentd`, `kafka`, `splunk` and `syslog` are supported (string)
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Labels for Project Logging object (map)
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The name of the Project Logging config (string)
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace id from Project logging (string)
	NamespaceId pulumi.StringPtrOutput `pulumi:"namespaceId"`
	// How often buffered logs would be flushed. Default: `3` seconds (int)
	OutputFlushInterval pulumi.IntPtrOutput `pulumi:"outputFlushInterval"`
	// The output tags for Project Logging (map)
	OutputTags pulumi.MapOutput `pulumi:"outputTags"`
	// The project id to configure logging (string)
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The splunk config for Project Logging. For `kind = splunk`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `syslogConfig` (list maxitems:1)
	SplunkConfig ProjectLoggingSplunkConfigPtrOutput `pulumi:"splunkConfig"`
	// The syslog config for Project Logging. For `kind = syslog`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `splunkConfig` (list maxitems:1)
	SyslogConfig ProjectLoggingSyslogConfigPtrOutput `pulumi:"syslogConfig"`
}

// NewProjectLogging registers a new resource with the given unique name, arguments, and options.
func NewProjectLogging(ctx *pulumi.Context,
	name string, args *ProjectLoggingArgs, opts ...pulumi.ResourceOption) (*ProjectLogging, error) {
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil {
		args = &ProjectLoggingArgs{}
	}
	var resource ProjectLogging
	err := ctx.RegisterResource("rancher2:index/projectLogging:ProjectLogging", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectLogging gets an existing ProjectLogging resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectLogging(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectLoggingState, opts ...pulumi.ResourceOption) (*ProjectLogging, error) {
	var resource ProjectLogging
	err := ctx.ReadResource("rancher2:index/projectLogging:ProjectLogging", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectLogging resources.
type projectLoggingState struct {
	// Annotations for Project Logging object (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// The custom target config for Cluster Logging. For `kind = custom`. Conflicts with `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	CustomTargetConfig *ProjectLoggingCustomTargetConfig `pulumi:"customTargetConfig"`
	// The elasticsearch config for Project Logging. For `kind = elasticsearch`. Conflicts with `customTargetConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	ElasticsearchConfig *ProjectLoggingElasticsearchConfig `pulumi:"elasticsearchConfig"`
	// The fluentd config for Project Logging. For `kind = fluentd`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	FluentdConfig *ProjectLoggingFluentdConfig `pulumi:"fluentdConfig"`
	// The kafka config for Project Logging. For `kind = kafka`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	KafkaConfig *ProjectLoggingKafkaConfig `pulumi:"kafkaConfig"`
	// The kind of the Project Logging. `elasticsearch`, `fluentd`, `kafka`, `splunk` and `syslog` are supported (string)
	Kind *string `pulumi:"kind"`
	// Labels for Project Logging object (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The name of the Project Logging config (string)
	Name *string `pulumi:"name"`
	// The namespace id from Project logging (string)
	NamespaceId *string `pulumi:"namespaceId"`
	// How often buffered logs would be flushed. Default: `3` seconds (int)
	OutputFlushInterval *int `pulumi:"outputFlushInterval"`
	// The output tags for Project Logging (map)
	OutputTags map[string]interface{} `pulumi:"outputTags"`
	// The project id to configure logging (string)
	ProjectId *string `pulumi:"projectId"`
	// The splunk config for Project Logging. For `kind = splunk`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `syslogConfig` (list maxitems:1)
	SplunkConfig *ProjectLoggingSplunkConfig `pulumi:"splunkConfig"`
	// The syslog config for Project Logging. For `kind = syslog`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `splunkConfig` (list maxitems:1)
	SyslogConfig *ProjectLoggingSyslogConfig `pulumi:"syslogConfig"`
}

type ProjectLoggingState struct {
	// Annotations for Project Logging object (map)
	Annotations pulumi.MapInput
	// The custom target config for Cluster Logging. For `kind = custom`. Conflicts with `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	CustomTargetConfig ProjectLoggingCustomTargetConfigPtrInput
	// The elasticsearch config for Project Logging. For `kind = elasticsearch`. Conflicts with `customTargetConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	ElasticsearchConfig ProjectLoggingElasticsearchConfigPtrInput
	// The fluentd config for Project Logging. For `kind = fluentd`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	FluentdConfig ProjectLoggingFluentdConfigPtrInput
	// The kafka config for Project Logging. For `kind = kafka`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	KafkaConfig ProjectLoggingKafkaConfigPtrInput
	// The kind of the Project Logging. `elasticsearch`, `fluentd`, `kafka`, `splunk` and `syslog` are supported (string)
	Kind pulumi.StringPtrInput
	// Labels for Project Logging object (map)
	Labels pulumi.MapInput
	// The name of the Project Logging config (string)
	Name pulumi.StringPtrInput
	// The namespace id from Project logging (string)
	NamespaceId pulumi.StringPtrInput
	// How often buffered logs would be flushed. Default: `3` seconds (int)
	OutputFlushInterval pulumi.IntPtrInput
	// The output tags for Project Logging (map)
	OutputTags pulumi.MapInput
	// The project id to configure logging (string)
	ProjectId pulumi.StringPtrInput
	// The splunk config for Project Logging. For `kind = splunk`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `syslogConfig` (list maxitems:1)
	SplunkConfig ProjectLoggingSplunkConfigPtrInput
	// The syslog config for Project Logging. For `kind = syslog`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `splunkConfig` (list maxitems:1)
	SyslogConfig ProjectLoggingSyslogConfigPtrInput
}

func (ProjectLoggingState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLoggingState)(nil)).Elem()
}

type projectLoggingArgs struct {
	// Annotations for Project Logging object (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// The custom target config for Cluster Logging. For `kind = custom`. Conflicts with `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	CustomTargetConfig *ProjectLoggingCustomTargetConfig `pulumi:"customTargetConfig"`
	// The elasticsearch config for Project Logging. For `kind = elasticsearch`. Conflicts with `customTargetConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	ElasticsearchConfig *ProjectLoggingElasticsearchConfig `pulumi:"elasticsearchConfig"`
	// The fluentd config for Project Logging. For `kind = fluentd`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	FluentdConfig *ProjectLoggingFluentdConfig `pulumi:"fluentdConfig"`
	// The kafka config for Project Logging. For `kind = kafka`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	KafkaConfig *ProjectLoggingKafkaConfig `pulumi:"kafkaConfig"`
	// The kind of the Project Logging. `elasticsearch`, `fluentd`, `kafka`, `splunk` and `syslog` are supported (string)
	Kind string `pulumi:"kind"`
	// Labels for Project Logging object (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The name of the Project Logging config (string)
	Name *string `pulumi:"name"`
	// The namespace id from Project logging (string)
	NamespaceId *string `pulumi:"namespaceId"`
	// How often buffered logs would be flushed. Default: `3` seconds (int)
	OutputFlushInterval *int `pulumi:"outputFlushInterval"`
	// The output tags for Project Logging (map)
	OutputTags map[string]interface{} `pulumi:"outputTags"`
	// The project id to configure logging (string)
	ProjectId string `pulumi:"projectId"`
	// The splunk config for Project Logging. For `kind = splunk`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `syslogConfig` (list maxitems:1)
	SplunkConfig *ProjectLoggingSplunkConfig `pulumi:"splunkConfig"`
	// The syslog config for Project Logging. For `kind = syslog`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `splunkConfig` (list maxitems:1)
	SyslogConfig *ProjectLoggingSyslogConfig `pulumi:"syslogConfig"`
}

// The set of arguments for constructing a ProjectLogging resource.
type ProjectLoggingArgs struct {
	// Annotations for Project Logging object (map)
	Annotations pulumi.MapInput
	// The custom target config for Cluster Logging. For `kind = custom`. Conflicts with `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	CustomTargetConfig ProjectLoggingCustomTargetConfigPtrInput
	// The elasticsearch config for Project Logging. For `kind = elasticsearch`. Conflicts with `customTargetConfig`, `fluentdConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	ElasticsearchConfig ProjectLoggingElasticsearchConfigPtrInput
	// The fluentd config for Project Logging. For `kind = fluentd`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `kafkaConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	FluentdConfig ProjectLoggingFluentdConfigPtrInput
	// The kafka config for Project Logging. For `kind = kafka`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `splunkConfig` and `syslogConfig` (list maxitems:1)
	KafkaConfig ProjectLoggingKafkaConfigPtrInput
	// The kind of the Project Logging. `elasticsearch`, `fluentd`, `kafka`, `splunk` and `syslog` are supported (string)
	Kind pulumi.StringInput
	// Labels for Project Logging object (map)
	Labels pulumi.MapInput
	// The name of the Project Logging config (string)
	Name pulumi.StringPtrInput
	// The namespace id from Project logging (string)
	NamespaceId pulumi.StringPtrInput
	// How often buffered logs would be flushed. Default: `3` seconds (int)
	OutputFlushInterval pulumi.IntPtrInput
	// The output tags for Project Logging (map)
	OutputTags pulumi.MapInput
	// The project id to configure logging (string)
	ProjectId pulumi.StringInput
	// The splunk config for Project Logging. For `kind = splunk`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `syslogConfig` (list maxitems:1)
	SplunkConfig ProjectLoggingSplunkConfigPtrInput
	// The syslog config for Project Logging. For `kind = syslog`. Conflicts with `customTargetConfig`, `elasticsearchConfig`, `fluentdConfig`, `kafkaConfig`, and `splunkConfig` (list maxitems:1)
	SyslogConfig ProjectLoggingSyslogConfigPtrInput
}

func (ProjectLoggingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectLoggingArgs)(nil)).Elem()
}


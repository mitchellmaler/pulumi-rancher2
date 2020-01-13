// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a Rancher v2 cluster alert group.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rancher2 from "@pulumi/rancher2";
 * 
 * const foo = rancher2.getClusterAlertGroup({
 *     clusterId: "<cluster_id>",
 *     name: "<cluster_alert_group_name>",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/cluster_alert_group.html.markdown.
 */
export function getClusterAlertGroup(args: GetClusterAlertGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterAlertGroupResult> & GetClusterAlertGroupResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetClusterAlertGroupResult> = pulumi.runtime.invoke("rancher2:index/getClusterAlertGroup:getClusterAlertGroup", {
        "clusterId": args.clusterId,
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getClusterAlertGroup.
 */
export interface GetClusterAlertGroupArgs {
    /**
     * The cluster id where create cluster alert group (string)
     */
    readonly clusterId: string;
    /**
     * The cluster alert group name (string)
     */
    readonly name: string;
}

/**
 * A collection of values returned by getClusterAlertGroup.
 */
export interface GetClusterAlertGroupResult {
    /**
     * (Computed) The cluster alert group annotations (map)
     */
    readonly annotations: {[key: string]: any};
    readonly clusterId: string;
    /**
     * (Computed) The cluster alert group description (string)
     */
    readonly description: string;
    /**
     * (Computed) The cluster alert group interval seconds. Default: `180` (int)
     */
    readonly groupIntervalSeconds: number;
    /**
     * (Computed) The cluster alert group wait seconds. Default: `180` (int)
     */
    readonly groupWaitSeconds: number;
    /**
     * (Computed) The cluster alert group labels (map)
     */
    readonly labels: {[key: string]: any};
    readonly name: string;
    /**
     * (Computed) The cluster alert group recipients (list)
     */
    readonly recipients: outputs.GetClusterAlertGroupRecipient[];
    /**
     * (Computed) The cluster alert group wait seconds. Default: `3600` (int)
     */
    readonly repeatIntervalSeconds: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
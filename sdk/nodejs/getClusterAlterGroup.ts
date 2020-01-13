// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getClusterAlterGroup(args: GetClusterAlterGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterAlterGroupResult> & GetClusterAlterGroupResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetClusterAlterGroupResult> = pulumi.runtime.invoke("rancher2:index/getClusterAlterGroup:getClusterAlterGroup", {
        "clusterId": args.clusterId,
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getClusterAlterGroup.
 */
export interface GetClusterAlterGroupArgs {
    readonly clusterId: string;
    readonly name: string;
}

/**
 * A collection of values returned by getClusterAlterGroup.
 */
export interface GetClusterAlterGroupResult {
    readonly annotations: {[key: string]: any};
    readonly clusterId: string;
    readonly description: string;
    readonly groupIntervalSeconds: number;
    readonly groupWaitSeconds: number;
    readonly labels: {[key: string]: any};
    readonly name: string;
    readonly recipients: outputs.GetClusterAlterGroupRecipient[];
    readonly repeatIntervalSeconds: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
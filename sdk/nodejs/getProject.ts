// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/d/project.html.markdown.
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> & GetProjectResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetProjectResult> = pulumi.runtime.invoke("rancher2:index/getProject:getProject", {
        "clusterId": args.clusterId,
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * ID of the Rancher 2 cluster (string)
     */
    readonly clusterId: string;
    /**
     * The project name (string)
     */
    readonly name: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * (Computed) Annotations of the rancher2 project (map)
     */
    readonly annotations: {[key: string]: any};
    readonly clusterId: string;
    /**
     * (Computed) Default containers resource limits on project (List maxitem:1)
     */
    readonly containerResourceLimit: outputs.GetProjectContainerResourceLimit;
    /**
     * (Computed) The project's description (string)
     */
    readonly description: string;
    /**
     * (Computed) Enable built-in project monitoring. Default `false` (bool)
     */
    readonly enableProjectMonitoring: boolean;
    /**
     * (Computed) Labels of the rancher2 project (map)
     */
    readonly labels: {[key: string]: any};
    readonly name: string;
    /**
     * (Computed) Default Pod Security Policy ID for the project (string)
     */
    readonly podSecurityPolicyTemplateId: string;
    /**
     * (Computed) Resource quota for project. Rancher v2.1.x or higher (list maxitems:1)
     */
    readonly resourceQuota: outputs.GetProjectResourceQuota;
    /**
     * (Computed) UUID of the project as stored by Rancher 2 (string)
     */
    readonly uuid: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

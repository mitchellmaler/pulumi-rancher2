// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class NodePool extends pulumi.CustomResource {
    /**
     * Get an existing NodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodePoolState, opts?: pulumi.CustomResourceOptions): NodePool {
        return new NodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rancher2:index/nodePool:NodePool';

    /**
     * Returns true if the given object is an instance of NodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodePool.__pulumiType;
    }

    /**
     * Annotations for Node Pool object (map)
     */
    public readonly annotations!: pulumi.Output<{[key: string]: any}>;
    /**
     * The RKE cluster id to use Node Pool (string)
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * RKE control plane role for created nodes (bool)
     */
    public readonly controlPlane!: pulumi.Output<boolean | undefined>;
    /**
     * Delete not ready node after secs. For Rancher v2.3.3 or above. Default `0` (int)
     */
    public readonly deleteNotReadyAfterSecs!: pulumi.Output<number | undefined>;
    /**
     * RKE etcd role for created nodes (bool)
     */
    public readonly etcd!: pulumi.Output<boolean | undefined>;
    /**
     * The prefix for created nodes of the Node Pool (string)
     */
    public readonly hostnamePrefix!: pulumi.Output<string>;
    /**
     * Labels for Node Pool object (map)
     */
    public readonly labels!: pulumi.Output<{[key: string]: any}>;
    /**
     * The name of the Node Pool (string)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Node taints. For Rancher v2.3.3 or above (List)
     */
    public readonly nodeTaints!: pulumi.Output<outputs.NodePoolNodeTaint[] | undefined>;
    /**
     * The Node Template ID to use for node creation (string)
     */
    public readonly nodeTemplateId!: pulumi.Output<string>;
    /**
     * The number of nodes to create on Node Pool. Default `1`. Only values >= 1 allowed (int)
     */
    public readonly quantity!: pulumi.Output<number | undefined>;
    /**
     * RKE role role for created nodes (bool)
     */
    public readonly worker!: pulumi.Output<boolean | undefined>;

    /**
     * Create a NodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodePoolArgs | NodePoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NodePoolState | undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["controlPlane"] = state ? state.controlPlane : undefined;
            inputs["deleteNotReadyAfterSecs"] = state ? state.deleteNotReadyAfterSecs : undefined;
            inputs["etcd"] = state ? state.etcd : undefined;
            inputs["hostnamePrefix"] = state ? state.hostnamePrefix : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeTaints"] = state ? state.nodeTaints : undefined;
            inputs["nodeTemplateId"] = state ? state.nodeTemplateId : undefined;
            inputs["quantity"] = state ? state.quantity : undefined;
            inputs["worker"] = state ? state.worker : undefined;
        } else {
            const args = argsOrState as NodePoolArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.hostnamePrefix === undefined) {
                throw new Error("Missing required property 'hostnamePrefix'");
            }
            if (!args || args.nodeTemplateId === undefined) {
                throw new Error("Missing required property 'nodeTemplateId'");
            }
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["controlPlane"] = args ? args.controlPlane : undefined;
            inputs["deleteNotReadyAfterSecs"] = args ? args.deleteNotReadyAfterSecs : undefined;
            inputs["etcd"] = args ? args.etcd : undefined;
            inputs["hostnamePrefix"] = args ? args.hostnamePrefix : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeTaints"] = args ? args.nodeTaints : undefined;
            inputs["nodeTemplateId"] = args ? args.nodeTemplateId : undefined;
            inputs["quantity"] = args ? args.quantity : undefined;
            inputs["worker"] = args ? args.worker : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NodePool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodePool resources.
 */
export interface NodePoolState {
    /**
     * Annotations for Node Pool object (map)
     */
    readonly annotations?: pulumi.Input<{[key: string]: any}>;
    /**
     * The RKE cluster id to use Node Pool (string)
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * RKE control plane role for created nodes (bool)
     */
    readonly controlPlane?: pulumi.Input<boolean>;
    /**
     * Delete not ready node after secs. For Rancher v2.3.3 or above. Default `0` (int)
     */
    readonly deleteNotReadyAfterSecs?: pulumi.Input<number>;
    /**
     * RKE etcd role for created nodes (bool)
     */
    readonly etcd?: pulumi.Input<boolean>;
    /**
     * The prefix for created nodes of the Node Pool (string)
     */
    readonly hostnamePrefix?: pulumi.Input<string>;
    /**
     * Labels for Node Pool object (map)
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the Node Pool (string)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Node taints. For Rancher v2.3.3 or above (List)
     */
    readonly nodeTaints?: pulumi.Input<pulumi.Input<inputs.NodePoolNodeTaint>[]>;
    /**
     * The Node Template ID to use for node creation (string)
     */
    readonly nodeTemplateId?: pulumi.Input<string>;
    /**
     * The number of nodes to create on Node Pool. Default `1`. Only values >= 1 allowed (int)
     */
    readonly quantity?: pulumi.Input<number>;
    /**
     * RKE role role for created nodes (bool)
     */
    readonly worker?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a NodePool resource.
 */
export interface NodePoolArgs {
    /**
     * Annotations for Node Pool object (map)
     */
    readonly annotations?: pulumi.Input<{[key: string]: any}>;
    /**
     * The RKE cluster id to use Node Pool (string)
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * RKE control plane role for created nodes (bool)
     */
    readonly controlPlane?: pulumi.Input<boolean>;
    /**
     * Delete not ready node after secs. For Rancher v2.3.3 or above. Default `0` (int)
     */
    readonly deleteNotReadyAfterSecs?: pulumi.Input<number>;
    /**
     * RKE etcd role for created nodes (bool)
     */
    readonly etcd?: pulumi.Input<boolean>;
    /**
     * The prefix for created nodes of the Node Pool (string)
     */
    readonly hostnamePrefix: pulumi.Input<string>;
    /**
     * Labels for Node Pool object (map)
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the Node Pool (string)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Node taints. For Rancher v2.3.3 or above (List)
     */
    readonly nodeTaints?: pulumi.Input<pulumi.Input<inputs.NodePoolNodeTaint>[]>;
    /**
     * The Node Template ID to use for node creation (string)
     */
    readonly nodeTemplateId: pulumi.Input<string>;
    /**
     * The number of nodes to create on Node Pool. Default `1`. Only values >= 1 allowed (int)
     */
    readonly quantity?: pulumi.Input<number>;
    /**
     * RKE role role for created nodes (bool)
     */
    readonly worker?: pulumi.Input<boolean>;
}

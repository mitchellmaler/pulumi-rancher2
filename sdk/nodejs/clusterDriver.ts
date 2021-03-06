// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a Rancher v2 Cluster Driver resource. This can be used to create Cluster Driver for Rancher v2.2.x Kontainer Engine clusters and retrieve their information.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rancher2 from "@pulumi/rancher2";
 * 
 * // Create a new Rancher2 Cluster Driver
 * const foo = new rancher2.ClusterDriver("foo", {
 *     active: true,
 *     builtin: false,
 *     checksum: "0x0",
 *     description: "Foo description",
 *     externalId: "fooExternal",
 *     uiUrl: "local://ui",
 *     url: "local://",
 *     whitelistDomains: ["*.foo.com"],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/clusterDriver.html.markdown.
 */
export class ClusterDriver extends pulumi.CustomResource {
    /**
     * Get an existing ClusterDriver resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterDriverState, opts?: pulumi.CustomResourceOptions): ClusterDriver {
        return new ClusterDriver(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rancher2:index/clusterDriver:ClusterDriver';

    /**
     * Returns true if the given object is an instance of ClusterDriver.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterDriver {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterDriver.__pulumiType;
    }

    /**
     * Specify the cluster driver state (bool)
     */
    public readonly active!: pulumi.Output<boolean>;
    /**
     * Actual url of the cluster driver (string)
     */
    public readonly actualUrl!: pulumi.Output<string | undefined>;
    /**
     * Annotations of the resource (map)
     */
    public readonly annotations!: pulumi.Output<{[key: string]: any}>;
    /**
     * Specify whether the cluster driver is an internal cluster driver or not (bool)
     */
    public readonly builtin!: pulumi.Output<boolean>;
    /**
     * Verify that the downloaded driver matches the expected checksum (string)
     */
    public readonly checksum!: pulumi.Output<string | undefined>;
    /**
     * Labels of the resource (map)
     */
    public readonly labels!: pulumi.Output<{[key: string]: any}>;
    /**
     * Name of the cluster driver (string)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL to load for customized Add Clusters screen for this driver (string)
     */
    public readonly uiUrl!: pulumi.Output<string | undefined>;
    /**
     * The URL to download the machine driver binary for 64-bit Linux (string)
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Domains to whitelist for the ui (list)
     */
    public readonly whitelistDomains!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ClusterDriver resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterDriverArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterDriverArgs | ClusterDriverState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterDriverState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["actualUrl"] = state ? state.actualUrl : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["builtin"] = state ? state.builtin : undefined;
            inputs["checksum"] = state ? state.checksum : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["uiUrl"] = state ? state.uiUrl : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["whitelistDomains"] = state ? state.whitelistDomains : undefined;
        } else {
            const args = argsOrState as ClusterDriverArgs | undefined;
            if (!args || args.active === undefined) {
                throw new Error("Missing required property 'active'");
            }
            if (!args || args.builtin === undefined) {
                throw new Error("Missing required property 'builtin'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["actualUrl"] = args ? args.actualUrl : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["builtin"] = args ? args.builtin : undefined;
            inputs["checksum"] = args ? args.checksum : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["uiUrl"] = args ? args.uiUrl : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["whitelistDomains"] = args ? args.whitelistDomains : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClusterDriver.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterDriver resources.
 */
export interface ClusterDriverState {
    /**
     * Specify the cluster driver state (bool)
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * Actual url of the cluster driver (string)
     */
    readonly actualUrl?: pulumi.Input<string>;
    /**
     * Annotations of the resource (map)
     */
    readonly annotations?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specify whether the cluster driver is an internal cluster driver or not (bool)
     */
    readonly builtin?: pulumi.Input<boolean>;
    /**
     * Verify that the downloaded driver matches the expected checksum (string)
     */
    readonly checksum?: pulumi.Input<string>;
    /**
     * Labels of the resource (map)
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Name of the cluster driver (string)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The URL to load for customized Add Clusters screen for this driver (string)
     */
    readonly uiUrl?: pulumi.Input<string>;
    /**
     * The URL to download the machine driver binary for 64-bit Linux (string)
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Domains to whitelist for the ui (list)
     */
    readonly whitelistDomains?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ClusterDriver resource.
 */
export interface ClusterDriverArgs {
    /**
     * Specify the cluster driver state (bool)
     */
    readonly active: pulumi.Input<boolean>;
    /**
     * Actual url of the cluster driver (string)
     */
    readonly actualUrl?: pulumi.Input<string>;
    /**
     * Annotations of the resource (map)
     */
    readonly annotations?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specify whether the cluster driver is an internal cluster driver or not (bool)
     */
    readonly builtin: pulumi.Input<boolean>;
    /**
     * Verify that the downloaded driver matches the expected checksum (string)
     */
    readonly checksum?: pulumi.Input<string>;
    /**
     * Labels of the resource (map)
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Name of the cluster driver (string)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The URL to load for customized Add Clusters screen for this driver (string)
     */
    readonly uiUrl?: pulumi.Input<string>;
    /**
     * The URL to download the machine driver binary for 64-bit Linux (string)
     */
    readonly url: pulumi.Input<string>;
    /**
     * Domains to whitelist for the ui (list)
     */
    readonly whitelistDomains?: pulumi.Input<pulumi.Input<string>[]>;
}

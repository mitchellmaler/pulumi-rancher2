// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/bootstrap.html.markdown.
 */
export class Bootstrap extends pulumi.CustomResource {
    /**
     * Get an existing Bootstrap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BootstrapState, opts?: pulumi.CustomResourceOptions): Bootstrap {
        return new Bootstrap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rancher2:index/bootstrap:Bootstrap';

    /**
     * Returns true if the given object is an instance of Bootstrap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bootstrap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bootstrap.__pulumiType;
    }

    /**
     * Current password for Admin user. Just needed for recover if admin password has been changed from other resources and token is expired (string)
     */
    public readonly currentPassword!: pulumi.Output<string>;
    /**
     * Password for Admin user or random generated if empty (string)
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Send telemetry anonymous data. Default: `false` (bool)
     */
    public readonly telemetry!: pulumi.Output<boolean | undefined>;
    /**
     * (Computed) Generated API temporary token as helper. Should be empty (string)
     */
    public /*out*/ readonly tempToken!: pulumi.Output<string>;
    /**
     * (Computed) Generated API temporary token id as helper. Should be empty (string)
     */
    public /*out*/ readonly tempTokenId!: pulumi.Output<string>;
    /**
     * (Computed) Generated API token for Admin User (string)
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * (Computed) Generated API token id for Admin User (string)
     */
    public /*out*/ readonly tokenId!: pulumi.Output<string>;
    /**
     * TTL in seconds for generated admin token. Default: `0`  (int)
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * Regenerate admin token. Default: `false` (bool)
     */
    public readonly tokenUpdate!: pulumi.Output<boolean | undefined>;
    /**
     * (Computed) URL set as server-url (string)
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * (Computed) Admin username (string)
     */
    public /*out*/ readonly user!: pulumi.Output<string>;

    /**
     * Create a Bootstrap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BootstrapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BootstrapArgs | BootstrapState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BootstrapState | undefined;
            inputs["currentPassword"] = state ? state.currentPassword : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["telemetry"] = state ? state.telemetry : undefined;
            inputs["tempToken"] = state ? state.tempToken : undefined;
            inputs["tempTokenId"] = state ? state.tempTokenId : undefined;
            inputs["token"] = state ? state.token : undefined;
            inputs["tokenId"] = state ? state.tokenId : undefined;
            inputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            inputs["tokenUpdate"] = state ? state.tokenUpdate : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as BootstrapArgs | undefined;
            inputs["currentPassword"] = args ? args.currentPassword : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["telemetry"] = args ? args.telemetry : undefined;
            inputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            inputs["tokenUpdate"] = args ? args.tokenUpdate : undefined;
            inputs["tempToken"] = undefined /*out*/;
            inputs["tempTokenId"] = undefined /*out*/;
            inputs["token"] = undefined /*out*/;
            inputs["tokenId"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
            inputs["user"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Bootstrap.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Bootstrap resources.
 */
export interface BootstrapState {
    /**
     * Current password for Admin user. Just needed for recover if admin password has been changed from other resources and token is expired (string)
     */
    readonly currentPassword?: pulumi.Input<string>;
    /**
     * Password for Admin user or random generated if empty (string)
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Send telemetry anonymous data. Default: `false` (bool)
     */
    readonly telemetry?: pulumi.Input<boolean>;
    /**
     * (Computed) Generated API temporary token as helper. Should be empty (string)
     */
    readonly tempToken?: pulumi.Input<string>;
    /**
     * (Computed) Generated API temporary token id as helper. Should be empty (string)
     */
    readonly tempTokenId?: pulumi.Input<string>;
    /**
     * (Computed) Generated API token for Admin User (string)
     */
    readonly token?: pulumi.Input<string>;
    /**
     * (Computed) Generated API token id for Admin User (string)
     */
    readonly tokenId?: pulumi.Input<string>;
    /**
     * TTL in seconds for generated admin token. Default: `0`  (int)
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * Regenerate admin token. Default: `false` (bool)
     */
    readonly tokenUpdate?: pulumi.Input<boolean>;
    /**
     * (Computed) URL set as server-url (string)
     */
    readonly url?: pulumi.Input<string>;
    /**
     * (Computed) Admin username (string)
     */
    readonly user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Bootstrap resource.
 */
export interface BootstrapArgs {
    /**
     * Current password for Admin user. Just needed for recover if admin password has been changed from other resources and token is expired (string)
     */
    readonly currentPassword?: pulumi.Input<string>;
    /**
     * Password for Admin user or random generated if empty (string)
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Send telemetry anonymous data. Default: `false` (bool)
     */
    readonly telemetry?: pulumi.Input<boolean>;
    /**
     * TTL in seconds for generated admin token. Default: `0`  (int)
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * Regenerate admin token. Default: `false` (bool)
     */
    readonly tokenUpdate?: pulumi.Input<boolean>;
}
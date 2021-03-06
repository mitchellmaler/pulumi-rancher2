// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Rancher v2 Auth Config KeyCloak resource. This can be used to configure and enable Auth Config KeyCloak for Rancher v2 RKE clusters and retrieve their information.
 * 
 * In addition to the built-in local auth, only one external auth config provider can be enabled at a time.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rancher2 from "@pulumi/rancher2";
 * 
 * // Create a new rancher2 Auth Config KeyCloak
 * const keycloak = new rancher2.AuthConfigKeycloak("keycloak", {
 *     displayNameField: "<DISPLAY_NAME_FIELD>",
 *     groupsField: "<GROUPS_FIELD>",
 *     idpMetadataContent: "<IDP_METADATA_CONTENT>",
 *     rancherApiHost: "https://<RANCHER_API_HOST>",
 *     spCert: "<SP_CERT>",
 *     spKey: "<SP_KEY>",
 *     uidField: "<UID_FIELD>",
 *     userNameField: "<USER_NAME_FIELD>",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/authConfigKeyCloak.html.markdown.
 */
export class AuthConfigKeycloak extends pulumi.CustomResource {
    /**
     * Get an existing AuthConfigKeycloak resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthConfigKeycloakState, opts?: pulumi.CustomResourceOptions): AuthConfigKeycloak {
        return new AuthConfigKeycloak(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rancher2:index/authConfigKeycloak:AuthConfigKeycloak';

    /**
     * Returns true if the given object is an instance of AuthConfigKeycloak.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthConfigKeycloak {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthConfigKeycloak.__pulumiType;
    }

    /**
     * Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
     */
    public readonly accessMode!: pulumi.Output<string | undefined>;
    /**
     * Allowed principal ids for auth. Required if `accessMode` is `required` or `restricted`. Ex: `keycloak_user://<USER_ID>`  `keycloak_group://<GROUP_ID>` (list)
     */
    public readonly allowedPrincipalIds!: pulumi.Output<string[] | undefined>;
    /**
     * Annotations of the resource (map)
     */
    public readonly annotations!: pulumi.Output<{[key: string]: any}>;
    /**
     * KeyCloak display name field (string)
     */
    public readonly displayNameField!: pulumi.Output<string>;
    /**
     * Enable auth config provider. Default `true` (bool)
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * KeyCloak group field (string)
     */
    public readonly groupsField!: pulumi.Output<string>;
    /**
     * KeyCloak IDP metadata content (string)
     */
    public readonly idpMetadataContent!: pulumi.Output<string>;
    /**
     * Labels of the resource (map)
     */
    public readonly labels!: pulumi.Output<{[key: string]: any}>;
    /**
     * (Computed) The name of the resource (string)
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Rancher url. Schema needs to be specified, `https://<RANCHER_API_HOST>` (string)
     */
    public readonly rancherApiHost!: pulumi.Output<string>;
    /**
     * KeyCloak SP cert (string)
     */
    public readonly spCert!: pulumi.Output<string>;
    /**
     * KeyCloak SP key (string)
     */
    public readonly spKey!: pulumi.Output<string>;
    /**
     * (Computed) The type of the resource (string)
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * KeyCloak UID field (string)
     */
    public readonly uidField!: pulumi.Output<string>;
    /**
     * KeyCloak user name field (string)
     */
    public readonly userNameField!: pulumi.Output<string>;

    /**
     * Create a AuthConfigKeycloak resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthConfigKeycloakArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthConfigKeycloakArgs | AuthConfigKeycloakState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthConfigKeycloakState | undefined;
            inputs["accessMode"] = state ? state.accessMode : undefined;
            inputs["allowedPrincipalIds"] = state ? state.allowedPrincipalIds : undefined;
            inputs["annotations"] = state ? state.annotations : undefined;
            inputs["displayNameField"] = state ? state.displayNameField : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["groupsField"] = state ? state.groupsField : undefined;
            inputs["idpMetadataContent"] = state ? state.idpMetadataContent : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rancherApiHost"] = state ? state.rancherApiHost : undefined;
            inputs["spCert"] = state ? state.spCert : undefined;
            inputs["spKey"] = state ? state.spKey : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["uidField"] = state ? state.uidField : undefined;
            inputs["userNameField"] = state ? state.userNameField : undefined;
        } else {
            const args = argsOrState as AuthConfigKeycloakArgs | undefined;
            if (!args || args.displayNameField === undefined) {
                throw new Error("Missing required property 'displayNameField'");
            }
            if (!args || args.groupsField === undefined) {
                throw new Error("Missing required property 'groupsField'");
            }
            if (!args || args.idpMetadataContent === undefined) {
                throw new Error("Missing required property 'idpMetadataContent'");
            }
            if (!args || args.rancherApiHost === undefined) {
                throw new Error("Missing required property 'rancherApiHost'");
            }
            if (!args || args.spCert === undefined) {
                throw new Error("Missing required property 'spCert'");
            }
            if (!args || args.spKey === undefined) {
                throw new Error("Missing required property 'spKey'");
            }
            if (!args || args.uidField === undefined) {
                throw new Error("Missing required property 'uidField'");
            }
            if (!args || args.userNameField === undefined) {
                throw new Error("Missing required property 'userNameField'");
            }
            inputs["accessMode"] = args ? args.accessMode : undefined;
            inputs["allowedPrincipalIds"] = args ? args.allowedPrincipalIds : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["displayNameField"] = args ? args.displayNameField : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["groupsField"] = args ? args.groupsField : undefined;
            inputs["idpMetadataContent"] = args ? args.idpMetadataContent : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["rancherApiHost"] = args ? args.rancherApiHost : undefined;
            inputs["spCert"] = args ? args.spCert : undefined;
            inputs["spKey"] = args ? args.spKey : undefined;
            inputs["uidField"] = args ? args.uidField : undefined;
            inputs["userNameField"] = args ? args.userNameField : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthConfigKeycloak.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthConfigKeycloak resources.
 */
export interface AuthConfigKeycloakState {
    /**
     * Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
     */
    readonly accessMode?: pulumi.Input<string>;
    /**
     * Allowed principal ids for auth. Required if `accessMode` is `required` or `restricted`. Ex: `keycloak_user://<USER_ID>`  `keycloak_group://<GROUP_ID>` (list)
     */
    readonly allowedPrincipalIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Annotations of the resource (map)
     */
    readonly annotations?: pulumi.Input<{[key: string]: any}>;
    /**
     * KeyCloak display name field (string)
     */
    readonly displayNameField?: pulumi.Input<string>;
    /**
     * Enable auth config provider. Default `true` (bool)
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * KeyCloak group field (string)
     */
    readonly groupsField?: pulumi.Input<string>;
    /**
     * KeyCloak IDP metadata content (string)
     */
    readonly idpMetadataContent?: pulumi.Input<string>;
    /**
     * Labels of the resource (map)
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Computed) The name of the resource (string)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Rancher url. Schema needs to be specified, `https://<RANCHER_API_HOST>` (string)
     */
    readonly rancherApiHost?: pulumi.Input<string>;
    /**
     * KeyCloak SP cert (string)
     */
    readonly spCert?: pulumi.Input<string>;
    /**
     * KeyCloak SP key (string)
     */
    readonly spKey?: pulumi.Input<string>;
    /**
     * (Computed) The type of the resource (string)
     */
    readonly type?: pulumi.Input<string>;
    /**
     * KeyCloak UID field (string)
     */
    readonly uidField?: pulumi.Input<string>;
    /**
     * KeyCloak user name field (string)
     */
    readonly userNameField?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthConfigKeycloak resource.
 */
export interface AuthConfigKeycloakArgs {
    /**
     * Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
     */
    readonly accessMode?: pulumi.Input<string>;
    /**
     * Allowed principal ids for auth. Required if `accessMode` is `required` or `restricted`. Ex: `keycloak_user://<USER_ID>`  `keycloak_group://<GROUP_ID>` (list)
     */
    readonly allowedPrincipalIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Annotations of the resource (map)
     */
    readonly annotations?: pulumi.Input<{[key: string]: any}>;
    /**
     * KeyCloak display name field (string)
     */
    readonly displayNameField: pulumi.Input<string>;
    /**
     * Enable auth config provider. Default `true` (bool)
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * KeyCloak group field (string)
     */
    readonly groupsField: pulumi.Input<string>;
    /**
     * KeyCloak IDP metadata content (string)
     */
    readonly idpMetadataContent: pulumi.Input<string>;
    /**
     * Labels of the resource (map)
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Rancher url. Schema needs to be specified, `https://<RANCHER_API_HOST>` (string)
     */
    readonly rancherApiHost: pulumi.Input<string>;
    /**
     * KeyCloak SP cert (string)
     */
    readonly spCert: pulumi.Input<string>;
    /**
     * KeyCloak SP key (string)
     */
    readonly spKey: pulumi.Input<string>;
    /**
     * KeyCloak UID field (string)
     */
    readonly uidField: pulumi.Input<string>;
    /**
     * KeyCloak user name field (string)
     */
    readonly userNameField: pulumi.Input<string>;
}

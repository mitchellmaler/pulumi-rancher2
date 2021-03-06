// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 Auth Config OpenLdap resource. This can be used to configure and enable Auth Config OpenLdap for Rancher v2 RKE clusters and retrieve their information.
//
// In addition to the built-in local auth, only one external auth config provider can be enabled at a time.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/authConfigOpenLdap.html.markdown.
type AuthConfigOpenLdap struct {
	pulumi.CustomResourceState

	// Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
	AccessMode pulumi.StringPtrOutput `pulumi:"accessMode"`
	// Allowed principal ids for auth. Required if `accessMode` is `required` or `restricted`. Ex: `openldap_user://<DN>`  `openldap_group://<DN>` (list)
	AllowedPrincipalIds pulumi.StringArrayOutput `pulumi:"allowedPrincipalIds"`
	// Annotations of the resource (map)
	Annotations pulumi.MapOutput `pulumi:"annotations"`
	// Base64 encoded CA certificate for TLS if self-signed. Use filebase64(<FILE>) for encoding file (string)
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// OpenLdap connection timeout. Default `5000` (int)
	ConnectionTimeout pulumi.IntPtrOutput `pulumi:"connectionTimeout"`
	// Enable auth config provider. Default `true` (bool)
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Group DN attribute. Default `entryDN` (string)
	GroupDnAttribute pulumi.StringOutput `pulumi:"groupDnAttribute"`
	// Group member mapping attribute. Default `member` (string)
	GroupMemberMappingAttribute pulumi.StringOutput `pulumi:"groupMemberMappingAttribute"`
	// Group member user attribute. Default `entryDN` (string)
	GroupMemberUserAttribute pulumi.StringOutput `pulumi:"groupMemberUserAttribute"`
	// Group name attribute. Default `cn` (string)
	GroupNameAttribute pulumi.StringOutput `pulumi:"groupNameAttribute"`
	// Group object class. Default `groupOfNames` (string)
	GroupObjectClass pulumi.StringOutput `pulumi:"groupObjectClass"`
	// Group search attribute. Default `cn` (string)
	GroupSearchAttribute pulumi.StringOutput `pulumi:"groupSearchAttribute"`
	// Group search base (string)
	GroupSearchBase pulumi.StringOutput `pulumi:"groupSearchBase"`
	// Labels of the resource (map)
	Labels pulumi.MapOutput `pulumi:"labels"`
	// (Computed) The name of the resource (string)
	Name pulumi.StringOutput `pulumi:"name"`
	// Nested group membership enable. Default `false` (bool)
	NestedGroupMembershipEnabled pulumi.BoolOutput `pulumi:"nestedGroupMembershipEnabled"`
	// OpenLdap port. Default `389` (int)
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// OpenLdap servers list (list)
	Servers pulumi.StringArrayOutput `pulumi:"servers"`
	// Service account DN for access OpenLdap service (string)
	ServiceAccountDistinguishedName pulumi.StringOutput `pulumi:"serviceAccountDistinguishedName"`
	// Service account password for access OpenLdap service (string)
	ServiceAccountPassword pulumi.StringOutput `pulumi:"serviceAccountPassword"`
	// Enable TLS connection (bool)
	Tls pulumi.BoolOutput `pulumi:"tls"`
	// (Computed) The type of the resource (string)
	Type pulumi.StringOutput `pulumi:"type"`
	// User disabled bit mask (int)
	UserDisabledBitMask pulumi.IntOutput `pulumi:"userDisabledBitMask"`
	// User enable attribute (string)
	UserEnabledAttribute pulumi.StringOutput `pulumi:"userEnabledAttribute"`
	// User login attribute. Default `uid` (string)
	UserLoginAttribute pulumi.StringOutput `pulumi:"userLoginAttribute"`
	// User member attribute. Default `memberOf` (string)
	UserMemberAttribute pulumi.StringOutput `pulumi:"userMemberAttribute"`
	// User name attribute. Default `givenName` (string)
	UserNameAttribute pulumi.StringOutput `pulumi:"userNameAttribute"`
	// User object class. Default `inetorgperson` (string)
	UserObjectClass pulumi.StringOutput `pulumi:"userObjectClass"`
	// User search attribute. Default `uid|sn|givenName` (string)
	UserSearchAttribute pulumi.StringOutput `pulumi:"userSearchAttribute"`
	// User search base DN (string)
	UserSearchBase pulumi.StringOutput `pulumi:"userSearchBase"`
}

// NewAuthConfigOpenLdap registers a new resource with the given unique name, arguments, and options.
func NewAuthConfigOpenLdap(ctx *pulumi.Context,
	name string, args *AuthConfigOpenLdapArgs, opts ...pulumi.ResourceOption) (*AuthConfigOpenLdap, error) {
	if args == nil || args.Servers == nil {
		return nil, errors.New("missing required argument 'Servers'")
	}
	if args == nil || args.ServiceAccountDistinguishedName == nil {
		return nil, errors.New("missing required argument 'ServiceAccountDistinguishedName'")
	}
	if args == nil || args.ServiceAccountPassword == nil {
		return nil, errors.New("missing required argument 'ServiceAccountPassword'")
	}
	if args == nil || args.UserSearchBase == nil {
		return nil, errors.New("missing required argument 'UserSearchBase'")
	}
	if args == nil {
		args = &AuthConfigOpenLdapArgs{}
	}
	var resource AuthConfigOpenLdap
	err := ctx.RegisterResource("rancher2:index/authConfigOpenLdap:AuthConfigOpenLdap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthConfigOpenLdap gets an existing AuthConfigOpenLdap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthConfigOpenLdap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthConfigOpenLdapState, opts ...pulumi.ResourceOption) (*AuthConfigOpenLdap, error) {
	var resource AuthConfigOpenLdap
	err := ctx.ReadResource("rancher2:index/authConfigOpenLdap:AuthConfigOpenLdap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthConfigOpenLdap resources.
type authConfigOpenLdapState struct {
	// Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
	AccessMode *string `pulumi:"accessMode"`
	// Allowed principal ids for auth. Required if `accessMode` is `required` or `restricted`. Ex: `openldap_user://<DN>`  `openldap_group://<DN>` (list)
	AllowedPrincipalIds []string `pulumi:"allowedPrincipalIds"`
	// Annotations of the resource (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// Base64 encoded CA certificate for TLS if self-signed. Use filebase64(<FILE>) for encoding file (string)
	Certificate *string `pulumi:"certificate"`
	// OpenLdap connection timeout. Default `5000` (int)
	ConnectionTimeout *int `pulumi:"connectionTimeout"`
	// Enable auth config provider. Default `true` (bool)
	Enabled *bool `pulumi:"enabled"`
	// Group DN attribute. Default `entryDN` (string)
	GroupDnAttribute *string `pulumi:"groupDnAttribute"`
	// Group member mapping attribute. Default `member` (string)
	GroupMemberMappingAttribute *string `pulumi:"groupMemberMappingAttribute"`
	// Group member user attribute. Default `entryDN` (string)
	GroupMemberUserAttribute *string `pulumi:"groupMemberUserAttribute"`
	// Group name attribute. Default `cn` (string)
	GroupNameAttribute *string `pulumi:"groupNameAttribute"`
	// Group object class. Default `groupOfNames` (string)
	GroupObjectClass *string `pulumi:"groupObjectClass"`
	// Group search attribute. Default `cn` (string)
	GroupSearchAttribute *string `pulumi:"groupSearchAttribute"`
	// Group search base (string)
	GroupSearchBase *string `pulumi:"groupSearchBase"`
	// Labels of the resource (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// (Computed) The name of the resource (string)
	Name *string `pulumi:"name"`
	// Nested group membership enable. Default `false` (bool)
	NestedGroupMembershipEnabled *bool `pulumi:"nestedGroupMembershipEnabled"`
	// OpenLdap port. Default `389` (int)
	Port *int `pulumi:"port"`
	// OpenLdap servers list (list)
	Servers []string `pulumi:"servers"`
	// Service account DN for access OpenLdap service (string)
	ServiceAccountDistinguishedName *string `pulumi:"serviceAccountDistinguishedName"`
	// Service account password for access OpenLdap service (string)
	ServiceAccountPassword *string `pulumi:"serviceAccountPassword"`
	// Enable TLS connection (bool)
	Tls *bool `pulumi:"tls"`
	// (Computed) The type of the resource (string)
	Type *string `pulumi:"type"`
	// User disabled bit mask (int)
	UserDisabledBitMask *int `pulumi:"userDisabledBitMask"`
	// User enable attribute (string)
	UserEnabledAttribute *string `pulumi:"userEnabledAttribute"`
	// User login attribute. Default `uid` (string)
	UserLoginAttribute *string `pulumi:"userLoginAttribute"`
	// User member attribute. Default `memberOf` (string)
	UserMemberAttribute *string `pulumi:"userMemberAttribute"`
	// User name attribute. Default `givenName` (string)
	UserNameAttribute *string `pulumi:"userNameAttribute"`
	// User object class. Default `inetorgperson` (string)
	UserObjectClass *string `pulumi:"userObjectClass"`
	// User search attribute. Default `uid|sn|givenName` (string)
	UserSearchAttribute *string `pulumi:"userSearchAttribute"`
	// User search base DN (string)
	UserSearchBase *string `pulumi:"userSearchBase"`
}

type AuthConfigOpenLdapState struct {
	// Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
	AccessMode pulumi.StringPtrInput
	// Allowed principal ids for auth. Required if `accessMode` is `required` or `restricted`. Ex: `openldap_user://<DN>`  `openldap_group://<DN>` (list)
	AllowedPrincipalIds pulumi.StringArrayInput
	// Annotations of the resource (map)
	Annotations pulumi.MapInput
	// Base64 encoded CA certificate for TLS if self-signed. Use filebase64(<FILE>) for encoding file (string)
	Certificate pulumi.StringPtrInput
	// OpenLdap connection timeout. Default `5000` (int)
	ConnectionTimeout pulumi.IntPtrInput
	// Enable auth config provider. Default `true` (bool)
	Enabled pulumi.BoolPtrInput
	// Group DN attribute. Default `entryDN` (string)
	GroupDnAttribute pulumi.StringPtrInput
	// Group member mapping attribute. Default `member` (string)
	GroupMemberMappingAttribute pulumi.StringPtrInput
	// Group member user attribute. Default `entryDN` (string)
	GroupMemberUserAttribute pulumi.StringPtrInput
	// Group name attribute. Default `cn` (string)
	GroupNameAttribute pulumi.StringPtrInput
	// Group object class. Default `groupOfNames` (string)
	GroupObjectClass pulumi.StringPtrInput
	// Group search attribute. Default `cn` (string)
	GroupSearchAttribute pulumi.StringPtrInput
	// Group search base (string)
	GroupSearchBase pulumi.StringPtrInput
	// Labels of the resource (map)
	Labels pulumi.MapInput
	// (Computed) The name of the resource (string)
	Name pulumi.StringPtrInput
	// Nested group membership enable. Default `false` (bool)
	NestedGroupMembershipEnabled pulumi.BoolPtrInput
	// OpenLdap port. Default `389` (int)
	Port pulumi.IntPtrInput
	// OpenLdap servers list (list)
	Servers pulumi.StringArrayInput
	// Service account DN for access OpenLdap service (string)
	ServiceAccountDistinguishedName pulumi.StringPtrInput
	// Service account password for access OpenLdap service (string)
	ServiceAccountPassword pulumi.StringPtrInput
	// Enable TLS connection (bool)
	Tls pulumi.BoolPtrInput
	// (Computed) The type of the resource (string)
	Type pulumi.StringPtrInput
	// User disabled bit mask (int)
	UserDisabledBitMask pulumi.IntPtrInput
	// User enable attribute (string)
	UserEnabledAttribute pulumi.StringPtrInput
	// User login attribute. Default `uid` (string)
	UserLoginAttribute pulumi.StringPtrInput
	// User member attribute. Default `memberOf` (string)
	UserMemberAttribute pulumi.StringPtrInput
	// User name attribute. Default `givenName` (string)
	UserNameAttribute pulumi.StringPtrInput
	// User object class. Default `inetorgperson` (string)
	UserObjectClass pulumi.StringPtrInput
	// User search attribute. Default `uid|sn|givenName` (string)
	UserSearchAttribute pulumi.StringPtrInput
	// User search base DN (string)
	UserSearchBase pulumi.StringPtrInput
}

func (AuthConfigOpenLdapState) ElementType() reflect.Type {
	return reflect.TypeOf((*authConfigOpenLdapState)(nil)).Elem()
}

type authConfigOpenLdapArgs struct {
	// Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
	AccessMode *string `pulumi:"accessMode"`
	// Allowed principal ids for auth. Required if `accessMode` is `required` or `restricted`. Ex: `openldap_user://<DN>`  `openldap_group://<DN>` (list)
	AllowedPrincipalIds []string `pulumi:"allowedPrincipalIds"`
	// Annotations of the resource (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	// Base64 encoded CA certificate for TLS if self-signed. Use filebase64(<FILE>) for encoding file (string)
	Certificate *string `pulumi:"certificate"`
	// OpenLdap connection timeout. Default `5000` (int)
	ConnectionTimeout *int `pulumi:"connectionTimeout"`
	// Enable auth config provider. Default `true` (bool)
	Enabled *bool `pulumi:"enabled"`
	// Group DN attribute. Default `entryDN` (string)
	GroupDnAttribute *string `pulumi:"groupDnAttribute"`
	// Group member mapping attribute. Default `member` (string)
	GroupMemberMappingAttribute *string `pulumi:"groupMemberMappingAttribute"`
	// Group member user attribute. Default `entryDN` (string)
	GroupMemberUserAttribute *string `pulumi:"groupMemberUserAttribute"`
	// Group name attribute. Default `cn` (string)
	GroupNameAttribute *string `pulumi:"groupNameAttribute"`
	// Group object class. Default `groupOfNames` (string)
	GroupObjectClass *string `pulumi:"groupObjectClass"`
	// Group search attribute. Default `cn` (string)
	GroupSearchAttribute *string `pulumi:"groupSearchAttribute"`
	// Group search base (string)
	GroupSearchBase *string `pulumi:"groupSearchBase"`
	// Labels of the resource (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// Nested group membership enable. Default `false` (bool)
	NestedGroupMembershipEnabled *bool `pulumi:"nestedGroupMembershipEnabled"`
	// OpenLdap port. Default `389` (int)
	Port *int `pulumi:"port"`
	// OpenLdap servers list (list)
	Servers []string `pulumi:"servers"`
	// Service account DN for access OpenLdap service (string)
	ServiceAccountDistinguishedName string `pulumi:"serviceAccountDistinguishedName"`
	// Service account password for access OpenLdap service (string)
	ServiceAccountPassword string `pulumi:"serviceAccountPassword"`
	// Enable TLS connection (bool)
	Tls *bool `pulumi:"tls"`
	// User disabled bit mask (int)
	UserDisabledBitMask *int `pulumi:"userDisabledBitMask"`
	// User enable attribute (string)
	UserEnabledAttribute *string `pulumi:"userEnabledAttribute"`
	// User login attribute. Default `uid` (string)
	UserLoginAttribute *string `pulumi:"userLoginAttribute"`
	// User member attribute. Default `memberOf` (string)
	UserMemberAttribute *string `pulumi:"userMemberAttribute"`
	// User name attribute. Default `givenName` (string)
	UserNameAttribute *string `pulumi:"userNameAttribute"`
	// User object class. Default `inetorgperson` (string)
	UserObjectClass *string `pulumi:"userObjectClass"`
	// User search attribute. Default `uid|sn|givenName` (string)
	UserSearchAttribute *string `pulumi:"userSearchAttribute"`
	// User search base DN (string)
	UserSearchBase string `pulumi:"userSearchBase"`
}

// The set of arguments for constructing a AuthConfigOpenLdap resource.
type AuthConfigOpenLdapArgs struct {
	// Access mode for auth. `required`, `restricted`, `unrestricted` are supported. Default `unrestricted` (string)
	AccessMode pulumi.StringPtrInput
	// Allowed principal ids for auth. Required if `accessMode` is `required` or `restricted`. Ex: `openldap_user://<DN>`  `openldap_group://<DN>` (list)
	AllowedPrincipalIds pulumi.StringArrayInput
	// Annotations of the resource (map)
	Annotations pulumi.MapInput
	// Base64 encoded CA certificate for TLS if self-signed. Use filebase64(<FILE>) for encoding file (string)
	Certificate pulumi.StringPtrInput
	// OpenLdap connection timeout. Default `5000` (int)
	ConnectionTimeout pulumi.IntPtrInput
	// Enable auth config provider. Default `true` (bool)
	Enabled pulumi.BoolPtrInput
	// Group DN attribute. Default `entryDN` (string)
	GroupDnAttribute pulumi.StringPtrInput
	// Group member mapping attribute. Default `member` (string)
	GroupMemberMappingAttribute pulumi.StringPtrInput
	// Group member user attribute. Default `entryDN` (string)
	GroupMemberUserAttribute pulumi.StringPtrInput
	// Group name attribute. Default `cn` (string)
	GroupNameAttribute pulumi.StringPtrInput
	// Group object class. Default `groupOfNames` (string)
	GroupObjectClass pulumi.StringPtrInput
	// Group search attribute. Default `cn` (string)
	GroupSearchAttribute pulumi.StringPtrInput
	// Group search base (string)
	GroupSearchBase pulumi.StringPtrInput
	// Labels of the resource (map)
	Labels pulumi.MapInput
	// Nested group membership enable. Default `false` (bool)
	NestedGroupMembershipEnabled pulumi.BoolPtrInput
	// OpenLdap port. Default `389` (int)
	Port pulumi.IntPtrInput
	// OpenLdap servers list (list)
	Servers pulumi.StringArrayInput
	// Service account DN for access OpenLdap service (string)
	ServiceAccountDistinguishedName pulumi.StringInput
	// Service account password for access OpenLdap service (string)
	ServiceAccountPassword pulumi.StringInput
	// Enable TLS connection (bool)
	Tls pulumi.BoolPtrInput
	// User disabled bit mask (int)
	UserDisabledBitMask pulumi.IntPtrInput
	// User enable attribute (string)
	UserEnabledAttribute pulumi.StringPtrInput
	// User login attribute. Default `uid` (string)
	UserLoginAttribute pulumi.StringPtrInput
	// User member attribute. Default `memberOf` (string)
	UserMemberAttribute pulumi.StringPtrInput
	// User name attribute. Default `givenName` (string)
	UserNameAttribute pulumi.StringPtrInput
	// User object class. Default `inetorgperson` (string)
	UserObjectClass pulumi.StringPtrInput
	// User search attribute. Default `uid|sn|givenName` (string)
	UserSearchAttribute pulumi.StringPtrInput
	// User search base DN (string)
	UserSearchBase pulumi.StringInput
}

func (AuthConfigOpenLdapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authConfigOpenLdapArgs)(nil)).Elem()
}


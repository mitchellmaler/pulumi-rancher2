// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rancher2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Rancher v2 User resource. This can be used to create Users for Rancher v2 environments and retrieve their information.
//
// When a Rancher User is created, it doesn't have a global role binding. At least, `user-base` global role binding in needed in order to enable user login.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rancher2/blob/master/website/docs/r/user.html.markdown.
type User struct {
	pulumi.CustomResourceState

	// Annotations for global role binding (map)
	Annotations pulumi.MapOutput `pulumi:"annotations"`
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Labels for global role binding (map)
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The user full name (string)
	Name pulumi.StringOutput `pulumi:"name"`
	// The user password (string)
	Password pulumi.StringOutput `pulumi:"password"`
	// (Computed) The user principal IDs (list)
	PrincipalIds pulumi.StringArrayOutput `pulumi:"principalIds"`
	// The user username (string)
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("rancher2:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("rancher2:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Annotations for global role binding (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	Enabled *bool `pulumi:"enabled"`
	// Labels for global role binding (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The user full name (string)
	Name *string `pulumi:"name"`
	// The user password (string)
	Password *string `pulumi:"password"`
	// (Computed) The user principal IDs (list)
	PrincipalIds []string `pulumi:"principalIds"`
	// The user username (string)
	Username *string `pulumi:"username"`
}

type UserState struct {
	// Annotations for global role binding (map)
	Annotations pulumi.MapInput
	Enabled pulumi.BoolPtrInput
	// Labels for global role binding (map)
	Labels pulumi.MapInput
	// The user full name (string)
	Name pulumi.StringPtrInput
	// The user password (string)
	Password pulumi.StringPtrInput
	// (Computed) The user principal IDs (list)
	PrincipalIds pulumi.StringArrayInput
	// The user username (string)
	Username pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Annotations for global role binding (map)
	Annotations map[string]interface{} `pulumi:"annotations"`
	Enabled *bool `pulumi:"enabled"`
	// Labels for global role binding (map)
	Labels map[string]interface{} `pulumi:"labels"`
	// The user full name (string)
	Name *string `pulumi:"name"`
	// The user password (string)
	Password string `pulumi:"password"`
	// The user username (string)
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Annotations for global role binding (map)
	Annotations pulumi.MapInput
	Enabled pulumi.BoolPtrInput
	// Labels for global role binding (map)
	Labels pulumi.MapInput
	// The user full name (string)
	Name pulumi.StringPtrInput
	// The user password (string)
	Password pulumi.StringInput
	// The user username (string)
	Username pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}


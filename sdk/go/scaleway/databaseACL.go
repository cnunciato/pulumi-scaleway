// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Database instance autorized IPs.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
//
// ## Examples
//
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := scaleway.NewDatabaseAcl(ctx, "main", &scaleway.DatabaseAclArgs{
// 			InstanceId: pulumi.Any(scaleway_rdb_instance.Main.Id),
// 			AclRules: DatabaseAclAclRuleArray{
// 				&DatabaseAclAclRuleArgs{
// 					Ip:          pulumi.String("1.2.3.4/32"),
// 					Description: pulumi.String("foo"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Database Instance can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//  $ pulumi import scaleway:index/databaseAcl:DatabaseAcl acl01 fr-par/11111111-1111-1111-1111-111111111111
// ```
type DatabaseAcl struct {
	pulumi.CustomResourceState

	// A list of ACLs (structure is described below)
	AclRules DatabaseAclAclRuleArrayOutput `pulumi:"aclRules"`
	// The instance on which to create the ACL.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewDatabaseAcl registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAcl(ctx *pulumi.Context,
	name string, args *DatabaseAclArgs, opts ...pulumi.ResourceOption) (*DatabaseAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclRules == nil {
		return nil, errors.New("invalid value for required argument 'AclRules'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource DatabaseAcl
	err := ctx.RegisterResource("scaleway:index/databaseAcl:DatabaseAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAcl gets an existing DatabaseAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAclState, opts ...pulumi.ResourceOption) (*DatabaseAcl, error) {
	var resource DatabaseAcl
	err := ctx.ReadResource("scaleway:index/databaseAcl:DatabaseAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAcl resources.
type databaseAclState struct {
	// A list of ACLs (structure is described below)
	AclRules []DatabaseAclAclRule `pulumi:"aclRules"`
	// The instance on which to create the ACL.
	InstanceId *string `pulumi:"instanceId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
}

type DatabaseAclState struct {
	// A list of ACLs (structure is described below)
	AclRules DatabaseAclAclRuleArrayInput
	// The instance on which to create the ACL.
	InstanceId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
}

func (DatabaseAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAclState)(nil)).Elem()
}

type databaseAclArgs struct {
	// A list of ACLs (structure is described below)
	AclRules []DatabaseAclAclRule `pulumi:"aclRules"`
	// The instance on which to create the ACL.
	InstanceId string `pulumi:"instanceId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a DatabaseAcl resource.
type DatabaseAclArgs struct {
	// A list of ACLs (structure is described below)
	AclRules DatabaseAclAclRuleArrayInput
	// The instance on which to create the ACL.
	InstanceId pulumi.StringInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
}

func (DatabaseAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAclArgs)(nil)).Elem()
}

type DatabaseAclInput interface {
	pulumi.Input

	ToDatabaseAclOutput() DatabaseAclOutput
	ToDatabaseAclOutputWithContext(ctx context.Context) DatabaseAclOutput
}

func (*DatabaseAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAcl)(nil)).Elem()
}

func (i *DatabaseAcl) ToDatabaseAclOutput() DatabaseAclOutput {
	return i.ToDatabaseAclOutputWithContext(context.Background())
}

func (i *DatabaseAcl) ToDatabaseAclOutputWithContext(ctx context.Context) DatabaseAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAclOutput)
}

type DatabaseAclOutput struct{ *pulumi.OutputState }

func (DatabaseAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAcl)(nil)).Elem()
}

func (o DatabaseAclOutput) ToDatabaseAclOutput() DatabaseAclOutput {
	return o
}

func (o DatabaseAclOutput) ToDatabaseAclOutputWithContext(ctx context.Context) DatabaseAclOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseAclInput)(nil)).Elem(), &DatabaseAcl{})
	pulumi.RegisterOutputType(DatabaseAclOutput{})
}

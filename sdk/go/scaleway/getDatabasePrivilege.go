// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabasePrivilege(ctx *pulumi.Context, args *LookupDatabasePrivilegeArgs, opts ...pulumi.InvokeOption) (*LookupDatabasePrivilegeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupDatabasePrivilegeResult
	err := ctx.Invoke("scaleway:index/getDatabasePrivilege:getDatabasePrivilege", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabasePrivilege.
type LookupDatabasePrivilegeArgs struct {
	DatabaseName string `pulumi:"databaseName"`
	InstanceId   string `pulumi:"instanceId"`
	UserName     string `pulumi:"userName"`
}

// A collection of values returned by getDatabasePrivilege.
type LookupDatabasePrivilegeResult struct {
	DatabaseName string `pulumi:"databaseName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	Permission string `pulumi:"permission"`
	UserName   string `pulumi:"userName"`
}

func LookupDatabasePrivilegeOutput(ctx *pulumi.Context, args LookupDatabasePrivilegeOutputArgs, opts ...pulumi.InvokeOption) LookupDatabasePrivilegeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabasePrivilegeResult, error) {
			args := v.(LookupDatabasePrivilegeArgs)
			r, err := LookupDatabasePrivilege(ctx, &args, opts...)
			var s LookupDatabasePrivilegeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabasePrivilegeResultOutput)
}

// A collection of arguments for invoking getDatabasePrivilege.
type LookupDatabasePrivilegeOutputArgs struct {
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	InstanceId   pulumi.StringInput `pulumi:"instanceId"`
	UserName     pulumi.StringInput `pulumi:"userName"`
}

func (LookupDatabasePrivilegeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasePrivilegeArgs)(nil)).Elem()
}

// A collection of values returned by getDatabasePrivilege.
type LookupDatabasePrivilegeResultOutput struct{ *pulumi.OutputState }

func (LookupDatabasePrivilegeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabasePrivilegeResult)(nil)).Elem()
}

func (o LookupDatabasePrivilegeResultOutput) ToLookupDatabasePrivilegeResultOutput() LookupDatabasePrivilegeResultOutput {
	return o
}

func (o LookupDatabasePrivilegeResultOutput) ToLookupDatabasePrivilegeResultOutputWithContext(ctx context.Context) LookupDatabasePrivilegeResultOutput {
	return o
}

func (o LookupDatabasePrivilegeResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatabasePrivilegeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatabasePrivilegeResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o LookupDatabasePrivilegeResultOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.Permission }).(pulumi.StringOutput)
}

func (o LookupDatabasePrivilegeResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabasePrivilegeResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabasePrivilegeResultOutput{})
}

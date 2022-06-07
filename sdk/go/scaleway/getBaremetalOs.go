// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBaremetalOs(ctx *pulumi.Context, args *GetBaremetalOsArgs, opts ...pulumi.InvokeOption) (*GetBaremetalOsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetBaremetalOsResult
	err := ctx.Invoke("scaleway:index/getBaremetalOs:getBaremetalOs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBaremetalOs.
type GetBaremetalOsArgs struct {
	Name    *string `pulumi:"name"`
	OsId    *string `pulumi:"osId"`
	Version *string `pulumi:"version"`
	Zone    *string `pulumi:"zone"`
}

// A collection of values returned by getBaremetalOs.
type GetBaremetalOsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Name    *string `pulumi:"name"`
	OsId    *string `pulumi:"osId"`
	Version *string `pulumi:"version"`
	Zone    string  `pulumi:"zone"`
}

func GetBaremetalOsOutput(ctx *pulumi.Context, args GetBaremetalOsOutputArgs, opts ...pulumi.InvokeOption) GetBaremetalOsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBaremetalOsResult, error) {
			args := v.(GetBaremetalOsArgs)
			r, err := GetBaremetalOs(ctx, &args, opts...)
			var s GetBaremetalOsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBaremetalOsResultOutput)
}

// A collection of arguments for invoking getBaremetalOs.
type GetBaremetalOsOutputArgs struct {
	Name    pulumi.StringPtrInput `pulumi:"name"`
	OsId    pulumi.StringPtrInput `pulumi:"osId"`
	Version pulumi.StringPtrInput `pulumi:"version"`
	Zone    pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetBaremetalOsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaremetalOsArgs)(nil)).Elem()
}

// A collection of values returned by getBaremetalOs.
type GetBaremetalOsResultOutput struct{ *pulumi.OutputState }

func (GetBaremetalOsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaremetalOsResult)(nil)).Elem()
}

func (o GetBaremetalOsResultOutput) ToGetBaremetalOsResultOutput() GetBaremetalOsResultOutput {
	return o
}

func (o GetBaremetalOsResultOutput) ToGetBaremetalOsResultOutputWithContext(ctx context.Context) GetBaremetalOsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetBaremetalOsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBaremetalOsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaremetalOsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetBaremetalOsResultOutput) OsId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaremetalOsResult) *string { return v.OsId }).(pulumi.StringPtrOutput)
}

func (o GetBaremetalOsResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaremetalOsResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func (o GetBaremetalOsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBaremetalOsResultOutput{})
}

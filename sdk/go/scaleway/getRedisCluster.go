// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Redis cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.LookupRedisCluster(ctx, &GetRedisClusterArgs{
//				ClusterId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRedisCluster(ctx *pulumi.Context, args *LookupRedisClusterArgs, opts ...pulumi.InvokeOption) (*LookupRedisClusterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRedisClusterResult
	err := ctx.Invoke("scaleway:index/getRedisCluster:getRedisCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRedisCluster.
type LookupRedisClusterArgs struct {
	// The Redis cluster ID.
	// Only one of the `name` and `clusterId` should be specified.
	ClusterId *string `pulumi:"clusterId"`
	// The name of the Redis cluster.
	// Only one of the `name` and `clusterId` should be specified.
	Name *string `pulumi:"name"`
	// `region`) The zone in which the server exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getRedisCluster.
type LookupRedisClusterResult struct {
	Acls        []GetRedisClusterAcl `pulumi:"acls"`
	Certificate string               `pulumi:"certificate"`
	ClusterId   *string              `pulumi:"clusterId"`
	ClusterSize int                  `pulumi:"clusterSize"`
	CreatedAt   string               `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                          `pulumi:"id"`
	Name            *string                         `pulumi:"name"`
	NodeType        string                          `pulumi:"nodeType"`
	Password        string                          `pulumi:"password"`
	PrivateNetworks []GetRedisClusterPrivateNetwork `pulumi:"privateNetworks"`
	ProjectId       string                          `pulumi:"projectId"`
	PublicNetworks  []GetRedisClusterPublicNetwork  `pulumi:"publicNetworks"`
	Settings        map[string]string               `pulumi:"settings"`
	Tags            []string                        `pulumi:"tags"`
	TlsEnabled      bool                            `pulumi:"tlsEnabled"`
	UpdatedAt       string                          `pulumi:"updatedAt"`
	UserName        string                          `pulumi:"userName"`
	Version         string                          `pulumi:"version"`
	Zone            *string                         `pulumi:"zone"`
}

func LookupRedisClusterOutput(ctx *pulumi.Context, args LookupRedisClusterOutputArgs, opts ...pulumi.InvokeOption) LookupRedisClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRedisClusterResult, error) {
			args := v.(LookupRedisClusterArgs)
			r, err := LookupRedisCluster(ctx, &args, opts...)
			var s LookupRedisClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRedisClusterResultOutput)
}

// A collection of arguments for invoking getRedisCluster.
type LookupRedisClusterOutputArgs struct {
	// The Redis cluster ID.
	// Only one of the `name` and `clusterId` should be specified.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// The name of the Redis cluster.
	// Only one of the `name` and `clusterId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `region`) The zone in which the server exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupRedisClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisClusterArgs)(nil)).Elem()
}

// A collection of values returned by getRedisCluster.
type LookupRedisClusterResultOutput struct{ *pulumi.OutputState }

func (LookupRedisClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisClusterResult)(nil)).Elem()
}

func (o LookupRedisClusterResultOutput) ToLookupRedisClusterResultOutput() LookupRedisClusterResultOutput {
	return o
}

func (o LookupRedisClusterResultOutput) ToLookupRedisClusterResultOutputWithContext(ctx context.Context) LookupRedisClusterResultOutput {
	return o
}

func (o LookupRedisClusterResultOutput) Acls() GetRedisClusterAclArrayOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) []GetRedisClusterAcl { return v.Acls }).(GetRedisClusterAclArrayOutput)
}

func (o LookupRedisClusterResultOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) string { return v.Certificate }).(pulumi.StringOutput)
}

func (o LookupRedisClusterResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o LookupRedisClusterResultOutput) ClusterSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) int { return v.ClusterSize }).(pulumi.IntOutput)
}

func (o LookupRedisClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRedisClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRedisClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupRedisClusterResultOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) string { return v.NodeType }).(pulumi.StringOutput)
}

func (o LookupRedisClusterResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupRedisClusterResultOutput) PrivateNetworks() GetRedisClusterPrivateNetworkArrayOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) []GetRedisClusterPrivateNetwork { return v.PrivateNetworks }).(GetRedisClusterPrivateNetworkArrayOutput)
}

func (o LookupRedisClusterResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupRedisClusterResultOutput) PublicNetworks() GetRedisClusterPublicNetworkArrayOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) []GetRedisClusterPublicNetwork { return v.PublicNetworks }).(GetRedisClusterPublicNetworkArrayOutput)
}

func (o LookupRedisClusterResultOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) map[string]string { return v.Settings }).(pulumi.StringMapOutput)
}

func (o LookupRedisClusterResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupRedisClusterResultOutput) TlsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) bool { return v.TlsEnabled }).(pulumi.BoolOutput)
}

func (o LookupRedisClusterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupRedisClusterResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) string { return v.UserName }).(pulumi.StringOutput)
}

func (o LookupRedisClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupRedisClusterResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisClusterResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRedisClusterResultOutput{})
}

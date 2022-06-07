// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpcPublicGatewayPatRule struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the PAT rule
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the gateway this PAT rule is applied to
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The private IP used in the PAT rule
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The private port used in the PAT rule
	PrivatePort pulumi.IntOutput `pulumi:"privatePort"`
	// The protocol used in the PAT rule
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The public port used in the PAT rule
	PublicPort pulumi.IntOutput `pulumi:"publicPort"`
	// The date and time of the last update of the PAT rule
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPublicGatewayPatRule registers a new resource with the given unique name, arguments, and options.
func NewVpcPublicGatewayPatRule(ctx *pulumi.Context,
	name string, args *VpcPublicGatewayPatRuleArgs, opts ...pulumi.ResourceOption) (*VpcPublicGatewayPatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.PrivateIp == nil {
		return nil, errors.New("invalid value for required argument 'PrivateIp'")
	}
	if args.PrivatePort == nil {
		return nil, errors.New("invalid value for required argument 'PrivatePort'")
	}
	if args.PublicPort == nil {
		return nil, errors.New("invalid value for required argument 'PublicPort'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VpcPublicGatewayPatRule
	err := ctx.RegisterResource("scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPublicGatewayPatRule gets an existing VpcPublicGatewayPatRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPublicGatewayPatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPublicGatewayPatRuleState, opts ...pulumi.ResourceOption) (*VpcPublicGatewayPatRule, error) {
	var resource VpcPublicGatewayPatRule
	err := ctx.ReadResource("scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPublicGatewayPatRule resources.
type vpcPublicGatewayPatRuleState struct {
	// The date and time of the creation of the PAT rule
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the gateway this PAT rule is applied to
	GatewayId *string `pulumi:"gatewayId"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// The private IP used in the PAT rule
	PrivateIp *string `pulumi:"privateIp"`
	// The private port used in the PAT rule
	PrivatePort *int `pulumi:"privatePort"`
	// The protocol used in the PAT rule
	Protocol *string `pulumi:"protocol"`
	// The public port used in the PAT rule
	PublicPort *int `pulumi:"publicPort"`
	// The date and time of the last update of the PAT rule
	UpdatedAt *string `pulumi:"updatedAt"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type VpcPublicGatewayPatRuleState struct {
	// The date and time of the creation of the PAT rule
	CreatedAt pulumi.StringPtrInput
	// The ID of the gateway this PAT rule is applied to
	GatewayId pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// The private IP used in the PAT rule
	PrivateIp pulumi.StringPtrInput
	// The private port used in the PAT rule
	PrivatePort pulumi.IntPtrInput
	// The protocol used in the PAT rule
	Protocol pulumi.StringPtrInput
	// The public port used in the PAT rule
	PublicPort pulumi.IntPtrInput
	// The date and time of the last update of the PAT rule
	UpdatedAt pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayPatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayPatRuleState)(nil)).Elem()
}

type vpcPublicGatewayPatRuleArgs struct {
	// The ID of the gateway this PAT rule is applied to
	GatewayId string `pulumi:"gatewayId"`
	// The private IP used in the PAT rule
	PrivateIp string `pulumi:"privateIp"`
	// The private port used in the PAT rule
	PrivatePort int `pulumi:"privatePort"`
	// The protocol used in the PAT rule
	Protocol *string `pulumi:"protocol"`
	// The public port used in the PAT rule
	PublicPort int `pulumi:"publicPort"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPublicGatewayPatRule resource.
type VpcPublicGatewayPatRuleArgs struct {
	// The ID of the gateway this PAT rule is applied to
	GatewayId pulumi.StringInput
	// The private IP used in the PAT rule
	PrivateIp pulumi.StringInput
	// The private port used in the PAT rule
	PrivatePort pulumi.IntInput
	// The protocol used in the PAT rule
	Protocol pulumi.StringPtrInput
	// The public port used in the PAT rule
	PublicPort pulumi.IntInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayPatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayPatRuleArgs)(nil)).Elem()
}

type VpcPublicGatewayPatRuleInput interface {
	pulumi.Input

	ToVpcPublicGatewayPatRuleOutput() VpcPublicGatewayPatRuleOutput
	ToVpcPublicGatewayPatRuleOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleOutput
}

func (*VpcPublicGatewayPatRule) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayPatRule)(nil)).Elem()
}

func (i *VpcPublicGatewayPatRule) ToVpcPublicGatewayPatRuleOutput() VpcPublicGatewayPatRuleOutput {
	return i.ToVpcPublicGatewayPatRuleOutputWithContext(context.Background())
}

func (i *VpcPublicGatewayPatRule) ToVpcPublicGatewayPatRuleOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayPatRuleOutput)
}

type VpcPublicGatewayPatRuleOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayPatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayPatRule)(nil)).Elem()
}

func (o VpcPublicGatewayPatRuleOutput) ToVpcPublicGatewayPatRuleOutput() VpcPublicGatewayPatRuleOutput {
	return o
}

func (o VpcPublicGatewayPatRuleOutput) ToVpcPublicGatewayPatRuleOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleOutput {
	return o
}

// The date and time of the creation of the PAT rule
func (o VpcPublicGatewayPatRuleOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the gateway this PAT rule is applied to
func (o VpcPublicGatewayPatRuleOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// The organization_id you want to attach the resource to
func (o VpcPublicGatewayPatRuleOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The private IP used in the PAT rule
func (o VpcPublicGatewayPatRuleOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// The private port used in the PAT rule
func (o VpcPublicGatewayPatRuleOutput) PrivatePort() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.IntOutput { return v.PrivatePort }).(pulumi.IntOutput)
}

// The protocol used in the PAT rule
func (o VpcPublicGatewayPatRuleOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The public port used in the PAT rule
func (o VpcPublicGatewayPatRuleOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.IntOutput { return v.PublicPort }).(pulumi.IntOutput)
}

// The date and time of the last update of the PAT rule
func (o VpcPublicGatewayPatRuleOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The zone you want to attach the resource to
func (o VpcPublicGatewayPatRuleOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayPatRuleInput)(nil)).Elem(), &VpcPublicGatewayPatRule{})
	pulumi.RegisterOutputType(VpcPublicGatewayPatRuleOutput{})
}

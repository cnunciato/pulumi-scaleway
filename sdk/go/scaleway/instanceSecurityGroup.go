// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceSecurityGroup struct {
	pulumi.CustomResourceState

	// The description of the security group
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Enable blocking of SMTP on IPv4 and IPv6
	EnableDefaultSecurity pulumi.BoolPtrOutput `pulumi:"enableDefaultSecurity"`
	ExternalRules         pulumi.BoolPtrOutput `pulumi:"externalRules"`
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy pulumi.StringPtrOutput `pulumi:"inboundDefaultPolicy"`
	// Inbound rules for this security group
	InboundRules InstanceSecurityGroupInboundRuleArrayOutput `pulumi:"inboundRules"`
	// The name of the security group
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy pulumi.StringPtrOutput `pulumi:"outboundDefaultPolicy"`
	// Outbound rules for this security group
	OutboundRules InstanceSecurityGroupOutboundRuleArrayOutput `pulumi:"outboundRules"`
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The stateful value of the security group
	Stateful pulumi.BoolPtrOutput `pulumi:"stateful"`
	// The tags associated with the security group
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceSecurityGroup(ctx *pulumi.Context,
	name string, args *InstanceSecurityGroupArgs, opts ...pulumi.ResourceOption) (*InstanceSecurityGroup, error) {
	if args == nil {
		args = &InstanceSecurityGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceSecurityGroup
	err := ctx.RegisterResource("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceSecurityGroup gets an existing InstanceSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceSecurityGroupState, opts ...pulumi.ResourceOption) (*InstanceSecurityGroup, error) {
	var resource InstanceSecurityGroup
	err := ctx.ReadResource("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceSecurityGroup resources.
type instanceSecurityGroupState struct {
	// The description of the security group
	Description *string `pulumi:"description"`
	// Enable blocking of SMTP on IPv4 and IPv6
	EnableDefaultSecurity *bool `pulumi:"enableDefaultSecurity"`
	ExternalRules         *bool `pulumi:"externalRules"`
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy *string `pulumi:"inboundDefaultPolicy"`
	// Inbound rules for this security group
	InboundRules []InstanceSecurityGroupInboundRule `pulumi:"inboundRules"`
	// The name of the security group
	Name *string `pulumi:"name"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy *string `pulumi:"outboundDefaultPolicy"`
	// Outbound rules for this security group
	OutboundRules []InstanceSecurityGroupOutboundRule `pulumi:"outboundRules"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The stateful value of the security group
	Stateful *bool `pulumi:"stateful"`
	// The tags associated with the security group
	Tags []string `pulumi:"tags"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type InstanceSecurityGroupState struct {
	// The description of the security group
	Description pulumi.StringPtrInput
	// Enable blocking of SMTP on IPv4 and IPv6
	EnableDefaultSecurity pulumi.BoolPtrInput
	ExternalRules         pulumi.BoolPtrInput
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy pulumi.StringPtrInput
	// Inbound rules for this security group
	InboundRules InstanceSecurityGroupInboundRuleArrayInput
	// The name of the security group
	Name pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy pulumi.StringPtrInput
	// Outbound rules for this security group
	OutboundRules InstanceSecurityGroupOutboundRuleArrayInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The stateful value of the security group
	Stateful pulumi.BoolPtrInput
	// The tags associated with the security group
	Tags pulumi.StringArrayInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSecurityGroupState)(nil)).Elem()
}

type instanceSecurityGroupArgs struct {
	// The description of the security group
	Description *string `pulumi:"description"`
	// Enable blocking of SMTP on IPv4 and IPv6
	EnableDefaultSecurity *bool `pulumi:"enableDefaultSecurity"`
	ExternalRules         *bool `pulumi:"externalRules"`
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy *string `pulumi:"inboundDefaultPolicy"`
	// Inbound rules for this security group
	InboundRules []InstanceSecurityGroupInboundRule `pulumi:"inboundRules"`
	// The name of the security group
	Name *string `pulumi:"name"`
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy *string `pulumi:"outboundDefaultPolicy"`
	// Outbound rules for this security group
	OutboundRules []InstanceSecurityGroupOutboundRule `pulumi:"outboundRules"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The stateful value of the security group
	Stateful *bool `pulumi:"stateful"`
	// The tags associated with the security group
	Tags []string `pulumi:"tags"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceSecurityGroup resource.
type InstanceSecurityGroupArgs struct {
	// The description of the security group
	Description pulumi.StringPtrInput
	// Enable blocking of SMTP on IPv4 and IPv6
	EnableDefaultSecurity pulumi.BoolPtrInput
	ExternalRules         pulumi.BoolPtrInput
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy pulumi.StringPtrInput
	// Inbound rules for this security group
	InboundRules InstanceSecurityGroupInboundRuleArrayInput
	// The name of the security group
	Name pulumi.StringPtrInput
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy pulumi.StringPtrInput
	// Outbound rules for this security group
	OutboundRules InstanceSecurityGroupOutboundRuleArrayInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The stateful value of the security group
	Stateful pulumi.BoolPtrInput
	// The tags associated with the security group
	Tags pulumi.StringArrayInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (InstanceSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceSecurityGroupArgs)(nil)).Elem()
}

type InstanceSecurityGroupInput interface {
	pulumi.Input

	ToInstanceSecurityGroupOutput() InstanceSecurityGroupOutput
	ToInstanceSecurityGroupOutputWithContext(ctx context.Context) InstanceSecurityGroupOutput
}

func (*InstanceSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceSecurityGroup)(nil)).Elem()
}

func (i *InstanceSecurityGroup) ToInstanceSecurityGroupOutput() InstanceSecurityGroupOutput {
	return i.ToInstanceSecurityGroupOutputWithContext(context.Background())
}

func (i *InstanceSecurityGroup) ToInstanceSecurityGroupOutputWithContext(ctx context.Context) InstanceSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecurityGroupOutput)
}

type InstanceSecurityGroupOutput struct{ *pulumi.OutputState }

func (InstanceSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceSecurityGroup)(nil)).Elem()
}

func (o InstanceSecurityGroupOutput) ToInstanceSecurityGroupOutput() InstanceSecurityGroupOutput {
	return o
}

func (o InstanceSecurityGroupOutput) ToInstanceSecurityGroupOutputWithContext(ctx context.Context) InstanceSecurityGroupOutput {
	return o
}

// The description of the security group
func (o InstanceSecurityGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Enable blocking of SMTP on IPv4 and IPv6
func (o InstanceSecurityGroupOutput) EnableDefaultSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.BoolPtrOutput { return v.EnableDefaultSecurity }).(pulumi.BoolPtrOutput)
}

func (o InstanceSecurityGroupOutput) ExternalRules() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.BoolPtrOutput { return v.ExternalRules }).(pulumi.BoolPtrOutput)
}

// Default inbound traffic policy for this security group
func (o InstanceSecurityGroupOutput) InboundDefaultPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringPtrOutput { return v.InboundDefaultPolicy }).(pulumi.StringPtrOutput)
}

// Inbound rules for this security group
func (o InstanceSecurityGroupOutput) InboundRules() InstanceSecurityGroupInboundRuleArrayOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) InstanceSecurityGroupInboundRuleArrayOutput { return v.InboundRules }).(InstanceSecurityGroupInboundRuleArrayOutput)
}

// The name of the security group
func (o InstanceSecurityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization_id you want to attach the resource to
func (o InstanceSecurityGroupOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Default outbound traffic policy for this security group
func (o InstanceSecurityGroupOutput) OutboundDefaultPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringPtrOutput { return v.OutboundDefaultPolicy }).(pulumi.StringPtrOutput)
}

// Outbound rules for this security group
func (o InstanceSecurityGroupOutput) OutboundRules() InstanceSecurityGroupOutboundRuleArrayOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) InstanceSecurityGroupOutboundRuleArrayOutput { return v.OutboundRules }).(InstanceSecurityGroupOutboundRuleArrayOutput)
}

// The project_id you want to attach the resource to
func (o InstanceSecurityGroupOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The stateful value of the security group
func (o InstanceSecurityGroupOutput) Stateful() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.BoolPtrOutput { return v.Stateful }).(pulumi.BoolPtrOutput)
}

// The tags associated with the security group
func (o InstanceSecurityGroupOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The zone you want to attach the resource to
func (o InstanceSecurityGroupOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceSecurityGroup) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecurityGroupInput)(nil)).Elem(), &InstanceSecurityGroup{})
	pulumi.RegisterOutputType(InstanceSecurityGroupOutput{})
}

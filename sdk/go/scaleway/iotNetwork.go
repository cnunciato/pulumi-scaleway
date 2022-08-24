// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// IoT Networks can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/iotNetwork:IotNetwork net01 fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type IotNetwork struct {
	pulumi.CustomResourceState

	// The date and time the Network was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The endpoint to use when interacting with the network.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The hub ID to which the Network will be attached to.
	HubId pulumi.StringOutput `pulumi:"hubId"`
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name pulumi.StringOutput `pulumi:"name"`
	// The endpoint key to keep secret.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix pulumi.StringPtrOutput `pulumi:"topicPrefix"`
	// The network type to create (e.g. `sigfox`).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotNetwork registers a new resource with the given unique name, arguments, and options.
func NewIotNetwork(ctx *pulumi.Context,
	name string, args *IotNetworkArgs, opts ...pulumi.ResourceOption) (*IotNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubId == nil {
		return nil, errors.New("invalid value for required argument 'HubId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IotNetwork
	err := ctx.RegisterResource("scaleway:index/iotNetwork:IotNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotNetwork gets an existing IotNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotNetworkState, opts ...pulumi.ResourceOption) (*IotNetwork, error) {
	var resource IotNetwork
	err := ctx.ReadResource("scaleway:index/iotNetwork:IotNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotNetwork resources.
type iotNetworkState struct {
	// The date and time the Network was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The endpoint to use when interacting with the network.
	Endpoint *string `pulumi:"endpoint"`
	// The hub ID to which the Network will be attached to.
	HubId *string `pulumi:"hubId"`
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name *string `pulumi:"name"`
	// The endpoint key to keep secret.
	Secret *string `pulumi:"secret"`
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix *string `pulumi:"topicPrefix"`
	// The network type to create (e.g. `sigfox`).
	Type *string `pulumi:"type"`
}

type IotNetworkState struct {
	// The date and time the Network was created.
	CreatedAt pulumi.StringPtrInput
	// The endpoint to use when interacting with the network.
	Endpoint pulumi.StringPtrInput
	// The hub ID to which the Network will be attached to.
	HubId pulumi.StringPtrInput
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name pulumi.StringPtrInput
	// The endpoint key to keep secret.
	Secret pulumi.StringPtrInput
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix pulumi.StringPtrInput
	// The network type to create (e.g. `sigfox`).
	Type pulumi.StringPtrInput
}

func (IotNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotNetworkState)(nil)).Elem()
}

type iotNetworkArgs struct {
	// The hub ID to which the Network will be attached to.
	HubId string `pulumi:"hubId"`
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name *string `pulumi:"name"`
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix *string `pulumi:"topicPrefix"`
	// The network type to create (e.g. `sigfox`).
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a IotNetwork resource.
type IotNetworkArgs struct {
	// The hub ID to which the Network will be attached to.
	HubId pulumi.StringInput
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name pulumi.StringPtrInput
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix pulumi.StringPtrInput
	// The network type to create (e.g. `sigfox`).
	Type pulumi.StringInput
}

func (IotNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotNetworkArgs)(nil)).Elem()
}

type IotNetworkInput interface {
	pulumi.Input

	ToIotNetworkOutput() IotNetworkOutput
	ToIotNetworkOutputWithContext(ctx context.Context) IotNetworkOutput
}

func (*IotNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**IotNetwork)(nil)).Elem()
}

func (i *IotNetwork) ToIotNetworkOutput() IotNetworkOutput {
	return i.ToIotNetworkOutputWithContext(context.Background())
}

func (i *IotNetwork) ToIotNetworkOutputWithContext(ctx context.Context) IotNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotNetworkOutput)
}

// IotNetworkArrayInput is an input type that accepts IotNetworkArray and IotNetworkArrayOutput values.
// You can construct a concrete instance of `IotNetworkArrayInput` via:
//
//	IotNetworkArray{ IotNetworkArgs{...} }
type IotNetworkArrayInput interface {
	pulumi.Input

	ToIotNetworkArrayOutput() IotNetworkArrayOutput
	ToIotNetworkArrayOutputWithContext(context.Context) IotNetworkArrayOutput
}

type IotNetworkArray []IotNetworkInput

func (IotNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotNetwork)(nil)).Elem()
}

func (i IotNetworkArray) ToIotNetworkArrayOutput() IotNetworkArrayOutput {
	return i.ToIotNetworkArrayOutputWithContext(context.Background())
}

func (i IotNetworkArray) ToIotNetworkArrayOutputWithContext(ctx context.Context) IotNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotNetworkArrayOutput)
}

// IotNetworkMapInput is an input type that accepts IotNetworkMap and IotNetworkMapOutput values.
// You can construct a concrete instance of `IotNetworkMapInput` via:
//
//	IotNetworkMap{ "key": IotNetworkArgs{...} }
type IotNetworkMapInput interface {
	pulumi.Input

	ToIotNetworkMapOutput() IotNetworkMapOutput
	ToIotNetworkMapOutputWithContext(context.Context) IotNetworkMapOutput
}

type IotNetworkMap map[string]IotNetworkInput

func (IotNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotNetwork)(nil)).Elem()
}

func (i IotNetworkMap) ToIotNetworkMapOutput() IotNetworkMapOutput {
	return i.ToIotNetworkMapOutputWithContext(context.Background())
}

func (i IotNetworkMap) ToIotNetworkMapOutputWithContext(ctx context.Context) IotNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotNetworkMapOutput)
}

type IotNetworkOutput struct{ *pulumi.OutputState }

func (IotNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotNetwork)(nil)).Elem()
}

func (o IotNetworkOutput) ToIotNetworkOutput() IotNetworkOutput {
	return o
}

func (o IotNetworkOutput) ToIotNetworkOutputWithContext(ctx context.Context) IotNetworkOutput {
	return o
}

// The date and time the Network was created.
func (o IotNetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IotNetwork) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The endpoint to use when interacting with the network.
func (o IotNetworkOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *IotNetwork) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The hub ID to which the Network will be attached to.
func (o IotNetworkOutput) HubId() pulumi.StringOutput {
	return o.ApplyT(func(v *IotNetwork) pulumi.StringOutput { return v.HubId }).(pulumi.StringOutput)
}

// The name of the IoT Network you want to create (e.g. `my-net`).
func (o IotNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The endpoint key to keep secret.
func (o IotNetworkOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *IotNetwork) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

// The prefix that will be prepended to all topics for this Network.
func (o IotNetworkOutput) TopicPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotNetwork) pulumi.StringPtrOutput { return v.TopicPrefix }).(pulumi.StringPtrOutput)
}

// The network type to create (e.g. `sigfox`).
func (o IotNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IotNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type IotNetworkArrayOutput struct{ *pulumi.OutputState }

func (IotNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotNetwork)(nil)).Elem()
}

func (o IotNetworkArrayOutput) ToIotNetworkArrayOutput() IotNetworkArrayOutput {
	return o
}

func (o IotNetworkArrayOutput) ToIotNetworkArrayOutputWithContext(ctx context.Context) IotNetworkArrayOutput {
	return o
}

func (o IotNetworkArrayOutput) Index(i pulumi.IntInput) IotNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IotNetwork {
		return vs[0].([]*IotNetwork)[vs[1].(int)]
	}).(IotNetworkOutput)
}

type IotNetworkMapOutput struct{ *pulumi.OutputState }

func (IotNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotNetwork)(nil)).Elem()
}

func (o IotNetworkMapOutput) ToIotNetworkMapOutput() IotNetworkMapOutput {
	return o
}

func (o IotNetworkMapOutput) ToIotNetworkMapOutputWithContext(ctx context.Context) IotNetworkMapOutput {
	return o
}

func (o IotNetworkMapOutput) MapIndex(k pulumi.StringInput) IotNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IotNetwork {
		return vs[0].(map[string]*IotNetwork)[vs[1].(string)]
	}).(IotNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IotNetworkInput)(nil)).Elem(), &IotNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotNetworkArrayInput)(nil)).Elem(), IotNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotNetworkMapInput)(nil)).Elem(), IotNetworkMap{})
	pulumi.RegisterOutputType(IotNetworkOutput{})
	pulumi.RegisterOutputType(IotNetworkArrayOutput{})
	pulumi.RegisterOutputType(IotNetworkMapOutput{})
}

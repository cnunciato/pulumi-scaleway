// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Container Registry.
// For more information see [the documentation](https://developers.scaleway.com/en/products/registry/api/).
//
// ## Examples
//
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.NewRegistryNamespace(ctx, "main", &scaleway.RegistryNamespaceArgs{
//				Description: pulumi.String("Main container registry"),
//				IsPublic:    pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Namespaces can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/registryNamespace:RegistryNamespace main fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type RegistryNamespace struct {
	pulumi.CustomResourceState

	// The description of the namespace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Endpoint reachable by Docker.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// The unique name of the namespace.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the namespace is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewRegistryNamespace registers a new resource with the given unique name, arguments, and options.
func NewRegistryNamespace(ctx *pulumi.Context,
	name string, args *RegistryNamespaceArgs, opts ...pulumi.ResourceOption) (*RegistryNamespace, error) {
	if args == nil {
		args = &RegistryNamespaceArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RegistryNamespace
	err := ctx.RegisterResource("scaleway:index/registryNamespace:RegistryNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryNamespace gets an existing RegistryNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryNamespaceState, opts ...pulumi.ResourceOption) (*RegistryNamespace, error) {
	var resource RegistryNamespace
	err := ctx.ReadResource("scaleway:index/registryNamespace:RegistryNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryNamespace resources.
type registryNamespaceState struct {
	// The description of the namespace.
	Description *string `pulumi:"description"`
	// Endpoint reachable by Docker.
	Endpoint *string `pulumi:"endpoint"`
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic *bool `pulumi:"isPublic"`
	// The unique name of the namespace.
	Name *string `pulumi:"name"`
	// The organization ID the namespace is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the namespace should be created.
	Region *string `pulumi:"region"`
}

type RegistryNamespaceState struct {
	// The description of the namespace.
	Description pulumi.StringPtrInput
	// Endpoint reachable by Docker.
	Endpoint pulumi.StringPtrInput
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic pulumi.BoolPtrInput
	// The unique name of the namespace.
	Name pulumi.StringPtrInput
	// The organization ID the namespace is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringPtrInput
}

func (RegistryNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryNamespaceState)(nil)).Elem()
}

type registryNamespaceArgs struct {
	// The description of the namespace.
	Description *string `pulumi:"description"`
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic *bool `pulumi:"isPublic"`
	// The unique name of the namespace.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the namespace should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RegistryNamespace resource.
type RegistryNamespaceArgs struct {
	// The description of the namespace.
	Description pulumi.StringPtrInput
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic pulumi.BoolPtrInput
	// The unique name of the namespace.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringPtrInput
}

func (RegistryNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryNamespaceArgs)(nil)).Elem()
}

type RegistryNamespaceInput interface {
	pulumi.Input

	ToRegistryNamespaceOutput() RegistryNamespaceOutput
	ToRegistryNamespaceOutputWithContext(ctx context.Context) RegistryNamespaceOutput
}

func (*RegistryNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryNamespace)(nil)).Elem()
}

func (i *RegistryNamespace) ToRegistryNamespaceOutput() RegistryNamespaceOutput {
	return i.ToRegistryNamespaceOutputWithContext(context.Background())
}

func (i *RegistryNamespace) ToRegistryNamespaceOutputWithContext(ctx context.Context) RegistryNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryNamespaceOutput)
}

// RegistryNamespaceArrayInput is an input type that accepts RegistryNamespaceArray and RegistryNamespaceArrayOutput values.
// You can construct a concrete instance of `RegistryNamespaceArrayInput` via:
//
//	RegistryNamespaceArray{ RegistryNamespaceArgs{...} }
type RegistryNamespaceArrayInput interface {
	pulumi.Input

	ToRegistryNamespaceArrayOutput() RegistryNamespaceArrayOutput
	ToRegistryNamespaceArrayOutputWithContext(context.Context) RegistryNamespaceArrayOutput
}

type RegistryNamespaceArray []RegistryNamespaceInput

func (RegistryNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryNamespace)(nil)).Elem()
}

func (i RegistryNamespaceArray) ToRegistryNamespaceArrayOutput() RegistryNamespaceArrayOutput {
	return i.ToRegistryNamespaceArrayOutputWithContext(context.Background())
}

func (i RegistryNamespaceArray) ToRegistryNamespaceArrayOutputWithContext(ctx context.Context) RegistryNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryNamespaceArrayOutput)
}

// RegistryNamespaceMapInput is an input type that accepts RegistryNamespaceMap and RegistryNamespaceMapOutput values.
// You can construct a concrete instance of `RegistryNamespaceMapInput` via:
//
//	RegistryNamespaceMap{ "key": RegistryNamespaceArgs{...} }
type RegistryNamespaceMapInput interface {
	pulumi.Input

	ToRegistryNamespaceMapOutput() RegistryNamespaceMapOutput
	ToRegistryNamespaceMapOutputWithContext(context.Context) RegistryNamespaceMapOutput
}

type RegistryNamespaceMap map[string]RegistryNamespaceInput

func (RegistryNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryNamespace)(nil)).Elem()
}

func (i RegistryNamespaceMap) ToRegistryNamespaceMapOutput() RegistryNamespaceMapOutput {
	return i.ToRegistryNamespaceMapOutputWithContext(context.Background())
}

func (i RegistryNamespaceMap) ToRegistryNamespaceMapOutputWithContext(ctx context.Context) RegistryNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryNamespaceMapOutput)
}

type RegistryNamespaceOutput struct{ *pulumi.OutputState }

func (RegistryNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryNamespace)(nil)).Elem()
}

func (o RegistryNamespaceOutput) ToRegistryNamespaceOutput() RegistryNamespaceOutput {
	return o
}

func (o RegistryNamespaceOutput) ToRegistryNamespaceOutputWithContext(ctx context.Context) RegistryNamespaceOutput {
	return o
}

// The description of the namespace.
func (o RegistryNamespaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryNamespace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Endpoint reachable by Docker.
func (o RegistryNamespaceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryNamespace) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Whether the images stored in the namespace should be downloadable publicly (docker pull).
func (o RegistryNamespaceOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RegistryNamespace) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

// The unique name of the namespace.
func (o RegistryNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the namespace is associated with.
func (o RegistryNamespaceOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryNamespace) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the namespace is associated with.
func (o RegistryNamespaceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryNamespace) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`). The region in which the namespace should be created.
func (o RegistryNamespaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryNamespace) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type RegistryNamespaceArrayOutput struct{ *pulumi.OutputState }

func (RegistryNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryNamespace)(nil)).Elem()
}

func (o RegistryNamespaceArrayOutput) ToRegistryNamespaceArrayOutput() RegistryNamespaceArrayOutput {
	return o
}

func (o RegistryNamespaceArrayOutput) ToRegistryNamespaceArrayOutputWithContext(ctx context.Context) RegistryNamespaceArrayOutput {
	return o
}

func (o RegistryNamespaceArrayOutput) Index(i pulumi.IntInput) RegistryNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegistryNamespace {
		return vs[0].([]*RegistryNamespace)[vs[1].(int)]
	}).(RegistryNamespaceOutput)
}

type RegistryNamespaceMapOutput struct{ *pulumi.OutputState }

func (RegistryNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryNamespace)(nil)).Elem()
}

func (o RegistryNamespaceMapOutput) ToRegistryNamespaceMapOutput() RegistryNamespaceMapOutput {
	return o
}

func (o RegistryNamespaceMapOutput) ToRegistryNamespaceMapOutputWithContext(ctx context.Context) RegistryNamespaceMapOutput {
	return o
}

func (o RegistryNamespaceMapOutput) MapIndex(k pulumi.StringInput) RegistryNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegistryNamespace {
		return vs[0].(map[string]*RegistryNamespace)[vs[1].(string)]
	}).(RegistryNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryNamespaceInput)(nil)).Elem(), &RegistryNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryNamespaceArrayInput)(nil)).Elem(), RegistryNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryNamespaceMapInput)(nil)).Elem(), RegistryNamespaceMap{})
	pulumi.RegisterOutputType(RegistryNamespaceOutput{})
	pulumi.RegisterOutputType(RegistryNamespaceArrayOutput{})
	pulumi.RegisterOutputType(RegistryNamespaceMapOutput{})
}

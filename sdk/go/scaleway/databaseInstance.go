// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseInstance struct {
	pulumi.CustomResourceState

	// Boolean to store logical backups in the same region as the database instance
	BackupSameRegion pulumi.BoolOutput `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours
	BackupScheduleFrequency pulumi.IntOutput `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days
	BackupScheduleRetention pulumi.IntOutput `pulumi:"backupScheduleRetention"`
	// Certificate of the database instance
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Disable automated backup for the database instance
	DisableBackup pulumi.BoolPtrOutput `pulumi:"disableBackup"`
	// Endpoint IP of the database instance
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp pulumi.StringOutput `pulumi:"endpointIp"`
	// Endpoint port of the database instance
	EndpointPort pulumi.IntOutput `pulumi:"endpointPort"`
	// Database's engine version id
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Enable or disable high availability for the database instance
	IsHaCluster pulumi.BoolPtrOutput `pulumi:"isHaCluster"`
	// Load balancer of the database instance
	LoadBalancers DatabaseInstanceLoadBalancerArrayOutput `pulumi:"loadBalancers"`
	// Name of the database instance
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of database instance you want to create
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Password for the first user of the database instance
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// List of private network to expose your database instance
	PrivateNetwork DatabaseInstancePrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Read replicas of the database instance
	ReadReplicas DatabaseInstanceReadReplicaArrayOutput `pulumi:"readReplicas"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// Map of engine settings to be set.
	Settings pulumi.StringMapOutput `pulumi:"settings"`
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Identifier for the first user of the database instance
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
	// Volume size (in GB) when volume_type is not lssd
	VolumeSizeInGb pulumi.IntOutput `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored
	VolumeType pulumi.StringPtrOutput `pulumi:"volumeType"`
}

// NewDatabaseInstance registers a new resource with the given unique name, arguments, and options.
func NewDatabaseInstance(ctx *pulumi.Context,
	name string, args *DatabaseInstanceArgs, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DatabaseInstance
	err := ctx.RegisterResource("scaleway:index/databaseInstance:DatabaseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseInstance gets an existing DatabaseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseInstanceState, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	var resource DatabaseInstance
	err := ctx.ReadResource("scaleway:index/databaseInstance:DatabaseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseInstance resources.
type databaseInstanceState struct {
	// Boolean to store logical backups in the same region as the database instance
	BackupSameRegion *bool `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours
	BackupScheduleFrequency *int `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days
	BackupScheduleRetention *int `pulumi:"backupScheduleRetention"`
	// Certificate of the database instance
	Certificate *string `pulumi:"certificate"`
	// Disable automated backup for the database instance
	DisableBackup *bool `pulumi:"disableBackup"`
	// Endpoint IP of the database instance
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp *string `pulumi:"endpointIp"`
	// Endpoint port of the database instance
	EndpointPort *int `pulumi:"endpointPort"`
	// Database's engine version id
	Engine *string `pulumi:"engine"`
	// Enable or disable high availability for the database instance
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// Load balancer of the database instance
	LoadBalancers []DatabaseInstanceLoadBalancer `pulumi:"loadBalancers"`
	// Name of the database instance
	Name *string `pulumi:"name"`
	// The type of database instance you want to create
	NodeType *string `pulumi:"nodeType"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// Password for the first user of the database instance
	Password *string `pulumi:"password"`
	// List of private network to expose your database instance
	PrivateNetwork *DatabaseInstancePrivateNetwork `pulumi:"privateNetwork"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// Read replicas of the database instance
	ReadReplicas []DatabaseInstanceReadReplica `pulumi:"readReplicas"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// Map of engine settings to be set.
	Settings map[string]string `pulumi:"settings"`
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance
	UserName *string `pulumi:"userName"`
	// Volume size (in GB) when volume_type is not lssd
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored
	VolumeType *string `pulumi:"volumeType"`
}

type DatabaseInstanceState struct {
	// Boolean to store logical backups in the same region as the database instance
	BackupSameRegion pulumi.BoolPtrInput
	// Backup schedule frequency in hours
	BackupScheduleFrequency pulumi.IntPtrInput
	// Backup schedule retention in days
	BackupScheduleRetention pulumi.IntPtrInput
	// Certificate of the database instance
	Certificate pulumi.StringPtrInput
	// Disable automated backup for the database instance
	DisableBackup pulumi.BoolPtrInput
	// Endpoint IP of the database instance
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp pulumi.StringPtrInput
	// Endpoint port of the database instance
	EndpointPort pulumi.IntPtrInput
	// Database's engine version id
	Engine pulumi.StringPtrInput
	// Enable or disable high availability for the database instance
	IsHaCluster pulumi.BoolPtrInput
	// Load balancer of the database instance
	LoadBalancers DatabaseInstanceLoadBalancerArrayInput
	// Name of the database instance
	Name pulumi.StringPtrInput
	// The type of database instance you want to create
	NodeType pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// Password for the first user of the database instance
	Password pulumi.StringPtrInput
	// List of private network to expose your database instance
	PrivateNetwork DatabaseInstancePrivateNetworkPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// Read replicas of the database instance
	ReadReplicas DatabaseInstanceReadReplicaArrayInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// Map of engine settings to be set.
	Settings pulumi.StringMapInput
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance
	UserName pulumi.StringPtrInput
	// Volume size (in GB) when volume_type is not lssd
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored
	VolumeType pulumi.StringPtrInput
}

func (DatabaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceState)(nil)).Elem()
}

type databaseInstanceArgs struct {
	// Boolean to store logical backups in the same region as the database instance
	BackupSameRegion *bool `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours
	BackupScheduleFrequency *int `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days
	BackupScheduleRetention *int `pulumi:"backupScheduleRetention"`
	// Disable automated backup for the database instance
	DisableBackup *bool `pulumi:"disableBackup"`
	// Database's engine version id
	Engine string `pulumi:"engine"`
	// Enable or disable high availability for the database instance
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// Name of the database instance
	Name *string `pulumi:"name"`
	// The type of database instance you want to create
	NodeType string `pulumi:"nodeType"`
	// Password for the first user of the database instance
	Password *string `pulumi:"password"`
	// List of private network to expose your database instance
	PrivateNetwork *DatabaseInstancePrivateNetwork `pulumi:"privateNetwork"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// Map of engine settings to be set.
	Settings map[string]string `pulumi:"settings"`
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance
	UserName *string `pulumi:"userName"`
	// Volume size (in GB) when volume_type is not lssd
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored
	VolumeType *string `pulumi:"volumeType"`
}

// The set of arguments for constructing a DatabaseInstance resource.
type DatabaseInstanceArgs struct {
	// Boolean to store logical backups in the same region as the database instance
	BackupSameRegion pulumi.BoolPtrInput
	// Backup schedule frequency in hours
	BackupScheduleFrequency pulumi.IntPtrInput
	// Backup schedule retention in days
	BackupScheduleRetention pulumi.IntPtrInput
	// Disable automated backup for the database instance
	DisableBackup pulumi.BoolPtrInput
	// Database's engine version id
	Engine pulumi.StringInput
	// Enable or disable high availability for the database instance
	IsHaCluster pulumi.BoolPtrInput
	// Name of the database instance
	Name pulumi.StringPtrInput
	// The type of database instance you want to create
	NodeType pulumi.StringInput
	// Password for the first user of the database instance
	Password pulumi.StringPtrInput
	// List of private network to expose your database instance
	PrivateNetwork DatabaseInstancePrivateNetworkPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// Map of engine settings to be set.
	Settings pulumi.StringMapInput
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance
	UserName pulumi.StringPtrInput
	// Volume size (in GB) when volume_type is not lssd
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored
	VolumeType pulumi.StringPtrInput
}

func (DatabaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceArgs)(nil)).Elem()
}

type DatabaseInstanceInput interface {
	pulumi.Input

	ToDatabaseInstanceOutput() DatabaseInstanceOutput
	ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput
}

func (*DatabaseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil)).Elem()
}

func (i *DatabaseInstance) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return i.ToDatabaseInstanceOutputWithContext(context.Background())
}

func (i *DatabaseInstance) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceOutput)
}

type DatabaseInstanceOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return o
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return o
}

// Boolean to store logical backups in the same region as the database instance
func (o DatabaseInstanceOutput) BackupSameRegion() pulumi.BoolOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.BoolOutput { return v.BackupSameRegion }).(pulumi.BoolOutput)
}

// Backup schedule frequency in hours
func (o DatabaseInstanceOutput) BackupScheduleFrequency() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.IntOutput { return v.BackupScheduleFrequency }).(pulumi.IntOutput)
}

// Backup schedule retention in days
func (o DatabaseInstanceOutput) BackupScheduleRetention() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.IntOutput { return v.BackupScheduleRetention }).(pulumi.IntOutput)
}

// Certificate of the database instance
func (o DatabaseInstanceOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Disable automated backup for the database instance
func (o DatabaseInstanceOutput) DisableBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.BoolPtrOutput { return v.DisableBackup }).(pulumi.BoolPtrOutput)
}

// Endpoint IP of the database instance
//
// Deprecated: Please use the private_network or the load_balancer attribute
func (o DatabaseInstanceOutput) EndpointIp() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.EndpointIp }).(pulumi.StringOutput)
}

// Endpoint port of the database instance
func (o DatabaseInstanceOutput) EndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.IntOutput { return v.EndpointPort }).(pulumi.IntOutput)
}

// Database's engine version id
func (o DatabaseInstanceOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Enable or disable high availability for the database instance
func (o DatabaseInstanceOutput) IsHaCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.BoolPtrOutput { return v.IsHaCluster }).(pulumi.BoolPtrOutput)
}

// Load balancer of the database instance
func (o DatabaseInstanceOutput) LoadBalancers() DatabaseInstanceLoadBalancerArrayOutput {
	return o.ApplyT(func(v *DatabaseInstance) DatabaseInstanceLoadBalancerArrayOutput { return v.LoadBalancers }).(DatabaseInstanceLoadBalancerArrayOutput)
}

// Name of the database instance
func (o DatabaseInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of database instance you want to create
func (o DatabaseInstanceOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// The organization_id you want to attach the resource to
func (o DatabaseInstanceOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Password for the first user of the database instance
func (o DatabaseInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// List of private network to expose your database instance
func (o DatabaseInstanceOutput) PrivateNetwork() DatabaseInstancePrivateNetworkPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) DatabaseInstancePrivateNetworkPtrOutput { return v.PrivateNetwork }).(DatabaseInstancePrivateNetworkPtrOutput)
}

// The project_id you want to attach the resource to
func (o DatabaseInstanceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Read replicas of the database instance
func (o DatabaseInstanceOutput) ReadReplicas() DatabaseInstanceReadReplicaArrayOutput {
	return o.ApplyT(func(v *DatabaseInstance) DatabaseInstanceReadReplicaArrayOutput { return v.ReadReplicas }).(DatabaseInstanceReadReplicaArrayOutput)
}

// The region you want to attach the resource to
func (o DatabaseInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Map of engine settings to be set.
func (o DatabaseInstanceOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringMapOutput { return v.Settings }).(pulumi.StringMapOutput)
}

// List of tags ["tag1", "tag2", ...] attached to a database instance
func (o DatabaseInstanceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Identifier for the first user of the database instance
func (o DatabaseInstanceOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

// Volume size (in GB) when volume_type is not lssd
func (o DatabaseInstanceOutput) VolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.IntOutput { return v.VolumeSizeInGb }).(pulumi.IntOutput)
}

// Type of volume where data are stored
func (o DatabaseInstanceOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceInput)(nil)).Elem(), &DatabaseInstance{})
	pulumi.RegisterOutputType(DatabaseInstanceOutput{})
}

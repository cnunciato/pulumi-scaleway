// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Compute Instance Volumes.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39).
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const serverVolume = new scaleway.InstanceVolume("server_volume", {
 *     sizeInGb: 20,
 *     type: "l_ssd",
 * });
 * ```
 *
 * ## Import
 *
 * volumes can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/instanceVolume:InstanceVolume server_volume fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class InstanceVolume extends pulumi.CustomResource {
    /**
     * Get an existing InstanceVolume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceVolumeState, opts?: pulumi.CustomResourceOptions): InstanceVolume {
        return new InstanceVolume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/instanceVolume:InstanceVolume';

    /**
     * Returns true if the given object is an instance of InstanceVolume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceVolume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceVolume.__pulumiType;
    }

    /**
     * Create a volume based on a image
     */
    public readonly fromSnapshotId!: pulumi.Output<string | undefined>;
    /**
     * If set, the new volume will be copied from this volume. Only one of `sizeInGb`, `fromVolumeId` and `fromSnapshotId` should be specified.
     */
    public readonly fromVolumeId!: pulumi.Output<string | undefined>;
    /**
     * The name of the volume. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the volume is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the volume is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The id of the associated server.
     */
    public /*out*/ readonly serverId!: pulumi.Output<string>;
    /**
     * The size of the volume. Only one of `sizeInGb`, `fromVolumeId` and `fromVolumeId` should be specified.
     */
    public readonly sizeInGb!: pulumi.Output<number | undefined>;
    /**
     * A list of tags to apply to the volume.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The type of the volume. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD).
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the volume should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceVolume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceVolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceVolumeArgs | InstanceVolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceVolumeState | undefined;
            resourceInputs["fromSnapshotId"] = state ? state.fromSnapshotId : undefined;
            resourceInputs["fromVolumeId"] = state ? state.fromVolumeId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["serverId"] = state ? state.serverId : undefined;
            resourceInputs["sizeInGb"] = state ? state.sizeInGb : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceVolumeArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["fromSnapshotId"] = args ? args.fromSnapshotId : undefined;
            resourceInputs["fromVolumeId"] = args ? args.fromVolumeId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["sizeInGb"] = args ? args.sizeInGb : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["serverId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceVolume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceVolume resources.
 */
export interface InstanceVolumeState {
    /**
     * Create a volume based on a image
     */
    fromSnapshotId?: pulumi.Input<string>;
    /**
     * If set, the new volume will be copied from this volume. Only one of `sizeInGb`, `fromVolumeId` and `fromSnapshotId` should be specified.
     */
    fromVolumeId?: pulumi.Input<string>;
    /**
     * The name of the volume. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the volume is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the volume is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The id of the associated server.
     */
    serverId?: pulumi.Input<string>;
    /**
     * The size of the volume. Only one of `sizeInGb`, `fromVolumeId` and `fromVolumeId` should be specified.
     */
    sizeInGb?: pulumi.Input<number>;
    /**
     * A list of tags to apply to the volume.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the volume. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD).
     */
    type?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the volume should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceVolume resource.
 */
export interface InstanceVolumeArgs {
    /**
     * Create a volume based on a image
     */
    fromSnapshotId?: pulumi.Input<string>;
    /**
     * If set, the new volume will be copied from this volume. Only one of `sizeInGb`, `fromVolumeId` and `fromSnapshotId` should be specified.
     */
    fromVolumeId?: pulumi.Input<string>;
    /**
     * The name of the volume. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the volume is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The size of the volume. Only one of `sizeInGb`, `fromVolumeId` and `fromVolumeId` should be specified.
     */
    sizeInGb?: pulumi.Input<number>;
    /**
     * A list of tags to apply to the volume.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the volume. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD).
     */
    type: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the volume should be created.
     */
    zone?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Redis Clusters.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.RedisCluster("main", {
 *     acls: [{
 *         description: "Allow all",
 *         ip: "0.0.0.0/0",
 *     }],
 *     clusterSize: 1,
 *     nodeType: "RED1-MICRO",
 *     password: "thiZ_is_v&ry_s3cret",
 *     tags: [
 *         "test",
 *         "redis",
 *     ],
 *     tlsEnabled: true,
 *     userName: "my_initial_user",
 *     version: "6.2.6",
 * });
 * ```
 *
 * ### With settings
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.RedisCluster("main", {
 *     nodeType: "RED1-MICRO",
 *     password: "thiZ_is_v&ry_s3cret",
 *     settings: {
 *         maxclients: "1000",
 *         "tcp-keepalive": "120",
 *     },
 *     userName: "my_initial_user",
 *     version: "6.2.6",
 * });
 * ```
 *
 * ### With a private network
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const pn = new scaleway.VpcPrivateNetwork("pn", {});
 * const main = new scaleway.RedisCluster("main", {
 *     version: "6.2.6",
 *     nodeType: "RED1-MICRO",
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     clusterSize: 1,
 *     privateNetworks: [{
 *         id: pn.id,
 *         serviceIps: ["10.12.1.1/20"],
 *     }],
 * }, {
 *     dependsOn: [pn],
 * });
 * ```
 *
 * ## Import
 *
 * Redis Cluster can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/redisCluster:RedisCluster redis01 fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class RedisCluster extends pulumi.CustomResource {
    /**
     * Get an existing RedisCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RedisClusterState, opts?: pulumi.CustomResourceOptions): RedisCluster {
        return new RedisCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/redisCluster:RedisCluster';

    /**
     * Returns true if the given object is an instance of RedisCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RedisCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RedisCluster.__pulumiType;
    }

    /**
     * List of acl rules, this is cluster's authorized IPs.
     */
    public readonly acls!: pulumi.Output<outputs.RedisClusterAcl[] | undefined>;
    /**
     * The PEM of the certificate used by redis, only when `tlsEnabled` is true
     */
    public /*out*/ readonly certificate!: pulumi.Output<string>;
    /**
     * The number of nodes in the Redis Cluster.
     */
    public readonly clusterSize!: pulumi.Output<number>;
    /**
     * The date and time of creation of the Redis Cluster.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The name of the Redis Cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of Redis Cluster you want to create (e.g. `RED1-M`).
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * Password for the first user of the Redis Cluster.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Describes the private network you want to connect to your cluster. If not set, a public network will be provided.
     */
    public readonly privateNetworks!: pulumi.Output<outputs.RedisClusterPrivateNetwork[] | undefined>;
    /**
     * `projectId`) The ID of the project the Redis Cluster is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Public network specs details
     */
    public readonly publicNetwork!: pulumi.Output<outputs.RedisClusterPublicNetwork>;
    /**
     * Map of settings for redis cluster. Available settings can be found by listing redis versions with scaleway API or CLI
     */
    public readonly settings!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The tags associated with the Redis Cluster.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Whether TLS is enabled or not.
     */
    public readonly tlsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The date and time of the last update of the Redis Cluster.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Identifier for the first user of the Redis Cluster.
     */
    public readonly userName!: pulumi.Output<string>;
    /**
     * Redis's Cluster version (e.g. `6.2.6`).
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the Redis Cluster should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a RedisCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RedisClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RedisClusterArgs | RedisClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RedisClusterState | undefined;
            resourceInputs["acls"] = state ? state.acls : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["clusterSize"] = state ? state.clusterSize : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["privateNetworks"] = state ? state.privateNetworks : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["publicNetwork"] = state ? state.publicNetwork : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tlsEnabled"] = state ? state.tlsEnabled : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as RedisClusterArgs | undefined;
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["acls"] = args ? args.acls : undefined;
            resourceInputs["clusterSize"] = args ? args.clusterSize : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["privateNetworks"] = args ? args.privateNetworks : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["publicNetwork"] = args ? args.publicNetwork : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tlsEnabled"] = args ? args.tlsEnabled : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["certificate"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RedisCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RedisCluster resources.
 */
export interface RedisClusterState {
    /**
     * List of acl rules, this is cluster's authorized IPs.
     */
    acls?: pulumi.Input<pulumi.Input<inputs.RedisClusterAcl>[]>;
    /**
     * The PEM of the certificate used by redis, only when `tlsEnabled` is true
     */
    certificate?: pulumi.Input<string>;
    /**
     * The number of nodes in the Redis Cluster.
     */
    clusterSize?: pulumi.Input<number>;
    /**
     * The date and time of creation of the Redis Cluster.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The name of the Redis Cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of Redis Cluster you want to create (e.g. `RED1-M`).
     */
    nodeType?: pulumi.Input<string>;
    /**
     * Password for the first user of the Redis Cluster.
     */
    password?: pulumi.Input<string>;
    /**
     * Describes the private network you want to connect to your cluster. If not set, a public network will be provided.
     */
    privateNetworks?: pulumi.Input<pulumi.Input<inputs.RedisClusterPrivateNetwork>[]>;
    /**
     * `projectId`) The ID of the project the Redis Cluster is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Public network specs details
     */
    publicNetwork?: pulumi.Input<inputs.RedisClusterPublicNetwork>;
    /**
     * Map of settings for redis cluster. Available settings can be found by listing redis versions with scaleway API or CLI
     */
    settings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The tags associated with the Redis Cluster.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether TLS is enabled or not.
     */
    tlsEnabled?: pulumi.Input<boolean>;
    /**
     * The date and time of the last update of the Redis Cluster.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Identifier for the first user of the Redis Cluster.
     */
    userName?: pulumi.Input<string>;
    /**
     * Redis's Cluster version (e.g. `6.2.6`).
     */
    version?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the Redis Cluster should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RedisCluster resource.
 */
export interface RedisClusterArgs {
    /**
     * List of acl rules, this is cluster's authorized IPs.
     */
    acls?: pulumi.Input<pulumi.Input<inputs.RedisClusterAcl>[]>;
    /**
     * The number of nodes in the Redis Cluster.
     */
    clusterSize?: pulumi.Input<number>;
    /**
     * The name of the Redis Cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of Redis Cluster you want to create (e.g. `RED1-M`).
     */
    nodeType: pulumi.Input<string>;
    /**
     * Password for the first user of the Redis Cluster.
     */
    password: pulumi.Input<string>;
    /**
     * Describes the private network you want to connect to your cluster. If not set, a public network will be provided.
     */
    privateNetworks?: pulumi.Input<pulumi.Input<inputs.RedisClusterPrivateNetwork>[]>;
    /**
     * `projectId`) The ID of the project the Redis Cluster is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Public network specs details
     */
    publicNetwork?: pulumi.Input<inputs.RedisClusterPublicNetwork>;
    /**
     * Map of settings for redis cluster. Available settings can be found by listing redis versions with scaleway API or CLI
     */
    settings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The tags associated with the Redis Cluster.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether TLS is enabled or not.
     */
    tlsEnabled?: pulumi.Input<boolean>;
    /**
     * Identifier for the first user of the Redis Cluster.
     */
    userName: pulumi.Input<string>;
    /**
     * Redis's Cluster version (e.g. `6.2.6`).
     */
    version: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the Redis Cluster should be created.
     */
    zone?: pulumi.Input<string>;
}

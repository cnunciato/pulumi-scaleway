// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway object storage buckets.
 * For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const someBucket = new scaleway.ObjectBucket("some_bucket", {
 *     acl: "private",
 *     tags: {
 *         key: "value",
 *     },
 * });
 * ```
 * ### Using object lifecycle
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.ObjectBucket("main", {
 *     acl: "private",
 *     lifecycleRules: [
 *         // This lifecycle configuration rule will make that all objects that got a filter key that start with (path1/) be transferred
 *         // from their default storage class (STANDARD, ONEZONE_IA) to GLACIER after 120 days counting 
 *         // from their creation and then 365 days after that they will be expired and deleted.
 *         {
 *             enabled: true,
 *             expiration: {
 *                 days: 365,
 *             },
 *             id: "id1",
 *             prefix: "path1/",
 *             transitions: [{
 *                 days: 120,
 *                 storageClass: "GLACIER",
 *             }],
 *         },
 *         // This lifecycle configuration rule specifies that all objects (identified by the key name prefix (path2/) in the rule)
 *         // from their creation and then 50 days after that they will be expired and deleted.
 *         {
 *             enabled: true,
 *             expiration: {
 *                 days: 50,
 *             },
 *             id: "id2",
 *             prefix: "path2/",
 *         },
 *         // This lifecycle configuration rule remove any object with (path3/) prefix that match
 *         // with the tags one day after creation.
 *         {
 *             enabled: false,
 *             expiration: {
 *                 days: 1,
 *             },
 *             id: "id3",
 *             prefix: "path3/",
 *             tags: {
 *                 tagKey: "tagValue",
 *                 terraform: "hashicorp",
 *             },
 *         },
 *         // This lifecycle configuration rule specifies a tag-based filter (tag1/value1).
 *         // This rule directs Scaleway S3 to transition objects S3 Glacier class soon after creation.
 *         // It is also disable temporaly.
 *         {
 *             enabled: true,
 *             id: "id4",
 *             tags: {
 *                 tag1: "value1",
 *             },
 *             transitions: [{
 *                 days: 0,
 *                 storageClass: "GLACIER",
 *             }],
 *         },
 *         // This lifecycle configuration rule specifies with the AbortIncompleteMultipartUpload action to 
 *         // stop incomplete multipart uploads (identified by the key name prefix (path5/) in the rule)
 *         // if they aren't completed within a specified number of days after initiation.
 *         // Note: It's not recommended using prefix/ for AbortIncompleteMultipartUpload as any incomplete multipart upload will be billed
 *         {
 *             abortIncompleteMultipartUploadDays: 30,
 *             //  prefix  = "path5/"
 *             enabled: true,
 *         },
 *     ],
 *     region: "fr-par",
 * });
 * ```
 * ## The ACL
 *
 * Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl)
 *
 * The `CORS` object supports the following:
 *
 * * `allowedHeaders` (Optional) Specifies which headers are allowed.
 * * `allowedMethods` (Required) Specifies which methods are allowed. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
 * * `allowedOrigins` (Required) Specifies which origins are allowed.
 * * `exposeHeaders` (Optional) Specifies expose header in the response.
 * * `maxAgeSeconds` (Optional) Specifies time in seconds that browser can cache the response for a preflight request.
 *
 * The `lifecycleRule` (Optional) object supports the following:
 *
 * * `id` - (Optional) Unique identifier for the rule. Must be less than or equal to 255 characters in length.
 * * `prefix` - (Optional) Object key prefix identifying one or more objects to which the rule applies.
 * * `tags` - (Optional) Specifies object tags key and value.
 * * `enabled` - (Required) The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
 *
 * * `abortIncompleteMultipartUploadDays` (Optional) Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
 *   
 *     * > **Important:** It's not recommended using `prefix` for `AbortIncompleteMultipartUpload` as any incomplete multipart upload will be billed
 *
 * * `expiration` - (Optional) Specifies a period in the object's expire (documented below).
 * * `transition` - (Optional) Specifies a period in the object's transitions (documented below).
 *
 * At least one of `abortIncompleteMultipartUploadDays`, `expiration`, `transition` must be specified.
 *
 * The `expiration` object supports the following
 *
 * * `days` (Optional) Specifies the number of days after object creation when the specific rule action takes effect.
 *
 * > **Important:**  If versioning is enabled, this rule only deletes the current version of an object.
 *
 * The `transition` object supports the following
 *
 * * `days` (Optional) Specifies the number of days after object creation when the specific rule action takes effect.
 * * `storageClass` (Required) Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA`  to which you want the object to transition.
 *
 * > **Important:**  `ONEZONE_IA` is only available in `fr-par` region. The storage class `GLACIER` is not available in `pl-waw` region.
 *
 * The `versioning` object supports the following:
 *
 * * `enabled` - (Optional) Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
 *
 * ## Import
 *
 * Buckets can be imported using the `{region}/{bucketName}` identifier, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/objectBucket:ObjectBucket some_bucket fr-par/some-bucket
 * ```
 */
export class ObjectBucket extends pulumi.CustomResource {
    /**
     * Get an existing ObjectBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectBucketState, opts?: pulumi.CustomResourceOptions): ObjectBucket {
        return new ObjectBucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/objectBucket:ObjectBucket';

    /**
     * Returns true if the given object is an instance of ObjectBucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectBucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectBucket.__pulumiType;
    }

    /**
     * The canned ACL you want to apply to the bucket.
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    public readonly corsRules!: pulumi.Output<outputs.ObjectBucketCorsRule[] | undefined>;
    /**
     * The endpoint URL of the bucket
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
     */
    public readonly lifecycleRules!: pulumi.Output<outputs.ObjectBucketLifecycleRule[] | undefined>;
    /**
     * The name of the bucket.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A list of tags (key / value) for the bucket.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    public readonly versioning!: pulumi.Output<outputs.ObjectBucketVersioning>;

    /**
     * Create a ObjectBucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ObjectBucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectBucketArgs | ObjectBucketState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectBucketState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["corsRules"] = state ? state.corsRules : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["lifecycleRules"] = state ? state.lifecycleRules : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["versioning"] = state ? state.versioning : undefined;
        } else {
            const args = argsOrState as ObjectBucketArgs | undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["corsRules"] = args ? args.corsRules : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["lifecycleRules"] = args ? args.lifecycleRules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["versioning"] = args ? args.versioning : undefined;
            resourceInputs["endpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectBucket.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectBucket resources.
 */
export interface ObjectBucketState {
    /**
     * The canned ACL you want to apply to the bucket.
     */
    acl?: pulumi.Input<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    corsRules?: pulumi.Input<pulumi.Input<inputs.ObjectBucketCorsRule>[]>;
    /**
     * The endpoint URL of the bucket
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.ObjectBucketLifecycleRule>[]>;
    /**
     * The name of the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of tags (key / value) for the bucket.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    versioning?: pulumi.Input<inputs.ObjectBucketVersioning>;
}

/**
 * The set of arguments for constructing a ObjectBucket resource.
 */
export interface ObjectBucketArgs {
    /**
     * The canned ACL you want to apply to the bucket.
     */
    acl?: pulumi.Input<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    corsRules?: pulumi.Input<pulumi.Input<inputs.ObjectBucketCorsRule>[]>;
    /**
     * Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.ObjectBucketLifecycleRule>[]>;
    /**
     * The name of the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of tags (key / value) for the bucket.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    versioning?: pulumi.Input<inputs.ObjectBucketVersioning>;
}

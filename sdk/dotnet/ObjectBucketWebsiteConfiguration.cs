// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    /// <summary>
    /// Provides an Object bucket website configuration resource.
    /// For more information, see [Hosting Websites on Object bucket](https://www.scaleway.com/en/docs/storage/object/how-to/use-bucket-website/).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainObjectBucket = new Scaleway.ObjectBucket("mainObjectBucket", new()
    ///     {
    ///         Acl = "public-read",
    ///     });
    /// 
    ///     var mainObjectBucketWebsiteConfiguration = new Scaleway.ObjectBucketWebsiteConfiguration("mainObjectBucketWebsiteConfiguration", new()
    ///     {
    ///         Bucket = mainObjectBucket.Name,
    ///         IndexDocument = new Scaleway.Inputs.ObjectBucketWebsiteConfigurationIndexDocumentArgs
    ///         {
    ///             Suffix = "index.html",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Example with `policy`
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainObjectBucket = new Scaleway.ObjectBucket("mainObjectBucket", new()
    ///     {
    ///         Acl = "public-read",
    ///     });
    /// 
    ///     var mainObjectBucketPolicy = new Scaleway.ObjectBucketPolicy("mainObjectBucketPolicy", new()
    ///     {
    ///         Bucket = mainObjectBucket.Name,
    ///         Policy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Version"] = "2012-10-17",
    ///             ["Id"] = "MyPolicy",
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Sid"] = "GrantToEveryone",
    ///                     ["Effect"] = "Allow",
    ///                     ["Principal"] = "*",
    ///                     ["Action"] = new[]
    ///                     {
    ///                         "s3:GetObject",
    ///                     },
    ///                     ["Resource"] = new[]
    ///                     {
    ///                         "&lt;bucket-name&gt;/*",
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    ///     var mainObjectBucketWebsiteConfiguration = new Scaleway.ObjectBucketWebsiteConfiguration("mainObjectBucketWebsiteConfiguration", new()
    ///     {
    ///         Bucket = mainObjectBucket.Name,
    ///         IndexDocument = new Scaleway.Inputs.ObjectBucketWebsiteConfigurationIndexDocumentArgs
    ///         {
    ///             Suffix = "index.html",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## error_document
    /// 
    /// The error_document configuration block supports the following arguments:
    /// 
    /// * `key` - (Required) The object key name to use when a 4XX class error occurs.
    /// 
    /// ## index_document
    /// 
    /// The `index_document` configuration block supports the following arguments:
    /// 
    /// * `suffix` - (Required) A suffix that is appended to a request that is for a directory on the website endpoint.
    /// 
    /// &gt; **Important:** The suffix must not be empty and must not include a slash character. The routing is not supported.
    /// 
    /// In addition to all above arguments, the following attribute is exported:
    /// 
    /// * `id` - The bucket and region separated by a slash (/)
    /// * `website_domain` - The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/dns-cloud/how-to/manage-dns-records).
    /// * `website_endpoint` - The website endpoint.
    /// 
    /// &gt; **Important:** Please check our concepts section to know more about the [endpoint](https://www.scaleway.com/en/docs/storage/object/concepts/#endpoint).
    /// 
    /// ## Import
    /// 
    /// Website configuration Bucket can be imported using the `{region}/{bucketName}` identifier, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration some_bucket fr-par/some-bucket
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration")]
    public partial class ObjectBucketWebsiteConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Required, Forces new resource) The name of the bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// (Optional) The name of the error document for the website detailed below.
        /// </summary>
        [Output("errorDocument")]
        public Output<Outputs.ObjectBucketWebsiteConfigurationErrorDocument?> ErrorDocument { get; private set; } = null!;

        /// <summary>
        /// (Optional) The name of the index document for the website detailed below.
        /// </summary>
        [Output("indexDocument")]
        public Output<Outputs.ObjectBucketWebsiteConfigurationIndexDocument?> IndexDocument { get; private set; } = null!;

        /// <summary>
        /// The website endpoint.
        /// </summary>
        [Output("websiteDomain")]
        public Output<string> WebsiteDomain { get; private set; } = null!;

        /// <summary>
        /// The domain of the website endpoint.
        /// </summary>
        [Output("websiteEndpoint")]
        public Output<string> WebsiteEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectBucketWebsiteConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectBucketWebsiteConfiguration(string name, ObjectBucketWebsiteConfigurationArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration", name, args ?? new ObjectBucketWebsiteConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectBucketWebsiteConfiguration(string name, Input<string> id, ObjectBucketWebsiteConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ObjectBucketWebsiteConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectBucketWebsiteConfiguration Get(string name, Input<string> id, ObjectBucketWebsiteConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectBucketWebsiteConfiguration(name, id, state, options);
        }
    }

    public sealed class ObjectBucketWebsiteConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Required, Forces new resource) The name of the bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// (Optional) The name of the error document for the website detailed below.
        /// </summary>
        [Input("errorDocument")]
        public Input<Inputs.ObjectBucketWebsiteConfigurationErrorDocumentArgs>? ErrorDocument { get; set; }

        /// <summary>
        /// (Optional) The name of the index document for the website detailed below.
        /// </summary>
        [Input("indexDocument")]
        public Input<Inputs.ObjectBucketWebsiteConfigurationIndexDocumentArgs>? IndexDocument { get; set; }

        public ObjectBucketWebsiteConfigurationArgs()
        {
        }
        public static new ObjectBucketWebsiteConfigurationArgs Empty => new ObjectBucketWebsiteConfigurationArgs();
    }

    public sealed class ObjectBucketWebsiteConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Required, Forces new resource) The name of the bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// (Optional) The name of the error document for the website detailed below.
        /// </summary>
        [Input("errorDocument")]
        public Input<Inputs.ObjectBucketWebsiteConfigurationErrorDocumentGetArgs>? ErrorDocument { get; set; }

        /// <summary>
        /// (Optional) The name of the index document for the website detailed below.
        /// </summary>
        [Input("indexDocument")]
        public Input<Inputs.ObjectBucketWebsiteConfigurationIndexDocumentGetArgs>? IndexDocument { get; set; }

        /// <summary>
        /// The website endpoint.
        /// </summary>
        [Input("websiteDomain")]
        public Input<string>? WebsiteDomain { get; set; }

        /// <summary>
        /// The domain of the website endpoint.
        /// </summary>
        [Input("websiteEndpoint")]
        public Input<string>? WebsiteEndpoint { get; set; }

        public ObjectBucketWebsiteConfigurationState()
        {
        }
        public static new ObjectBucketWebsiteConfigurationState Empty => new ObjectBucketWebsiteConfigurationState();
    }
}

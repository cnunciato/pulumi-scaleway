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
    /// Creates and manages Scaleway Database Users.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
    /// 
    /// ## Examples
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Random = Pulumi.Random;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dbPassword = new Random.RandomPassword("dbPassword", new()
    ///     {
    ///         Length = 16,
    ///         Special = true,
    ///     });
    /// 
    ///     var dbAdmin = new Scaleway.DatabaseUser("dbAdmin", new()
    ///     {
    ///         InstanceId = scaleway_rdb_instance.Main.Id,
    ///         Password = dbPassword.Result,
    ///         IsAdmin = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database User can be imported using `{region}/{instance_id}/{name}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/databaseUser:DatabaseUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/databaseUser:DatabaseUser")]
    public partial class DatabaseUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The instance on which to create the user.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Grant admin permissions to the Database User.
        /// </summary>
        [Output("isAdmin")]
        public Output<bool?> IsAdmin { get; private set; } = null!;

        /// <summary>
        /// Database User name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Database User password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseUser(string name, DatabaseUserArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/databaseUser:DatabaseUser", name, args ?? new DatabaseUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseUser(string name, Input<string> id, DatabaseUserState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/databaseUser:DatabaseUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatabaseUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseUser Get(string name, Input<string> id, DatabaseUserState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseUser(name, id, state, options);
        }
    }

    public sealed class DatabaseUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance on which to create the user.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Grant admin permissions to the Database User.
        /// </summary>
        [Input("isAdmin")]
        public Input<bool>? IsAdmin { get; set; }

        /// <summary>
        /// Database User name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Database User password.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DatabaseUserArgs()
        {
        }
        public static new DatabaseUserArgs Empty => new DatabaseUserArgs();
    }

    public sealed class DatabaseUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance on which to create the user.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Grant admin permissions to the Database User.
        /// </summary>
        [Input("isAdmin")]
        public Input<bool>? IsAdmin { get; set; }

        /// <summary>
        /// Database User name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Database User password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DatabaseUserState()
        {
        }
        public static new DatabaseUserState Empty => new DatabaseUserState();
    }
}

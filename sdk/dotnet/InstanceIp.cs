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
    /// Creates and manages Scaleway Compute Instance IPs. For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#ips-268151).
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
    ///     var serverIp = new Scaleway.InstanceIp("serverIp");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IPs can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/instanceIp:InstanceIp server_ip fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/instanceIp:InstanceIp")]
    public partial class InstanceIp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IP address.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// The organization ID the IP is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the IP is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The reverse dns attached to this IP
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// The server associated with this IP
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the IP.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceIp(string name, InstanceIpArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceIp:InstanceIp", name, args ?? new InstanceIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceIp(string name, Input<string> id, InstanceIpState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceIp:InstanceIp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceIp Get(string name, Input<string> id, InstanceIpState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceIp(name, id, state, options);
        }
    }

    public sealed class InstanceIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceIpArgs()
        {
        }
        public static new InstanceIpArgs Empty => new InstanceIpArgs();
    }

    public sealed class InstanceIpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The organization ID the IP is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The reverse dns attached to this IP
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// The server associated with this IP
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceIpState()
        {
        }
        public static new InstanceIpState Empty => new InstanceIpState();
    }
}

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
    /// ## Import
    /// 
    /// Instance security group can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/instanceSecurityGroup:InstanceSecurityGroup web fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup")]
    public partial class InstanceSecurityGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the security group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        /// </summary>
        [Output("enableDefaultSecurity")]
        public Output<bool?> EnableDefaultSecurity { get; private set; } = null!;

        /// <summary>
        /// A boolean to specify whether to use instance_security_group_rules.
        /// If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        /// </summary>
        [Output("externalRules")]
        public Output<bool?> ExternalRules { get; private set; } = null!;

        /// <summary>
        /// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        /// </summary>
        [Output("inboundDefaultPolicy")]
        public Output<string?> InboundDefaultPolicy { get; private set; } = null!;

        /// <summary>
        /// A list of inbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        [Output("inboundRules")]
        public Output<ImmutableArray<Outputs.InstanceSecurityGroupInboundRule>> InboundRules { get; private set; } = null!;

        /// <summary>
        /// The name of the security group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the security group is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        /// </summary>
        [Output("outboundDefaultPolicy")]
        public Output<string?> OutboundDefaultPolicy { get; private set; } = null!;

        /// <summary>
        /// A list of outbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        [Output("outboundRules")]
        public Output<ImmutableArray<Outputs.InstanceSecurityGroupOutboundRule>> OutboundRules { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the security group is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// A boolean to specify whether the security group should be stateful or not.
        /// </summary>
        [Output("stateful")]
        public Output<bool?> Stateful { get; private set; } = null!;

        /// <summary>
        /// The tags of the security group.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the security group should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceSecurityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceSecurityGroup(string name, InstanceSecurityGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, args ?? new InstanceSecurityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceSecurityGroup(string name, Input<string> id, InstanceSecurityGroupState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceSecurityGroup:InstanceSecurityGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceSecurityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceSecurityGroup Get(string name, Input<string> id, InstanceSecurityGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceSecurityGroup(name, id, state, options);
        }
    }

    public sealed class InstanceSecurityGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the security group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        /// </summary>
        [Input("enableDefaultSecurity")]
        public Input<bool>? EnableDefaultSecurity { get; set; }

        /// <summary>
        /// A boolean to specify whether to use instance_security_group_rules.
        /// If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        /// </summary>
        [Input("externalRules")]
        public Input<bool>? ExternalRules { get; set; }

        /// <summary>
        /// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        /// </summary>
        [Input("inboundDefaultPolicy")]
        public Input<string>? InboundDefaultPolicy { get; set; }

        [Input("inboundRules")]
        private InputList<Inputs.InstanceSecurityGroupInboundRuleArgs>? _inboundRules;

        /// <summary>
        /// A list of inbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public InputList<Inputs.InstanceSecurityGroupInboundRuleArgs> InboundRules
        {
            get => _inboundRules ?? (_inboundRules = new InputList<Inputs.InstanceSecurityGroupInboundRuleArgs>());
            set => _inboundRules = value;
        }

        /// <summary>
        /// The name of the security group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        /// </summary>
        [Input("outboundDefaultPolicy")]
        public Input<string>? OutboundDefaultPolicy { get; set; }

        [Input("outboundRules")]
        private InputList<Inputs.InstanceSecurityGroupOutboundRuleArgs>? _outboundRules;

        /// <summary>
        /// A list of outbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public InputList<Inputs.InstanceSecurityGroupOutboundRuleArgs> OutboundRules
        {
            get => _outboundRules ?? (_outboundRules = new InputList<Inputs.InstanceSecurityGroupOutboundRuleArgs>());
            set => _outboundRules = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the security group is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// A boolean to specify whether the security group should be stateful or not.
        /// </summary>
        [Input("stateful")]
        public Input<bool>? Stateful { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags of the security group.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the security group should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceSecurityGroupArgs()
        {
        }
        public static new InstanceSecurityGroupArgs Empty => new InstanceSecurityGroupArgs();
    }

    public sealed class InstanceSecurityGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the security group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        /// </summary>
        [Input("enableDefaultSecurity")]
        public Input<bool>? EnableDefaultSecurity { get; set; }

        /// <summary>
        /// A boolean to specify whether to use instance_security_group_rules.
        /// If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        /// </summary>
        [Input("externalRules")]
        public Input<bool>? ExternalRules { get; set; }

        /// <summary>
        /// The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        /// </summary>
        [Input("inboundDefaultPolicy")]
        public Input<string>? InboundDefaultPolicy { get; set; }

        [Input("inboundRules")]
        private InputList<Inputs.InstanceSecurityGroupInboundRuleGetArgs>? _inboundRules;

        /// <summary>
        /// A list of inbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public InputList<Inputs.InstanceSecurityGroupInboundRuleGetArgs> InboundRules
        {
            get => _inboundRules ?? (_inboundRules = new InputList<Inputs.InstanceSecurityGroupInboundRuleGetArgs>());
            set => _inboundRules = value;
        }

        /// <summary>
        /// The name of the security group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the security group is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        /// </summary>
        [Input("outboundDefaultPolicy")]
        public Input<string>? OutboundDefaultPolicy { get; set; }

        [Input("outboundRules")]
        private InputList<Inputs.InstanceSecurityGroupOutboundRuleGetArgs>? _outboundRules;

        /// <summary>
        /// A list of outbound rule to add to the security group. (Structure is documented below.)
        /// </summary>
        public InputList<Inputs.InstanceSecurityGroupOutboundRuleGetArgs> OutboundRules
        {
            get => _outboundRules ?? (_outboundRules = new InputList<Inputs.InstanceSecurityGroupOutboundRuleGetArgs>());
            set => _outboundRules = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the security group is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// A boolean to specify whether the security group should be stateful or not.
        /// </summary>
        [Input("stateful")]
        public Input<bool>? Stateful { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags of the security group.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the security group should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceSecurityGroupState()
        {
        }
        public static new InstanceSecurityGroupState Empty => new InstanceSecurityGroupState();
    }
}

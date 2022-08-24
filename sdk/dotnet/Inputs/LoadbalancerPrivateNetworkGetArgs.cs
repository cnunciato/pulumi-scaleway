// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Inputs
{

    public sealed class LoadbalancerPrivateNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Set to true if you want to let DHCP assign IP addresses. See below.
        /// </summary>
        [Input("dhcpConfig")]
        public Input<bool>? DhcpConfig { get; set; }

        /// <summary>
        /// (Required) The ID of the Private Network to associate.
        /// </summary>
        [Input("privateNetworkId", required: true)]
        public Input<string> PrivateNetworkId { get; set; } = null!;

        [Input("staticConfigs")]
        private InputList<string>? _staticConfigs;

        /// <summary>
        /// (Optional) Define two local ip address of your choice for each load balancer instance. See below.
        /// </summary>
        public InputList<string> StaticConfigs
        {
            get => _staticConfigs ?? (_staticConfigs = new InputList<string>());
            set => _staticConfigs = value;
        }

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public LoadbalancerPrivateNetworkGetArgs()
        {
        }
        public static new LoadbalancerPrivateNetworkGetArgs Empty => new LoadbalancerPrivateNetworkGetArgs();
    }
}

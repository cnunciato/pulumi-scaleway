// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Inputs
{

    public sealed class InstanceSecurityGroupInboundRuleArgs : Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("ipRange")]
        public Input<string>? IpRange { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public InstanceSecurityGroupInboundRuleArgs()
        {
        }
    }
}

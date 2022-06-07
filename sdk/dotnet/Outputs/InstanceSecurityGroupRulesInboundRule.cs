// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Outputs
{

    [OutputType]
    public sealed class InstanceSecurityGroupRulesInboundRule
    {
        public readonly string Action;
        public readonly string? Ip;
        public readonly string? IpRange;
        public readonly int? Port;
        public readonly string? PortRange;
        public readonly string? Protocol;

        [OutputConstructor]
        private InstanceSecurityGroupRulesInboundRule(
            string action,

            string? ip,

            string? ipRange,

            int? port,

            string? portRange,

            string? protocol)
        {
            Action = action;
            Ip = ip;
            IpRange = ipRange;
            Port = port;
            PortRange = portRange;
            Protocol = protocol;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Inputs
{

    public sealed class DatabaseInstanceLoadBalancerGetArgs : Pulumi.ResourceArgs
    {
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        public DatabaseInstanceLoadBalancerGetArgs()
        {
        }
    }
}

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

    public sealed class DatabaseInstanceLoadBalancerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the endpoint of the private network.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// Name of the endpoint.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IP of the endpoint.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Port of the endpoint.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public DatabaseInstanceLoadBalancerArgs()
        {
        }
        public static new DatabaseInstanceLoadBalancerArgs Empty => new DatabaseInstanceLoadBalancerArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetRedisClusterPrivateNetworkResult
    {
        public readonly string EndpointId;
        public readonly string Id;
        public readonly ImmutableArray<string> ServiceIps;
        /// <summary>
        /// `region`) The zone in which the server exists.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetRedisClusterPrivateNetworkResult(
            string endpointId,

            string id,

            ImmutableArray<string> serviceIps,

            string zone)
        {
            EndpointId = endpointId;
            Id = id;
            ServiceIps = serviceIps;
            Zone = zone;
        }
    }
}

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
    public sealed class DomainRecordWeighted
    {
        public readonly string Ip;
        public readonly int Weight;

        [OutputConstructor]
        private DomainRecordWeighted(
            string ip,

            int weight)
        {
            Ip = ip;
            Weight = weight;
        }
    }
}

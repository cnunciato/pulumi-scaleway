// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Inputs
{

    public sealed class LoadbalancerBackendHealthCheckHttpsGetArgs : Pulumi.ResourceArgs
    {
        [Input("code")]
        public Input<int>? Code { get; set; }

        [Input("method")]
        public Input<string>? Method { get; set; }

        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public LoadbalancerBackendHealthCheckHttpsGetArgs()
        {
        }
    }
}

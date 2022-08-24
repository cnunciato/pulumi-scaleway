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
    public sealed class IotDeviceMessageFilters
    {
        /// <summary>
        /// Rules used to restrict topics the device can publish to.
        /// </summary>
        public readonly Outputs.IotDeviceMessageFiltersPublish? Publish;
        /// <summary>
        /// Rules used to restrict topics the device can subscribe to.
        /// </summary>
        public readonly Outputs.IotDeviceMessageFiltersSubscribe? Subscribe;

        [OutputConstructor]
        private IotDeviceMessageFilters(
            Outputs.IotDeviceMessageFiltersPublish? publish,

            Outputs.IotDeviceMessageFiltersSubscribe? subscribe)
        {
            Publish = publish;
            Subscribe = subscribe;
        }
    }
}

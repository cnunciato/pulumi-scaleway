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
    public static class GetDomainRecord
    {
        /// <summary>
        /// Gets information about a domain record.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byContent = Scaleway.GetDomainRecord.Invoke(new()
        ///     {
        ///         Data = "1.2.3.4",
        ///         DnsZone = "domain.tld",
        ///         Name = "www",
        ///         Type = "A",
        ///     });
        /// 
        ///     var byId = Scaleway.GetDomainRecord.Invoke(new()
        ///     {
        ///         DnsZone = "domain.tld",
        ///         RecordId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDomainRecordResult> InvokeAsync(GetDomainRecordArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainRecordResult>("scaleway:index/getDomainRecord:getDomainRecord", args ?? new GetDomainRecordArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a domain record.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byContent = Scaleway.GetDomainRecord.Invoke(new()
        ///     {
        ///         Data = "1.2.3.4",
        ///         DnsZone = "domain.tld",
        ///         Name = "www",
        ///         Type = "A",
        ///     });
        /// 
        ///     var byId = Scaleway.GetDomainRecord.Invoke(new()
        ///     {
        ///         DnsZone = "domain.tld",
        ///         RecordId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDomainRecordResult> Invoke(GetDomainRecordInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDomainRecordResult>("scaleway:index/getDomainRecord:getDomainRecord", args ?? new GetDomainRecordInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainRecordArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
        /// Cannot be used with `record_id`.
        /// </summary>
        [Input("data")]
        public string? Data { get; set; }

        /// <summary>
        /// The IP address.
        /// </summary>
        [Input("dnsZone")]
        public string? DnsZone { get; set; }

        /// <summary>
        /// The name of the record (can be an empty string for a root record).
        /// Cannot be used with `record_id`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The record ID.
        /// Cannot be used with `name`, `type` and `data`.
        /// </summary>
        [Input("recordId")]
        public string? RecordId { get; set; }

        /// <summary>
        /// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
        /// Cannot be used with `record_id`.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetDomainRecordArgs()
        {
        }
        public static new GetDomainRecordArgs Empty => new GetDomainRecordArgs();
    }

    public sealed class GetDomainRecordInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
        /// Cannot be used with `record_id`.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// The IP address.
        /// </summary>
        [Input("dnsZone")]
        public Input<string>? DnsZone { get; set; }

        /// <summary>
        /// The name of the record (can be an empty string for a root record).
        /// Cannot be used with `record_id`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The record ID.
        /// Cannot be used with `name`, `type` and `data`.
        /// </summary>
        [Input("recordId")]
        public Input<string>? RecordId { get; set; }

        /// <summary>
        /// The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
        /// Cannot be used with `record_id`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetDomainRecordInvokeArgs()
        {
        }
        public static new GetDomainRecordInvokeArgs Empty => new GetDomainRecordInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainRecordResult
    {
        public readonly string? Data;
        public readonly string? DnsZone;
        /// <summary>
        /// Dynamic record base on user geolocalisation (More information about dynamic records)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainRecordGeoIpResult> GeoIps;
        /// <summary>
        /// Dynamic record base on URL resolve (More information about dynamic records)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainRecordHttpServiceResult> HttpServices;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool KeepEmptyZone;
        public readonly string? Name;
        /// <summary>
        /// The priority of the record (mostly used with an `MX` record)
        /// </summary>
        public readonly int Priority;
        public readonly string ProjectId;
        public readonly string? RecordId;
        public readonly bool RootZone;
        /// <summary>
        /// Time To Live of the record in seconds.
        /// </summary>
        public readonly int Ttl;
        public readonly string? Type;
        /// <summary>
        /// Dynamic record based on the client’s (resolver) subnet (More information about dynamic records)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainRecordViewResult> Views;
        /// <summary>
        /// Dynamic record base on IP weights (More information about dynamic records)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainRecordWeightedResult> Weighteds;

        [OutputConstructor]
        private GetDomainRecordResult(
            string? data,

            string? dnsZone,

            ImmutableArray<Outputs.GetDomainRecordGeoIpResult> geoIps,

            ImmutableArray<Outputs.GetDomainRecordHttpServiceResult> httpServices,

            string id,

            bool keepEmptyZone,

            string? name,

            int priority,

            string projectId,

            string? recordId,

            bool rootZone,

            int ttl,

            string? type,

            ImmutableArray<Outputs.GetDomainRecordViewResult> views,

            ImmutableArray<Outputs.GetDomainRecordWeightedResult> weighteds)
        {
            Data = data;
            DnsZone = dnsZone;
            GeoIps = geoIps;
            HttpServices = httpServices;
            Id = id;
            KeepEmptyZone = keepEmptyZone;
            Name = name;
            Priority = priority;
            ProjectId = projectId;
            RecordId = recordId;
            RootZone = rootZone;
            Ttl = ttl;
            Type = type;
            Views = views;
            Weighteds = weighteds;
        }
    }
}

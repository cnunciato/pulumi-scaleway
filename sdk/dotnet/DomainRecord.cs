// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    [ScalewayResourceType("scaleway:index/domainRecord:DomainRecord")]
    public partial class DomainRecord : Pulumi.CustomResource
    {
        /// <summary>
        /// The data of the record
        /// </summary>
        [Output("data")]
        public Output<string> Data { get; private set; } = null!;

        /// <summary>
        /// The zone you want to add the record in
        /// </summary>
        [Output("dnsZone")]
        public Output<string> DnsZone { get; private set; } = null!;

        /// <summary>
        /// Return record based on client localisation
        /// </summary>
        [Output("geoIp")]
        public Output<Outputs.DomainRecordGeoIp?> GeoIp { get; private set; } = null!;

        /// <summary>
        /// Return record based on client localisation
        /// </summary>
        [Output("httpService")]
        public Output<Outputs.DomainRecordHttpService?> HttpService { get; private set; } = null!;

        /// <summary>
        /// When destroy a resource record, if a zone have only NS, delete the zone
        /// </summary>
        [Output("keepEmptyZone")]
        public Output<bool?> KeepEmptyZone { get; private set; } = null!;

        /// <summary>
        /// The name of the record
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The priority of the record
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Does the DNS zone is the root zone or not
        /// </summary>
        [Output("rootZone")]
        public Output<bool> RootZone { get; private set; } = null!;

        /// <summary>
        /// The ttl of the record
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of the record
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Return record based on client subnet
        /// </summary>
        [Output("views")]
        public Output<ImmutableArray<Outputs.DomainRecordView>> Views { get; private set; } = null!;

        /// <summary>
        /// Return record based on weight
        /// </summary>
        [Output("weighteds")]
        public Output<ImmutableArray<Outputs.DomainRecordWeighted>> Weighteds { get; private set; } = null!;


        /// <summary>
        /// Create a DomainRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainRecord(string name, DomainRecordArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/domainRecord:DomainRecord", name, args ?? new DomainRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainRecord(string name, Input<string> id, DomainRecordState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/domainRecord:DomainRecord", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/jaxxstorm/pulumi-scaleway/releases/download/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainRecord Get(string name, Input<string> id, DomainRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainRecord(name, id, state, options);
        }
    }

    public sealed class DomainRecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data of the record
        /// </summary>
        [Input("data", required: true)]
        public Input<string> Data { get; set; } = null!;

        /// <summary>
        /// The zone you want to add the record in
        /// </summary>
        [Input("dnsZone", required: true)]
        public Input<string> DnsZone { get; set; } = null!;

        /// <summary>
        /// Return record based on client localisation
        /// </summary>
        [Input("geoIp")]
        public Input<Inputs.DomainRecordGeoIpArgs>? GeoIp { get; set; }

        /// <summary>
        /// Return record based on client localisation
        /// </summary>
        [Input("httpService")]
        public Input<Inputs.DomainRecordHttpServiceArgs>? HttpService { get; set; }

        /// <summary>
        /// When destroy a resource record, if a zone have only NS, delete the zone
        /// </summary>
        [Input("keepEmptyZone")]
        public Input<bool>? KeepEmptyZone { get; set; }

        /// <summary>
        /// The name of the record
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority of the record
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The ttl of the record
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of the record
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("views")]
        private InputList<Inputs.DomainRecordViewArgs>? _views;

        /// <summary>
        /// Return record based on client subnet
        /// </summary>
        public InputList<Inputs.DomainRecordViewArgs> Views
        {
            get => _views ?? (_views = new InputList<Inputs.DomainRecordViewArgs>());
            set => _views = value;
        }

        [Input("weighteds")]
        private InputList<Inputs.DomainRecordWeightedArgs>? _weighteds;

        /// <summary>
        /// Return record based on weight
        /// </summary>
        public InputList<Inputs.DomainRecordWeightedArgs> Weighteds
        {
            get => _weighteds ?? (_weighteds = new InputList<Inputs.DomainRecordWeightedArgs>());
            set => _weighteds = value;
        }

        public DomainRecordArgs()
        {
        }
    }

    public sealed class DomainRecordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data of the record
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// The zone you want to add the record in
        /// </summary>
        [Input("dnsZone")]
        public Input<string>? DnsZone { get; set; }

        /// <summary>
        /// Return record based on client localisation
        /// </summary>
        [Input("geoIp")]
        public Input<Inputs.DomainRecordGeoIpGetArgs>? GeoIp { get; set; }

        /// <summary>
        /// Return record based on client localisation
        /// </summary>
        [Input("httpService")]
        public Input<Inputs.DomainRecordHttpServiceGetArgs>? HttpService { get; set; }

        /// <summary>
        /// When destroy a resource record, if a zone have only NS, delete the zone
        /// </summary>
        [Input("keepEmptyZone")]
        public Input<bool>? KeepEmptyZone { get; set; }

        /// <summary>
        /// The name of the record
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority of the record
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Does the DNS zone is the root zone or not
        /// </summary>
        [Input("rootZone")]
        public Input<bool>? RootZone { get; set; }

        /// <summary>
        /// The ttl of the record
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of the record
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("views")]
        private InputList<Inputs.DomainRecordViewGetArgs>? _views;

        /// <summary>
        /// Return record based on client subnet
        /// </summary>
        public InputList<Inputs.DomainRecordViewGetArgs> Views
        {
            get => _views ?? (_views = new InputList<Inputs.DomainRecordViewGetArgs>());
            set => _views = value;
        }

        [Input("weighteds")]
        private InputList<Inputs.DomainRecordWeightedGetArgs>? _weighteds;

        /// <summary>
        /// Return record based on weight
        /// </summary>
        public InputList<Inputs.DomainRecordWeightedGetArgs> Weighteds
        {
            get => _weighteds ?? (_weighteds = new InputList<Inputs.DomainRecordWeightedGetArgs>());
            set => _weighteds = value;
        }

        public DomainRecordState()
        {
        }
    }
}

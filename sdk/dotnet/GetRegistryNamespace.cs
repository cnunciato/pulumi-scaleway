// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    public static class GetRegistryNamespace
    {
        public static Task<GetRegistryNamespaceResult> InvokeAsync(GetRegistryNamespaceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryNamespaceResult>("scaleway:index/getRegistryNamespace:getRegistryNamespace", args ?? new GetRegistryNamespaceArgs(), options.WithDefaults());

        public static Output<GetRegistryNamespaceResult> Invoke(GetRegistryNamespaceInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegistryNamespaceResult>("scaleway:index/getRegistryNamespace:getRegistryNamespace", args ?? new GetRegistryNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistryNamespaceArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("namespaceId")]
        public string? NamespaceId { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        public GetRegistryNamespaceArgs()
        {
        }
    }

    public sealed class GetRegistryNamespaceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetRegistryNamespaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryNamespaceResult
    {
        public readonly string Description;
        public readonly string Endpoint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsPublic;
        public readonly string? Name;
        public readonly string? NamespaceId;
        public readonly string OrganizationId;
        public readonly string ProjectId;
        public readonly string? Region;

        [OutputConstructor]
        private GetRegistryNamespaceResult(
            string description,

            string endpoint,

            string id,

            bool isPublic,

            string? name,

            string? namespaceId,

            string organizationId,

            string projectId,

            string? region)
        {
            Description = description;
            Endpoint = endpoint;
            Id = id;
            IsPublic = isPublic;
            Name = name;
            NamespaceId = namespaceId;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
        }
    }
}

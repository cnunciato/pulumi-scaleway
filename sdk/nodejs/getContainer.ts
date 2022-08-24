// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about the Scaleway Container.
 *
 * For more information consult the [documentation](https://www.scaleway.com/en/docs/faq/serverless-containers/).
 *
 * For more details about the limitation check [containers-limitations](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/).
 *
 * You can check also our [containers guide](https://www.scaleway.com/en/docs/compute/containers/concepts/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const mainContainerNamespace = new scaleway.ContainerNamespace("mainContainerNamespace", {});
 * const mainContainer = new scaleway.Container("mainContainer", {namespaceId: mainContainerNamespace.id});
 * const byName = scaleway.getContainerOutput({
 *     namespaceId: mainContainerNamespace.id,
 *     name: mainContainer.name,
 * });
 * const byId = scaleway.getContainerOutput({
 *     namespaceId: mainContainerNamespace.id,
 *     containerId: mainContainer.id,
 * });
 * ```
 */
export function getContainer(args: GetContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway:index/getContainer:getContainer", {
        "containerId": args.containerId,
        "name": args.name,
        "namespaceId": args.namespaceId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainer.
 */
export interface GetContainerArgs {
    containerId?: string;
    /**
     * The unique name of the container name.
     */
    name?: string;
    /**
     * The container namespace ID of the container.
     */
    namespaceId: string;
    /**
     * (Defaults to provider `region`) The region in which the container was created.
     */
    region?: string;
}

/**
 * A collection of values returned by getContainer.
 */
export interface GetContainerResult {
    readonly containerId?: string;
    /**
     * The amount of vCPU computing resources to allocate to each container. Defaults  to 70.
     */
    readonly cpuLimit: number;
    /**
     * The cron status of the container.
     */
    readonly cronStatus: string;
    /**
     * Boolean indicating whether the container is on a production environment.
     */
    readonly deploy: boolean;
    readonly description: string;
    /**
     * The container domain name.
     */
    readonly domainName: string;
    /**
     * The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
     */
    readonly environmentVariables: {[key: string]: string};
    /**
     * The error message of the container.
     */
    readonly errorMessage: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
     */
    readonly maxConcurrency: number;
    /**
     * The maximum of number of instances this container can scale to. Default to 20.
     */
    readonly maxScale: number;
    /**
     * The memory computing resources in MB to allocate to each container. Defaults to 128.
     */
    readonly memoryLimit: number;
    /**
     * The minimum of running container instances continuously. Defaults to 0.
     */
    readonly minScale: number;
    readonly name?: string;
    readonly namespaceId: string;
    /**
     * The port to expose the container. Defaults to 8080.
     */
    readonly port: number;
    /**
     * The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
     */
    readonly privacy: string;
    /**
     * The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
     */
    readonly protocol: string;
    /**
     * (Defaults to provider `region`) The region in which the container was created.
     */
    readonly region?: string;
    /**
     * The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
     */
    readonly registryImage: string;
    /**
     * The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
     */
    readonly registrySha256: string;
    /**
     * The container status.
     */
    readonly status: string;
    /**
     * The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
     */
    readonly timeout: number;
}

export function getContainerOutput(args: GetContainerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContainerResult> {
    return pulumi.output(args).apply(a => getContainer(a, opts))
}

/**
 * A collection of arguments for invoking getContainer.
 */
export interface GetContainerOutputArgs {
    containerId?: pulumi.Input<string>;
    /**
     * The unique name of the container name.
     */
    name?: pulumi.Input<string>;
    /**
     * The container namespace ID of the container.
     */
    namespaceId: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region in which the container was created.
     */
    region?: pulumi.Input<string>;
}

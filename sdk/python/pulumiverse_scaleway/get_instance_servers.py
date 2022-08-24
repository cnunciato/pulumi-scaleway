# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetInstanceServersResult',
    'AwaitableGetInstanceServersResult',
    'get_instance_servers',
    'get_instance_servers_output',
]

@pulumi.output_type
class GetInstanceServersResult:
    """
    A collection of values returned by getInstanceServers.
    """
    def __init__(__self__, id=None, name=None, organization_id=None, project_id=None, servers=None, tags=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if servers and not isinstance(servers, list):
            raise TypeError("Expected argument 'servers' to be a list")
        pulumi.set(__self__, "servers", servers)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the server.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The organization ID the server is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The ID of the project the server is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def servers(self) -> Sequence['outputs.GetInstanceServersServerResult']:
        """
        List of found servers
        """
        return pulumi.get(self, "servers")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        The tags associated with the server.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The zone in which the server is.
        """
        return pulumi.get(self, "zone")


class AwaitableGetInstanceServersResult(GetInstanceServersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceServersResult(
            id=self.id,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            servers=self.servers,
            tags=self.tags,
            zone=self.zone)


def get_instance_servers(name: Optional[str] = None,
                         project_id: Optional[str] = None,
                         tags: Optional[Sequence[str]] = None,
                         zone: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceServersResult:
    """
    Gets information about multiple instance servers.

    ## Examples

    ### Basic

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_instance_servers(name="myserver",
        zone="fr-par-2")
    ```


    :param str name: The server name used as filter. Servers with a name like it are listed.
    :param str project_id: The ID of the project the server is associated with.
    :param Sequence[str] tags: List of tags used as filter. Servers with these exact tags are listed.
    :param str zone: `zone`) The zone in which servers exist.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['tags'] = tags
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getInstanceServers:getInstanceServers', __args__, opts=opts, typ=GetInstanceServersResult).value

    return AwaitableGetInstanceServersResult(
        id=__ret__.id,
        name=__ret__.name,
        organization_id=__ret__.organization_id,
        project_id=__ret__.project_id,
        servers=__ret__.servers,
        tags=__ret__.tags,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_instance_servers)
def get_instance_servers_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                zone: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceServersResult]:
    """
    Gets information about multiple instance servers.

    ## Examples

    ### Basic

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_instance_servers(name="myserver",
        zone="fr-par-2")
    ```


    :param str name: The server name used as filter. Servers with a name like it are listed.
    :param str project_id: The ID of the project the server is associated with.
    :param Sequence[str] tags: List of tags used as filter. Servers with these exact tags are listed.
    :param str zone: `zone`) The zone in which servers exist.
    """
    ...

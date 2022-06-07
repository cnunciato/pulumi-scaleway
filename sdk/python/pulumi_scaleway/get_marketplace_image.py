# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetMarketplaceImageResult',
    'AwaitableGetMarketplaceImageResult',
    'get_marketplace_image',
    'get_marketplace_image_output',
]

@pulumi.output_type
class GetMarketplaceImageResult:
    """
    A collection of values returned by getMarketplaceImage.
    """
    def __init__(__self__, id=None, instance_type=None, label=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        pulumi.set(__self__, "label", label)
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
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def label(self) -> str:
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetMarketplaceImageResult(GetMarketplaceImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMarketplaceImageResult(
            id=self.id,
            instance_type=self.instance_type,
            label=self.label,
            zone=self.zone)


def get_marketplace_image(instance_type: Optional[str] = None,
                          label: Optional[str] = None,
                          zone: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMarketplaceImageResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['instanceType'] = instance_type
    __args__['label'] = label
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('scaleway:index/getMarketplaceImage:getMarketplaceImage', __args__, opts=opts, typ=GetMarketplaceImageResult).value

    return AwaitableGetMarketplaceImageResult(
        id=__ret__.id,
        instance_type=__ret__.instance_type,
        label=__ret__.label,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_marketplace_image)
def get_marketplace_image_output(instance_type: Optional[pulumi.Input[Optional[str]]] = None,
                                 label: Optional[pulumi.Input[str]] = None,
                                 zone: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMarketplaceImageResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...

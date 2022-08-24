# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AccountSshKeyArgs', 'AccountSshKey']

@pulumi.input_type
class AccountSshKeyArgs:
    def __init__(__self__, *,
                 public_key: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AccountSshKey resource.
        :param pulumi.Input[str] public_key: The public SSH key to be added.
        :param pulumi.Input[str] name: The name of the SSH key.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the SSH key is associated with.
        """
        pulumi.set(__self__, "public_key", public_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Input[str]:
        """
        The public SSH key to be added.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SSH key.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the SSH key is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _AccountSshKeyState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccountSshKey resources.
        :param pulumi.Input[str] name: The name of the SSH key.
        :param pulumi.Input[str] organization_id: The organization ID the SSH key is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the SSH key is associated with.
        :param pulumi.Input[str] public_key: The public SSH key to be added.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SSH key.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the SSH key is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the SSH key is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        The public SSH key to be added.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)


class AccountSshKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages user SSH keys to access servers provisioned on Scaleway.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.AccountSshKey("main", public_key="<YOUR-PUBLIC-SSH-KEY>")
        ```

        ## Import

        SSH keys can be imported using the `id`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/accountSshKey:AccountSshKey main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the SSH key.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the SSH key is associated with.
        :param pulumi.Input[str] public_key: The public SSH key to be added.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountSshKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages user SSH keys to access servers provisioned on Scaleway.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.AccountSshKey("main", public_key="<YOUR-PUBLIC-SSH-KEY>")
        ```

        ## Import

        SSH keys can be imported using the `id`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/accountSshKey:AccountSshKey main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param AccountSshKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountSshKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountSshKeyArgs.__new__(AccountSshKeyArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            if public_key is None and not opts.urn:
                raise TypeError("Missing required property 'public_key'")
            __props__.__dict__["public_key"] = public_key
            __props__.__dict__["organization_id"] = None
        super(AccountSshKey, __self__).__init__(
            'scaleway:index/accountSshKey:AccountSshKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None) -> 'AccountSshKey':
        """
        Get an existing AccountSshKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the SSH key.
        :param pulumi.Input[str] organization_id: The organization ID the SSH key is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the SSH key is associated with.
        :param pulumi.Input[str] public_key: The public SSH key to be added.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountSshKeyState.__new__(_AccountSshKeyState)

        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["public_key"] = public_key
        return AccountSshKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the SSH key.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the SSH key is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the SSH key is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        The public SSH key to be added.
        """
        return pulumi.get(self, "public_key")


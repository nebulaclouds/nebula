# coding: utf-8

"""
    nebulaidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class CoreGPUAccelerator(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'device': 'str',
        'unpartitioned': 'bool',
        'partition_size': 'str'
    }

    attribute_map = {
        'device': 'device',
        'unpartitioned': 'unpartitioned',
        'partition_size': 'partition_size'
    }

    def __init__(self, device=None, unpartitioned=None, partition_size=None):  # noqa: E501
        """CoreGPUAccelerator - a model defined in Swagger"""  # noqa: E501

        self._device = None
        self._unpartitioned = None
        self._partition_size = None
        self.discriminator = None

        if device is not None:
            self.device = device
        if unpartitioned is not None:
            self.unpartitioned = unpartitioned
        if partition_size is not None:
            self.partition_size = partition_size

    @property
    def device(self):
        """Gets the device of this CoreGPUAccelerator.  # noqa: E501

        This can be any arbitrary string, and should be informed by the labels or taints associated with the nodes in question. Default cloud provider labels typically use the following values: `nvidia-tesla-t4`, `nvidia-tesla-a100`, etc.  # noqa: E501

        :return: The device of this CoreGPUAccelerator.  # noqa: E501
        :rtype: str
        """
        return self._device

    @device.setter
    def device(self, device):
        """Sets the device of this CoreGPUAccelerator.

        This can be any arbitrary string, and should be informed by the labels or taints associated with the nodes in question. Default cloud provider labels typically use the following values: `nvidia-tesla-t4`, `nvidia-tesla-a100`, etc.  # noqa: E501

        :param device: The device of this CoreGPUAccelerator.  # noqa: E501
        :type: str
        """

        self._device = device

    @property
    def unpartitioned(self):
        """Gets the unpartitioned of this CoreGPUAccelerator.  # noqa: E501


        :return: The unpartitioned of this CoreGPUAccelerator.  # noqa: E501
        :rtype: bool
        """
        return self._unpartitioned

    @unpartitioned.setter
    def unpartitioned(self, unpartitioned):
        """Sets the unpartitioned of this CoreGPUAccelerator.


        :param unpartitioned: The unpartitioned of this CoreGPUAccelerator.  # noqa: E501
        :type: bool
        """

        self._unpartitioned = unpartitioned

    @property
    def partition_size(self):
        """Gets the partition_size of this CoreGPUAccelerator.  # noqa: E501

        Like `device`, this can be any arbitrary string, and should be informed by the labels or taints associated with the nodes in question. Default cloud provider labels typically use the following values: `1g.5gb`, `2g.10gb`, etc.  # noqa: E501

        :return: The partition_size of this CoreGPUAccelerator.  # noqa: E501
        :rtype: str
        """
        return self._partition_size

    @partition_size.setter
    def partition_size(self, partition_size):
        """Sets the partition_size of this CoreGPUAccelerator.

        Like `device`, this can be any arbitrary string, and should be informed by the labels or taints associated with the nodes in question. Default cloud provider labels typically use the following values: `1g.5gb`, `2g.10gb`, etc.  # noqa: E501

        :param partition_size: The partition_size of this CoreGPUAccelerator.  # noqa: E501
        :type: str
        """

        self._partition_size = partition_size

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CoreGPUAccelerator, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreGPUAccelerator):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
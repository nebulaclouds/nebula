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


class CoreWorkflowMetadataDefaults(object):
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
        'interruptible': 'bool'
    }

    attribute_map = {
        'interruptible': 'interruptible'
    }

    def __init__(self, interruptible=None):  # noqa: E501
        """CoreWorkflowMetadataDefaults - a model defined in Swagger"""  # noqa: E501

        self._interruptible = None
        self.discriminator = None

        if interruptible is not None:
            self.interruptible = interruptible

    @property
    def interruptible(self):
        """Gets the interruptible of this CoreWorkflowMetadataDefaults.  # noqa: E501

        Whether child nodes of the workflow are interruptible.  # noqa: E501

        :return: The interruptible of this CoreWorkflowMetadataDefaults.  # noqa: E501
        :rtype: bool
        """
        return self._interruptible

    @interruptible.setter
    def interruptible(self, interruptible):
        """Sets the interruptible of this CoreWorkflowMetadataDefaults.

        Whether child nodes of the workflow are interruptible.  # noqa: E501

        :param interruptible: The interruptible of this CoreWorkflowMetadataDefaults.  # noqa: E501
        :type: bool
        """

        self._interruptible = interruptible

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
        if issubclass(CoreWorkflowMetadataDefaults, dict):
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
        if not isinstance(other, CoreWorkflowMetadataDefaults):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

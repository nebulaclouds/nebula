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


class AdminClusterResourceAttributes(object):
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
        'attributes': 'dict(str, str)'
    }

    attribute_map = {
        'attributes': 'attributes'
    }

    def __init__(self, attributes=None):  # noqa: E501
        """AdminClusterResourceAttributes - a model defined in Swagger"""  # noqa: E501

        self._attributes = None
        self.discriminator = None

        if attributes is not None:
            self.attributes = attributes

    @property
    def attributes(self):
        """Gets the attributes of this AdminClusterResourceAttributes.  # noqa: E501

        Custom resource attributes which will be applied in cluster resource creation (e.g. quotas). Map keys are the *case-sensitive* names of variables in templatized resource files. Map values should be the custom values which get substituted during resource creation.  # noqa: E501

        :return: The attributes of this AdminClusterResourceAttributes.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._attributes

    @attributes.setter
    def attributes(self, attributes):
        """Sets the attributes of this AdminClusterResourceAttributes.

        Custom resource attributes which will be applied in cluster resource creation (e.g. quotas). Map keys are the *case-sensitive* names of variables in templatized resource files. Map values should be the custom values which get substituted during resource creation.  # noqa: E501

        :param attributes: The attributes of this AdminClusterResourceAttributes.  # noqa: E501
        :type: dict(str, str)
        """

        self._attributes = attributes

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
        if issubclass(AdminClusterResourceAttributes, dict):
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
        if not isinstance(other, AdminClusterResourceAttributes):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

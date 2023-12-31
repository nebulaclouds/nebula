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


class CoreContainerPort(object):
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
        'container_port': 'int'
    }

    attribute_map = {
        'container_port': 'container_port'
    }

    def __init__(self, container_port=None):  # noqa: E501
        """CoreContainerPort - a model defined in Swagger"""  # noqa: E501

        self._container_port = None
        self.discriminator = None

        if container_port is not None:
            self.container_port = container_port

    @property
    def container_port(self):
        """Gets the container_port of this CoreContainerPort.  # noqa: E501

        Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.  # noqa: E501

        :return: The container_port of this CoreContainerPort.  # noqa: E501
        :rtype: int
        """
        return self._container_port

    @container_port.setter
    def container_port(self, container_port):
        """Sets the container_port of this CoreContainerPort.

        Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.  # noqa: E501

        :param container_port: The container_port of this CoreContainerPort.  # noqa: E501
        :type: int
        """

        self._container_port = container_port

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
        if issubclass(CoreContainerPort, dict):
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
        if not isinstance(other, CoreContainerPort):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

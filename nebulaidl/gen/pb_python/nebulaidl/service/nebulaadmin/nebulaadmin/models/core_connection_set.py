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

from nebulaadmin.models.connection_set_id_list import ConnectionSetIdList  # noqa: F401,E501


class CoreConnectionSet(object):
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
        'downstream': 'dict(str, ConnectionSetIdList)',
        'upstream': 'dict(str, ConnectionSetIdList)'
    }

    attribute_map = {
        'downstream': 'downstream',
        'upstream': 'upstream'
    }

    def __init__(self, downstream=None, upstream=None):  # noqa: E501
        """CoreConnectionSet - a model defined in Swagger"""  # noqa: E501

        self._downstream = None
        self._upstream = None
        self.discriminator = None

        if downstream is not None:
            self.downstream = downstream
        if upstream is not None:
            self.upstream = upstream

    @property
    def downstream(self):
        """Gets the downstream of this CoreConnectionSet.  # noqa: E501


        :return: The downstream of this CoreConnectionSet.  # noqa: E501
        :rtype: dict(str, ConnectionSetIdList)
        """
        return self._downstream

    @downstream.setter
    def downstream(self, downstream):
        """Sets the downstream of this CoreConnectionSet.


        :param downstream: The downstream of this CoreConnectionSet.  # noqa: E501
        :type: dict(str, ConnectionSetIdList)
        """

        self._downstream = downstream

    @property
    def upstream(self):
        """Gets the upstream of this CoreConnectionSet.  # noqa: E501


        :return: The upstream of this CoreConnectionSet.  # noqa: E501
        :rtype: dict(str, ConnectionSetIdList)
        """
        return self._upstream

    @upstream.setter
    def upstream(self, upstream):
        """Sets the upstream of this CoreConnectionSet.


        :param upstream: The upstream of this CoreConnectionSet.  # noqa: E501
        :type: dict(str, ConnectionSetIdList)
        """

        self._upstream = upstream

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
        if issubclass(CoreConnectionSet, dict):
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
        if not isinstance(other, CoreConnectionSet):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

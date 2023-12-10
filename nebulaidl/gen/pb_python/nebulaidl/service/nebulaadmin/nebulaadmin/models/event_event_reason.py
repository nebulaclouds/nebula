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


class EventEventReason(object):
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
        'reason': 'str',
        'occurred_at': 'datetime'
    }

    attribute_map = {
        'reason': 'reason',
        'occurred_at': 'occurred_at'
    }

    def __init__(self, reason=None, occurred_at=None):  # noqa: E501
        """EventEventReason - a model defined in Swagger"""  # noqa: E501

        self._reason = None
        self._occurred_at = None
        self.discriminator = None

        if reason is not None:
            self.reason = reason
        if occurred_at is not None:
            self.occurred_at = occurred_at

    @property
    def reason(self):
        """Gets the reason of this EventEventReason.  # noqa: E501


        :return: The reason of this EventEventReason.  # noqa: E501
        :rtype: str
        """
        return self._reason

    @reason.setter
    def reason(self, reason):
        """Sets the reason of this EventEventReason.


        :param reason: The reason of this EventEventReason.  # noqa: E501
        :type: str
        """

        self._reason = reason

    @property
    def occurred_at(self):
        """Gets the occurred_at of this EventEventReason.  # noqa: E501


        :return: The occurred_at of this EventEventReason.  # noqa: E501
        :rtype: datetime
        """
        return self._occurred_at

    @occurred_at.setter
    def occurred_at(self, occurred_at):
        """Sets the occurred_at of this EventEventReason.


        :param occurred_at: The occurred_at of this EventEventReason.  # noqa: E501
        :type: datetime
        """

        self._occurred_at = occurred_at

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
        if issubclass(EventEventReason, dict):
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
        if not isinstance(other, EventEventReason):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
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


class AdminSlackNotification(object):
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
        'recipients_email': 'list[str]'
    }

    attribute_map = {
        'recipients_email': 'recipients_email'
    }

    def __init__(self, recipients_email=None):  # noqa: E501
        """AdminSlackNotification - a model defined in Swagger"""  # noqa: E501

        self._recipients_email = None
        self.discriminator = None

        if recipients_email is not None:
            self.recipients_email = recipients_email

    @property
    def recipients_email(self):
        """Gets the recipients_email of this AdminSlackNotification.  # noqa: E501


        :return: The recipients_email of this AdminSlackNotification.  # noqa: E501
        :rtype: list[str]
        """
        return self._recipients_email

    @recipients_email.setter
    def recipients_email(self, recipients_email):
        """Sets the recipients_email of this AdminSlackNotification.


        :param recipients_email: The recipients_email of this AdminSlackNotification.  # noqa: E501
        :type: list[str]
        """

        self._recipients_email = recipients_email

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
        if issubclass(AdminSlackNotification, dict):
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
        if not isinstance(other, AdminSlackNotification):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

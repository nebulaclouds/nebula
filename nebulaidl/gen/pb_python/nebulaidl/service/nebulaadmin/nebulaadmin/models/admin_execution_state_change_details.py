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

from nebulaadmin.models.admin_execution_state import AdminExecutionState  # noqa: F401,E501


class AdminExecutionStateChangeDetails(object):
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
        'state': 'AdminExecutionState',
        'occurred_at': 'datetime',
        'principal': 'str'
    }

    attribute_map = {
        'state': 'state',
        'occurred_at': 'occurred_at',
        'principal': 'principal'
    }

    def __init__(self, state=None, occurred_at=None, principal=None):  # noqa: E501
        """AdminExecutionStateChangeDetails - a model defined in Swagger"""  # noqa: E501

        self._state = None
        self._occurred_at = None
        self._principal = None
        self.discriminator = None

        if state is not None:
            self.state = state
        if occurred_at is not None:
            self.occurred_at = occurred_at
        if principal is not None:
            self.principal = principal

    @property
    def state(self):
        """Gets the state of this AdminExecutionStateChangeDetails.  # noqa: E501

        The state of the execution is used to control its visibility in the UI/CLI.  # noqa: E501

        :return: The state of this AdminExecutionStateChangeDetails.  # noqa: E501
        :rtype: AdminExecutionState
        """
        return self._state

    @state.setter
    def state(self, state):
        """Sets the state of this AdminExecutionStateChangeDetails.

        The state of the execution is used to control its visibility in the UI/CLI.  # noqa: E501

        :param state: The state of this AdminExecutionStateChangeDetails.  # noqa: E501
        :type: AdminExecutionState
        """

        self._state = state

    @property
    def occurred_at(self):
        """Gets the occurred_at of this AdminExecutionStateChangeDetails.  # noqa: E501

        This timestamp represents when the state changed.  # noqa: E501

        :return: The occurred_at of this AdminExecutionStateChangeDetails.  # noqa: E501
        :rtype: datetime
        """
        return self._occurred_at

    @occurred_at.setter
    def occurred_at(self, occurred_at):
        """Sets the occurred_at of this AdminExecutionStateChangeDetails.

        This timestamp represents when the state changed.  # noqa: E501

        :param occurred_at: The occurred_at of this AdminExecutionStateChangeDetails.  # noqa: E501
        :type: datetime
        """

        self._occurred_at = occurred_at

    @property
    def principal(self):
        """Gets the principal of this AdminExecutionStateChangeDetails.  # noqa: E501


        :return: The principal of this AdminExecutionStateChangeDetails.  # noqa: E501
        :rtype: str
        """
        return self._principal

    @principal.setter
    def principal(self, principal):
        """Sets the principal of this AdminExecutionStateChangeDetails.


        :param principal: The principal of this AdminExecutionStateChangeDetails.  # noqa: E501
        :type: str
        """

        self._principal = principal

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
        if issubclass(AdminExecutionStateChangeDetails, dict):
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
        if not isinstance(other, AdminExecutionStateChangeDetails):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

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

from nebulaadmin.models.core_workflow_execution_identifier import CoreWorkflowExecutionIdentifier  # noqa: F401,E501


class CoreNodeExecutionIdentifier(object):
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
        'node_id': 'str',
        'execution_id': 'CoreWorkflowExecutionIdentifier'
    }

    attribute_map = {
        'node_id': 'node_id',
        'execution_id': 'execution_id'
    }

    def __init__(self, node_id=None, execution_id=None):  # noqa: E501
        """CoreNodeExecutionIdentifier - a model defined in Swagger"""  # noqa: E501

        self._node_id = None
        self._execution_id = None
        self.discriminator = None

        if node_id is not None:
            self.node_id = node_id
        if execution_id is not None:
            self.execution_id = execution_id

    @property
    def node_id(self):
        """Gets the node_id of this CoreNodeExecutionIdentifier.  # noqa: E501


        :return: The node_id of this CoreNodeExecutionIdentifier.  # noqa: E501
        :rtype: str
        """
        return self._node_id

    @node_id.setter
    def node_id(self, node_id):
        """Sets the node_id of this CoreNodeExecutionIdentifier.


        :param node_id: The node_id of this CoreNodeExecutionIdentifier.  # noqa: E501
        :type: str
        """

        self._node_id = node_id

    @property
    def execution_id(self):
        """Gets the execution_id of this CoreNodeExecutionIdentifier.  # noqa: E501


        :return: The execution_id of this CoreNodeExecutionIdentifier.  # noqa: E501
        :rtype: CoreWorkflowExecutionIdentifier
        """
        return self._execution_id

    @execution_id.setter
    def execution_id(self, execution_id):
        """Sets the execution_id of this CoreNodeExecutionIdentifier.


        :param execution_id: The execution_id of this CoreNodeExecutionIdentifier.  # noqa: E501
        :type: CoreWorkflowExecutionIdentifier
        """

        self._execution_id = execution_id

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
        if issubclass(CoreNodeExecutionIdentifier, dict):
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
        if not isinstance(other, CoreNodeExecutionIdentifier):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

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

from nebulaadmin.models.core_error import CoreError  # noqa: F401,E501
from nebulaadmin.models.core_if_block import CoreIfBlock  # noqa: F401,E501
from nebulaadmin.models.core_node import CoreNode  # noqa: F401,E501


class CoreIfElseBlock(object):
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
        'case': 'CoreIfBlock',
        'other': 'list[CoreIfBlock]',
        'else_node': 'CoreNode',
        'error': 'CoreError'
    }

    attribute_map = {
        'case': 'case',
        'other': 'other',
        'else_node': 'else_node',
        'error': 'error'
    }

    def __init__(self, case=None, other=None, else_node=None, error=None):  # noqa: E501
        """CoreIfElseBlock - a model defined in Swagger"""  # noqa: E501

        self._case = None
        self._other = None
        self._else_node = None
        self._error = None
        self.discriminator = None

        if case is not None:
            self.case = case
        if other is not None:
            self.other = other
        if else_node is not None:
            self.else_node = else_node
        if error is not None:
            self.error = error

    @property
    def case(self):
        """Gets the case of this CoreIfElseBlock.  # noqa: E501

        +required. First condition to evaluate.  # noqa: E501

        :return: The case of this CoreIfElseBlock.  # noqa: E501
        :rtype: CoreIfBlock
        """
        return self._case

    @case.setter
    def case(self, case):
        """Sets the case of this CoreIfElseBlock.

        +required. First condition to evaluate.  # noqa: E501

        :param case: The case of this CoreIfElseBlock.  # noqa: E501
        :type: CoreIfBlock
        """

        self._case = case

    @property
    def other(self):
        """Gets the other of this CoreIfElseBlock.  # noqa: E501

        +optional. Additional branches to evaluate.  # noqa: E501

        :return: The other of this CoreIfElseBlock.  # noqa: E501
        :rtype: list[CoreIfBlock]
        """
        return self._other

    @other.setter
    def other(self, other):
        """Sets the other of this CoreIfElseBlock.

        +optional. Additional branches to evaluate.  # noqa: E501

        :param other: The other of this CoreIfElseBlock.  # noqa: E501
        :type: list[CoreIfBlock]
        """

        self._other = other

    @property
    def else_node(self):
        """Gets the else_node of this CoreIfElseBlock.  # noqa: E501

        The node to execute in case none of the branches were taken.  # noqa: E501

        :return: The else_node of this CoreIfElseBlock.  # noqa: E501
        :rtype: CoreNode
        """
        return self._else_node

    @else_node.setter
    def else_node(self, else_node):
        """Sets the else_node of this CoreIfElseBlock.

        The node to execute in case none of the branches were taken.  # noqa: E501

        :param else_node: The else_node of this CoreIfElseBlock.  # noqa: E501
        :type: CoreNode
        """

        self._else_node = else_node

    @property
    def error(self):
        """Gets the error of this CoreIfElseBlock.  # noqa: E501

        An error to throw in case none of the branches were taken.  # noqa: E501

        :return: The error of this CoreIfElseBlock.  # noqa: E501
        :rtype: CoreError
        """
        return self._error

    @error.setter
    def error(self, error):
        """Sets the error of this CoreIfElseBlock.

        An error to throw in case none of the branches were taken.  # noqa: E501

        :param error: The error of this CoreIfElseBlock.  # noqa: E501
        :type: CoreError
        """

        self._error = error

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
        if issubclass(CoreIfElseBlock, dict):
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
        if not isinstance(other, CoreIfElseBlock):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

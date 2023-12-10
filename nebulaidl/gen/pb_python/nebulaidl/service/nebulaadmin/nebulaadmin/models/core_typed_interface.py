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

from nebulaadmin.models.core_variable_map import CoreVariableMap  # noqa: F401,E501


class CoreTypedInterface(object):
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
        'inputs': 'CoreVariableMap',
        'outputs': 'CoreVariableMap'
    }

    attribute_map = {
        'inputs': 'inputs',
        'outputs': 'outputs'
    }

    def __init__(self, inputs=None, outputs=None):  # noqa: E501
        """CoreTypedInterface - a model defined in Swagger"""  # noqa: E501

        self._inputs = None
        self._outputs = None
        self.discriminator = None

        if inputs is not None:
            self.inputs = inputs
        if outputs is not None:
            self.outputs = outputs

    @property
    def inputs(self):
        """Gets the inputs of this CoreTypedInterface.  # noqa: E501


        :return: The inputs of this CoreTypedInterface.  # noqa: E501
        :rtype: CoreVariableMap
        """
        return self._inputs

    @inputs.setter
    def inputs(self, inputs):
        """Sets the inputs of this CoreTypedInterface.


        :param inputs: The inputs of this CoreTypedInterface.  # noqa: E501
        :type: CoreVariableMap
        """

        self._inputs = inputs

    @property
    def outputs(self):
        """Gets the outputs of this CoreTypedInterface.  # noqa: E501


        :return: The outputs of this CoreTypedInterface.  # noqa: E501
        :rtype: CoreVariableMap
        """
        return self._outputs

    @outputs.setter
    def outputs(self, outputs):
        """Sets the outputs of this CoreTypedInterface.


        :param outputs: The outputs of this CoreTypedInterface.  # noqa: E501
        :type: CoreVariableMap
        """

        self._outputs = outputs

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
        if issubclass(CoreTypedInterface, dict):
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
        if not isinstance(other, CoreTypedInterface):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

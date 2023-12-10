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

from nebulaadmin.models.core_compiled_task import CoreCompiledTask  # noqa: F401,E501
from nebulaadmin.models.core_compiled_workflow import CoreCompiledWorkflow  # noqa: F401,E501


class CoreCompiledWorkflowClosure(object):
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
        'primary': 'CoreCompiledWorkflow',
        'sub_workflows': 'list[CoreCompiledWorkflow]',
        'tasks': 'list[CoreCompiledTask]'
    }

    attribute_map = {
        'primary': 'primary',
        'sub_workflows': 'sub_workflows',
        'tasks': 'tasks'
    }

    def __init__(self, primary=None, sub_workflows=None, tasks=None):  # noqa: E501
        """CoreCompiledWorkflowClosure - a model defined in Swagger"""  # noqa: E501

        self._primary = None
        self._sub_workflows = None
        self._tasks = None
        self.discriminator = None

        if primary is not None:
            self.primary = primary
        if sub_workflows is not None:
            self.sub_workflows = sub_workflows
        if tasks is not None:
            self.tasks = tasks

    @property
    def primary(self):
        """Gets the primary of this CoreCompiledWorkflowClosure.  # noqa: E501


        :return: The primary of this CoreCompiledWorkflowClosure.  # noqa: E501
        :rtype: CoreCompiledWorkflow
        """
        return self._primary

    @primary.setter
    def primary(self, primary):
        """Sets the primary of this CoreCompiledWorkflowClosure.


        :param primary: The primary of this CoreCompiledWorkflowClosure.  # noqa: E501
        :type: CoreCompiledWorkflow
        """

        self._primary = primary

    @property
    def sub_workflows(self):
        """Gets the sub_workflows of this CoreCompiledWorkflowClosure.  # noqa: E501


        :return: The sub_workflows of this CoreCompiledWorkflowClosure.  # noqa: E501
        :rtype: list[CoreCompiledWorkflow]
        """
        return self._sub_workflows

    @sub_workflows.setter
    def sub_workflows(self, sub_workflows):
        """Sets the sub_workflows of this CoreCompiledWorkflowClosure.


        :param sub_workflows: The sub_workflows of this CoreCompiledWorkflowClosure.  # noqa: E501
        :type: list[CoreCompiledWorkflow]
        """

        self._sub_workflows = sub_workflows

    @property
    def tasks(self):
        """Gets the tasks of this CoreCompiledWorkflowClosure.  # noqa: E501


        :return: The tasks of this CoreCompiledWorkflowClosure.  # noqa: E501
        :rtype: list[CoreCompiledTask]
        """
        return self._tasks

    @tasks.setter
    def tasks(self, tasks):
        """Sets the tasks of this CoreCompiledWorkflowClosure.


        :param tasks: The tasks of this CoreCompiledWorkflowClosure.  # noqa: E501
        :type: list[CoreCompiledTask]
        """

        self._tasks = tasks

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
        if issubclass(CoreCompiledWorkflowClosure, dict):
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
        if not isinstance(other, CoreCompiledWorkflowClosure):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
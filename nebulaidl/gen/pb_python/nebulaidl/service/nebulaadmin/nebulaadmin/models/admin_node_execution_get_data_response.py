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

from nebulaadmin.models.admin_nebula_ur_ls import AdminNebulaURLs  # noqa: F401,E501
from nebulaadmin.models.admin_url_blob import AdminUrlBlob  # noqa: F401,E501
from nebulaadmin.models.core_literal_map import CoreLiteralMap  # noqa: F401,E501
from nebulaadmin.models.nebulaidladmin_dynamic_workflow_node_metadata import NebulaidladminDynamicWorkflowNodeMetadata  # noqa: F401,E501


class AdminNodeExecutionGetDataResponse(object):
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
        'inputs': 'AdminUrlBlob',
        'outputs': 'AdminUrlBlob',
        'full_inputs': 'CoreLiteralMap',
        'full_outputs': 'CoreLiteralMap',
        'dynamic_workflow': 'NebulaidladminDynamicWorkflowNodeMetadata',
        'nebula_urls': 'AdminNebulaURLs'
    }

    attribute_map = {
        'inputs': 'inputs',
        'outputs': 'outputs',
        'full_inputs': 'full_inputs',
        'full_outputs': 'full_outputs',
        'dynamic_workflow': 'dynamic_workflow',
        'nebula_urls': 'nebula_urls'
    }

    def __init__(self, inputs=None, outputs=None, full_inputs=None, full_outputs=None, dynamic_workflow=None, nebula_urls=None):  # noqa: E501
        """AdminNodeExecutionGetDataResponse - a model defined in Swagger"""  # noqa: E501

        self._inputs = None
        self._outputs = None
        self._full_inputs = None
        self._full_outputs = None
        self._dynamic_workflow = None
        self._nebula_urls = None
        self.discriminator = None

        if inputs is not None:
            self.inputs = inputs
        if outputs is not None:
            self.outputs = outputs
        if full_inputs is not None:
            self.full_inputs = full_inputs
        if full_outputs is not None:
            self.full_outputs = full_outputs
        if dynamic_workflow is not None:
            self.dynamic_workflow = dynamic_workflow
        if nebula_urls is not None:
            self.nebula_urls = nebula_urls

    @property
    def inputs(self):
        """Gets the inputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501

        Signed url to fetch a core.LiteralMap of node execution inputs. Deprecated: Please use full_inputs instead.  # noqa: E501

        :return: The inputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :rtype: AdminUrlBlob
        """
        return self._inputs

    @inputs.setter
    def inputs(self, inputs):
        """Sets the inputs of this AdminNodeExecutionGetDataResponse.

        Signed url to fetch a core.LiteralMap of node execution inputs. Deprecated: Please use full_inputs instead.  # noqa: E501

        :param inputs: The inputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :type: AdminUrlBlob
        """

        self._inputs = inputs

    @property
    def outputs(self):
        """Gets the outputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501

        Signed url to fetch a core.LiteralMap of node execution outputs. Deprecated: Please use full_outputs instead.  # noqa: E501

        :return: The outputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :rtype: AdminUrlBlob
        """
        return self._outputs

    @outputs.setter
    def outputs(self, outputs):
        """Sets the outputs of this AdminNodeExecutionGetDataResponse.

        Signed url to fetch a core.LiteralMap of node execution outputs. Deprecated: Please use full_outputs instead.  # noqa: E501

        :param outputs: The outputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :type: AdminUrlBlob
        """

        self._outputs = outputs

    @property
    def full_inputs(self):
        """Gets the full_inputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501

        Full_inputs will only be populated if they are under a configured size threshold.  # noqa: E501

        :return: The full_inputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :rtype: CoreLiteralMap
        """
        return self._full_inputs

    @full_inputs.setter
    def full_inputs(self, full_inputs):
        """Sets the full_inputs of this AdminNodeExecutionGetDataResponse.

        Full_inputs will only be populated if they are under a configured size threshold.  # noqa: E501

        :param full_inputs: The full_inputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :type: CoreLiteralMap
        """

        self._full_inputs = full_inputs

    @property
    def full_outputs(self):
        """Gets the full_outputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501

        Full_outputs will only be populated if they are under a configured size threshold.  # noqa: E501

        :return: The full_outputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :rtype: CoreLiteralMap
        """
        return self._full_outputs

    @full_outputs.setter
    def full_outputs(self, full_outputs):
        """Sets the full_outputs of this AdminNodeExecutionGetDataResponse.

        Full_outputs will only be populated if they are under a configured size threshold.  # noqa: E501

        :param full_outputs: The full_outputs of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :type: CoreLiteralMap
        """

        self._full_outputs = full_outputs

    @property
    def dynamic_workflow(self):
        """Gets the dynamic_workflow of this AdminNodeExecutionGetDataResponse.  # noqa: E501

        Optional Workflow closure for a dynamically generated workflow, in the case this node yields a dynamic workflow we return its structure here.  # noqa: E501

        :return: The dynamic_workflow of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :rtype: NebulaidladminDynamicWorkflowNodeMetadata
        """
        return self._dynamic_workflow

    @dynamic_workflow.setter
    def dynamic_workflow(self, dynamic_workflow):
        """Sets the dynamic_workflow of this AdminNodeExecutionGetDataResponse.

        Optional Workflow closure for a dynamically generated workflow, in the case this node yields a dynamic workflow we return its structure here.  # noqa: E501

        :param dynamic_workflow: The dynamic_workflow of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :type: NebulaidladminDynamicWorkflowNodeMetadata
        """

        self._dynamic_workflow = dynamic_workflow

    @property
    def nebula_urls(self):
        """Gets the nebula_urls of this AdminNodeExecutionGetDataResponse.  # noqa: E501


        :return: The nebula_urls of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :rtype: AdminNebulaURLs
        """
        return self._nebula_urls

    @nebula_urls.setter
    def nebula_urls(self, nebula_urls):
        """Sets the nebula_urls of this AdminNodeExecutionGetDataResponse.


        :param nebula_urls: The nebula_urls of this AdminNodeExecutionGetDataResponse.  # noqa: E501
        :type: AdminNebulaURLs
        """

        self._nebula_urls = nebula_urls

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
        if issubclass(AdminNodeExecutionGetDataResponse, dict):
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
        if not isinstance(other, AdminNodeExecutionGetDataResponse):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
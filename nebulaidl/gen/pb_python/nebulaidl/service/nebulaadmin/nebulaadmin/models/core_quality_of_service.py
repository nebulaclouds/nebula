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

from nebulaadmin.models.core_quality_of_service_spec import CoreQualityOfServiceSpec  # noqa: F401,E501
from nebulaadmin.models.quality_of_service_tier import QualityOfServiceTier  # noqa: F401,E501


class CoreQualityOfService(object):
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
        'tier': 'QualityOfServiceTier',
        'spec': 'CoreQualityOfServiceSpec'
    }

    attribute_map = {
        'tier': 'tier',
        'spec': 'spec'
    }

    def __init__(self, tier=None, spec=None):  # noqa: E501
        """CoreQualityOfService - a model defined in Swagger"""  # noqa: E501

        self._tier = None
        self._spec = None
        self.discriminator = None

        if tier is not None:
            self.tier = tier
        if spec is not None:
            self.spec = spec

    @property
    def tier(self):
        """Gets the tier of this CoreQualityOfService.  # noqa: E501


        :return: The tier of this CoreQualityOfService.  # noqa: E501
        :rtype: QualityOfServiceTier
        """
        return self._tier

    @tier.setter
    def tier(self, tier):
        """Sets the tier of this CoreQualityOfService.


        :param tier: The tier of this CoreQualityOfService.  # noqa: E501
        :type: QualityOfServiceTier
        """

        self._tier = tier

    @property
    def spec(self):
        """Gets the spec of this CoreQualityOfService.  # noqa: E501


        :return: The spec of this CoreQualityOfService.  # noqa: E501
        :rtype: CoreQualityOfServiceSpec
        """
        return self._spec

    @spec.setter
    def spec(self, spec):
        """Sets the spec of this CoreQualityOfService.


        :param spec: The spec of this CoreQualityOfService.  # noqa: E501
        :type: CoreQualityOfServiceSpec
        """

        self._spec = spec

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
        if issubclass(CoreQualityOfService, dict):
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
        if not isinstance(other, CoreQualityOfService):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

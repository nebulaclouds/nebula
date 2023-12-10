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


class AdminVersion(object):
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
        'build': 'str',
        'version': 'str',
        'build_time': 'str'
    }

    attribute_map = {
        'build': 'Build',
        'version': 'Version',
        'build_time': 'BuildTime'
    }

    def __init__(self, build=None, version=None, build_time=None):  # noqa: E501
        """AdminVersion - a model defined in Swagger"""  # noqa: E501

        self._build = None
        self._version = None
        self._build_time = None
        self.discriminator = None

        if build is not None:
            self.build = build
        if version is not None:
            self.version = version
        if build_time is not None:
            self.build_time = build_time

    @property
    def build(self):
        """Gets the build of this AdminVersion.  # noqa: E501


        :return: The build of this AdminVersion.  # noqa: E501
        :rtype: str
        """
        return self._build

    @build.setter
    def build(self, build):
        """Sets the build of this AdminVersion.


        :param build: The build of this AdminVersion.  # noqa: E501
        :type: str
        """

        self._build = build

    @property
    def version(self):
        """Gets the version of this AdminVersion.  # noqa: E501


        :return: The version of this AdminVersion.  # noqa: E501
        :rtype: str
        """
        return self._version

    @version.setter
    def version(self, version):
        """Sets the version of this AdminVersion.


        :param version: The version of this AdminVersion.  # noqa: E501
        :type: str
        """

        self._version = version

    @property
    def build_time(self):
        """Gets the build_time of this AdminVersion.  # noqa: E501


        :return: The build_time of this AdminVersion.  # noqa: E501
        :rtype: str
        """
        return self._build_time

    @build_time.setter
    def build_time(self, build_time):
        """Sets the build_time of this AdminVersion.


        :param build_time: The build_time of this AdminVersion.  # noqa: E501
        :type: str
        """

        self._build_time = build_time

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
        if issubclass(AdminVersion, dict):
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
        if not isinstance(other, AdminVersion):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

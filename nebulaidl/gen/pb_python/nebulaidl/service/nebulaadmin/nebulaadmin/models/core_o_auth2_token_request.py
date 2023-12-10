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

from nebulaadmin.models.core_o_auth2_client import CoreOAuth2Client  # noqa: F401,E501
from nebulaadmin.models.core_o_auth2_token_request_type import CoreOAuth2TokenRequestType  # noqa: F401,E501


class CoreOAuth2TokenRequest(object):
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
        'name': 'str',
        'type': 'CoreOAuth2TokenRequestType',
        'client': 'CoreOAuth2Client',
        'idp_discovery_endpoint': 'str',
        'token_endpoint': 'str'
    }

    attribute_map = {
        'name': 'name',
        'type': 'type',
        'client': 'client',
        'idp_discovery_endpoint': 'idp_discovery_endpoint',
        'token_endpoint': 'token_endpoint'
    }

    def __init__(self, name=None, type=None, client=None, idp_discovery_endpoint=None, token_endpoint=None):  # noqa: E501
        """CoreOAuth2TokenRequest - a model defined in Swagger"""  # noqa: E501

        self._name = None
        self._type = None
        self._client = None
        self._idp_discovery_endpoint = None
        self._token_endpoint = None
        self.discriminator = None

        if name is not None:
            self.name = name
        if type is not None:
            self.type = type
        if client is not None:
            self.client = client
        if idp_discovery_endpoint is not None:
            self.idp_discovery_endpoint = idp_discovery_endpoint
        if token_endpoint is not None:
            self.token_endpoint = token_endpoint

    @property
    def name(self):
        """Gets the name of this CoreOAuth2TokenRequest.  # noqa: E501


        :return: The name of this CoreOAuth2TokenRequest.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this CoreOAuth2TokenRequest.


        :param name: The name of this CoreOAuth2TokenRequest.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def type(self):
        """Gets the type of this CoreOAuth2TokenRequest.  # noqa: E501


        :return: The type of this CoreOAuth2TokenRequest.  # noqa: E501
        :rtype: CoreOAuth2TokenRequestType
        """
        return self._type

    @type.setter
    def type(self, type):
        """Sets the type of this CoreOAuth2TokenRequest.


        :param type: The type of this CoreOAuth2TokenRequest.  # noqa: E501
        :type: CoreOAuth2TokenRequestType
        """

        self._type = type

    @property
    def client(self):
        """Gets the client of this CoreOAuth2TokenRequest.  # noqa: E501


        :return: The client of this CoreOAuth2TokenRequest.  # noqa: E501
        :rtype: CoreOAuth2Client
        """
        return self._client

    @client.setter
    def client(self, client):
        """Sets the client of this CoreOAuth2TokenRequest.


        :param client: The client of this CoreOAuth2TokenRequest.  # noqa: E501
        :type: CoreOAuth2Client
        """

        self._client = client

    @property
    def idp_discovery_endpoint(self):
        """Gets the idp_discovery_endpoint of this CoreOAuth2TokenRequest.  # noqa: E501


        :return: The idp_discovery_endpoint of this CoreOAuth2TokenRequest.  # noqa: E501
        :rtype: str
        """
        return self._idp_discovery_endpoint

    @idp_discovery_endpoint.setter
    def idp_discovery_endpoint(self, idp_discovery_endpoint):
        """Sets the idp_discovery_endpoint of this CoreOAuth2TokenRequest.


        :param idp_discovery_endpoint: The idp_discovery_endpoint of this CoreOAuth2TokenRequest.  # noqa: E501
        :type: str
        """

        self._idp_discovery_endpoint = idp_discovery_endpoint

    @property
    def token_endpoint(self):
        """Gets the token_endpoint of this CoreOAuth2TokenRequest.  # noqa: E501


        :return: The token_endpoint of this CoreOAuth2TokenRequest.  # noqa: E501
        :rtype: str
        """
        return self._token_endpoint

    @token_endpoint.setter
    def token_endpoint(self, token_endpoint):
        """Sets the token_endpoint of this CoreOAuth2TokenRequest.


        :param token_endpoint: The token_endpoint of this CoreOAuth2TokenRequest.  # noqa: E501
        :type: str
        """

        self._token_endpoint = token_endpoint

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
        if issubclass(CoreOAuth2TokenRequest, dict):
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
        if not isinstance(other, CoreOAuth2TokenRequest):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

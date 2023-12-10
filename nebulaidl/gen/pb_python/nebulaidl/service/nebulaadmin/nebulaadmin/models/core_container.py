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

from nebulaadmin.models.container_architecture import ContainerArchitecture  # noqa: F401,E501
from nebulaadmin.models.core_container_port import CoreContainerPort  # noqa: F401,E501
from nebulaadmin.models.core_data_loading_config import CoreDataLoadingConfig  # noqa: F401,E501
from nebulaadmin.models.core_key_value_pair import CoreKeyValuePair  # noqa: F401,E501
from nebulaadmin.models.core_resources import CoreResources  # noqa: F401,E501


class CoreContainer(object):
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
        'image': 'str',
        'command': 'list[str]',
        'args': 'list[str]',
        'resources': 'CoreResources',
        'env': 'list[CoreKeyValuePair]',
        'config': 'list[CoreKeyValuePair]',
        'ports': 'list[CoreContainerPort]',
        'data_config': 'CoreDataLoadingConfig',
        'architecture': 'ContainerArchitecture'
    }

    attribute_map = {
        'image': 'image',
        'command': 'command',
        'args': 'args',
        'resources': 'resources',
        'env': 'env',
        'config': 'config',
        'ports': 'ports',
        'data_config': 'data_config',
        'architecture': 'architecture'
    }

    def __init__(self, image=None, command=None, args=None, resources=None, env=None, config=None, ports=None, data_config=None, architecture=None):  # noqa: E501
        """CoreContainer - a model defined in Swagger"""  # noqa: E501

        self._image = None
        self._command = None
        self._args = None
        self._resources = None
        self._env = None
        self._config = None
        self._ports = None
        self._data_config = None
        self._architecture = None
        self.discriminator = None

        if image is not None:
            self.image = image
        if command is not None:
            self.command = command
        if args is not None:
            self.args = args
        if resources is not None:
            self.resources = resources
        if env is not None:
            self.env = env
        if config is not None:
            self.config = config
        if ports is not None:
            self.ports = ports
        if data_config is not None:
            self.data_config = data_config
        if architecture is not None:
            self.architecture = architecture

    @property
    def image(self):
        """Gets the image of this CoreContainer.  # noqa: E501


        :return: The image of this CoreContainer.  # noqa: E501
        :rtype: str
        """
        return self._image

    @image.setter
    def image(self, image):
        """Sets the image of this CoreContainer.


        :param image: The image of this CoreContainer.  # noqa: E501
        :type: str
        """

        self._image = image

    @property
    def command(self):
        """Gets the command of this CoreContainer.  # noqa: E501

        Command to be executed, if not provided, the default entrypoint in the container image will be used.  # noqa: E501

        :return: The command of this CoreContainer.  # noqa: E501
        :rtype: list[str]
        """
        return self._command

    @command.setter
    def command(self, command):
        """Sets the command of this CoreContainer.

        Command to be executed, if not provided, the default entrypoint in the container image will be used.  # noqa: E501

        :param command: The command of this CoreContainer.  # noqa: E501
        :type: list[str]
        """

        self._command = command

    @property
    def args(self):
        """Gets the args of this CoreContainer.  # noqa: E501

        These will default to Nebula given paths. If provided, the system will not append known paths. If the task still needs nebula's inputs and outputs path, add $(NEBULA_INPUT_FILE), $(NEBULA_OUTPUT_FILE) wherever makes sense and the system will populate these before executing the container.  # noqa: E501

        :return: The args of this CoreContainer.  # noqa: E501
        :rtype: list[str]
        """
        return self._args

    @args.setter
    def args(self, args):
        """Sets the args of this CoreContainer.

        These will default to Nebula given paths. If provided, the system will not append known paths. If the task still needs nebula's inputs and outputs path, add $(NEBULA_INPUT_FILE), $(NEBULA_OUTPUT_FILE) wherever makes sense and the system will populate these before executing the container.  # noqa: E501

        :param args: The args of this CoreContainer.  # noqa: E501
        :type: list[str]
        """

        self._args = args

    @property
    def resources(self):
        """Gets the resources of this CoreContainer.  # noqa: E501

        Container resources requirement as specified by the container engine.  # noqa: E501

        :return: The resources of this CoreContainer.  # noqa: E501
        :rtype: CoreResources
        """
        return self._resources

    @resources.setter
    def resources(self, resources):
        """Sets the resources of this CoreContainer.

        Container resources requirement as specified by the container engine.  # noqa: E501

        :param resources: The resources of this CoreContainer.  # noqa: E501
        :type: CoreResources
        """

        self._resources = resources

    @property
    def env(self):
        """Gets the env of this CoreContainer.  # noqa: E501

        Environment variables will be set as the container is starting up.  # noqa: E501

        :return: The env of this CoreContainer.  # noqa: E501
        :rtype: list[CoreKeyValuePair]
        """
        return self._env

    @env.setter
    def env(self, env):
        """Sets the env of this CoreContainer.

        Environment variables will be set as the container is starting up.  # noqa: E501

        :param env: The env of this CoreContainer.  # noqa: E501
        :type: list[CoreKeyValuePair]
        """

        self._env = env

    @property
    def config(self):
        """Gets the config of this CoreContainer.  # noqa: E501

        Allows extra configs to be available for the container. TODO: elaborate on how configs will become available. Deprecated, please use TaskTemplate.config instead.  # noqa: E501

        :return: The config of this CoreContainer.  # noqa: E501
        :rtype: list[CoreKeyValuePair]
        """
        return self._config

    @config.setter
    def config(self, config):
        """Sets the config of this CoreContainer.

        Allows extra configs to be available for the container. TODO: elaborate on how configs will become available. Deprecated, please use TaskTemplate.config instead.  # noqa: E501

        :param config: The config of this CoreContainer.  # noqa: E501
        :type: list[CoreKeyValuePair]
        """

        self._config = config

    @property
    def ports(self):
        """Gets the ports of this CoreContainer.  # noqa: E501


        :return: The ports of this CoreContainer.  # noqa: E501
        :rtype: list[CoreContainerPort]
        """
        return self._ports

    @ports.setter
    def ports(self, ports):
        """Sets the ports of this CoreContainer.


        :param ports: The ports of this CoreContainer.  # noqa: E501
        :type: list[CoreContainerPort]
        """

        self._ports = ports

    @property
    def data_config(self):
        """Gets the data_config of this CoreContainer.  # noqa: E501


        :return: The data_config of this CoreContainer.  # noqa: E501
        :rtype: CoreDataLoadingConfig
        """
        return self._data_config

    @data_config.setter
    def data_config(self, data_config):
        """Sets the data_config of this CoreContainer.


        :param data_config: The data_config of this CoreContainer.  # noqa: E501
        :type: CoreDataLoadingConfig
        """

        self._data_config = data_config

    @property
    def architecture(self):
        """Gets the architecture of this CoreContainer.  # noqa: E501


        :return: The architecture of this CoreContainer.  # noqa: E501
        :rtype: ContainerArchitecture
        """
        return self._architecture

    @architecture.setter
    def architecture(self, architecture):
        """Sets the architecture of this CoreContainer.


        :param architecture: The architecture of this CoreContainer.  # noqa: E501
        :type: ContainerArchitecture
        """

        self._architecture = architecture

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
        if issubclass(CoreContainer, dict):
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
        if not isinstance(other, CoreContainer):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other

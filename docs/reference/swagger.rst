.. _reference-swagger:

#############################
Nebula API Playground: Swagger
#############################

.. tags:: Basic

Nebula services expose gRPC services for efficient/low latency communication across all services as well as for external clients (NebulaCTL, NebulaConsole, Nebulakit Remote, etc.).

The service definitions are defined `here <https://github.com/nebulaclouds/nebulaidl/tree/master/protos/nebulaidl/service>`__.
NebulaIDL also houses open API schema definitions for the exposed services:

- `Admin <https://github.com/nebulaclouds/nebulaidl/blob/master/gen/pb-go/nebulaidl/service/admin.swagger.json>`__
- `Auth <https://github.com/nebulaclouds/nebulaidl/blob/master/gen/pb-go/nebulaidl/service/auth.swagger.json>`__
- `Identity <https://github.com/nebulaclouds/nebulaidl/blob/master/gen/pb-go/nebulaidl/service/identity.swagger.json>`__

To view the UI, run the following command:

.. prompt:: bash $

   nebulactl demo start

Once sandbox setup is complete, a ready-to-explore message is shown:

.. prompt::

   👨‍💻 Nebula is ready! Nebula UI is available at http://localhost:30081/console 🚀 🚀 🎉


Visit ``http://localhost:30080/api/v1/openapi`` to view the swagger documentation of the payload fields.

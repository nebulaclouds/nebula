######################
Protocol Documentation
######################




.. _ref_nebulaidl/service/admin.proto:

nebulaidl/service/admin.proto
==================================================================




..
   end messages


..
   end enums


..
   end HasExtensions



.. _ref_nebulaidl.service.AdminService:

AdminService
------------------------------------------------------------------

The following defines an RPC service that is also served over HTTP via grpc-gateway.
Standard response codes for both are defined here: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go

.. csv-table:: AdminService service methods
   :header: "Method Name", "Request Type", "Response Type", "Description"
   :widths: auto

   "CreateTask", ":ref:`ref_nebulaidl.admin.TaskCreateRequest`", ":ref:`ref_nebulaidl.admin.TaskCreateResponse`", "Create and upload a :ref:`ref_nebulaidl.admin.Task` definition"
   "GetTask", ":ref:`ref_nebulaidl.admin.ObjectGetRequest`", ":ref:`ref_nebulaidl.admin.Task`", "Fetch a :ref:`ref_nebulaidl.admin.Task` definition."
   "ListTaskIds", ":ref:`ref_nebulaidl.admin.NamedEntityIdentifierListRequest`", ":ref:`ref_nebulaidl.admin.NamedEntityIdentifierList`", "Fetch a list of :ref:`ref_nebulaidl.admin.NamedEntityIdentifier` of task objects."
   "ListTasks", ":ref:`ref_nebulaidl.admin.ResourceListRequest`", ":ref:`ref_nebulaidl.admin.TaskList`", "Fetch a list of :ref:`ref_nebulaidl.admin.Task` definitions."
   "CreateWorkflow", ":ref:`ref_nebulaidl.admin.WorkflowCreateRequest`", ":ref:`ref_nebulaidl.admin.WorkflowCreateResponse`", "Create and upload a :ref:`ref_nebulaidl.admin.Workflow` definition"
   "GetWorkflow", ":ref:`ref_nebulaidl.admin.ObjectGetRequest`", ":ref:`ref_nebulaidl.admin.Workflow`", "Fetch a :ref:`ref_nebulaidl.admin.Workflow` definition."
   "ListWorkflowIds", ":ref:`ref_nebulaidl.admin.NamedEntityIdentifierListRequest`", ":ref:`ref_nebulaidl.admin.NamedEntityIdentifierList`", "Fetch a list of :ref:`ref_nebulaidl.admin.NamedEntityIdentifier` of workflow objects."
   "ListWorkflows", ":ref:`ref_nebulaidl.admin.ResourceListRequest`", ":ref:`ref_nebulaidl.admin.WorkflowList`", "Fetch a list of :ref:`ref_nebulaidl.admin.Workflow` definitions."
   "CreateLaunchPlan", ":ref:`ref_nebulaidl.admin.LaunchPlanCreateRequest`", ":ref:`ref_nebulaidl.admin.LaunchPlanCreateResponse`", "Create and upload a :ref:`ref_nebulaidl.admin.LaunchPlan` definition"
   "GetLaunchPlan", ":ref:`ref_nebulaidl.admin.ObjectGetRequest`", ":ref:`ref_nebulaidl.admin.LaunchPlan`", "Fetch a :ref:`ref_nebulaidl.admin.LaunchPlan` definition."
   "GetActiveLaunchPlan", ":ref:`ref_nebulaidl.admin.ActiveLaunchPlanRequest`", ":ref:`ref_nebulaidl.admin.LaunchPlan`", "Fetch the active version of a :ref:`ref_nebulaidl.admin.LaunchPlan`."
   "ListActiveLaunchPlans", ":ref:`ref_nebulaidl.admin.ActiveLaunchPlanListRequest`", ":ref:`ref_nebulaidl.admin.LaunchPlanList`", "List active versions of :ref:`ref_nebulaidl.admin.LaunchPlan`."
   "ListLaunchPlanIds", ":ref:`ref_nebulaidl.admin.NamedEntityIdentifierListRequest`", ":ref:`ref_nebulaidl.admin.NamedEntityIdentifierList`", "Fetch a list of :ref:`ref_nebulaidl.admin.NamedEntityIdentifier` of launch plan objects."
   "ListLaunchPlans", ":ref:`ref_nebulaidl.admin.ResourceListRequest`", ":ref:`ref_nebulaidl.admin.LaunchPlanList`", "Fetch a list of :ref:`ref_nebulaidl.admin.LaunchPlan` definitions."
   "UpdateLaunchPlan", ":ref:`ref_nebulaidl.admin.LaunchPlanUpdateRequest`", ":ref:`ref_nebulaidl.admin.LaunchPlanUpdateResponse`", "Updates the status of a registered :ref:`ref_nebulaidl.admin.LaunchPlan`."
   "CreateExecution", ":ref:`ref_nebulaidl.admin.ExecutionCreateRequest`", ":ref:`ref_nebulaidl.admin.ExecutionCreateResponse`", "Triggers the creation of a :ref:`ref_nebulaidl.admin.Execution`"
   "RelaunchExecution", ":ref:`ref_nebulaidl.admin.ExecutionRelaunchRequest`", ":ref:`ref_nebulaidl.admin.ExecutionCreateResponse`", "Triggers the creation of an identical :ref:`ref_nebulaidl.admin.Execution`"
   "RecoverExecution", ":ref:`ref_nebulaidl.admin.ExecutionRecoverRequest`", ":ref:`ref_nebulaidl.admin.ExecutionCreateResponse`", "Recreates a previously-run workflow execution that will only start executing from the last known failure point. In Recover mode, users cannot change any input parameters or update the version of the execution. This is extremely useful to recover from system errors and byzantine faults like - Loss of K8s cluster, bugs in platform or instability, machine failures, downstream system failures (downstream services), or simply to recover executions that failed because of retry exhaustion and should complete if tried again. See :ref:`ref_nebulaidl.admin.ExecutionRecoverRequest` for more details."
   "GetExecution", ":ref:`ref_nebulaidl.admin.WorkflowExecutionGetRequest`", ":ref:`ref_nebulaidl.admin.Execution`", "Fetches a :ref:`ref_nebulaidl.admin.Execution`."
   "UpdateExecution", ":ref:`ref_nebulaidl.admin.ExecutionUpdateRequest`", ":ref:`ref_nebulaidl.admin.ExecutionUpdateResponse`", "Update execution belonging to project domain :ref:`ref_nebulaidl.admin.Execution`."
   "GetExecutionData", ":ref:`ref_nebulaidl.admin.WorkflowExecutionGetDataRequest`", ":ref:`ref_nebulaidl.admin.WorkflowExecutionGetDataResponse`", "Fetches input and output data for a :ref:`ref_nebulaidl.admin.Execution`."
   "ListExecutions", ":ref:`ref_nebulaidl.admin.ResourceListRequest`", ":ref:`ref_nebulaidl.admin.ExecutionList`", "Fetch a list of :ref:`ref_nebulaidl.admin.Execution`."
   "TerminateExecution", ":ref:`ref_nebulaidl.admin.ExecutionTerminateRequest`", ":ref:`ref_nebulaidl.admin.ExecutionTerminateResponse`", "Terminates an in-progress :ref:`ref_nebulaidl.admin.Execution`."
   "GetNodeExecution", ":ref:`ref_nebulaidl.admin.NodeExecutionGetRequest`", ":ref:`ref_nebulaidl.admin.NodeExecution`", "Fetches a :ref:`ref_nebulaidl.admin.NodeExecution`."
   "ListNodeExecutions", ":ref:`ref_nebulaidl.admin.NodeExecutionListRequest`", ":ref:`ref_nebulaidl.admin.NodeExecutionList`", "Fetch a list of :ref:`ref_nebulaidl.admin.NodeExecution`."
   "ListNodeExecutionsForTask", ":ref:`ref_nebulaidl.admin.NodeExecutionForTaskListRequest`", ":ref:`ref_nebulaidl.admin.NodeExecutionList`", "Fetch a list of :ref:`ref_nebulaidl.admin.NodeExecution` launched by the reference :ref:`ref_nebulaidl.admin.TaskExecution`."
   "GetNodeExecutionData", ":ref:`ref_nebulaidl.admin.NodeExecutionGetDataRequest`", ":ref:`ref_nebulaidl.admin.NodeExecutionGetDataResponse`", "Fetches input and output data for a :ref:`ref_nebulaidl.admin.NodeExecution`."
   "RegisterProject", ":ref:`ref_nebulaidl.admin.ProjectRegisterRequest`", ":ref:`ref_nebulaidl.admin.ProjectRegisterResponse`", "Registers a :ref:`ref_nebulaidl.admin.Project` with the Nebula deployment."
   "UpdateProject", ":ref:`ref_nebulaidl.admin.Project`", ":ref:`ref_nebulaidl.admin.ProjectUpdateResponse`", "Updates an existing :ref:`ref_nebulaidl.admin.Project` nebulaidl.admin.Project should be passed but the domains property should be empty; it will be ignored in the handler as domains cannot be updated via this API."
   "ListProjects", ":ref:`ref_nebulaidl.admin.ProjectListRequest`", ":ref:`ref_nebulaidl.admin.Projects`", "Fetches a list of :ref:`ref_nebulaidl.admin.Project`"
   "CreateWorkflowEvent", ":ref:`ref_nebulaidl.admin.WorkflowExecutionEventRequest`", ":ref:`ref_nebulaidl.admin.WorkflowExecutionEventResponse`", "Indicates a :ref:`ref_nebulaidl.event.WorkflowExecutionEvent` has occurred."
   "CreateNodeEvent", ":ref:`ref_nebulaidl.admin.NodeExecutionEventRequest`", ":ref:`ref_nebulaidl.admin.NodeExecutionEventResponse`", "Indicates a :ref:`ref_nebulaidl.event.NodeExecutionEvent` has occurred."
   "CreateTaskEvent", ":ref:`ref_nebulaidl.admin.TaskExecutionEventRequest`", ":ref:`ref_nebulaidl.admin.TaskExecutionEventResponse`", "Indicates a :ref:`ref_nebulaidl.event.TaskExecutionEvent` has occurred."
   "GetTaskExecution", ":ref:`ref_nebulaidl.admin.TaskExecutionGetRequest`", ":ref:`ref_nebulaidl.admin.TaskExecution`", "Fetches a :ref:`ref_nebulaidl.admin.TaskExecution`."
   "ListTaskExecutions", ":ref:`ref_nebulaidl.admin.TaskExecutionListRequest`", ":ref:`ref_nebulaidl.admin.TaskExecutionList`", "Fetches a list of :ref:`ref_nebulaidl.admin.TaskExecution`."
   "GetTaskExecutionData", ":ref:`ref_nebulaidl.admin.TaskExecutionGetDataRequest`", ":ref:`ref_nebulaidl.admin.TaskExecutionGetDataResponse`", "Fetches input and output data for a :ref:`ref_nebulaidl.admin.TaskExecution`."
   "UpdateProjectDomainAttributes", ":ref:`ref_nebulaidl.admin.ProjectDomainAttributesUpdateRequest`", ":ref:`ref_nebulaidl.admin.ProjectDomainAttributesUpdateResponse`", "Creates or updates custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` for a project and domain."
   "GetProjectDomainAttributes", ":ref:`ref_nebulaidl.admin.ProjectDomainAttributesGetRequest`", ":ref:`ref_nebulaidl.admin.ProjectDomainAttributesGetResponse`", "Fetches custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` for a project and domain."
   "DeleteProjectDomainAttributes", ":ref:`ref_nebulaidl.admin.ProjectDomainAttributesDeleteRequest`", ":ref:`ref_nebulaidl.admin.ProjectDomainAttributesDeleteResponse`", "Deletes custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` for a project and domain."
   "UpdateProjectAttributes", ":ref:`ref_nebulaidl.admin.ProjectAttributesUpdateRequest`", ":ref:`ref_nebulaidl.admin.ProjectAttributesUpdateResponse`", "Creates or updates custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` at the project level"
   "GetProjectAttributes", ":ref:`ref_nebulaidl.admin.ProjectAttributesGetRequest`", ":ref:`ref_nebulaidl.admin.ProjectAttributesGetResponse`", "Fetches custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` for a project and domain."
   "DeleteProjectAttributes", ":ref:`ref_nebulaidl.admin.ProjectAttributesDeleteRequest`", ":ref:`ref_nebulaidl.admin.ProjectAttributesDeleteResponse`", "Deletes custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` for a project and domain."
   "UpdateWorkflowAttributes", ":ref:`ref_nebulaidl.admin.WorkflowAttributesUpdateRequest`", ":ref:`ref_nebulaidl.admin.WorkflowAttributesUpdateResponse`", "Creates or updates custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` for a project, domain and workflow."
   "GetWorkflowAttributes", ":ref:`ref_nebulaidl.admin.WorkflowAttributesGetRequest`", ":ref:`ref_nebulaidl.admin.WorkflowAttributesGetResponse`", "Fetches custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` for a project, domain and workflow."
   "DeleteWorkflowAttributes", ":ref:`ref_nebulaidl.admin.WorkflowAttributesDeleteRequest`", ":ref:`ref_nebulaidl.admin.WorkflowAttributesDeleteResponse`", "Deletes custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` for a project, domain and workflow."
   "ListMatchableAttributes", ":ref:`ref_nebulaidl.admin.ListMatchableAttributesRequest`", ":ref:`ref_nebulaidl.admin.ListMatchableAttributesResponse`", "Lists custom :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration` for a specific resource type."
   "ListNamedEntities", ":ref:`ref_nebulaidl.admin.NamedEntityListRequest`", ":ref:`ref_nebulaidl.admin.NamedEntityList`", "Returns a list of :ref:`ref_nebulaidl.admin.NamedEntity` objects."
   "GetNamedEntity", ":ref:`ref_nebulaidl.admin.NamedEntityGetRequest`", ":ref:`ref_nebulaidl.admin.NamedEntity`", "Returns a :ref:`ref_nebulaidl.admin.NamedEntity` object."
   "UpdateNamedEntity", ":ref:`ref_nebulaidl.admin.NamedEntityUpdateRequest`", ":ref:`ref_nebulaidl.admin.NamedEntityUpdateResponse`", "Updates a :ref:`ref_nebulaidl.admin.NamedEntity` object."
   "GetVersion", ":ref:`ref_nebulaidl.admin.GetVersionRequest`", ":ref:`ref_nebulaidl.admin.GetVersionResponse`", ""
   "GetDescriptionEntity", ":ref:`ref_nebulaidl.admin.ObjectGetRequest`", ":ref:`ref_nebulaidl.admin.DescriptionEntity`", "Fetch a :ref:`ref_nebulaidl.admin.DescriptionEntity` object."
   "ListDescriptionEntities", ":ref:`ref_nebulaidl.admin.DescriptionEntityListRequest`", ":ref:`ref_nebulaidl.admin.DescriptionEntityList`", "Fetch a list of :ref:`ref_nebulaidl.admin.DescriptionEntity` definitions."

..
   end services




.. _ref_nebulaidl/service/auth.proto:

nebulaidl/service/auth.proto
==================================================================





.. _ref_nebulaidl.service.OAuth2MetadataRequest:

OAuth2MetadataRequest
------------------------------------------------------------------










.. _ref_nebulaidl.service.OAuth2MetadataResponse:

OAuth2MetadataResponse
------------------------------------------------------------------

OAuth2MetadataResponse defines an RFC-Compliant response for /.well-known/oauth-authorization-server metadata
as defined in https://tools.ietf.org/html/rfc8414



.. csv-table:: OAuth2MetadataResponse type fields
   :header: "Field", "Type", "Label", "Description"
   :widths: auto

   "issuer", ":ref:`ref_string`", "", "Defines the issuer string in all JWT tokens this server issues. The issuer can be admin itself or an external issuer."
   "authorization_endpoint", ":ref:`ref_string`", "", "URL of the authorization server's authorization endpoint [RFC6749]. This is REQUIRED unless no grant types are supported that use the authorization endpoint."
   "token_endpoint", ":ref:`ref_string`", "", "URL of the authorization server's token endpoint [RFC6749]."
   "response_types_supported", ":ref:`ref_string`", "repeated", "Array containing a list of the OAuth 2.0 response_type values that this authorization server supports."
   "scopes_supported", ":ref:`ref_string`", "repeated", "JSON array containing a list of the OAuth 2.0 [RFC6749] scope values that this authorization server supports."
   "token_endpoint_auth_methods_supported", ":ref:`ref_string`", "repeated", "JSON array containing a list of client authentication methods supported by this token endpoint."
   "jwks_uri", ":ref:`ref_string`", "", "URL of the authorization server's JWK Set [JWK] document. The referenced document contains the signing key(s) the client uses to validate signatures from the authorization server."
   "code_challenge_methods_supported", ":ref:`ref_string`", "repeated", "JSON array containing a list of Proof Key for Code Exchange (PKCE) [RFC7636] code challenge methods supported by this authorization server."
   "grant_types_supported", ":ref:`ref_string`", "repeated", "JSON array containing a list of the OAuth 2.0 grant type values that this authorization server supports."
   "device_authorization_endpoint", ":ref:`ref_string`", "", "URL of the authorization server's device authorization endpoint, as defined in Section 3.1 of [RFC8628]"







.. _ref_nebulaidl.service.PublicClientAuthConfigRequest:

PublicClientAuthConfigRequest
------------------------------------------------------------------










.. _ref_nebulaidl.service.PublicClientAuthConfigResponse:

PublicClientAuthConfigResponse
------------------------------------------------------------------

NebulaClientResponse encapsulates public information that nebula clients (CLIs... etc.) can use to authenticate users.



.. csv-table:: PublicClientAuthConfigResponse type fields
   :header: "Field", "Type", "Label", "Description"
   :widths: auto

   "client_id", ":ref:`ref_string`", "", "client_id to use when initiating OAuth2 authorization requests."
   "redirect_uri", ":ref:`ref_string`", "", "redirect uri to use when initiating OAuth2 authorization requests."
   "scopes", ":ref:`ref_string`", "repeated", "scopes to request when initiating OAuth2 authorization requests."
   "authorization_metadata_key", ":ref:`ref_string`", "", "Authorization Header to use when passing Access Tokens to the server. If not provided, the client should use the default http `Authorization` header."
   "service_http_endpoint", ":ref:`ref_string`", "", "ServiceHttpEndpoint points to the http endpoint for the backend. If empty, clients can assume the endpoint used to configure the gRPC connection can be used for the http one respecting the insecure flag to choose between SSL or no SSL connections."
   "audience", ":ref:`ref_string`", "", "audience to use when initiating OAuth2 authorization requests."






..
   end messages


..
   end enums


..
   end HasExtensions



.. _ref_nebulaidl.service.AuthMetadataService:

AuthMetadataService
------------------------------------------------------------------

The following defines an RPC service that is also served over HTTP via grpc-gateway.
Standard response codes for both are defined here: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
RPCs defined in this service must be anonymously accessible.

.. csv-table:: AuthMetadataService service methods
   :header: "Method Name", "Request Type", "Response Type", "Description"
   :widths: auto

   "GetOAuth2Metadata", ":ref:`ref_nebulaidl.service.OAuth2MetadataRequest`", ":ref:`ref_nebulaidl.service.OAuth2MetadataResponse`", "Anonymously accessible. Retrieves local or external oauth authorization server metadata."
   "GetPublicClientConfig", ":ref:`ref_nebulaidl.service.PublicClientAuthConfigRequest`", ":ref:`ref_nebulaidl.service.PublicClientAuthConfigResponse`", "Anonymously accessible. Retrieves the client information clients should use when initiating OAuth2 authorization requests."

..
   end services




.. _ref_nebulaidl/service/dataproxy.proto:

nebulaidl/service/dataproxy.proto
==================================================================





.. _ref_nebulaidl.service.CreateDownloadLinkRequest:

CreateDownloadLinkRequest
------------------------------------------------------------------

CreateDownloadLinkRequest defines the request parameters to create a download link (signed url)



.. csv-table:: CreateDownloadLinkRequest type fields
   :header: "Field", "Type", "Label", "Description"
   :widths: auto

   "artifact_type", ":ref:`ref_nebulaidl.service.ArtifactType`", "", "ArtifactType of the artifact requested."
   "expires_in", ":ref:`ref_google.protobuf.Duration`", "", "ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this exceeds the platform allowed max. +optional. The default value comes from a global config."
   "node_execution_id", ":ref:`ref_nebulaidl.core.NodeExecutionIdentifier`", "", "NodeId is the unique identifier for the node execution. For a task node, this will retrieve the output of the most recent attempt of the task."







.. _ref_nebulaidl.service.CreateDownloadLinkResponse:

CreateDownloadLinkResponse
------------------------------------------------------------------

CreateDownloadLinkResponse defines the response for the generated links



.. csv-table:: CreateDownloadLinkResponse type fields
   :header: "Field", "Type", "Label", "Description"
   :widths: auto

   "signed_url", ":ref:`ref_string`", "repeated", "SignedUrl specifies the url to use to download content from (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)"
   "expires_at", ":ref:`ref_google.protobuf.Timestamp`", "", "ExpiresAt defines when will the signed URL expire."







.. _ref_nebulaidl.service.CreateDownloadLocationRequest:

CreateDownloadLocationRequest
------------------------------------------------------------------

CreateDownloadLocationRequest specified request for the CreateDownloadLocation API.



.. csv-table:: CreateDownloadLocationRequest type fields
   :header: "Field", "Type", "Label", "Description"
   :widths: auto

   "native_url", ":ref:`ref_string`", "", "NativeUrl specifies the url in the format of the configured storage provider (e.g. s3://my-bucket/randomstring/suffix.tar)"
   "expires_in", ":ref:`ref_google.protobuf.Duration`", "", "ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this exceeds the platform allowed max. +optional. The default value comes from a global config."







.. _ref_nebulaidl.service.CreateDownloadLocationResponse:

CreateDownloadLocationResponse
------------------------------------------------------------------





.. csv-table:: CreateDownloadLocationResponse type fields
   :header: "Field", "Type", "Label", "Description"
   :widths: auto

   "signed_url", ":ref:`ref_string`", "", "SignedUrl specifies the url to use to download content from (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)"
   "expires_at", ":ref:`ref_google.protobuf.Timestamp`", "", "ExpiresAt defines when will the signed URL expires."







.. _ref_nebulaidl.service.CreateUploadLocationRequest:

CreateUploadLocationRequest
------------------------------------------------------------------

CreateUploadLocationRequest specified request for the CreateUploadLocation API.



.. csv-table:: CreateUploadLocationRequest type fields
   :header: "Field", "Type", "Label", "Description"
   :widths: auto

   "project", ":ref:`ref_string`", "", "Project to create the upload location for +required"
   "domain", ":ref:`ref_string`", "", "Domain to create the upload location for. +required"
   "filename", ":ref:`ref_string`", "", "Filename specifies a desired suffix for the generated location. E.g. `file.py` or `pre/fix/file.zip`. +optional. By default, the service will generate a consistent name based on the provided parameters."
   "expires_in", ":ref:`ref_google.protobuf.Duration`", "", "ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this exceeds the platform allowed max. +optional. The default value comes from a global config."
   "content_md5", ":ref:`ref_bytes`", "", "ContentMD5 restricts the upload location to the specific MD5 provided. The ContentMD5 will also appear in the generated path. +required"







.. _ref_nebulaidl.service.CreateUploadLocationResponse:

CreateUploadLocationResponse
------------------------------------------------------------------





.. csv-table:: CreateUploadLocationResponse type fields
   :header: "Field", "Type", "Label", "Description"
   :widths: auto

   "signed_url", ":ref:`ref_string`", "", "SignedUrl specifies the url to use to upload content to (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)"
   "native_url", ":ref:`ref_string`", "", "NativeUrl specifies the url in the format of the configured storage provider (e.g. s3://my-bucket/randomstring/suffix.tar)"
   "expires_at", ":ref:`ref_google.protobuf.Timestamp`", "", "ExpiresAt defines when will the signed URL expires."






..
   end messages



.. _ref_nebulaidl.service.ArtifactType:

ArtifactType
------------------------------------------------------------------

ArtifactType

.. csv-table:: Enum ArtifactType values
   :header: "Name", "Number", "Description"
   :widths: auto

   "ARTIFACT_TYPE_UNDEFINED", "0", "ARTIFACT_TYPE_UNDEFINED is the default, often invalid, value for the enum."
   "ARTIFACT_TYPE_DECK", "1", "ARTIFACT_TYPE_DECK refers to the deck html file optionally generated after a task, a workflow or a launch plan finishes executing."


..
   end enums


..
   end HasExtensions



.. _ref_nebulaidl.service.DataProxyService:

DataProxyService
------------------------------------------------------------------

DataProxyService defines an RPC Service that allows access to user-data in a controlled manner.

.. csv-table:: DataProxyService service methods
   :header: "Method Name", "Request Type", "Response Type", "Description"
   :widths: auto

   "CreateUploadLocation", ":ref:`ref_nebulaidl.service.CreateUploadLocationRequest`", ":ref:`ref_nebulaidl.service.CreateUploadLocationResponse`", "CreateUploadLocation creates a signed url to upload artifacts to for a given project/domain."
   "CreateDownloadLocation", ":ref:`ref_nebulaidl.service.CreateDownloadLocationRequest`", ":ref:`ref_nebulaidl.service.CreateDownloadLocationResponse`", "CreateDownloadLocation creates a signed url to download artifacts."
   "CreateDownloadLink", ":ref:`ref_nebulaidl.service.CreateDownloadLinkRequest`", ":ref:`ref_nebulaidl.service.CreateDownloadLinkResponse`", "CreateDownloadLocation creates a signed url to download artifacts."

..
   end services




.. _ref_nebulaidl/service/identity.proto:

nebulaidl/service/identity.proto
==================================================================





.. _ref_nebulaidl.service.UserInfoRequest:

UserInfoRequest
------------------------------------------------------------------










.. _ref_nebulaidl.service.UserInfoResponse:

UserInfoResponse
------------------------------------------------------------------

See the OpenID Connect spec at https://openid.net/specs/openid-connect-core-1_0.html#UserInfoResponse for more information.



.. csv-table:: UserInfoResponse type fields
   :header: "Field", "Type", "Label", "Description"
   :widths: auto

   "subject", ":ref:`ref_string`", "", "Locally unique and never reassigned identifier within the Issuer for the End-User, which is intended to be consumed by the Client."
   "name", ":ref:`ref_string`", "", "Full name"
   "preferred_username", ":ref:`ref_string`", "", "Shorthand name by which the End-User wishes to be referred to"
   "given_name", ":ref:`ref_string`", "", "Given name(s) or first name(s)"
   "family_name", ":ref:`ref_string`", "", "Surname(s) or last name(s)"
   "email", ":ref:`ref_string`", "", "Preferred e-mail address"
   "picture", ":ref:`ref_string`", "", "Profile picture URL"






..
   end messages


..
   end enums


..
   end HasExtensions



.. _ref_nebulaidl.service.IdentityService:

IdentityService
------------------------------------------------------------------

IdentityService defines an RPC Service that interacts with user/app identities.

.. csv-table:: IdentityService service methods
   :header: "Method Name", "Request Type", "Response Type", "Description"
   :widths: auto

   "UserInfo", ":ref:`ref_nebulaidl.service.UserInfoRequest`", ":ref:`ref_nebulaidl.service.UserInfoResponse`", "Retrieves user information about the currently logged in user."

..
   end services




.. _ref_nebulaidl/service/signal.proto:

nebulaidl/service/signal.proto
==================================================================




..
   end messages


..
   end enums


..
   end HasExtensions



.. _ref_nebulaidl.service.SignalService:

SignalService
------------------------------------------------------------------

SignalService defines an RPC Service that may create, update, and retrieve signal(s).

.. csv-table:: SignalService service methods
   :header: "Method Name", "Request Type", "Response Type", "Description"
   :widths: auto

   "GetOrCreateSignal", ":ref:`ref_nebulaidl.admin.SignalGetOrCreateRequest`", ":ref:`ref_nebulaidl.admin.Signal`", "Fetches or creates a :ref:`ref_nebulaidl.admin.Signal`."
   "ListSignals", ":ref:`ref_nebulaidl.admin.SignalListRequest`", ":ref:`ref_nebulaidl.admin.SignalList`", "Fetch a list of :ref:`ref_nebulaidl.admin.Signal` definitions."
   "SetSignal", ":ref:`ref_nebulaidl.admin.SignalSetRequest`", ":ref:`ref_nebulaidl.admin.SignalSetResponse`", "Sets the value on a :ref:`ref_nebulaidl.admin.Signal` definition"

..
   end services



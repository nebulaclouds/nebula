package oauth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nebulaclouds/nebula/nebulaidl/clients/go/admin/mocks"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/service"
)

func TestGenerateClientConfig(t *testing.T) {
	ctx := context.Background()
	mockAuthClient := new(mocks.AuthMetadataServiceClient)
	nebulaClientResp := &service.PublicClientAuthConfigResponse{
		ClientId:    "dummyClient",
		RedirectUri: "dummyRedirectUri",
		Scopes:      []string{"dummyScopes"},
		Audience:    "dummyAudience",
	}
	oauthMetaDataResp := &service.OAuth2MetadataResponse{
		Issuer:                        "dummyIssuer",
		AuthorizationEndpoint:         "dummyAuthEndPoint",
		TokenEndpoint:                 "dummyTokenEndpoint",
		CodeChallengeMethodsSupported: []string{"dummyCodeChallenege"},
		DeviceAuthorizationEndpoint:   "dummyDeviceEndpoint",
	}
	mockAuthClient.OnGetPublicClientConfigMatch(ctx, mock.Anything).Return(nebulaClientResp, nil)
	mockAuthClient.OnGetOAuth2MetadataMatch(ctx, mock.Anything).Return(oauthMetaDataResp, nil)
	oauthConfig, err := BuildConfigFromMetadataService(ctx, mockAuthClient)
	assert.Nil(t, err)
	assert.NotNil(t, oauthConfig)
	assert.Equal(t, "dummyClient", oauthConfig.ClientID)
	assert.Equal(t, "dummyRedirectUri", oauthConfig.RedirectURL)
	assert.Equal(t, "dummyTokenEndpoint", oauthConfig.Endpoint.TokenURL)
	assert.Equal(t, "dummyAuthEndPoint", oauthConfig.Endpoint.AuthURL)
	assert.Equal(t, "dummyDeviceEndpoint", oauthConfig.DeviceEndpoint)
	assert.Equal(t, "dummyAudience", oauthConfig.Audience)
}

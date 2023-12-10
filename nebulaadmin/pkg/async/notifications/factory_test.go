package notifications

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/notifications/implementations"
	runtimeInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

var (
	scope               = promutils.NewScope("test_sandbox_processor")
	notificationsConfig = runtimeInterfaces.NotificationsConfig{
		Type: "sandbox",
	}
	testEmail = admin.EmailMessage{
		RecipientsEmail: []string{
			"a@example.com",
			"b@example.com",
		},
		SenderEmail: "no-reply@example.com",
		SubjectLine: "Test email",
		Body:        "This is a sample email.",
	}
)

func TestGetEmailer(t *testing.T) {
	defer func() { r := recover(); assert.NotNil(t, r) }()
	cfg := runtimeInterfaces.NotificationsConfig{
		NotificationsEmailerConfig: runtimeInterfaces.NotificationsEmailerConfig{
			EmailerConfig: runtimeInterfaces.EmailServerConfig{
				ServiceName: "unsupported",
			},
		},
	}

	GetEmailer(cfg, promutils.NewTestScope())

	// shouldn't reach here
	t.Errorf("did not panic")
}

func TestNewNotificationPublisherAndProcessor(t *testing.T) {
	testSandboxPublisher := NewNotificationsPublisher(notificationsConfig, scope)
	assert.IsType(t, testSandboxPublisher, &implementations.SandboxPublisher{})
	testSandboxProcessor := NewNotificationsProcessor(notificationsConfig, scope)
	assert.IsType(t, testSandboxProcessor, &implementations.SandboxProcessor{})

	go func() {
		testSandboxProcessor.StartProcessing()
	}()

	assert.Nil(t, testSandboxPublisher.Publish(context.Background(), "TEST_NOTIFICATION", &testEmail))

	assert.Nil(t, testSandboxProcessor.StopProcessing())
}

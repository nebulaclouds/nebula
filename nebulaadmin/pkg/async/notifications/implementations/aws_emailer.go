package implementations

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"google.golang.org/grpc/codes"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/async/notifications/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	runtimeInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type AwsEmailer struct {
	config        runtimeInterfaces.NotificationsConfig
	systemMetrics emailMetrics
	awsEmail      sesiface.SESAPI
}

func NebulaEmailToSesEmailInput(email admin.EmailMessage) ses.SendEmailInput {
	var toAddress []*string
	for _, toEmail := range email.RecipientsEmail {
		// SES email input takes an array of pointers to strings so we have to create a new one for each email
		//nolint:unconvert
		e := string(toEmail)
		toAddress = append(toAddress, &e)
	}

	return ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: toAddress,
		},
		// Currently use the senderEmail specified apart of the Emailer instead of the body.
		// Once a more generic way of setting the emailNotification is defined, remove this
		// workaround and defer back to email.SenderEmail
		Source: &email.SenderEmail,
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: &email.Body,
				},
			},
			Subject: &ses.Content{
				Data: &email.SubjectLine,
			},
		},
	}
}

func (e *AwsEmailer) SendEmail(ctx context.Context, email admin.EmailMessage) error {
	emailInput := NebulaEmailToSesEmailInput(email)
	_, err := e.awsEmail.SendEmail(&emailInput)
	e.systemMetrics.SendTotal.Inc()
	if err != nil {
		// TODO: If we see a certain set of AWS errors consistently, we can break the errors down based on type.
		logger.Errorf(ctx, "error in sending email [%s] via ses mailer with err: %s", email.String(), err)
		e.systemMetrics.SendError.Inc()
		return errors.NewNebulaAdminErrorf(codes.Internal, "errors were seen while sending emails")
	}
	logger.Debugf(ctx, "Sent email to %s sub: %s", email.RecipientsEmail, email.SubjectLine)
	e.systemMetrics.SendSuccess.Inc()
	return nil
}

func NewAwsEmailer(config runtimeInterfaces.NotificationsConfig, scope promutils.Scope, awsEmail sesiface.SESAPI) interfaces.Emailer {
	return &AwsEmailer{
		config:        config,
		systemMetrics: newEmailMetrics(scope.NewSubScope("aws_ses")),
		awsEmail:      awsEmail,
	}
}

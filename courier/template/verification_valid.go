package template

import (
	"encoding/json"

	"github.com/ory/kratos/driver/config"
)

type (
	VerificationValid struct {
		c *config.Config
		m *VerificationValidModel
	}
	VerificationValidModel struct {
		To              string
		VerificationURL string
		Identity        map[string]interface{}
	}
)

func NewVerificationValid(c *config.Config, m *VerificationValidModel) *VerificationValid {
	return &VerificationValid{c: c, m: m}
}

func (t *VerificationValid) EmailRecipient() (string, error) {
	return t.m.To, nil
}

func (t *VerificationValid) EmailSubject() (string, error) {
	return loadTextTemplate(t.c.CourierTemplatesRoot(), "verification/valid/email.subject.gotmpl", "verification/valid/email.subject*", t.m)
}

func (t *VerificationValid) EmailBody() (string, error) {
	return loadHTMLTemplate(t.c.CourierTemplatesRoot(), "verification/valid/email.body.gotmpl", "verification/valid/email.body*", t.m)
}

func (t *VerificationValid) EmailBodyPlaintext() (string, error) {
	return loadTextTemplate(t.c.CourierTemplatesRoot(), "verification/valid/email.body.plaintext.gotmpl", "verification/valid/email.body.plaintext*", t.m)
}

func (t *VerificationValid) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.m)
}

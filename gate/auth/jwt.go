package auth

import (
	microauth "github.com/micro/go-micro/v2/auth"
	iauth "github.com/nano-kit/goeasy/internal/auth"
)

var _ microauth.Auth = (*jwt)(nil)

type jwt struct {
	opts microauth.Options
}

func (j *jwt) String() string {
	return "jwt"
}

func (j *jwt) Init(opts ...microauth.Option) {
	for _, o := range opts {
		o(&j.opts)
	}
}

func (j *jwt) Options() microauth.Options {
	return j.opts
}

func (j *jwt) Generate(id string, opts ...microauth.GenerateOption) (*microauth.Account, error) {
	options := microauth.NewGenerateOptions(opts...)

	return &microauth.Account{
		ID:       id,
		Secret:   options.Secret,
		Metadata: options.Metadata,
		Scopes:   options.Scopes,
		Issuer:   j.Options().Namespace,
	}, nil
}

func (j *jwt) Token(opts ...microauth.TokenOption) (*microauth.Token, error) {
	return &microauth.Token{}, nil
}

func (j *jwt) Grant(rule *microauth.Rule) error {
	return nil
}

func (j *jwt) Revoke(rule *microauth.Rule) error {
	return nil
}

func (j *jwt) Rules(opts ...microauth.RulesOption) ([]*microauth.Rule, error) {
	return []*microauth.Rule{}, nil
}

func (j *jwt) Verify(acc *microauth.Account, res *microauth.Resource, opts ...microauth.VerifyOption) error {
	return nil
}

func (j *jwt) Inspect(token string, opts ...microauth.InspectOption) (*microauth.Account, error) {
	return iauth.AccountFromToken(token)
}

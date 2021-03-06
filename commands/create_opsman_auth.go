package commands

import (
	"fmt"

	"github.com/pivotalservices/pipeline-utilities/common"
	"github.com/pivotalservices/pipeline-utilities/opsman"
)

// CreateOpsmanAuth to create a deprecated auth.yml
type CreateOpsmanAuth struct {
	URL                  string `long:"url" env:"OPSMAN_URL" description:"OpsManager URL" required:"true"`
	SkipSSLValidation    bool   `long:"skip-ssl-validation" env:"OPSMAN_SKIP_SSL_VALIDATION" description:"Skip SSL Validation when interacting with OpsManager"`
	Username             string `long:"username" env:"OPSMAN_USERNAME" description:"OpsManager username"`
	Password             string `long:"password" env:"OPSMAN_PASSWORD" description:"OpsManager password"`
	ClientID             string `long:"client-id" env:"OPSMAN_CLIENTID" description:"OpsManager client id"`
	ClientSecret         string `long:"client-secret" env:"OPSMAN_CLIENT_SECRET" description:"OpsManager client secret"`
	DecryptionPassphrase string `long:"decryption-passphrase" env:"OPSMAN_DECRYPTION_PASSPHRASE" description:"OpsManager Decryption Passphrase"  required:"true"`
	OutputFile           string `long:"output-file" description:"output file for yaml" required:"true"`
}

// Execute - generates structs
func (c *CreateOpsmanAuth) Execute([]string) error {
	fmt.Println("******* WARNING DEPRECATED COMMAND - use env-file and auth-file ************")
	omAuthConfig := opsman.OmAuthConfig{
		OpsmanURL:            c.URL,
		SkipSSLValidation:    c.SkipSSLValidation,
		DecryptionPassphrase: c.DecryptionPassphrase,
	}

	omAuthConfig.Credentials.UserName = c.Username
	omAuthConfig.Credentials.Password = c.Password
	omAuthConfig.Credentials.ClientID = c.ClientID
	omAuthConfig.Credentials.ClientSecret = c.ClientSecret

	return common.WriteYamlFile(c.OutputFile, &omAuthConfig)
}

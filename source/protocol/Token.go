package protocol

import "crypto/ed25519"
import "encoding/base64"
import "encoding/json"
import "fmt"
import "tholian-core/protocol/id"
import "tholian-core/protocol/capabilities"
import "tholian-core/protocol/methods"
import "tholian-core/protocol/permissions"
import "tholian-core/types"
import "time"

type Token struct {
	Version        string                    `json:"version"`
	ID             string                    `json:"id"`
	Issuer         string                    `json:"issuer"`
	Subject        string                    `json:"subject"`
	System         string                    `json:"system"`
	Purpose        permissions.Purpose       `json:"purpose"`
	Grant          string                    `json:"grant"`
	Capabilities   []capabilities.Capability `json:"capabilities"`
	Scope          permissions.Scope         `json:"scope"`
	Method         methods.Method            `json:"method"`
	ParametersHash string                    `json:"parameters-hash"`
	IssuedAt       types.Datetime            `json:"issued-at"`
	ExpiresAt      types.Datetime            `json:"expires-at"`
	Nonce          string                    `json:"nonce"`
	Signature      string                    `json:"signature"`
}

type token_payload struct {
	Version        string                    `json:"version"`
	ID             string                    `json:"id"`
	Issuer         string                    `json:"issuer"`
	Subject        string                    `json:"subject"`
	System         string                    `json:"system"`
	Purpose        permissions.Purpose       `json:"purpose"`
	Grant          string                    `json:"grant"`
	Capabilities   []capabilities.Capability `json:"capabilities"`
	Scope          permissions.Scope         `json:"scope"`
	Method         methods.Method            `json:"method"`
	ParametersHash string                    `json:"parameters-hash"`
	IssuedAt       types.Datetime            `json:"issued-at"`
	ExpiresAt      types.Datetime            `json:"expires-at"`
	Nonce          string                    `json:"nonce"`
}

func NewToken(issuer string, subject string, system string, purpose permissions.Purpose, grant string, caps []capabilities.Capability, scope permissions.Scope, method methods.Method, parameters any, expires_in time.Duration) (Token, error) {

	var token Token

	now := time.Now()
	parameters_hash, err := HashParameters(parameters)

	if err != nil {
		return token, err
	}

	token.Version = Version
	token.ID = id.NewID()
	token.Issuer = issuer
	token.Subject = subject
	token.System = system
	token.Purpose = purpose
	token.Grant = grant
	token.Capabilities = append([]capabilities.Capability(nil), caps...)
	token.Scope = scope
	token.Method = method
	token.ParametersHash = parameters_hash
	token.IssuedAt = types.ToDatetime(now.Format(time.RFC3339))
	token.ExpiresAt = types.ToDatetime(now.Add(expires_in).Format(time.RFC3339))
	token.Nonce = id.NewNonce()

	return token, nil

}

func (token *Token) Sign(key ed25519.PrivateKey) error {

	payload, err := token.Payload()

	if err == nil {

		signature := ed25519.Sign(key, payload)
		token.Signature = base64.StdEncoding.EncodeToString(signature)

		return nil

	}

	return err

}

func (token *Token) Verify(key ed25519.PublicKey) error {

	if token == nil {
		return fmt.Errorf("capability token: missing token")
	}

	if token.Version != Version {
		return fmt.Errorf("capability token: unsupported version %q", token.Version)
	}

	signature, err := base64.StdEncoding.DecodeString(token.Signature)

	if err != nil {
		return fmt.Errorf("capability token: invalid signature encoding")
	}

	if len(key) != ed25519.PublicKeySize {
		return fmt.Errorf("capability token: invalid public key")
	}

	payload, err := token.Payload()

	if err != nil {
		return fmt.Errorf("capability token: invalid payload")
	}

	if ed25519.Verify(key, payload, signature) == false {
		return fmt.Errorf("capability token: signature verification failed")
	}

	return nil

}

func (token *Token) HasCapability(capability capabilities.Capability) bool {

	if token != nil {

		for _, granted := range token.Capabilities {

			if granted == capability {
				return true
			}

		}

	}

	return false

}

func (token *Token) IsExpired(now types.Datetime) bool {

	if token == nil {
		return true
	}

	return token.ExpiresAt.IsBefore(now)

}

func (token *Token) AllowsMethod(system string, method methods.Method, parameters any) bool {

	if token == nil {
		return false
	}

	if token.System != system || token.Method != method {
		return false
	}

	parameters_hash, err := HashParameters(parameters)

	if err != nil {
		return false
	}

	return token.ParametersHash == parameters_hash

}

func (token *Token) Payload() ([]byte, error) {

	payload := token_payload{
		Version:        token.Version,
		ID:             token.ID,
		Issuer:         token.Issuer,
		Subject:        token.Subject,
		System:         token.System,
		Purpose:        token.Purpose,
		Grant:          token.Grant,
		Capabilities:   token.Capabilities,
		Scope:          token.Scope,
		Method:         token.Method,
		ParametersHash: token.ParametersHash,
		IssuedAt:       token.IssuedAt,
		ExpiresAt:      token.ExpiresAt,
		Nonce:          token.Nonce,
	}

	return json.Marshal(payload)

}

package protocol

import "bytes"
import "crypto/ed25519"
import "crypto/sha256"
import "encoding/base64"
import "encoding/hex"
import "encoding/json"
import "fmt"
import "tholian-core/types"
import "time"

type CapabilityToken struct {
	Version      string         `json:"version"`
	ID           string         `json:"id"`
	Issuer       string         `json:"issuer"`
	Subject      string         `json:"subject"`
	SystemID     string         `json:"system"`
	Purpose      Purpose        `json:"purpose"`
	GrantID      string         `json:"grant"`
	Capabilities []Capability   `json:"capabilities"`
	Scope        Scope          `json:"scope"`
	Method       string         `json:"method"`
	ParamsHash   string         `json:"params-hash"`
	IssuedAt     types.Datetime `json:"issued-at"`
	ExpiresAt    types.Datetime `json:"expires-at"`
	Nonce        string         `json:"nonce"`
	Signature    string         `json:"signature"`
}

type tokenPayload struct {
	Version      string         `json:"version"`
	ID           string         `json:"id"`
	Issuer       string         `json:"issuer"`
	Subject      string         `json:"subject"`
	SystemID     string         `json:"system"`
	Purpose      Purpose        `json:"purpose"`
	GrantID      string         `json:"grant"`
	Capabilities []Capability   `json:"capabilities"`
	Scope        Scope          `json:"scope"`
	Method       string         `json:"method"`
	ParamsHash   string         `json:"params-hash"`
	IssuedAt     types.Datetime `json:"issued-at"`
	ExpiresAt    types.Datetime `json:"expires-at"`
	Nonce        string         `json:"nonce"`
}

func NewCapabilityToken(issuer string, subject string, system string, purpose Purpose, grant string, capabilities []Capability, scope Scope, method string, params any, expires_in time.Duration) (CapabilityToken, error) {

	var token CapabilityToken

	now := time.Now()
	params_hash, err := HashParams(params)

	if err != nil {
		return token, err
	}

	token.Version = Version
	token.ID = NewID()
	token.Issuer = issuer
	token.Subject = subject
	token.SystemID = system
	token.Purpose = purpose
	token.GrantID = grant
	token.Capabilities = append([]Capability(nil), capabilities...)
	token.Scope = scope
	token.Method = method
	token.ParamsHash = params_hash
	token.IssuedAt = types.ToDatetime(now.Format(time.RFC3339))
	token.ExpiresAt = types.ToDatetime(now.Add(expires_in).Format(time.RFC3339))
	token.Nonce = NewNonce()

	return token, nil

}

func (token *CapabilityToken) Payload() ([]byte, error) {

	payload := tokenPayload{
		Version:      token.Version,
		ID:           token.ID,
		Issuer:       token.Issuer,
		Subject:      token.Subject,
		SystemID:     token.SystemID,
		Purpose:      token.Purpose,
		GrantID:      token.GrantID,
		Capabilities: token.Capabilities,
		Scope:        token.Scope,
		Method:       token.Method,
		ParamsHash:   token.ParamsHash,
		IssuedAt:     token.IssuedAt,
		ExpiresAt:    token.ExpiresAt,
		Nonce:        token.Nonce,
	}

	return json.Marshal(payload)

}

func (token *CapabilityToken) Sign(key ed25519.PrivateKey) error {

	payload, err := token.Payload()

	if err == nil {

		signature := ed25519.Sign(key, payload)
		token.Signature = base64.StdEncoding.EncodeToString(signature)

		return nil

	}

	return err

}

func (token *CapabilityToken) Verify(key ed25519.PublicKey) error {

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

func (token *CapabilityToken) IsExpired(now types.Datetime) bool {

	if token == nil {
		return true
	}

	return token.ExpiresAt.IsBefore(now)

}

func (token *CapabilityToken) Allows(system string, method string, params any) bool {

	if token == nil {
		return false
	}

	if token.SystemID != system || token.Method != method {
		return false
	}

	params_hash, err := HashParams(params)

	if err != nil {
		return false
	}

	return token.ParamsHash == params_hash

}

func (token *CapabilityToken) HasCapability(capability Capability) bool {

	if token != nil {

		for _, granted := range token.Capabilities {
			if granted == capability {
				return true
			}
		}

	}

	return false

}

func HashParams(params any) (string, error) {

	var blob []byte
	var err error

	if params == nil {
		blob = []byte("{}")
	} else if raw, ok := params.(json.RawMessage); ok == true {
		blob = raw
	} else if raw, ok := params.([]byte); ok == true {
		blob = raw
	} else {
		blob, err = json.Marshal(params)
	}

	if err != nil {
		return "", err
	}

	var compact bytes.Buffer
	err0 := json.Compact(&compact, blob)

	if err0 != nil {
		return "", err0
	}

	checksum := sha256.Sum256(compact.Bytes())

	return hex.EncodeToString(checksum[:]), nil

}

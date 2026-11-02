package handshake

import "crypto/ed25519"
import "encoding/base64"

func SignWelcome(key ed25519.PrivateKey, nonce string, system_id string) string {

	message := []byte(PrefixWelcome.String() + ":" + nonce + ":" + system_id)
	signature := ed25519.Sign(key, message)

	return base64.StdEncoding.EncodeToString(signature)

}


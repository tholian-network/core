package handshake

import "crypto/ed25519"
import "encoding/base64"

func SignHello(key ed25519.PrivateKey, nonce string, node_id string) string {

	message := []byte(PrefixHello.String() + ":" + nonce + ":" + node_id)
	signature := ed25519.Sign(key, message)

	return base64.StdEncoding.EncodeToString(signature)

}


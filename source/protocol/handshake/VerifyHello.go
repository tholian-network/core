package handshake

import "crypto/ed25519"
import "encoding/base64"
import "fmt"

func VerifyHello(key ed25519.PublicKey, nonce string, node_id string, signature string) error {

	message := []byte(PrefixHello.String() + ":" + nonce + ":" + node_id)
	decoded, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return fmt.Errorf("handshake: invalid hello signature encoding")
	}

	if len(key) != ed25519.PublicKeySize {
		return fmt.Errorf("handshake: invalid public key")
	}

	if ed25519.Verify(key, message, decoded) == false {
		return fmt.Errorf("handshake: hello signature verification failed")
	}

	return nil

}


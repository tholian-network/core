package handshake

import "crypto/ed25519"
import "encoding/base64"
import "fmt"

func VerifyWelcome(key ed25519.PublicKey, nonce string, system_id string, signature string) error {

	message := []byte(PrefixWelcome.String() + ":" + nonce + ":" + system_id)
	decoded, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return fmt.Errorf("handshake: invalid welcome signature encoding")
	}

	if len(key) != ed25519.PublicKeySize {
		return fmt.Errorf("handshake: invalid public key")
	}

	if ed25519.Verify(key, message, decoded) == false {
		return fmt.Errorf("handshake: welcome signature verification failed")
	}

	return nil

}

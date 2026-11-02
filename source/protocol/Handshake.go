package protocol

import "crypto/ed25519"
import "encoding/base64"
import "fmt"

func SignHello(key ed25519.PrivateKey, nonce string, node_id string) string {

	message := []byte(HelloPrefix + ":" + nonce + ":" + node_id)
	signature := ed25519.Sign(key, message)

	return base64.StdEncoding.EncodeToString(signature)

}

func VerifyHello(key ed25519.PublicKey, nonce string, node_id string, signature string) error {

	message := []byte(HelloPrefix + ":" + nonce + ":" + node_id)
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

func SignWelcome(key ed25519.PrivateKey, nonce string, system_id string) string {

	message := []byte(WelcomePrefix + ":" + nonce + ":" + system_id)
	signature := ed25519.Sign(key, message)

	return base64.StdEncoding.EncodeToString(signature)

}

func VerifyWelcome(key ed25519.PublicKey, nonce string, system_id string, signature string) error {

	message := []byte(WelcomePrefix + ":" + nonce + ":" + system_id)
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

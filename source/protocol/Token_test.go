package protocol

import "crypto/ed25519"
import "crypto/rand"
import "testing"
import "tholian-core/protocol/capabilities"
import "tholian-core/protocol/methods"
import "tholian-core/protocol/permissions"
import "tholian-core/types"
import "time"

func tokenFixture(t *testing.T, expires_in time.Duration) (Token, ed25519.PublicKey) {

	t.Helper()

	public_key, private_key, err0 := ed25519.GenerateKey(rand.Reader)

	if err0 != nil {
		t.Fatalf("GenerateKey failed: %s", err0)
	}

	purpose := permissions.NewPurpose(permissions.PurposeInvestigation, "Investigate log4j")
	scope := permissions.NewUnrestrictedScope()

	token, err1 := NewToken(
		"watchtower-1",
		"operator-1",
		"system-1",
		purpose,
		"grant-1",
		[]capabilities.Capability{capabilities.ScoutWrite},
		scope,
		methods.ScoutUpdatePackages,
		PackageSearchFixture(),
		expires_in,
	)

	if err1 != nil {
		t.Fatalf("NewToken failed: %s", err1)
	}

	err2 := token.Sign(private_key)

	if err2 != nil {
		t.Fatalf("Sign failed: %s", err2)
	}

	return token, public_key

}

func TestTokenSignAndVerify(t *testing.T) {

	token, public_key := tokenFixture(t, time.Minute)

	err := token.Verify(public_key)

	if err != nil {
		t.Fatalf("Verify failed: %s", err)
	}

	if token.HasCapability(capabilities.ScoutWrite) == false {
		t.Fatal("capability was not preserved")
	}

	if token.HasCapability(capabilities.ScoutRead) == true {
		t.Fatal("token has a capability that was not granted")
	}

	if token.AllowsMethod("system-1", methods.ScoutUpdatePackages, PackageSearchFixture()) == false {
		t.Fatal("token did not allow the signed system and method")
	}

	if token.AllowsMethod("system-2", methods.ScoutUpdatePackages, PackageSearchFixture()) == true {
		t.Fatal("token allowed a different system")
	}

	if token.AllowsMethod("system-1", methods.FirewallForbidNetworks, PackageSearchFixture()) == true {
		t.Fatal("token allowed a different method")
	}

	other_params := PackageSearchFixture()
	other_params.Search.Name = "openssl"

	if token.AllowsMethod("system-1", methods.ScoutUpdatePackages, other_params) == true {
		t.Fatal("token allowed different parameters")
	}

}

func TestTokenDetectsTampering(t *testing.T) {

	token, public_key := tokenFixture(t, time.Minute)

	token.System = "system-2"

	err := token.Verify(public_key)

	if err == nil {
		t.Fatal("tampered token was accepted")
	}

}

func TestTokenExpiry(t *testing.T) {

	token, _ := tokenFixture(t, -1*time.Minute)
	now := types.ToDatetime(time.Now().Format(time.RFC3339))

	if token.IsExpired(now) == false {
		t.Fatal("expired token was considered valid")
	}

	valid, _ := tokenFixture(t, time.Minute)

	if valid.IsExpired(now) == true {
		t.Fatal("valid token was considered expired")
	}

}

func TestTokenWrongKey(t *testing.T) {

	token, _ := tokenFixture(t, time.Minute)
	other_public, _, err0 := ed25519.GenerateKey(rand.Reader)

	if err0 != nil {
		t.Fatalf("GenerateKey failed: %s", err0)
	}

	err1 := token.Verify(other_public)

	if err1 == nil {
		t.Fatal("token verified with the wrong public key")
	}

}

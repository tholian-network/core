package protocol

import "crypto/ed25519"
import "crypto/rand"
import "testing"
import "tholian-core/types"
import "time"

func tokenFixture(t *testing.T, expires_in time.Duration) (CapabilityToken, ed25519.PublicKey) {

	t.Helper()

	public_key, private_key, err0 := ed25519.GenerateKey(rand.Reader)

	if err0 != nil {
		t.Fatalf("GenerateKey failed: %s", err0)
	}

	purpose := NewPurpose(PurposeInvestigation, "Investigate log4j")
	scope := NewUnrestrictedScope()

	token, err1 := NewCapabilityToken(
		"watchtower-1",
		"operator-1",
		"system-1",
		purpose,
		"grant-1",
		[]Capability{CapabilityPackagesUpdate},
		scope,
		MethodPackagesUpdate,
		PackageSearchFixture(),
		expires_in,
	)

	if err1 != nil {
		t.Fatalf("NewCapabilityToken failed: %s", err1)
	}

	err2 := token.Sign(private_key)

	if err2 != nil {
		t.Fatalf("Sign failed: %s", err2)
	}

	return token, public_key

}

func TestCapabilityTokenSignAndVerify(t *testing.T) {

	token, public_key := tokenFixture(t, time.Minute)

	err := token.Verify(public_key)

	if err != nil {
		t.Fatalf("Verify failed: %s", err)
	}

	if token.HasCapability(CapabilityPackagesUpdate) == false {
		t.Fatal("capability was not preserved")
	}

	if token.Allows("system-1", MethodPackagesUpdate, PackageSearchFixture()) == false {
		t.Fatal("token did not allow the signed system and method")
	}

	if token.Allows("system-2", MethodPackagesUpdate, PackageSearchFixture()) == true {
		t.Fatal("token allowed a different system")
	}

	if token.Allows("system-1", MethodFirewallForbid, PackageSearchFixture()) == true {
		t.Fatal("token allowed a different method")
	}

}

func TestCapabilityTokenDetectsTampering(t *testing.T) {

	token, public_key := tokenFixture(t, time.Minute)

	token.SystemID = "system-2"

	err := token.Verify(public_key)

	if err == nil {
		t.Fatal("tampered token was accepted")
	}

}

func TestCapabilityTokenExpiry(t *testing.T) {

	token, _ := tokenFixture(t, -1*time.Minute)
	now := types.ToDatetime(time.Now().Format(time.RFC3339))

	if token.IsExpired(now) == false {
		t.Fatal("expired token was considered valid")
	}

}

func TestCapabilityTokenWrongKey(t *testing.T) {

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

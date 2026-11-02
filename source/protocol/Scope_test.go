package protocol

import "testing"
import "tholian-core/matchers"

func TestScopeAllowsMatchingPackages(t *testing.T) {

	var scope Scope

	allowed := matchers.NewPackage()
	allowed.SetName("log4j")

	scope.Packages = []matchers.Package{allowed}

	requested := matchers.NewPackage()
	requested.SetName("log4j")
	requested.SetVersion("2.14.1")

	if scope.AllowsPackages([]matchers.Package{requested}) == false {
		t.Fatal("matching package was rejected")
	}

	other := matchers.NewPackage()
	other.SetName("openssl")

	if scope.AllowsPackages([]matchers.Package{other}) == true {
		t.Fatal("non-matching package was accepted")
	}

}

func TestScopeDeniesWithoutMatchers(t *testing.T) {

	scope := NewScope()
	requested := matchers.NewPackage()
	requested.SetName("log4j")

	if scope.AllowsPackages([]matchers.Package{requested}) == true {
		t.Fatal("empty scope allowed a package")
	}

	if scope.AllowsConnections(nil) == true {
		t.Fatal("empty scope allowed an empty connection list")
	}

}

func TestUnrestrictedScopeAllowsEverything(t *testing.T) {

	scope := NewUnrestrictedScope()
	requested := matchers.NewPackage()
	requested.SetName("log4j")

	if scope.AllowsPackages([]matchers.Package{requested}) == false {
		t.Fatal("unrestricted scope rejected a package")
	}

	var connection matchers.Connection
	connection.SetHost("1.3.3.7")
	connection.SetPort(1337)

	if scope.AllowsConnections([]matchers.Connection{connection}) == false {
		t.Fatal("unrestricted scope rejected a connection")
	}

}

func TestScopeAllowsMatchingConnection(t *testing.T) {

	var scope Scope

	allowed := matchers.NewConnection()
	allowed.SetHost("1.3.3.7")
	allowed.SetPort(1337)

	scope.Connections = []matchers.Connection{allowed}

	requested := matchers.NewConnection()
	requested.SetHost("1.3.3.7")
	requested.SetPort(1337)

	if scope.AllowsConnections([]matchers.Connection{requested}) == false {
		t.Fatal("matching connection was rejected")
	}

	other := matchers.NewConnection()
	other.SetHost("1.3.3.8")
	other.SetPort(1337)

	if scope.AllowsConnections([]matchers.Connection{other}) == true {
		t.Fatal("non-matching connection was accepted")
	}

}

package protocol

import "testing"
import "tholian-core/protocol/methods"
import "tholian-core/protocol/parameters"

func TestNewRequest(t *testing.T) {

	envelope, err := NewRequest("id-1", methods.ScoutSearchPackages, PackageSearchFixture())

	if err != nil {
		t.Fatalf("NewRequest failed: %s", err)
	}

	if envelope.Version != Version {
		t.Fatalf("unexpected version: %q", envelope.Version)
	}

	if envelope.ID != "id-1" || envelope.Kind != KindRequest || envelope.Method != methods.ScoutSearchPackages {
		t.Fatal("request fields were not preserved")
	}

	if envelope.Metadata == nil {
		t.Fatal("request metadata was not initialized")
	}

	if len(envelope.Parameters) == 0 {
		t.Fatal("request params were not encoded")
	}

}

func TestNewResponse(t *testing.T) {

	result := map[string]string{"status": "ok"}
	envelope, err := NewResponse("id-1", methods.ScoutSearchPackages, result)

	if err != nil {
		t.Fatalf("NewResponse failed: %s", err)
	}

	if envelope.Version != Version {
		t.Fatalf("unexpected version: %q", envelope.Version)
	}

	if envelope.ID != "id-1" || envelope.Kind != KindResponse || envelope.Method != methods.ScoutSearchPackages {
		t.Fatal("response fields were not preserved")
	}

	if envelope.Metadata == nil {
		t.Fatal("response metadata was not initialized")
	}

	if len(envelope.Result) == 0 {
		t.Fatal("response result was not encoded")
	}

}

func TestNewEvent(t *testing.T) {

	envelope, err := NewEvent(methods.GuardOnline, map[string]string{"system": "system-1"})

	if err != nil {
		t.Fatalf("NewEvent failed: %s", err)
	}

	if envelope.Version != Version {
		t.Fatalf("unexpected version: %q", envelope.Version)
	}

	if envelope.Kind != KindEvent || envelope.Method != methods.GuardOnline {
		t.Fatal("event fields were not preserved")
	}

	if len(envelope.Parameters) == 0 {
		t.Fatal("event data was not encoded")
	}

}

func TestNewErrorResponse(t *testing.T) {

	failure := NewError(ErrorForbidden, "not allowed")
	envelope := NewErrorResponse("id-1", methods.WatchtowerListAudits, failure)

	if envelope.IsError() == false {
		t.Fatal("error response was not marked as error")
	}

	if envelope.ID != "id-1" || envelope.Method != methods.WatchtowerListAudits {
		t.Fatal("error response fields were not preserved")
	}

	if envelope.Error == nil || envelope.Error.Code != ErrorForbidden {
		t.Fatal("error details were not preserved")
	}

}

func TestEnvelopeKindChecks(t *testing.T) {

	var nil_envelope *Envelope

	if nil_envelope.IsRequest() || nil_envelope.IsResponse() || nil_envelope.IsEvent() || nil_envelope.IsError() {
		t.Fatal("nil envelope matched a kind")
	}

	request := Envelope{Kind: KindRequest}
	response := Envelope{Kind: KindResponse}
	event := Envelope{Kind: KindEvent}
	error_response := Envelope{Kind: KindError}

	if request.IsRequest() == false || request.IsResponse() || request.IsEvent() || request.IsError() {
		t.Fatal("request envelope matched the wrong kind")
	}

	if response.IsResponse() == false || response.IsRequest() || response.IsEvent() || response.IsError() {
		t.Fatal("response envelope matched the wrong kind")
	}

	if event.IsEvent() == false || event.IsRequest() || event.IsResponse() || event.IsError() {
		t.Fatal("event envelope matched the wrong kind")
	}

	if error_response.IsError() == false || error_response.IsRequest() || error_response.IsResponse() || error_response.IsEvent() {
		t.Fatal("error envelope matched the wrong kind")
	}

}

func TestEnvelopeMetadata(t *testing.T) {

	var envelope Envelope

	if envelope.GetMetadata("missing") != "" {
		t.Fatal("uninitialized metadata returned a value")
	}

	envelope.SetMetadata("session", "session-1")

	if envelope.GetMetadata("session") != "session-1" {
		t.Fatal("metadata was not set")
	}

	envelope.Metadata = nil

	if envelope.GetMetadata("session") != "" {
		t.Fatal("nil metadata returned a value")
	}

}

func PackageSearchFixture() parameters.ScoutSearchPackages {

	var params parameters.ScoutSearchPackages

	params.System = "system-1"
	params.Search.Name = "log4j"

	return params

}

package tholian

import "encoding/json"
import "testing"
import "tholian-core/protocols/tholian/methods"
import "tholian-core/protocols/tholian/parameters"

func TestNewPacket(t *testing.T) {

	packet := NewPacket()

	if packet.Version != Version {
		t.Fatalf("unexpected version: %q", packet.Version)
	}

	if packet.Type != TypeRequest {
		t.Fatalf("unexpected type: %q", packet.Type)
	}

	if packet.Metadata == nil {
		t.Fatal("metadata was not initialized")
	}

	if packet.ID != "" || packet.Method != methods.Method("") {
		t.Fatal("packet was not empty")
	}

	if packet.Parameters != nil || packet.Result != nil || packet.Error != nil {
		t.Fatal("packet payload was not empty")
	}

}

func TestNewRequest(t *testing.T) {

	packet, err := NewRequest("id-1", methods.ScoutSearchPackages, PackageSearchFixture())

	if err != nil {
		t.Fatalf("NewRequest failed: %s", err)
	}

	if packet.Version != Version {
		t.Fatalf("unexpected version: %q", packet.Version)
	}

	if packet.ID != "id-1" || packet.Type != TypeRequest || packet.Method != methods.ScoutSearchPackages {
		t.Fatal("request fields were not preserved")
	}

	if packet.Metadata == nil {
		t.Fatal("request metadata was not initialized")
	}

	if len(packet.Parameters) == 0 {
		t.Fatal("request parameters were not encoded")
	}

	if packet.Result != nil || packet.Error != nil {
		t.Fatal("request contained response fields")
	}

}

func TestNewResponse(t *testing.T) {

	result := map[string]string{"status": "ok"}
	packet, err := NewResponse("id-1", methods.ScoutSearchPackages, result)

	if err != nil {
		t.Fatalf("NewResponse failed: %s", err)
	}

	if packet.Version != Version {
		t.Fatalf("unexpected version: %q", packet.Version)
	}

	if packet.ID != "id-1" || packet.Type != TypeResponse || packet.Method != methods.ScoutSearchPackages {
		t.Fatal("response fields were not preserved")
	}

	if packet.Metadata == nil {
		t.Fatal("response metadata was not initialized")
	}

	if len(packet.Result) == 0 {
		t.Fatal("response result was not encoded")
	}

	if packet.Parameters != nil || packet.Error != nil {
		t.Fatal("response contained request fields")
	}

}

func TestNewEvent(t *testing.T) {

	packet, err := NewEvent(methods.GuardOnline, map[string]string{"system": "system-1"})

	if err != nil {
		t.Fatalf("NewEvent failed: %s", err)
	}

	if packet.Version != Version {
		t.Fatalf("unexpected version: %q", packet.Version)
	}

	if packet.Type != TypeEvent || packet.Method != methods.GuardOnline {
		t.Fatal("event fields were not preserved")
	}

	if len(packet.Parameters) == 0 {
		t.Fatal("event data was not encoded")
	}

	if packet.Result != nil || packet.Error != nil {
		t.Fatal("event contained response fields")
	}

}

func TestNewErrorResponse(t *testing.T) {

	failure := NewError(ErrorForbidden, "not allowed")
	packet := NewErrorResponse("id-1", methods.WatchtowerListAudits, failure)

	if packet.IsError() == false {
		t.Fatal("error response was not marked as error")
	}

	if packet.ID != "id-1" || packet.Method != methods.WatchtowerListAudits {
		t.Fatal("error response fields were not preserved")
	}

	if packet.Error == nil || packet.Error.Code != ErrorForbidden {
		t.Fatal("error details were not preserved")
	}

}

func TestPacketTypeChecks(t *testing.T) {

	var nil_packet *Packet

	if nil_packet.IsRequest() || nil_packet.IsResponse() || nil_packet.IsEvent() || nil_packet.IsError() {
		t.Fatal("nil packet matched a type")
	}

	request := Packet{Type: TypeRequest}
	response := Packet{Type: TypeResponse}
	event := Packet{Type: TypeEvent}
	error_response := Packet{Type: TypeError}

	if request.IsRequest() == false || request.IsResponse() || request.IsEvent() || request.IsError() {
		t.Fatal("request packet matched the wrong type")
	}

	if response.IsResponse() == false || response.IsRequest() || response.IsEvent() || response.IsError() {
		t.Fatal("response packet matched the wrong type")
	}

	if event.IsEvent() == false || event.IsRequest() || event.IsResponse() || event.IsError() {
		t.Fatal("event packet matched the wrong type")
	}

	if error_response.IsError() == false || error_response.IsRequest() || error_response.IsResponse() || error_response.IsEvent() {
		t.Fatal("error packet matched the wrong type")
	}

}

func TestPacketSetType(t *testing.T) {

	packet := NewPacket()

	packet.SetType(TypeResponse)

	if packet.Type != TypeResponse {
		t.Fatalf("unexpected type: %q", packet.Type)
	}

	packet.SetType("nonsense")

	if packet.Type != TypeResponse {
		t.Fatalf("invalid type was accepted: %q", packet.Type)
	}

}

func TestPacketSetParametersAndResult(t *testing.T) {

	packet := NewPacket()

	err0 := packet.SetParameters(map[string]string{"value": "request"})

	if err0 != nil {
		t.Fatalf("SetParameters failed: %s", err0)
	}

	err1 := packet.SetResult(map[string]string{"value": "response"})

	if err1 != nil {
		t.Fatalf("SetResult failed: %s", err1)
	}

	if len(packet.Parameters) == 0 || len(packet.Result) == 0 {
		t.Fatal("packet payload was not encoded")
	}

	packet.SetParameters(json.RawMessage(`{"raw":true}`))
	packet.SetResult([]byte(`{"raw":true}`))

	if string(packet.Parameters) != `{"raw":true}` {
		t.Fatalf("raw parameters were not preserved: %s", packet.Parameters)
	}

	if string(packet.Result) != `{"raw":true}` {
		t.Fatalf("raw result was not preserved: %s", packet.Result)
	}

}

func TestPacketMetadata(t *testing.T) {

	var packet Packet

	if packet.GetMetadata("missing") != "" {
		t.Fatal("uninitialized metadata returned a value")
	}

	packet.SetMetadata("session", "session-1")

	if packet.GetMetadata("session") != "session-1" {
		t.Fatal("metadata was not set")
	}

	packet.Metadata = nil

	if packet.GetMetadata("session") != "" {
		t.Fatal("nil metadata returned a value")
	}

}

func TestPacketSetters(t *testing.T) {

	packet := NewPacket()

	packet.SetVersion("2")
	packet.SetID("id-2")
	packet.SetMethod(methods.Ping)

	if packet.Version != "2" || packet.ID != "id-2" || packet.Method != methods.Ping {
		t.Fatal("packet setters did not update fields")
	}

}

func PackageSearchFixture() parameters.ScoutSearchPackages {

	var params parameters.ScoutSearchPackages

	params.System = "system-1"
	params.Search.Name = "log4j"

	return params

}

package structs

import "bytes"
import "hash/crc32"
import "strconv"
import "strings"
import "tholian-core/types"
import "time"

// Evidence MIME type identifiers. The value is intentionally a plain string to
// follow the existing structs package convention (Incident.Type,
// Connection.Type, Subnet.Type) while remaining compatible with future data
// and file formats.
const (
	EvidenceTypePCAP     = "application/vnd.tcpdump.pcap"
	EvidenceTypePCAPGzip = "application/vnd.tcpdump.pcap+gzip"
	EvidenceTypePCAPNG   = "application/x-pcapng"
	EvidenceTypeJSON     = "application/json"
	EvidenceTypeText     = "text/plain"
)

// Evidence is a reusable recording attached to an incident or response. The
// Data field is the raw recording; JSON encoding uses base64 automatically.
type Evidence struct {
	Type     string         `json:"type"`
	Datetime types.Datetime `json:"datetime"`
	Data     []byte         `json:"data"`
}

func NewEvidence(typ string, data []byte) Evidence {

	var evidence Evidence

	evidence.SetType(typ)
	evidence.SetDatetime(time.Now().Format(time.RFC3339))
	evidence.SetData(data)

	return evidence

}

func (evidence *Evidence) Equal(value Evidence) bool {

	if evidence.Type == value.Type &&
		evidence.Datetime.Equal(value.Datetime) &&
		bytes.Equal(evidence.Data, value.Data) {
		return true
	}

	return false

}

func (evidence *Evidence) IsValid() bool {

	if evidence.Type != "" &&
		evidence.Datetime.IsValid() &&
		len(evidence.Data) > 0 {
		return true
	}

	return false

}

func (evidence *Evidence) SetType(value string) {
	evidence.Type = strings.TrimSpace(value)
}

func (evidence *Evidence) SetDatetime(value string) {

	datetime := types.ToDatetime(value)

	if datetime.IsValid() {
		evidence.Datetime = datetime
	}

}

func (evidence *Evidence) SetData(value []byte) {

	if len(value) > 0 {
		evidence.Data = append([]byte(nil), value...)
	} else {
		evidence.Data = nil
	}

}

func (evidence *Evidence) Hash() string {

	hash := ""

	if evidence.IsValid() {

		checksum := crc32.ChecksumIEEE(evidence.Data)

		hash = toHash(
			evidence.Type,
			evidence.Datetime.String(),
			strconv.FormatUint(uint64(checksum), 10),
			strconv.Itoa(len(evidence.Data)),
		)

	}

	return hash

}

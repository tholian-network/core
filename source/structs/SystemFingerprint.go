package structs

type SystemFingerprint struct {
	Country  string `json:"country"`
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	Token    string `json:"token"`
}

func (fingerprint *SystemFingerprint) Equal(value SystemFingerprint) bool {

	if fingerprint.Country == value.Country &&
		fingerprint.Locale == value.Locale &&
		fingerprint.Timezone == value.Timezone &&
		fingerprint.Token == value.Token {
		return true
	}

	return false

}

func (fingerprint *SystemFingerprint) Hash() string {

	hash := ""

	if fingerprint.Country != "" || fingerprint.Locale != "" || fingerprint.Timezone != "" || fingerprint.Token != "" {
		hash = toHash(fingerprint.Country, fingerprint.Locale, fingerprint.Timezone, fingerprint.Token)
	}

	return hash

}

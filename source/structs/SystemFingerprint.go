package structs

import "strings"

type SystemFingerprint struct {
	Country  string `json:"country"`
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	Token    string `json:"token"`
}

func NewSystemFingerprint() SystemFingerprint {
	return SystemFingerprint{}
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

func (fingerprint *SystemFingerprint) IsValid() bool {

	if fingerprint.Country != "" && fingerprint.Locale != "" && fingerprint.Timezone != "" {
		return true
	}

	return false

}

func (fingerprint *SystemFingerprint) SetCountry(value string) {

	value = strings.TrimSpace(value)

	if value != "" {
		fingerprint.Country = value
	}

}

func (fingerprint *SystemFingerprint) SetLocale(value string) {

	value = strings.TrimSpace(value)

	if value != "" {
		fingerprint.Locale = value
	}

}

func (fingerprint *SystemFingerprint) SetTimezone(value string) {

	value = strings.TrimSpace(value)

	if value != "" {
		fingerprint.Timezone = value
	}

}

func (fingerprint *SystemFingerprint) SetToken(value string) {

	value = strings.TrimSpace(value)

	if value != "" {
		fingerprint.Token = value
	}

}

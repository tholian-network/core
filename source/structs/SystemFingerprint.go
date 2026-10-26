package structs

type SystemFingerprint struct {
	Country  string `json:"country"`
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	Token    string `json:"token"`
}

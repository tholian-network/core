package parameters

type Welcome struct {
	Server    string `json:"server"`
	PublicKey string `json:"public-key"`
	Signature string `json:"signature"`
	System    string `json:"system,omitempty"`
	Session   string `json:"session,omitempty"`
}

package handshake

type Prefix string

const (
	PrefixHello   Prefix = "tholian-hello-v1"
	PrefixWelcome Prefix = "tholian-welcome-v1"
)

func (prefix Prefix) String() string {
	return string(prefix)
}


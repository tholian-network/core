package scout

import "tholian-core/types"

type SearchUsers struct {
	Users []types.User `json:"users"`
}

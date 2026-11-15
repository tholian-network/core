package scout

import "tholian-core/types"

type CollectUsers struct {
	Users []types.User `json:"users"`
}

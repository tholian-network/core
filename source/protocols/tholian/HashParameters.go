package tholian

import "bytes"
import "crypto/sha256"
import "encoding/hex"
import "encoding/json"

func HashParameters(params any) (string, error) {

	blob := make([]byte, 0)
	err := error(nil)

	if params == nil {
		blob = []byte("{}")
	} else if raw, ok := params.(json.RawMessage); ok == true {
		blob = raw
	} else if raw, ok := params.([]byte); ok == true {
		blob = raw
	} else {
		blob, err = json.Marshal(params)
	}

	if err != nil {
		return "", err
	}

	compact := bytes.Buffer{}
	err0 := json.Compact(&compact, blob)

	if err0 != nil {
		return "", err0
	}

	checksum := sha256.Sum256(compact.Bytes())

	return hex.EncodeToString(checksum[:]), nil

}

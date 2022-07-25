package handler

import "encoding/json"

func jsonError(msg string) []byte {
	res := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	r, err := json.Marshal(res)

	if err != nil {
		return []byte(err.Error())
	}

	return r
}

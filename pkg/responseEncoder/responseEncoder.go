package responseEncoder

import (
	"encoding/json"
	"galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdHttp "net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, v interface{}) error {
	reply := &Response{}
	reply.Code = stdHttp.StatusOK
	reply.Message = "success"

	codec, _ := http.CodecForRequest(r, "Accept")
	marshalData, err := codec.Marshal(v)
	if err != nil {
		return err
	}
	_ = json.Unmarshal(marshalData, &reply.Data)
	resp, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", codec.Name())
	w.WriteHeader(stdHttp.StatusOK)
	_, _ = w.Write(resp)
	return nil
}

func ErrorEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, err error) {
	codec, _ := http.CodecForRequest(r, "Accept")
	w.Header().Set("Content-Type", "application/"+codec.Name())
	// status code set 200
	err = errResponse.SetCustomizeErrInfo(err)
	se := errors.FromError(err)
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(body)
}

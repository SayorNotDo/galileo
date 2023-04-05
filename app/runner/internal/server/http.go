package server

import (
	"encoding/json"
	"galileo/app/runner/internal/conf"
	"galileo/app/runner/internal/interfaces"
	"galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdHttp "net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHTTPServer(c *conf.Server, ac *conf.Auth, runner *interfaces.RunnerUseCase, logger log.Logger) *http.Server {
	var opts []http.ServerOption
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	// add success custom json response
	opts = append(opts, http.ResponseEncoder(responseEncoder))
	// add error custom json response
	opts = append(opts, http.ErrorEncoder(errorEncoder))
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", interfaces.RegisterHTTPServer(ac, runner))
	return srv
}

func errorEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, err error) {
	codec, _ := http.CodecForRequest(r, "Accept")
	w.Header().Set("Content-Type", "application/"+codec.Name())
	err = errResponse.SetCustomizeErrInfo(err)
	se := errors.FromError(err)
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(body)
}

func responseEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, v interface{}) error {
	reply := &Response{}
	reply.Code = 20000
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

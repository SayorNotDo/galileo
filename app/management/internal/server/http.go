package server

import (
	projectV1 "galileo/api/management/project/v1"
	testCaseV1 "galileo/api/management/testcase/v1"
	"galileo/app/management/internal/conf"
	"galileo/app/management/internal/service"
	"galileo/pkg/responseEncoder"
	"github.com/go-kratos/kratos/v2/middleware/logging"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, project *service.ProjectService, testcase *service.TestcaseService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			logging.Server(logger),
			recovery.Recovery(),
		),
	}
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
	opts = append(opts, http.ResponseEncoder(responseEncoder.ResponseEncoder))
	// add error custom json response
	opts = append(opts, http.ErrorEncoder(responseEncoder.ErrorEncoder))
	srv := http.NewServer(opts...)
	projectV1.RegisterProjectHTTPServer(srv, project)
	testCaseV1.RegisterTestcaseHTTPServer(srv, testcase)
	return srv
}

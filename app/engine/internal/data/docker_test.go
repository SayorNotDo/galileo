package data_test

import (
	"context"
	"galileo/app/engine/internal/biz"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Docker", func() {
	var ro biz.DockerRepo
	composeFile := []byte(`
version: "3"
services:
 web:
   image: nginx:latest
   ports:
     - "80:80"
 db:
   image: postgres:latest
`)
	It("Parses the compose file", func() {
		_, err := ro.ParseComposeFile(context.Background(), composeFile)
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("Get Docker information", func() {
		ro.GetDockerInfo(context.Background())
	})

})

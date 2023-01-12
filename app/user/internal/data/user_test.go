package data_test

import (
	"galileo/app/user/internal/biz"
	"galileo/app/user/internal/data"
	"galileo/app/user/internal/pkg/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var ro biz.UserRepo
	var uD *biz.User
	BeforeEach(func() {
		ro = data.NewUserRepo(Db, nil)
		hashedPassword, _ := util.HashPassword("test1234")
		uD = &biz.User{
			Username:       "tester",
			HashedPassword: hashedPassword,
			Phone:          "16710790871",
			Email:          "tester@tester.com",
		}
	})
	It("Create", func() {
		u, err := ro.Create(ctx, uD)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(u.Phone).Should(Equal("16710790871"))
	})
})

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

	It("ListUser", func() {
		user, total, err := ro.List(ctx, 1, 10)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(user).ShouldNot(BeEmpty())
		Ω(total).Should(Equal(int32(1)))
		Ω(len(user)).Should(Equal(1))
		Ω(user[0].Phone).Should(Equal("16710790871"))
	})

	It("UpdateUser", func() {
		uD = &biz.User{
			ID:       1,
			Nickname: "Vince",
			Avatar:   "AvatarTest",
		}
		ok, err := ro.Update(ctx, uD)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(ok).Should(BeTrue())
	})

	It("ListUser", func() {
		user, total, err := ro.List(ctx, 1, 10)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(user).ShouldNot(BeEmpty())
		Ω(total).Should(Equal(int32(1)))
		Ω(len(user)).Should(Equal(1))
		Ω(user[0].Phone).Should(Equal("16710790871"))
	})
})

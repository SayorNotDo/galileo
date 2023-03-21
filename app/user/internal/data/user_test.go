package data_test

import (
	"context"
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
			Username: "tester",
			Password: hashedPassword,
			Phone:    "16710790871",
			Email:    "tester@tester.com",
		}
	})
	It("Create", func() {
		u, err := ro.Create(context.Background(), uD)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(u.Phone).Should(Equal("16710790871"))
	})

	It("GetById", func() {
		u, err := ro.GetById(context.Background(), 1)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(u.Phone).Should(Equal("16710790871"))
	})

	It("GetByUsername", func() {
		u, err := ro.GetByUsername(context.Background(), "tester")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(u.Username).Should(Equal("tester"))
	})

	It("ListUser", func() {
		user, total, err := ro.List(context.Background(), 1, 10)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(user).ShouldNot(BeEmpty())
		Ω(total).Should(Equal(int32(1)))
		Ω(len(user)).Should(Equal(1))
		Ω(user[0].Phone).Should(Equal("16710790871"))
	})

	//It("UpdateUser", func() {
	//	uD = &biz.User{
	//		Id:       1,
	//		Nickname: "Vince",
	//		Avatar:   "AvatarTest",
	//	}
	//	ok, err := ro.Update(ctx, uD)
	//	Ω(err).ShouldNot(HaveOccurred())
	//	Ω(ok).Should(BeTrue())
	//})

	//It("ListUser", func() {
	//	user, total, err := ro.List(ctx, 1, 10)
	//	Ω(err).ShouldNot(HaveOccurred())
	//	Ω(user).ShouldNot(BeEmpty())
	//	Ω(total).Should(Equal(int32(1)))
	//	Ω(len(user)).Should(Equal(1))
	//	Ω(user[0].Phone).Should(Equal("16710790871"))
	//})

	//It("CheckPassword", func() {
	//	password := "1is*down9sky"
	//	encryptedPassword := "$2a$10$2N1Hz4h26XbCB2epKk6pOu4/tGXXF1SKOwpIpPtaUxcCi1CxJP1oe"
	//	ok, err := ro.CheckPassword(ctx, password, encryptedPassword)
	//	Ω(err).ShouldNot(HaveOccurred())
	//	Ω(ok).Should(BeTrue())
	//	encryptedPassword1 := "$2a$10$2N1Hz4h26XbCB2epKk6pOu4/"
	//	ok, err = ro.CheckPassword(ctx, password, encryptedPassword1)
	//	Ω(err).ShouldNot(HaveOccurred())
	//	Ω(ok).Should(BeFalse())
	//})
	It("SoftDeleteUser", func() {
		ok, err := ro.SoftDeleteById(context.Background(), 1)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(ok).Should(BeTrue())
	})
	It("DeleteUser", func() {
		ok, err := ro.DeleteById(context.Background(), 1)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(ok).Should(BeTrue())
	})
})

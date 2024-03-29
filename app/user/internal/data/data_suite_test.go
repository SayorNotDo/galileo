package data_test

import (
	"context"
	"galileo/app/user/internal/conf"
	"galileo/app/user/internal/data"
	"galileo/ent"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"testing"
)

func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "test biz data")
}

var cleaner func()
var Db *data.Data
var ctx context.Context

func initialize(db *ent.Client) error {
	err := db.Schema.Create(context.Background())
	return errors.WithStack(err)
}

var _ = BeforeSuite(func() {
	con, f := data.DockerMysql("mariadb", "latest")
	cleaner = f
	config := &conf.Data{Database: &conf.Data_Database{Driver: "mysql", Source: con}}
	db, err := data.NewEntDB(config)
	if err != nil {
		return
	}
	mySQLDb, _, err := data.NewData(config, nil, db, nil)
	if err != nil {
		return
	}
	Db = mySQLDb
	err = initialize(db)
	if err != nil {
		return
	}
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	cleaner()
})

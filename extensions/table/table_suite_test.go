package table_test

import (
	. "github.com/allenluce/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Table Suite")
}

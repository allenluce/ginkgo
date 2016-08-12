package C_test

import (
	. "github.com/allenluce/ginkgo/integration/_fixtures/watch_fixtures/C"

	. "github.com/allenluce/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("C", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})

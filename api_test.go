package gobosh_test

import (
	. "gobosh"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api", func() {
	Describe("Test API", func() {
		var client *Client

		Describe("Test get stemcells", func() {
			BeforeEach(func() {
				setup(MockRoute{"GET", "/stemcells", stemcells, ""}, "basic")
				config := &Config{
					BOSHAddress: server.URL,
					Username:    "admin",
					Password:    "admin",
				}

				client, _ = NewClient(config)
			})

			AfterEach(func() {
				teardown()
			})

			It("can get stemcells", func() {
				stemcells, err := client.GetStemcells()
				Expect(err).Should(BeNil())
				Expect(stemcells[0].Name).Should(Equal("bosh-warden-boshlite-ubuntu-trusty-go_agent"))
				Expect(stemcells[0].OperatingSystem).Should(Equal("ubuntu-trusty"))
				Expect(stemcells[0].Version).Should(Equal("3126"))
				Expect(stemcells[0].CID).Should(Equal("c3705a0d-0dd3-4b67-52b5-50533a432244"))
			})
		})
	})
})

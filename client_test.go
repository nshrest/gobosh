package gobosh_test

import (
	. "gobosh"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	Describe("Testing Default Config", func() {
		config := DefaultConfig()

		It("returns default config", func() {
			Expect(config.BOSHAddress).Should(Equal("https://192.168.50.4:25555"))
			Expect(config.Username).Should(Equal("admin"))
			Expect(config.Password).Should(Equal("admin"))
		})
	})

	Describe("Test Creating basic auth client", func() {
		var client *Client
		BeforeEach(func() {
			setup(MockRoute{"GET", "/stemcells", `{}`, ""}, "basic")
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

		It("can get bosh info", func() {
			info, err := client.GetInfo()
			Expect(info.Name).Should(Equal("bosh2"))
			Expect(info.UUID).Should(Equal("2daf673a-9755-4b4f-aa6d-3632fbed8019"))
			Expect(info.Version).Should(Equal("1.3126.0 (00000000)"))
			Expect(info.User).Should(Equal("admin"))
			Expect(info.CPI).Should(Equal("warden_cpi"))

			Expect(err).Should(BeNil())
		})
	})

	Describe("Test Creating uaa auth client", func() {
		var client *Client

		BeforeEach(func() {
			setup(MockRoute{"GET", "/stemcells", `{}`, ""}, "uaa")
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

		It("can get bosh info", func() {
			info, err := client.GetInfo()
			Expect(info.Name).Should(Equal("bosh2"))
			Expect(info.UUID).Should(Equal("2daf673a-9755-4b4f-aa6d-3632fbed8019"))
			Expect(info.Version).Should(Equal("1.3126.0 (00000000)"))
			Expect(info.User).Should(Equal("admin"))
			Expect(info.CPI).Should(Equal("warden_cpi"))

			Expect(err).Should(BeNil())
		})
	})

	Describe("Test uaa auth", func() {
		var client *Client

		Context("when the refresh token has expired", func() {
			BeforeEach(func() {
				setup(MockRoute{"GET", "/stemcells", `{}`, ""}, "uaa")
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

			Context("and a clientRefresh occurs", func() {
				It("can get brand new uaa token", func() {
					token, err := client.GetToken()
					Expect(err).Should(BeNil())
					Consistently(token).Should(Equal("bearer foobar2"))
					_, err = client.GetInfo()
					Expect(err).Should(BeNil())
					token, err = client.GetToken()
					Expect(err).Should(BeNil())
					Consistently(token).Should(Equal("bearer foobar10"))
				})
			})

			Context("and a clientRefresh does not occur", func() {
				It("can get brand new uaa token", func() {
					token, err := client.GetToken()
					Expect(err).Should(BeNil())
					Consistently(token).Should(Equal("bearer foobar2"))
					token, err = client.GetToken()
					Expect(err).Should(MatchError("Error getting bearer token: oauth2: cannot fetch token: 401 Unauthorized\nResponse: {\"error\":\"invalid_token\",\"error_description\":\"Invalid refresh token (expired)\"}"))
					Consistently(token).Should(Equal(""))
				})
			})
		})

		Context("when the refresh token is valid", func() {
			BeforeEach(func() {
				setup(MockRoute{"GET", "/stemcells", `{}`, ""}, "uaa")
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

			It("can refresh its uaa token", func() {
				token, err := client.GetToken()
				Expect(err).Should(BeNil())
				Consistently(token).Should(Equal("bearer foobar2"))
			})
		})
	})
})

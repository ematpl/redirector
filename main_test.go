package main_test

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const leafGoImportMetaTag = `<meta name="go-import" content="em-go.cfapps.io/leaf git https://github.com/ematpl/leaf">`
const leafGoSourceMetaTag = `<meta name="go-source" content="em-go.cfapps.io/leaf https://github.com/ematpl/leaf https://github.com/ematpl/leaf/tree/master{/dir} https://github.com/ematpl/leaf/blob/master{/dir}/{file}#L{line}">`

const twigGoImportMetaTag = `<meta name="go-import" content="em-go.cfapps.io/twig git https://github.com/ematpl/twig">`
const twigGoSourceMetaTag = `<meta name="go-source" content="em-go.cfapps.io/twig https://github.com/ematpl/twig https://github.com/ematpl/twig/tree/master{/dir} https://github.com/ematpl/twig/blob/master{/dir}/{file}#L{line}">`

var _ = Describe("Redirector", func() {
	Context("when sent a request to the 'leaf' path", func() {
		It("responds with the 'leaf' redirect page", func() {
			resp, err := http.Get("http://" + redirectorAddress + "/leaf?go-get=1")
			Expect(err).NotTo(HaveOccurred())

			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			err = resp.Body.Close()
			Expect(err).NotTo(HaveOccurred())

			Expect(body).To(ContainSubstring(leafGoImportMetaTag))
			Expect(body).To(ContainSubstring(leafGoSourceMetaTag))
		})
	})

	Context("when sent a request to the 'leaf/vein' path", func() {
		It("responds with the 'leaf' redirect page", func() {
			resp, err := http.Get("http://" + redirectorAddress + "/leaf/vein?go-get=1")
			Expect(err).NotTo(HaveOccurred())

			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			err = resp.Body.Close()
			Expect(err).NotTo(HaveOccurred())

			Expect(body).To(ContainSubstring(leafGoImportMetaTag))
			Expect(body).To(ContainSubstring(leafGoSourceMetaTag))
		})
	})

	Context("when sent a request to the 'twig' path", func() {
		It("responds with the 'twig' redirect page", func() {
			resp, err := http.Get("http://" + redirectorAddress + "/twig?go-get=1")
			Expect(err).NotTo(HaveOccurred())

			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			err = resp.Body.Close()
			Expect(err).NotTo(HaveOccurred())

			Expect(body).To(ContainSubstring(twigGoImportMetaTag))
			Expect(body).To(ContainSubstring(twigGoSourceMetaTag))
		})
	})
})
package testsapi

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Markdown API", func() {

	Describe("Given a working Github API", func() {

		BeforeEach(func() {
			// TODO: INSERT CODE HERE TO VERIFY THAT THE GITHUB IS UP AND RUNNING
		})

		Context("When a call is made to the markdown endpoint", func() {

			BeforeEach(func() {
				// TODO: INSERT CODE HERE TO SEND A REQUEST TO THE MARKDOWN ENDPOINT
				// https://developer.github.com/v3/markdown/
			})

			Specify("Then the markdown endpoint will return the correct output", func() {

			})
		})
	})

	// TODO: Add more scenarios against the markdown API, feel free to keep adding more by uncommenting the below
	// and change the scenario descriptions to whatever you feel would be a useful test:

	// Describe("Given a working Github API", func() {

	// 	BeforeEach(func() {
	// 		// TODO: INSERT CODE HERE TO VERIFY THAT THE GITHUB IS UP AND RUNNING
	// 	})

	// 	Context("When a call is made to the markdown endpoint", func() {

	// 		BeforeEach(func() {
	// 			// TODO: INSERT CODE HERE TO SEND A REQUEST TO THE MARKDOWN ENDPOING
	// 			// https://developer.github.com/v3/markdown/
	// 		})

	// 		Specify("Then the markdown endpoint will return the correct output", func() {

	// 		})
	// 	})
	// })

	// Describe("Given a working Github API", func() {

	// 	BeforeEach(func() {
	// 		// TODO: INSERT CODE HERE TO VERIFY THAT THE GITHUB IS UP AND RUNNING
	// 	})

	// 	Context("When a call is made to the markdown endpoint", func() {

	// 		BeforeEach(func() {
	// 			// TODO: INSERT CODE HERE TO SEND A REQUEST TO THE MARKDOWN ENDPOING
	// 			// https://developer.github.com/v3/markdown/
	// 		})

	// 		Specify("Then the markdown endpoint will return the correct output", func() {

	// 		})
	// 	})
	// })
})

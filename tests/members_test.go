package testsapi

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Members API", func() {

	Describe("Given a working Github API", func() {

		// TODO: INSERT CODE HERE TO VERIFY THAT THE GITHUB IS UP AND RUNNING

		Context("When a call is made to the check members endpoint {TODO: ADD IN SPECIFIC TEST ex. marccampbell is a member of replicatedcom}", func() {

			// TODO: INSERT CODE HERE TO SEND A REQUEST TO THE CHECK MEMBERS ENDPOINT
			// Example: checking if `marccampbell` is a member of the organization `replicatedcom`
			// https://developer.github.com/v3/orgs/members/

			Specify("Then the members endpoint will verify if marccampbell is a member of replicatedcom", func() {

				// TODO: INSERT CODE HERE TO VERIFY QUALITY

			})
		})
	})

	// TODO: Add more scenarios against the members API, feel free to keep adding more by uncommenting the below
	// and change the scenario descriptions to whatever you feel would be a useful test:

	// 	Context("When a call is made to the check members endpoint {TODO: ADD IN SPECIFIC TEST ex. marccampbell is a member of replicatedcom}", func() {

	// 			// TODO: INSERT CODE HERE TO SEND A REQUEST TO THE CHECK MEMBERS ENDPOINT
	// 			// CHECKING IF `marccampbell` is a member of the organization `replicatedcom`
	// 			// https://developer.github.com/v3/orgs/members/

	// 		Specify("Then the members endpoint will verify if marccampbell is a member of replicatedcom", func() {

	// TODO: INSERT CODE HERE TO VERIFY QUALITY

	// 		})
	// 	})
	// })

	// 	Context("When a call is made to the check members endpoint {TODO: ADD IN SPECIFIC TEST ex. marccampbell is a member of replicatedcom}", func() {

	// 			// TODO: INSERT CODE HERE TO SEND A REQUEST TO THE CHECK MEMBERS ENDPOINT
	// 			// CHECKING IF `marccampbell` is a member of the organization `replicatedcom`
	// 			// https://developer.github.com/v3/orgs/members/

	// 		Specify("Then the members endpoint will verify if marccampbell is a member of replicatedcom", func() {

	// TODO: INSERT CODE HERE TO VERIFY QUALITY

	// 		})
	// 	})
	// })
})

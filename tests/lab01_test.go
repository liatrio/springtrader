package main

import (
	"log"

	. "github.com/liatrio/springtrader/tests/validate"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lab 1 Containers", func() {
	var failMessage string

	BeforeEach(func() {
		failMessage = ""
	})

	Context("Step 2", func() {
		It("should have a Dockerfile", func() {
			failMessage = "Dockerfile doesn't exist or is in the wrong location\n"
			Expect("../Dockerfile").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 3", func() {
		It("should have a skaffold.yaml file", func() {
			failMessage = "skaffold.yaml doesn't exist or is in the wrong location\n"
			Expect("../skaffold.yaml").To(BeAnExistingFile(), failMessage)
		})

		It("should have a valid skaffold.yaml", func() {
			skaffoldExpected := ExpectYamlToParse("../skaffold.yaml")
			skaffoldActual := ExpectYamlToParse("./validate/solution-data/lab01/step03-skaffold.yaml")
			failMessage = "skaffold.yaml has incorrect configuration"
			Expect(skaffoldActual).To(ValidateYamlObject(skaffoldExpected, &failMessage))
		})
	})

	Context("Step 6", func() {
		It("should have a deployment.yaml file", func() {
			failMessage = "deployment.yaml doesn't exist or is in the wrong location\n"
			Expect("../charts/springtrader/templates/deployment.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 7", func() {
		It("should have a statefulset.yaml file", func() {
			failMessage = "statefulset.yaml doesn't exist or is in the wrong location\n"
			Expect("../charts/springtrader/templates/statefulset.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 8", func() {
		It("should have a service.yaml file", func() {
			failMessage = "service.yaml doesn't exist or is in the wrong location\n"
			Expect("../charts/springtrader/templates/service.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 9", func() {
		It("should have a job.yaml file", func() {
			failMessage = "job.yaml doesn't exist or is in the wrong location\n"
			Expect("../charts/springtrader/templates/job.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 11", func() {
		It("skaffold file should have a profile section", func() {
			skaffoldExpected := ExpectYamlToParse("../skaffold.yaml")
			skaffoldActual := ExpectYamlToParse("./validate/solution-data/lab01/step11-skaffold.yaml")
			failMessage = "skaffold.yaml has incorrect configuration\n"
			Expect(skaffoldActual).To(ValidateYamlObject(skaffoldExpected, &failMessage))
		})
	})

	AfterEach(func() {
		log.Printf("%v\n", CurrentGinkgoTestDescription())
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})

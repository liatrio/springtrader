package main

import (
	"fmt"

	. "github.com/liatrio/springtrader/tests/validate"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lab 1", func() {
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
		It("should have a valid skaffold.yaml", func() {
			skaffoldExpected, errorMessage := ExpectYamlToParse("../skaffold.yaml")
			if errorMessage != "" {
				failMessage = errorMessage
			}
			Expect(errorMessage).To(BeEmpty(), failMessage)
			failMessage = fmt.Sprintf("Your skaffold.yaml seems to empty. Try again after configuring your file\n")
			Expect(skaffoldExpected).ToNot(BeNil(), failMessage)
			skaffoldActual, _ := ExpectYamlToParse("./validate/solution-data/lab01/step03-skaffold.yaml")
			_, err := ValidateYamlObject(skaffoldExpected, &failMessage).Match(skaffoldActual)
			if err != nil {
				failMessage = fmt.Sprintf("skaffold.yaml has incorrect configuration; %s\n", err.Error())
			}
			Expect(err).To(BeNil(), failMessage)
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
			skaffoldExpected, errorMessage := ExpectYamlToParse("../skaffold.yaml")
			if errorMessage != "" {
				failMessage = errorMessage
			}
			Expect(errorMessage).To(BeEmpty(), failMessage)
			failMessage = fmt.Sprintf("Your skaffold.yaml seems to empty. Try again after configuring your file\n")
			Expect(skaffoldExpected).ToNot(BeNil(), failMessage)
			skaffoldActual, _ := ExpectYamlToParse("./validate/solution-data/lab01/step11-skaffold.yaml")
			_, err := ValidateYamlObject(skaffoldExpected, &failMessage).Match(skaffoldActual)
			if err != nil {
				failMessage = fmt.Sprintf("skaffold.yaml has incorrect configuration; %s\n", err.Error())
			}
			Expect(err).To(BeNil())
		})
	})

	AfterEach(func() {
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage = ConcatenatedMessage + failMessage + "\n"
		}
	})
})

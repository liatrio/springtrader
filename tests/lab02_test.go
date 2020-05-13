package main

import (
	"fmt"

	. "github.com/liatrio/springtrader/tests/validate"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lab 2", func() {
	var failMessage string

	BeforeEach(func() {
		failMessage = ""
	})

	Context("Step 1", func() {
		It("should have a skaffold.yaml file", func() {
			failMessage = "skaffold.yaml doesn't exist or is in the wrong location\n"
			Expect("../skaffold.yaml").To(BeAnExistingFile(), failMessage)
		})
		It("should have a valid skaffold.yaml", func() {
			skaffoldExpected, errorMessage := ExpectYamlToParse("../skaffold.yaml")
			if errorMessage != "" {
				failMessage = errorMessage
			}
			Expect(errorMessage).To(BeEmpty(), failMessage)
			failMessage = fmt.Sprintf("Your skaffold.yaml seems to empty. Try again after configuring your file\n")
			Expect(skaffoldExpected).ToNot(BeNil(), failMessage)
			skaffoldActual, _ := ExpectYamlToParse("./validate/solution-data/lab02/step01-skaffold.yaml")
			_, err := ValidateYamlObject(skaffoldExpected, &failMessage).Match(skaffoldActual)
			if err != nil {
				failMessage = fmt.Sprintf("skaffold.yaml has incorrect configuration; %s\n", err.Error())
			}
			Expect(err).To(BeNil())
		})
	})

	Context("Step 6", func() {
		It("should have a Jenkinsfile", func() {
			failMessage = "Jenkinsfile doesn't exist or is in the wrong location\n"
			Expect("../Jenkinsfile").To(BeAnExistingFile(), failMessage)
		})
	})

	AfterEach(func() {
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})

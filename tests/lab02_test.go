package main

import (
	"fmt"

	. "github.com/liatrio/springtrader/tests/validate"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lab 2 jenkins", func() {
	var failMessage string

	BeforeEach(func() {
		failMessage = ""
	})

	Context("Step 1", func() {
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
			ConcatenatedMessage = ConcatenatedMessage + failMessage + "\n"
		}
	})
})

var _ = Describe("Lab 2 aws", func() {
	var failMessage string

	BeforeEach(func() {
		failMessage = ""
	})

	Context("Step 1", func() {
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

	Context("Step 3.2", func() {
		It("should have a valid Build buildspec", func() {
			buildspecExpected, errorMessage := ExpectYamlToParse("../buildspec-build.yaml")
			if errorMessage != "" {
				failMessage = errorMessage
			}
			Expect(errorMessage).To(BeEmpty(), failMessage)
			failMessage = fmt.Sprintf("Your buildspec-build.yaml seems to empty. Try again after configuring your file\n")
			Expect(buildspecExpected).ToNot(BeNil(), failMessage)
			buildspecActual, _ := ExpectYamlToParse("./validate/solution-data/lab02/step03-build-buildspec.yaml")
			_, err := ValidateYamlObject(buildspecExpected, &failMessage).Match(buildspecActual)
			if err != nil {
				failMessage = fmt.Sprintf("buildspec-build.yaml has incorrect configuration; %s\n", err.Error())
			}
			Expect(err).To(BeNil())
		})
	})

	Context("Step 3.3", func() {
		It("should have a valid Staging buildspec", func() {
			buildspecExpected, errorMessage := ExpectYamlToParse("../buildspec-staging.yaml")
			if errorMessage != "" {
				failMessage = errorMessage
			}
			Expect(errorMessage).To(BeEmpty(), failMessage)
			failMessage = fmt.Sprintf("Your buildspec-staging.yaml seems to empty. Try again after configuring your file\n")
			Expect(buildspecExpected).ToNot(BeNil(), failMessage)
			buildspecActual, _ := ExpectYamlToParse("./validate/solution-data/lab02/step03-staging-buildspec.yaml")
			_, err := ValidateYamlObject(buildspecExpected, &failMessage).Match(buildspecActual)
			if err != nil {
				failMessage = fmt.Sprintf("buildspec-staging.yaml has incorrect configuration; %s\n", err.Error())
			}
			Expect(err).To(BeNil())
		})
	})

	Context("Step 3.5", func() {
		It("should have a valid Prod buildspec", func() {
			buildspecExpected, errorMessage := ExpectYamlToParse("../buildspec-prod.yaml")
			if errorMessage != "" {
				failMessage = errorMessage
			}
			Expect(errorMessage).To(BeEmpty(), failMessage)
			failMessage = fmt.Sprintf("Your buildspec-prod.yaml seems to empty. Try again after configuring your file\n")
			Expect(buildspecExpected).ToNot(BeNil(), failMessage)
			buildspecActual, _ := ExpectYamlToParse("./validate/solution-data/lab02/step03-prod-buildspec.yaml")
			_, err := ValidateYamlObject(buildspecExpected, &failMessage).Match(buildspecActual)
			if err != nil {
				failMessage = fmt.Sprintf("buildspec-prod.yaml has incorrect configuration; %s\n", err.Error())
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

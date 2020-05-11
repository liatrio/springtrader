package validate

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lab 2 ", func() {
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
			skaffoldExpected := expectYamlToParse("../skaffold.yaml")
			skaffoldActual := expectYamlToParse("./solution-data/lab02/step01-skaffold.yaml")
			failMessage = "skaffold.yaml has incorrect configuration\n"
			Expect(skaffoldActual).To(ValidateYamlObject(skaffoldExpected), failMessage)
		})
	})

	Context("Step 6", func() {
		It("should have a Jenkinsfile", func() {
			failMessage = "Jenkinsfile doesn't exist or is in the wrong location\n"
			Expect("../Jenkinsfile").To(BeAnExistingFile(), failMessage)
		})
	})

	AfterEach(func() {
		log.Printf("%v\n", CurrentGinkgoTestDescription())
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})

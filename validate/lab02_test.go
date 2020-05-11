package validate

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lab 2 Continuous Delivery", func() {
	var failMessage string

	BeforeEach(func() {
		failMessage = ""
	})

	Context("Step 2", func() {
		It("should have a virtualservice.yaml", func() {
			failMessage = "virtualservice.yaml Doesn't Exist or is in the wrong location\n"
			Expect("../charts/springtrader/templates/virtualservice.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 3", func() {
		It("should have a Jenkinsfile file", func() {
			failMessage = "Jenkinsfile Doesn't Exist or is in the wrong location\n"
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

func expectYamlToParse(path string) interface{} {
	var output interface{}
	file, err := ioutil.ReadFile(path)
	failMessage := fmt.Sprintf("File at the path, %s, cannot be found. File may be in wrong location or misnamed.\n", path)
	Expect(err).To(BeNil(), failMessage)
	err = yaml.Unmarshal([]byte(file), &output)
	failMessage = fmt.Sprintf("File at the path, %s, could not be parsed as YAML.\n Error: %s\n", path, err)
	Expect(err).To(BeNil(), failMessage)
	return output
}

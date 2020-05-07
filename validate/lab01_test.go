package validate

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

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
			failMessage = "Dockerfile Doesn't Exist or is in the wrong location\n"
			Expect("../Dockerfile").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 3", func() {
		It("should have a skaffold.yaml file", func() {
			failMessage = "skaffold.yaml Doesn't Exist or is in the wrong location\n"
			Expect("../skaffold.yaml").To(BeAnExistingFile(), failMessage)
		})

		It("should have valid versions", func() {
			skaffoldExpected := YamlToInterface("../skaffold.yaml")
			skaffoldActual := YamlToInterface("./solution-data/lab01step03/skaffold-version.yaml")
			Expect(skaffoldActual).To(ValidateYamlObject(skaffoldExpected))
		})

		It("should have valid build and local sections", func() {
			skaffoldExpected := YamlToInterface("../skaffold.yaml")
			skaffoldActual := YamlToInterface("./solution-data/lab01step03/skaffold-build.yaml")
			Expect(skaffoldActual).To(ValidateYamlObject(skaffoldExpected))
		})
	})

	Context("Step 6", func() {
		It("should have a deployment.yaml file", func() {
			failMessage = "deployment.yaml Doesn't Exist or is in the wrong location\n"
			Expect("../charts/springtrader/templates/deployment.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 7", func() {
		It("should have a statefulset.yaml file", func() {
			failMessage = "statefulset.yaml Doesn't Exist or is in the wrong location\n"
			Expect("../charts/springtrader/templates/statefulset.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 8", func() {
		It("should have a service.yaml file", func() {
			failMessage = "service.yaml Doesn't Exist or is in the wrong location\n"
			Expect("../charts/springtrader/templates/service.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 9", func() {
		It("should have a job.yaml file", func() {
			failMessage = "job.yaml Doesn't Exist or is in the wrong location\n"
			Expect("../charts/springtrader/templates/job.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 11", func() {
		It("skaffold file should still have valid versions", func() {
			skaffoldExpected := YamlToInterface("../skaffold.yaml")
			skaffoldActual := YamlToInterface("./solution-data/lab01step03/skaffold-version.yaml")
			Expect(skaffoldActual).To(ValidateYamlObject(skaffoldExpected))
		})

		It("should still have valid build and local sections", func() {
			skaffoldExpected := YamlToInterface("../skaffold.yaml")
			skaffoldActual := YamlToInterface("./solution-data/lab01step03/skaffold-build.yaml")
			Expect(skaffoldActual).To(ValidateYamlObject(skaffoldExpected))
		})

		It("skaffold file should have a deploy section", func() {
			skaffoldExpected := YamlToInterface("../skaffold.yaml")
			skaffoldActual := YamlToInterface("./solution-data/lab01step11/skaffold-deploy.yaml")
			Expect(skaffoldActual).To(ValidateYamlObject(skaffoldExpected))
		})

		It("skaffold file should have a profile section", func() {
			skaffoldExpected := YamlToInterface("../skaffold.yaml")
			skaffoldActual := YamlToInterface("./solution-data/lab01step11/skaffold-profiles.yaml")
			Expect(skaffoldActual).To(ValidateYamlObject(skaffoldExpected))
		})
	})

	AfterEach(func() {
		log.Printf("%v\n", CurrentGinkgoTestDescription())
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})

func YamlToInterface(path string) interface{} {
	var output interface{}
	file, _ := ioutil.ReadFile(path)
	yaml.Unmarshal([]byte(file), &output)
	return output
}

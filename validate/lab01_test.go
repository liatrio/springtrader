package validate

import (
  "io/ioutil"
  "fmt"
  "os"

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
			Expect(fileExists("../Dockerfile")).To(Succeed(), failMessage)
		})
	})

	Context("Step 3", func() {
		It("should have a skaffold.yaml file", func() {
			failMessage = "skaffold.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../skaffold.yaml")).To(Succeed(), failMessage)
		})

		It("should be a valid skaffold configuration", func() {
			var skaffold interface{}
			skaffoldFile, err := ioutil.ReadFile("../skaffold.yaml")
			if err != nil {
				Skip("skaffold.yaml not found")
			}

			err = yaml.Unmarshal(skaffoldFile, &skaffold)
			Expect(err).ToNot(HaveOccurred())
			//failures := InterceptGomegaFailures(func() {

			failMessage = "Incorrect apiVersion in skaffold.yaml\n"
			Expect(treeValue(skaffold, []interface{}{"apiVersion"})).To(Equal("skaffold/v1beta12"), failMessage)
			failMessage = "First build artifact in skaffold.yaml should be \"springtrader\"\n"
			Expect(treeValue(skaffold, []interface{}{"build", "artifacts", 0, "image"})).To(Equal("springtrader"), failMessage)
			failMessage = "Second build artifact in skaffold.yaml should be \"sqlfdb\"\n"
			Expect(treeValue(skaffold, []interface{}{"build", "artifacts", 1, "image"})).To(Equal("sqlfdb"), failMessage)
		})
	})

    Context("Step 5", func() {
		It("should have a deployment.yaml file", func() {
			failMessage = "deployment.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../deployment.yaml")).To(Succeed(), failMessage)
		})

        It("should have a statefulset.yaml file", func() {
			failMessage = "statefulset.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../statefulset.yaml")).To(Succeed(), failMessage)
		})

        It("should have a service.yaml file", func() {
			failMessage = "service.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../service.yaml")).To(Succeed(), failMessage)
		})

        It("should have a job.yaml file", func() {
			failMessage = "job.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../job.yaml")).To(Succeed(), failMessage)
		})
    })

    Context("Step 11", func() {
		It("should have deploy and policy sections in the skaffold.yaml file", func() {
			var skaffold interface{}
			skaffoldFile, err := ioutil.ReadFile("../skaffold.yaml")
			if err != nil {
				Skip("skaffold.yaml not found")
			}

			err = yaml.Unmarshal(skaffoldFile, &skaffold)
			Expect(err).ToNot(HaveOccurred())

            failMessage = "Incorrect deploy section in skaffold.yaml\n"
			Expect(treeValue(skaffold, []interface{}{"deploy", "helm", "releases", 0, "name"})).To(Equal("springtrader"), failMessage)

            failMessage = "Incorrect profiles section in skaffold.yaml\n"
			Expect(treeValue(skaffold, []interface{}{"profiles", 0, "name"})).To(Equal("kind"), failMessage)
		})
	})


	AfterEach(func() {
		log.Printf("%v\n", CurrentGinkgoTestDescription())
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})

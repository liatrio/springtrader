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
			Expect(fileExists("../Dockerfile")).To(Succeed(), failMessage)
		})
	})

	Context("Step 3", func() {
		It("should have a skaffold.yaml file", func() {
			failMessage = "skaffold.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../skaffold.yaml")).To(Succeed(), failMessage)
		})

		It("should have a valid skaffold configuration", func() {
			var skaffoldFile interface{}
			var skaffoldExpected interface{}

			skaffoldExpectedFile, err := ioutil.ReadFile("./solution-data/lab01step03/skaffold.yaml")
			err = yaml.Unmarshal(skaffoldExpectedFile, &skaffoldExpected)

            skaffoldFile, err = ioutil.ReadFile("../skaffold.yaml")
			if err != nil {
				Skip("skaffold.yaml not found")
			}
			err = treeCompare(skaffoldFile, skaffoldExpected)
			Expect(err).To(BeNil())
		})
	})

	Context("Step 5", func() {
		It("should have a deployment.yaml file", func() {
			failMessage = "deployment.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../charts/springtrader/templates/deployment.yaml")).To(Succeed(), failMessage)
		})

        It("should have a valid deployment.yaml file", func() {
            failMessage = "deployment.yaml has incorrect contents\n"
            var deploymentFile interface{}
			var deploymentExpected interface{}

			deploymentExpectedFile, err := ioutil.ReadFile("./solution-data/lab01step05/deployment.yaml")
			err = yaml.Unmarshal(deploymentExpectedFile, &deploymentExpected)
			Expect(err).ToNot(HaveOccurred())

            deploymentFile, err = ioutil.ReadFile("../charts/springtrader/templates/deployment.yaml")
			if err != nil {
				Skip("deployment.yaml not found")
			}
			err = treeCompare(deploymentFile, deploymentExpected)
			Expect(err).ToNot(HaveOccurred())
        })

		It("should have a statefulset.yaml file", func() {
			failMessage = "statefulset.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../charts/springtrader/templates/statefulset.yaml")).To(Succeed(), failMessage)
		})

        It("should have a valid statefulset.yaml file", func() {
            failMessage = "statefulset.yaml has incorrect contents\n"
            var statefulsetFile interface{}
			var statefulsetExpected interface{}

			statefulsetExpectedFile, err := ioutil.ReadFile("./solution-data/lab01step05/statefulset.yaml")
			err = yaml.Unmarshal(statefulsetExpectedFile, &statefulsetExpected)
			Expect(err).ToNot(HaveOccurred())

            statefulsetFile, err = ioutil.ReadFile("../charts/springtrader/templates/statefulset.yaml")
			if err != nil {
				Skip("statefulset.yaml not found")
			}
			err = treeCompare(statefulsetFile, statefulsetExpected)
			Expect(err).ToNot(HaveOccurred())
        })

		It("should have a service.yaml file", func() {
			failMessage = "service.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../charts/springtrader/templates/service.yaml")).To(Succeed(), failMessage)
		})

        It("should have a valid service.yaml file", func() {
            failMessage = "service.yaml has incorrect contents\n"
            var serviceFile interface{}
			var serviceExpected interface{}

			serviceExpectedFile, err := ioutil.ReadFile("./solution-data/lab01step05/service.yaml")
			err = yaml.Unmarshal(serviceExpectedFile, &serviceExpected)
			Expect(err).ToNot(HaveOccurred())

            serviceFile, err = ioutil.ReadFile("../charts/springtrader/templates/service.yaml")
			if err != nil {
				Skip("service.yaml not found")
			}
			err = treeCompare(serviceFile, serviceExpected)
			Expect(err).ToNot(HaveOccurred())
        })

		It("should have a job.yaml file", func() {
			failMessage = "job.yaml Doesn't Exist or is in the wrong location\n"
			Expect(fileExists("../charts/springtrader/templates/job.yaml")).To(Succeed(), failMessage)
		})

        It("should have a valid job.yaml file", func() {
            failMessage = "job.yaml has incorrect contents\n"
            var jobFile interface{}
			var jobExpected interface{}

			jobExpectedFile, err := ioutil.ReadFile("./solution-data/lab01step05/job.yaml")
			err = yaml.Unmarshal(jobExpectedFile, &jobExpected)
			Expect(err).ToNot(HaveOccurred())

            jobFile, err = ioutil.ReadFile("../charts/springtrader/templates/job.yaml")
			if err != nil {
				Skip("job.yaml not found")
			}
			err = treeCompare(jobFile, jobExpected)
			Expect(err).ToNot(HaveOccurred())
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

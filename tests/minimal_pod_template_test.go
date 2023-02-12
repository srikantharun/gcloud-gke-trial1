package test

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
)


func TestPodTemplateRendersContainerImage(t *testing.T) {
	// Path to the helm chart we will test
	helmChartPath := "../app"

	// Setup the args. For this test, we will set the following input values:
	// - image=nginx:1.15.8
	options := &helm.Options{
		SetValues: map[string]string{
	            "image.repository": "nginx",
		    "image.tag"       : "1.15.8",	
		},
	}

	// Run RenderTemplate to render the template and capture the output.
	output := helm.RenderTemplate(t, options, helmChartPath, "pod", []string{"templates/deployment.yaml"})

	// Now we use kubernetes/client-go library to render the template output into the Pod struct. This will
	// ensure the Pod resource is rendered correctly.
	var pod corev1.Pod
	helm.UnmarshalK8SYaml(t, output, &pod)

	// Finally, we verify the pod spec is set to the expected container image value
	expectedContainerImage := "nginx"
        deploymentContainers := pod.spec.template.spec.containers
	//podContainers := pod.Spec.Containers
	if deploymentContainers[0].Image != expectedContainerImage {
		t.Fatalf("Rendered container image (%s) is not expected (%s)", deploymentContainers[0].Image, expectedContainerImage)
	}
}

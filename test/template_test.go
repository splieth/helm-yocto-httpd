package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	core "k8s.io/api/core/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
)

func TestTemplateRenders(t *testing.T) {
	helmChartPath := "../"
	expectedVersion := "felixb/yocto-httpd:latest"

	helmOptions := &helm.Options{
		SetValues: map[string]string{"image": "felixb/yocto-httpd:latest"},
	}

	renderedTemplate := helm.RenderTemplate(t, helmOptions, helmChartPath, []string{"templates/pod.yaml"})

	var pod core.Pod
	helm.UnmarshalK8SYaml(t, renderedTemplate, &pod)

	assert.Equal(t, expectedVersion, pod.Spec.Containers[0].Image)
}

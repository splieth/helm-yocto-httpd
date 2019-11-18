package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/helm"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		port      int
		expected  int
		withError bool
	}{
		{
			name:      "responds to /",
			path:      "/",
			port:      8080,
			expected:  200,
			withError: false,
		},
		{
			name:      "responds to other paths as well",
			path:      "/some-path",
			port:      8080,
			expected:  200,
			withError: false,
		},
		{
			name:      "fails on ther ports",
			path:      "/",
			port:      8081,
			expected:  503,
			withError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helmChartPath := "../"

			kubectlOptions := k8s.NewKubectlOptions("", "", "default")
			helmOptions := &helm.Options{
				SetValues: map[string]string{"image": "felixb/yocto-httpd:latest"},
			}

			releaseName := fmt.Sprintf("yocto-%s", strings.ToLower(random.UniqueId()))
			defer helm.Delete(t, helmOptions, releaseName, true)

			helm.Install(t, helmOptions, helmChartPath, releaseName)

			podName := fmt.Sprintf("%s-yocto-httpd", releaseName)
			if tt.withError {
				assert.NotNil(t, assertPodResponse(t, kubectlOptions, podName, tt.path, tt.port, tt.expected))
			} else {
				assert.Nil(t, assertPodResponse(t, kubectlOptions, podName, tt.path, tt.port, tt.expected))
			}
		})
	}
}

func assertPodResponse(t *testing.T, kubectlOptions *k8s.KubectlOptions, podName, path string, port, expected int) error {
	retries := 3
	sleep := 10 * time.Second
	k8s.WaitUntilPodAvailable(t, kubectlOptions, podName, retries, sleep)

	tunnel := k8s.NewTunnel(kubectlOptions, k8s.ResourceTypePod, podName, port, port)
	defer tunnel.Close()
	tunnel.ForwardPort(t)

	endpoint := fmt.Sprintf("http://%s%s", tunnel.Endpoint(), path)
	return http_helper.HttpGetWithRetryWithCustomValidationE(
		t,
		endpoint,
		nil,
		retries,
		sleep,
		func(statusCode int, body string) bool {
			return statusCode == expected && strings.Contains(body, "OK")
		},
	)
}

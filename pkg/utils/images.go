package utils

import "os"

type ComplianceComponent uint

const (
	OPENSCAP = iota
	OPERATOR
	DEFAULT_PROFILE_BUNDLES
)

var componentDefaults = []struct {
	defaultImage string
	envVar       string
}{
	{"quay.io/compliance-operator/openscap-ocp:1.3.3", "RELATED_IMAGE_OPENSCAP"},
	{"quay.io/compliance-operator/compliance-operator:latest", "RELATED_IMAGE_OPERATOR"},
	{"quay.io/complianceascode/ocp4:latest", "RELATED_IMAGE_PROFILE"},
}

// GetComponentImage returns a full image pull spec for a given component
// based on the component type
func GetComponentImage(component ComplianceComponent) string {
	comp := componentDefaults[component]

	imageTag := os.Getenv(comp.envVar)
	if imageTag == "" {
		imageTag = comp.defaultImage
	}
	return imageTag
}

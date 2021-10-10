package fns

import (
	"fmt"

	"github.com/yuekcc/fns/fnspec"
)

func generateContainerLabels(spec *fnspec.FunctionSpec) map[string]string {
	labels := map[string]string{
		"fns-spec-version": fmt.Sprintf("%d", fnspec.SPEC_VERSION),
		"fns-route":        fmt.Sprintf("/%s/v%d/%s", spec.Metadata.ServiceName, spec.Metadata.Version, spec.Metadata.Name),
		"traefik.enable": "true",
		"traefik.http.middlewares.response-compress.compress": "true",
	}

	return labels
}

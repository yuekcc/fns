package fnspec

const (
	SPEC_VERSION = 1
)

type Spec struct {
	SpecVersion int `json:"specVersion,omitempty"` // 默认值 1
}

type FunctionSpec struct {
	Spec
	Metadata struct {
		Name              string `json:"name"`              // 命名 ^[A-Za-z][A-za-z0-9_]*$
		Version           int    `json:"version"`           // 函数服务版本号
		ServiceName       string `json:"serviceName"`       // 命名 ^[A-Za-z][A-za-z0-9_]*$
		Sdk               string `json:"sdk"`               // 值 ^[nodejs|go|wasm]-\d*$
		LaunchCommandLine string `json:"launchCommandLine"` // 启动命令行
	}
}

package contree

type ExtensionFile string

const (
	YAML   ExtensionFile = "YAML_FORMAT"
	TOML   ExtensionFile = "TOML_FORMAT"
	JSON   ExtensionFile = "JSON_FORMAT"
	XML    ExtensionFile = "XML_FORMAT"
	DOTENV ExtensionFile = "DOTENV_FORMAT"
)

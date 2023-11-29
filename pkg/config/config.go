package config

type Config struct {
	Listen string `json:"listen"`

	GinMode string `json:"gin_mode"`

	StaticFileDir string `json:"static_file_dir"`

	LogDir string `json:"log_dir"`

	Debug bool `json:"debug"`
}

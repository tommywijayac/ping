package config

type Config struct {
	DisplayRoom map[int]DisplayRoom `yaml:"room"`
}

type DisplayRoom struct {
	Title    string `yaml:"title"`
	IconPath string `yaml:"icon_path"`
}

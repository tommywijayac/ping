package config

type Config struct {
	RoomConfig []RoomConfig `yaml:"rooms"`
}

type RoomConfig struct {
	ID       int    `yaml:"id"`
	Title    string `yaml:"title"`
	IconPath string `yaml:"icon_path"`
}

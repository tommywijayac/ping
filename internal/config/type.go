package config

type Config struct {
	RoomConfig []RoomConfig `yaml:"rooms"`
	PingDelay  int64        `yaml:"ping_delay" default:"5"`
}

type RoomConfig struct {
	ID       int    `yaml:"id"`
	Title    string `yaml:"title"`
	IconPath string `yaml:"icon_path" default:"default.png"` //empty icon path will cause client not working
}

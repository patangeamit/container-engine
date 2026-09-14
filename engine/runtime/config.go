package runtime

type Config struct {
	Name       string     `yaml:"name"`
	Namespaces Namespaces `yaml:"namespaces"`
	EntryPoint []string   `yaml:"entryPoint"`
}

type Namespaces struct {
	Hostname bool `yaml:"hostname"`
	Network  bool `yaml:"network"`
	Ipc      bool `yaml:"ipc"`
	Process  bool `yaml:"process"`
	Time     bool `yaml:"time"`
	Mount    bool `yaml:"mount"`
	Cgroup   bool `yaml:"cgroup"`
	User     bool `yaml:"user"`
}

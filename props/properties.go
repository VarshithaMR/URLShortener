package props

type Properties struct {
	Server ServerProps `yaml:"server"`
}

// ServerProps has the properties needed to create Server
type ServerProps struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	ContextRoot string `yaml:"contextroot"`
}

package swag

type Spec struct {
	Version, Host, BasePath, Title, Description, InfoInstanceName, SwaggerTemplate string
	Schemes                                                                        []string
}

func (s *Spec) InstanceName() string {
	if s.InfoInstanceName == "" {
		return "swagger"
	}
	return s.InfoInstanceName
}
func Register(string, *Spec) {}

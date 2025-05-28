package manager

var registrySchema map[string]AIService

func init() {
	registrySchema = make(map[string]AIService)
}

// RegisterAIService registers an ai creator function to the registry.
func RegisterAIService(creator AIService) {
	registrySchema[creator.Name()] = creator
}

func GetAIService(name string) (AIService, bool) {
	if service, ok := registrySchema[name]; ok {
		return service, true
	}
	return nil, false
}

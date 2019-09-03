package providers

type ProviderType int

const (
	ProviderGmailType ProviderType = 1 << iota
)

var instance *GmailProvider = nil


func GetProvider(t ProviderType) Provider {
	switch t {
	case ProviderGmailType:
		if instance == nil {
			instance = NewGmailInstance()
		}
		return instance
	}

	return nil
}

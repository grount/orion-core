package gmail

type ProviderType int

const (
	ProviderGmailType ProviderType = 1 << iota
)

func NewProvider(t ProviderType) Provider {
	switch t {
	case ProviderGmailType:
		return NewGmailProvider()
	}
}
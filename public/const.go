package public

const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"
	AdminSessionInfoKey = "AdminSessionInfoKey"

	LoadTypeHTTP = 0
	LoadTypeTCP = 1
	LoadTypeGRPC = 2

	HTTPPRuleTypePrefixURL = 0
	HTTPRuleTypeDomain = 1


)

var (
	LoadTypeMap =map[int]string{
		LoadTypeHTTP: "HTTP",
		LoadTypeTCP: "TCP",
		LoadTypeGRPC: "GRPC",
	}
)

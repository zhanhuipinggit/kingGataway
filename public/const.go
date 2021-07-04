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

	RedisFlowDayKey = "flow_day_count"
	RedisFlowHourKey = "flow_hour_count"

	FlowTotal = "flow_total"
	FlowCountServicePrefix = "flow_service_"
	FlowCountAppPrefix = "flow_app_"

)

var (
	LoadTypeMap =map[int]string{
		LoadTypeHTTP: "HTTP",
		LoadTypeTCP: "TCP",
		LoadTypeGRPC: "GRPC",
	}
)

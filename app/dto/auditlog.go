package dto

type AuditLogCreateValidation struct {
	ID          string `json:"id" form:"id" binding:"required,omitempty,uuid"`
	UserID      string `json:"user_id,omitempty"`
	IPAddress   string `json:"ip_address"`
	ServiceName string `json:"service_name"` // user
	MethodName  string `json:"method_name"`  // create user
	Level       string `json:"level"`        // info or error
	Metadata    string `json:"metadata"`     // response
}

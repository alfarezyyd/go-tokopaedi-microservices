package mapper

import "go-tokopaedi-microservices/model"

func FuncMapAuditable[S model.HasAuditable, R model.HasAuditableResponse](
	entityObject S,
	responseObject R,
) {
	responseObject.SetAuditableResponse(
		AuditableEntityIntoEntityResponse(entityObject.GetAuditable()),
	)
}

func FuncMapSimpleAuditable[S model.HasSimpleAuditable, R model.HasSimpleAuditableResponse](
	entityObject S,
	responseObject R,
) {
	responseObject.SetSimpleAuditableResponse(
		SimpleAuditableEntityIntoSimpleEntityResponse(entityObject.GetSimpleAuditable()),
	)
}

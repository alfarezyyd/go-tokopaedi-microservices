package model

import (
	"time"

	"gorm.io/gorm"
)

type Auditable struct {
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	CreatedBy string         `gorm:"<-:create" json:"created_by"`
	UpdatedAt time.Time      `json:"updated_at"`
	UpdatedBy string         `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	DeletedBy *string        `json:"deleted_by"`
}

type SimpleAuditable struct {
	CreatedAt time.Time `gorm:"<-:create" json:"created_at"`
	CreatedBy string    `gorm:"<-:create" json:"created_by"`
}
type AuditableResponse struct {
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedBy string `json:"updated_by"`
	UpdatedAt string `json:"updated_at"`
	DeletedBy string `json:"deleted_by"`
	DeletedAt string `json:"deleted_at"`
}

type SimpleAuditableResponse struct {
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
}

type HasAuditableResponse interface {
	GetAuditableResponse() *AuditableResponse
	SetAuditableResponse(auditableResponse *AuditableResponse)
}
type HasSimpleAuditableResponse interface {
	GetSimpleAuditableResponse() *SimpleAuditableResponse
	SetSimpleAuditableResponse(simpleAuditableResponse *SimpleAuditableResponse)
}

type HasAuditable interface {
	GetAuditable() *Auditable
}
type HasSimpleAuditable interface {
	GetSimpleAuditable() *SimpleAuditable
}

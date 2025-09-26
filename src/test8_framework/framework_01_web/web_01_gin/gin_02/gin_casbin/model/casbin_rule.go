package model

type CasbinRule struct {
	ID    uint `gorm:"primaryKey"`
	Ptype string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}

func (*CasbinRule) TableName() string {
	return "casbin_rule"
}

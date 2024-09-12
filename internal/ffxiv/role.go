package ffxiv

import (
	"reflect"
)

type Role int

const (
	DPS Role = iota
	Healer
	Tank
	Empty
)

type Roles struct {
	Roles []Role
}

func (rs Roles) Emoji() string {
	if reflect.DeepEqual(rs.Roles, []Role{DPS}) {
		return "<:dps:1283599161660280894>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{Healer}) {
		return "<:healer:1283599065648730144>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{Tank}) {
		return "<:tank:1283599150029733958>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{DPS, Healer}) {
		return "<:healerdps:1283599139912810587>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{DPS, Tank}) {
		return "<:tankdps:1283599118861864960>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{Healer, Tank}) {
		return "<:tankhealer:1283599086158614639>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{Healer, Tank, DPS}) {
		return "<:tankhealerdps:1283599100318842951>"
	}

	return "<:AnyRole:1283599173484150916>"
}

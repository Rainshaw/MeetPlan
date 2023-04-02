// Code generated by Validator v0.1.4. DO NOT EDIT.

package user

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *User) IsValid() error {
	if p.Id != nil {
		if *p.Id < int64(1) {
			return fmt.Errorf("field Id ge rule failed, current value: %v", *p.Id)
		}
	}
	if p.PkuId != nil {
		if len(*p.PkuId) < int(10) {
			return fmt.Errorf("field PkuId min_len rule failed, current value: %d", len(*p.PkuId))
		}
		if len(*p.PkuId) > int(10) {
			return fmt.Errorf("field PkuId max_len rule failed, current value: %d", len(*p.PkuId))
		}
		_src := "^[0-9]+$"
		if ok, _ := regexp.MatchString(_src, *p.PkuId); !ok {
			return fmt.Errorf("field PkuId pattern rule failed, current value: %v", *p.PkuId)
		}
	}
	return nil
}
func (p *LoginReq) IsValid() error {
	return nil
}
func (p *LoginResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *GetUserReq) IsValid() error {
	if p.Id != nil {
		if *p.Id < int64(1) {
			return fmt.Errorf("field Id ge rule failed, current value: %v", *p.Id)
		}
	}
	if p.PkuId != nil {
		if len(*p.PkuId) < int(10) {
			return fmt.Errorf("field PkuId min_len rule failed, current value: %d", len(*p.PkuId))
		}
		if len(*p.PkuId) > int(10) {
			return fmt.Errorf("field PkuId max_len rule failed, current value: %d", len(*p.PkuId))
		}
		_src := "^[0-9]+$"
		if ok, _ := regexp.MatchString(_src, *p.PkuId); !ok {
			return fmt.Errorf("field PkuId pattern rule failed, current value: %v", *p.PkuId)
		}
	}
	return nil
}
func (p *GetUserResp) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *MGetUserReq) IsValid() error {
	return nil
}
func (p *MGetUserResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *QueryUserReq) IsValid() error {
	if p.PageParam != nil {
		if err := p.PageParam.IsValid(); err != nil {
			return fmt.Errorf("filed PageParam not valid, %w", err)
		}
	}
	return nil
}
func (p *QueryUserResp) IsValid() error {
	if p.PageParam != nil {
		if err := p.PageParam.IsValid(); err != nil {
			return fmt.Errorf("filed PageParam not valid, %w", err)
		}
	}
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *UpdateUserReq) IsValid() error {
	if p.User == nil {
		return fmt.Errorf("field User not_nil rule failed")
	}
	if err := p.User.IsValid(); err != nil {
		return fmt.Errorf("filed User not valid, %w", err)
	}
	return nil
}
func (p *UpdateUserResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}

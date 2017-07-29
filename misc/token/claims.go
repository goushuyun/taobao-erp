/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/07/05 22:15
 */

package token

import (
	"time"

	"github.com/goushuyun/weixin-golang/seller/role"
)

type TokenType int64

const (
	_ TokenType = iota
	AppToken
	InterToken
	HospToken
)

var expires = map[TokenType]time.Duration{
	AppToken:   time.Hour * 2,
	InterToken: time.Hour * 3,
	HospToken:  time.Minute * 10,
}

var maxExpires = map[TokenType]time.Duration{
	AppToken:   time.Hour * 24 * 30,
	InterToken: time.Hour * 12,
	HospToken:  time.Minute * 30,
}

// WARN: if the Claims struct change, must update currentVersion
const currentVersion string = "1.0"

type Claims struct {
	UserId   string    `json:"sub,omitempty"`
	Mobile   string    `json:"mob,omitempty"`
	HospId   string    `json:"hsp,omitempty"` // self def
	Session  string    `json:"ses,omitempty"` // self def
	Scope    TokenType `json:"scp,omitempty"` // self def
	Role     role.Role `json:"rol,omitempty"` // self def
	IssuedAt int64     `json:"iat,omitempty"`
	SellerId string    `json:"seller_id,omitempty"`
	StoreId  string    `json:"store_id,omitempty"`
	Version  string    `json:"iss,omitempty"` // version of the claims

	//ExpiresAt int64  `json:"exp,omitempty"`
	//Audience  string `json:"aud,omitempty"` ?
	//Id        string `json:"jti,omitempty"` ?

	//NotBefore int64  `json:"nbf,omitempty"` this can be used for CanRefresh ?
	//Subject   string `json:"sub,omitempty"` this is in using `UserId`
}

// Valid is used for logical valid, but we do not use it, because we need control it behavior
func (c Claims) Valid() error {
	return nil
}

// VerifyIsExpired check if the token is expired
func (c *Claims) VerifyIsExpired() bool {
	return time.Now().Sub(time.Unix(c.IssuedAt, 0)) > expires[c.Scope]
}

// VerifyMaxExpire check if the token can refresh
func (c *Claims) VerifyCanRefresh() bool {
	return time.Now().Sub(time.Unix(c.IssuedAt, 0)) < maxExpires[c.Scope]
}

// VerifyVersion check if the version equal currentVersion
func (c *Claims) VerifyVersion() bool {
	return c.Version == currentVersion
}

// HasAllRole check if it matches all role
func (c *Claims) HasAllRole(roles ...role.Role) bool {
	return c.Role.HasAll(roles...)
}

// HasOneRole check if it matches at least one
func (c *Claims) HasOneRole(roles ...role.Role) bool {
	return c.Role.HasOne(roles...)
}

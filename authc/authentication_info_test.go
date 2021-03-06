// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/security source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package authc

import (
	"testing"

	"aahframework.org/test.v0/assert"
)

func TestAuthcAuthenticationInfo(t *testing.T) {
	a1 := NewAuthenticationInfo()
	p1 := &Principal{Value: "user@sample.com", IsPrimary: true}
	p2 := &Principal{Value: "userid"}

	a1.Principals = append(a1.Principals, p1, p2)
	assert.False(t, a1.IsLocked)
	assert.False(t, a1.IsExpired)
	assert.Equal(t, "AuthenticationInfo:: Principals[Realm: , Principal: user@sample.com, IsPrimary: true Realm: , Principal: userid, IsPrimary: false], Credential: *******, IsLocked: false, IsExpired: false", a1.String())

	p := a1.PrimaryPrincipal()
	assert.NotNil(t, p)
	assert.Equal(t, "user@sample.com", p.Value)
	assert.True(t, p.IsPrimary)
	assert.Equal(t, "Realm: , Principal: user@sample.com, IsPrimary: true", p.String())

	assert.NotNil(t, a1.Principals)
	assert.True(t, len(a1.Principals) == 2)

	ReleaseAuthenticationInfo(a1)
}

func TestAuthcAuthenticationInfoMerge(t *testing.T) {
	a1 := NewAuthenticationInfo()
	a2 := NewAuthenticationInfo()
	a1.Principals = append(a1.Principals, &Principal{Value: "user@sample.com"})
	a2.IsLocked = true
	a2.IsExpired = true

	a1.Merge(a2)
	assert.True(t, a1.IsLocked)
	assert.True(t, a1.IsExpired)
	assert.Nil(t, a1.PrimaryPrincipal())
}

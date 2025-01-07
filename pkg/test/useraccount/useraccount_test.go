package useraccount

import (
	"testing"

	murtest "github.com/codeready-toolchain/toolchain-common/pkg/test/masteruserrecord"
	"github.com/stretchr/testify/assert"
)

func TestUserAccountFromMur(t *testing.T) {
	t.Run("UserAccount from MUR should have same values", func(t *testing.T) {
		// given
		mur := murtest.NewMasterUserRecord(t, "john")
		userAcc := NewUserAccountFromMur(mur)

		// when & then
		assert.Equal(t, mur.Spec.PropagatedClaims, userAcc.Spec.PropagatedClaims)
		assert.Equal(t, mur.Spec.Disabled, userAcc.Spec.Disabled)
	})
}

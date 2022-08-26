//go:generate fielder -type=UserAccount -suffix=Column -embedded=true -tag=db -tag_strict=true -excluded=FullName

package models

import (
	"github.com/kpeu3i/fielder/examples/simple/models/common"
)

type UserAccount struct {
	common.Entity

	FirstName string `db:"name"`
	LastName  string `db:"surname"`
	Email     string `db:"email"`
	Password  string `db:"password"`

	FullName string
}

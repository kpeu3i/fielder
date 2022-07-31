//go:generate fielder -type=UserAccount -suffix=Column -embedded=true -tag=db -excluded=FullName

package example

import (
	"github.com/kpeu3i/fielder/example/common"
)

type UserAccount struct {
	common.Entity

	FirstName string `db:"name"`
	LastName  string `db:"surname"`
	Email     string `db:"email"`
	Password  string `db:"password"`

	FullName string
}

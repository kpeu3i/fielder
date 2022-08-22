// Code generated by "fielder -type=UserAccount -suffix=Column -embedded=true -tag=db -excluded=FullName"; DO NOT EDIT.

package models

// UserAccountColumn represents the values of the ENUM.
type UserAccountColumn string

const (
	UserAccountColumnID        UserAccountColumn = "id"
	UserAccountColumnCreatedAt UserAccountColumn = "created_at"
	UserAccountColumnUpdatedAt UserAccountColumn = "updated_at"
	UserAccountColumnFirstName UserAccountColumn = "name"
	UserAccountColumnLastName  UserAccountColumn = "surname"
	UserAccountColumnEmail     UserAccountColumn = "email"
	UserAccountColumnPassword  UserAccountColumn = "password"
)

var _UserAccountColumnValues = [...]UserAccountColumn{
	UserAccountColumnID,
	UserAccountColumnCreatedAt,
	UserAccountColumnUpdatedAt,
	UserAccountColumnFirstName,
	UserAccountColumnLastName,
	UserAccountColumnEmail,
	UserAccountColumnPassword,
}

// IsValid returns true if the value is a valid ENUM.
func (x UserAccountColumn) IsValid() bool {
	for _, v := range _UserAccountColumnValues {
		if v == x {
			return true
		}
	}

	return false
}

// String returns the string representation of the value.
func (x UserAccountColumn) String() string {
	return string(x)
}

// *********************************************************************************************************************

// UserAccountColumnList represents a collection of UserAccountColumn.
type UserAccountColumnList []UserAccountColumn

// Len returns the number of values in the collection.
func (l UserAccountColumnList) Len() int {
	return len(l)
}

// Contains returns true if the collection contains the value.
func (l UserAccountColumnList) Contains(v UserAccountColumn) bool {
	for _, x := range l {
		if x == v {
			return true
		}
	}

	return false
}

// Equals returns true if the two collections are equal.
func (l UserAccountColumnList) Equals(other UserAccountColumnList) bool {
	if len(l) != len(other) {
		return false
	}

	for i, x := range l {
		if x != other[i] {
			return false
		}
	}

	return true
}

// Similar returns true if the two collections contain the same values.
func (l UserAccountColumnList) Similar(other UserAccountColumnList) bool {
	if len(l) != len(other) {
		return false
	}

	for _, x := range l {
		if !other.Contains(x) {
			return false
		}
	}

	return true
}

// Add adds the values to the collection.
func (l *UserAccountColumnList) Add(v ...UserAccountColumn) *UserAccountColumnList {
	*l = append(*l, v...)

	return l
}

// AddIfNotContains adds the values to the collection if they are not already present.
func (l *UserAccountColumnList) AddIfNotContains(v ...UserAccountColumn) *UserAccountColumnList {
	for _, x := range v {
		if !l.Contains(x) {
			l.Add(x)
		}
	}

	return l
}

// Remove removes the values from the collection.
func (l *UserAccountColumnList) Remove(v ...UserAccountColumn) *UserAccountColumnList {
	for _, x := range v {
		for i, y := range *l {
			if y == x {
				*l = append((*l)[:i], (*l)[i+1:]...)

				break
			}
		}
	}

	return l
}

// Clear clears the collection.
func (l *UserAccountColumnList) Clear() *UserAccountColumnList {
	*l = []UserAccountColumn{}

	return l
}

// Strings returns a slice with all the strings of the collection items.
func (l UserAccountColumnList) Strings() []string {
	strings := make([]string, 0, len(l))
	for _, x := range l {
		strings = append(strings, x.String())
	}

	return strings
}

// *********************************************************************************************************************

// UserAccountColumnValues returns a slice with all the values of the ENUM.
func UserAccountColumnValues() []UserAccountColumn {
	result := make([]UserAccountColumn, len(_UserAccountColumnValues))
	copy(result, _UserAccountColumnValues[:])

	return result
}

// UserAccountColumnStrings returns a slice with all the strings of the ENUM.
func UserAccountColumnStrings() []string {
	strings := make([]string, 0, len(_UserAccountColumnValues))
	for _, v := range _UserAccountColumnValues {
		strings = append(strings, v.String())
	}

	return strings
}

// NewUserAccountColumnList returns a new UserAccountColumnList with all the values of the ENUM.
func NewUserAccountColumnList() UserAccountColumnList {
	result := make([]UserAccountColumn, len(_UserAccountColumnValues))
	copy(result, _UserAccountColumnValues[:])

	return result
}

// NewUserAccountColumnListWith returns a new UserAccountColumnList with the given values of the ENUM.
func NewUserAccountColumnListWith(v ...UserAccountColumn) UserAccountColumnList {
	result := UserAccountColumnList{}
	if len(v) > 0 {
		result.Add(v...)
	}

	return result
}

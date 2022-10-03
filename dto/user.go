package dto

type UserDto struct {
  ID uint64   `uri:"id" form:"id" binding:"required,uuid"`
  Name string `uri:"name" form:"name" binding:"required"`
  Email string `uri:"email" form:"email" binding:"required"`
  Password string `uri:"password, omitempty" form:"password, omitempty"`

}


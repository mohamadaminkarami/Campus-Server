package forms
type SignupUserData struct {
	// binding:"required" ensures that the field is provided
	StudentNumber string `json:"studentNumber" binding:"required,min=8,max=9,numeric"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	EntranceYear  int    `json:"entranceYear" binding:"required,gte=1390"`
	SchoolId      int    `json:"schoolId" binding:"required,doesSchoolExist,gte=1"`
}

type LoginUserData struct {
	StudentNumber string `json:"studentNumber" binding:"required,min=8,max=9,numeric"`
	Password      string `json:"password" binding:"required"`
}

type UpdateUserData struct {
	StudentNumber   string    `json:"studentNumber" binding:"required,min=8,max=9,numeric"`
	Password        string    `json:"password" binding:"-"`
	Email           string    `json:"email" binding:"required,email"`
	EntranceYear    int       `json:"entranceYear" binding:"required,gte=1390"`
	TakeCoursesTime int		  `json:"takeCoursesTime" binding:"required,isTimestamp"`
	SchoolId        int       `json:"schoolId" binding:"required,doesSchoolExist,gte=1"`
}

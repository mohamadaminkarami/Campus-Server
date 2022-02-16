package forms

type SignupUserData struct {
	// binding:"required" ensures that the field is provided
	StudentNumber string `json:"studentNumber" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	EntranceYear  int    `json:"entranceYear" binding:"required"`
	SchoolId      int    `json:"schoolId" binding:"required"`
}

type LoginUserData struct {
	StudentNumber string `json:"studentNumber" binding:"required"`
	Password      string `json:"password" binding:"required"`
}

type UpdateUserData struct {
	StudentNumber   string `json:"studentNumber" binding:"required"`
	Password        string `json:"password" binding:"-"`
	Email           string `json:"email" binding:"required"`
	EntranceYear    int    `json:"entranceYear" binding:"required"`
	TakeCoursesTime int    `json:"takeCoursesTime" binding:"required"`
	SchoolId        int    `json:"schoolId" binding:"required"`
}

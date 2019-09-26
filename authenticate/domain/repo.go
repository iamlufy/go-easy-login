package domain

type LoginUserDO struct {
	Username   string
	Password   string
	IsLock     bool
	UniqueCode string
	Mobile     string
}
type LoginUserRepo interface {
	Add(*LoginUserDO)
	GetOne(username string) *LoginUserDO
	Update()
	FindOne(username string) (*LoginUserDO, bool)
	FindSmsCode(mobile string) string
}

var repo LoginUserRepo

func NewRepo(loginUserRepo LoginUserRepo) {
	repo = loginUserRepo
}

func Add(userDO *LoginUserDO) {
	repo.Add(userDO)
}

func FindUser(username string) (*LoginUserDO, bool) {
	return repo.FindOne(username)
}

func GetUser(username string) *LoginUserDO {
	return repo.GetOne(username)
}

func FindSmsCode(mobile string) string {
	return repo.FindSmsCode(mobile)

}
func GetUniqueCode(username string) string {
	return repo.GetOne(username).UniqueCode
}

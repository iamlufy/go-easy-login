package service

import (
	"oneday-infrastructure/login/domain"
	"oneday-infrastructure/login/domain/common"
)

var loginService *LoginService

type LoginService struct {
	LoginUserRepo
	TokenService
}

type TokenService interface {
	Generate(uniqueCode string, effectiveSeconds int) string
	Verify(token string)
}

type LoginUserRepo interface {
	Add(*domain.LoginUserDO)
	GetOne(username, tenantId string) *domain.LoginUserDO
	Update()
	FindSmsCode(mobile string) string
}

func NewLoginService(repo LoginUserRepo, token TokenService) *LoginService {
	if loginService == nil {
		return &LoginService{
			LoginUserRepo: repo,
			TokenService:  token,
		}
	} else {
		return loginService
	}

}

//todo application should check tenant  and consider to cache result
//tenant属于更高层次的领域，loginUser属于低层次，不应该掌握高层的信息，只能被通知
func (service *LoginService) Login(loginCmd common.LoginCmd) (string, error) {
	userDO := service.GetOne(loginCmd.Username, loginCmd.TenantId)
	userE := common.ToLoginUserE(*userDO)
	userE.EncryptWay = domain.EncryptWay(loginCmd.EncryptWay)

	encryptCode := service.encryptCode(loginCmd.LoginWay, userDO)
	if _, err := userE.DoVerify(loginCmd.SourceCode, encryptCode); err != nil {
		return "", err
	}

	//todo add login event and callback
	return service.Generate(userE.UniqueCode, loginCmd.EffectiveSeconds), nil
}

func (service *LoginService) encryptCode(way string, userDO *domain.LoginUserDO) string {
	switch way {
	case "PASSWORD":
		return userDO.Password
	case "SMS":
		return service.FindSmsCode(userDO.Mobile)
	default:
		panic("unknown login way")
	}
}

func (service *LoginService) AddUser(cmd *common.AddLoginUserCmd) {
	loginUserDO := common.ToLoginUserDO(cmd)
	loginUserDO.Password = domain.EncryptWay(cmd.EncryptWay).EncryptHelper().Encrypt(loginUserDO.Password)
	service.Add(loginUserDO)
}

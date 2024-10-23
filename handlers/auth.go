// handlers/auth.go
package handlers

import (
	"go-login-server/auth"
	"go-login-server/common"
	"go-login-server/handlers/dto"
	"go-login-server/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register godoc
// @Summary 새로운 사용자를 등록합니다.
// @Description 사용자 이름과 비밀번호를 사용하여 새로운 사용자를 등록합니다.
// @Tags auth
// @Accept json
// @Produce json
// @Param registerInput body dto.RegisterInput true "회원가입 입력 정보"
// @Success 200 {object} common.Response "성공적으로 등록되었습니다."
// @Failure 400 {object} common.Response "잘못된 입력 또는 이미 존재하는 사용자 이름"
// @Failure 500 {object} common.Response "서버 내부 오류"
// @Router /register [post]
func Register(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var input dto.RegisterInput
        if err := c.ShouldBindJSON(&input); err != nil {
            common.ErrorResponse(c, &common.AppError{
                Code:       "VALIDATION_ERROR",
                Message:    "입력값이 유효하지 않습니다.",
                StatusCode: http.StatusBadRequest,
            })
            return
        }

        // 비밀번호 길이 검증
        if len(input.Password) < 8 {
            common.ErrorResponse(c, &common.AppError{
                Code:       "PASSWORD_TOO_SHORT",
                Message:    "비밀번호는 최소 8자 이상이어야 합니다.",
                StatusCode: http.StatusBadRequest,
            })
            return
        }

        // 비밀번호 해싱
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
        if err != nil {
            common.ErrorResponse(c, common.ErrHashingPassword)
            return
        }

        user := models.User{
            Username: input.Username,
            Password: string(hashedPassword),
        }

        if err := db.Create(&user).Error; err != nil {
            common.ErrorResponse(c, common.ErrUsernameExists)
            return
        }

        common.SuccessResponse(c, http.StatusOK, "회원가입이 성공적으로 완료되었습니다.", nil)
    }
}

// Login godoc
// @Summary 사용자를 로그인합니다.
// @Description 사용자 이름과 비밀번호를 사용하여 로그인을 수행합니다.
// @Tags auth
// @Accept json
// @Produce json
// @Param loginInput body dto.LoginInput true "로그인 입력 정보"
// @Success 200 {object} common.Response "로그인이 성공적으로 완료되었습니다."
// @Failure 400 {object} common.Response "잘못된 입력"
// @Failure 401 {object} common.Response "잘못된 사용자 이름 또는 비밀번호"
// @Failure 500 {object} common.Response "서버 내부 오류"
// @Router /login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var input dto.LoginInput
        if err := c.ShouldBindJSON(&input); err != nil {
            common.ErrorResponse(c, err)
            return
        }

        var user models.User
        if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
            common.ErrorResponse(c, common.ErrInvalidCredentials)
            return
        }

        // 비밀번호 검증
        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
            common.ErrorResponse(c, common.ErrInvalidCredentials)
            return
        }

        // JWT 생성
        token, err := auth.GenerateJWT(user.Username)
        if err != nil {
            common.ErrorResponse(c, common.ErrGeneratingToken)
            return
        }

        // 성공 응답
        responseData := gin.H{
            "token": token,
        }
        common.SuccessResponse(c, http.StatusOK, "로그인이 성공적으로 완료되었습니다.", responseData)
    }
}

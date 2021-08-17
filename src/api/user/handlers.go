package user

import (
	"net/http"

	cmmail "connpass-manager/common/mail"

	"github.com/labstack/echo/v4"

	cmerr "connpass-manager/common/error"
	"connpass-manager/common/session"
	"connpass-manager/db"
	"connpass-manager/usecase/user"
)

// Login ログイン処理を行う
func Login(c echo.Context) error {
	req := new(user.LoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	// ユースケース実行
	if err := user.NewLoginUseCase(db.GetConnection()).Execute(c, req); err != nil {
		return err
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// Logout ログアウト処理を行う
func Logout(c echo.Context) error {
	// ユースケース実行
	if err := user.NewLogoutUseCase(db.GetConnection()).Execute(c); err != nil {
		return err
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// Register 会員登録を行う
func Register(c echo.Context) error {
	req := new(user.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	// ユースケース実行
	if err := user.NewRegisterUseCase(db.GetConnection()).Execute(c, req); err != nil {
		return err
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// ForgotPassword パスワードリセットメールを送信する
func ForgotPassword(c echo.Context) error {
	req := new(user.ForgotPasswordRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	if err := user.NewForgotPasswordUseCase(db.GetConnection()).Execute(req); err != nil {
		return err
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// PasswordReset パスワードリセットを行う
func PasswordReset(c echo.Context) error {
	req := new(user.PasswordResetRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	if err := user.NewPasswordResetUseCase(db.GetConnection()).Execute(c, req); err != nil {
		return err
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// Withdraw 退会処理を行う
func Withdraw(c echo.Context) error {
	userID := session.GetUserID(c)
	if userID == nil {
		return cmerr.NewApplicationError(http.StatusInternalServerError, "ログインしていません")
	}
	if err := user.NewWithdrawUseCase(db.GetConnection()).Execute(*userID); err != nil {
		return err
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// SendMailRequest FIXME 後で消す
type SendMailRequest struct {
	Email   string `json:"email" validate:"required"`
	Subject string `json:"subject" validate:"required,max=255"`
	Body    string `json:"body" validate:"required,max=1024"`
}

func SendMail(c echo.Context) error {
	req := new(SendMailRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	toList := []string{req.Email}
	subject := req.Subject
	body := req.Body

	return cmmail.NewSender().SendTextMail(toList, subject, body)

}

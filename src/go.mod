module connpass-manager

go 1.15

require (
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gorilla/sessions v1.2.1
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo-contrib v0.11.0
	github.com/labstack/echo/v4 v4.3.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.1.1

package auth

import (
	"strings"

	"github.com/Team-IF/TeamWork_Backend/models/res"

	"github.com/Team-IF/TeamWork_Backend/models"
	"github.com/Team-IF/TeamWork_Backend/utils"
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	resutil.New(c).Response(resutil.R{})
}

func Refresh(c *gin.Context) {
	// r := resutil.New(c)
	r := resutil.New(c)

	// Parsing Token From Header
	clientToken := c.GetHeader("Authorization")
	if clientToken == "" {
		r.SendError(resutil.ERR_BAD_REQUEST, "Enter Authorization header")
		return
	}

	extractedToken := strings.Split(clientToken, "Bearer ")

	// Verify if the format of the token is correct
	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])
	} else {
		r.SendError(resutil.ERR_BAD_REQUEST, "Enter correct token type")
		return
	}

	// Parsing JWT To struct
	claims := &models.Claims{}
	_, errParseWithClaims := jwt.ParseWithClaims(clientToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetConfig().Etc.JwtSecret), nil
	})

	// Check Correct OR Has Error
	if errParseWithClaims != nil {
		if errParseWithClaims.Error() == jwt.ErrSignatureInvalid.Error() {
			r.SendError(resutil.ERR_BAD_REQUEST, "Enter jwt format is not correct")
		} else {
			token, err := utils.GetJwtToken(claims.ID)
			if err != nil {
				r.SendError(resutil.ERR_AUTH, "Error while create token")
			} else {
				r.Response(res.Token{
					Token: token,
				})
			}
		}
		return
	}

	r.Response(res.Token{
		Token: clientToken,
	})
}

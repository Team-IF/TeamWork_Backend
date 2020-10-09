package middlewares

import (
	"strings"

	"github.com/Team-IF/TeamWork_Backend/db"
	"github.com/Team-IF/TeamWork_Backend/models"
	"github.com/Team-IF/TeamWork_Backend/utils"
	resutil "github.com/Team-IF/TeamWork_Backend/utils/res"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
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
				r.SendError(resutil.ERR_AUTH, "not valid token")
			}
			return
		}

		user, err := db.FindUserByID(claims.ID)
		if err != nil {
			r.SendError(resutil.ERR_SERVER, "Error while load user data")
			return
		}

		c.Set("user", user)
		c.Set("userID", claims.ID)
		c.Next()
	}
}

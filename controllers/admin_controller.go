package controllers

import (
	"cn.anydevelop/go_recruit/common"
	"cn.anydevelop/go_recruit/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AdminController struct {
	beego.Controller
}

func (adminController *AdminController) LoginAdmin() {
	var admin models.Admin
	if err := json.Unmarshal(adminController.Ctx.Input.RequestBody, &admin); err != nil {
		adminController.Data["json"] = common.Fail(err.Error())
	} else {
		o := orm.NewOrm()
		table := o.QueryTable("admin")
		source := models.Admin{}
		err = table.Filter("name", admin.Name).One(&source)
		if err != nil {
			adminController.Data["json"] = common.Fail(err.Error())
		} else {
			err = bcrypt.CompareHashAndPassword([]byte(source.Password), []byte(admin.Password))
			if err != nil {
				adminController.Data["json"] = common.Fail(err.Error())
			} else {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"adminName": source.Name,
					"datetime":  time.Now(),
				})
				signedString, sigErr := token.SignedString([]byte(beego.AppConfig.String("token_secret_key")))
				if sigErr != nil {
					adminController.Data["json"] = common.Fail(err.Error())
				} else {
					common.HashPut(source.Name, "adminName", source.Name)
					common.HashPut(source.Name, "adminPicture", source.Picture)
					common.Expire(source.Name, 30*MINUTE)
					adminController.Data["json"] = common.Success(signedString)
				}
			}
		}
	}
	adminController.ServeJSON()
}

func (adminController *AdminController) LogoutAdmin() {
	tokenStr := adminController.Ctx.Request.Header.Get("Token")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\", tokenStr.Header[\"alg\"])")
		}
		return []byte(beego.AppConfig.String("token_secret_key")), nil
	})
	if err != nil {
		beego.Debug(err)
		adminController.Data["json"] = common.Fail(err.Error())
	} else {
		if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			beego.Debug(claim["adminName"].(string))
			beego.Debug(claim["datetime"].(string))
			common.Delete(claim["adminName"].(string))
			adminController.Data["json"] = common.Success(true)
		}
	}
	adminController.ServeJSON()
}

func (adminController *AdminController) AddAdmin() {
	var admin models.Admin
	if err := json.Unmarshal(adminController.Ctx.Input.RequestBody, &admin); err != nil {
		adminController.Data["json"] = common.Fail(err.Error())
	} else {
		var hash []byte
		hash, err = bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			adminController.Data["json"] = common.Fail(err.Error())
			beego.Debug(err.Error())
		} else {
			admin.Password = string(hash)
			o := orm.NewOrm()
			if _, err = o.Insert(&admin); err != nil {
				adminController.Data["json"] = common.Fail(err.Error())
				beego.Debug(err.Error())
			} else {
				adminController.Data["json"] = common.Success(true)
			}
		}
	}
	adminController.ServeJSON()
}

func (adminController *AdminController) DeleteAdmin() {

}

func (adminController *AdminController) ModifyAdmin() {

}

func (adminController *AdminController) SearchAdmin() {

}

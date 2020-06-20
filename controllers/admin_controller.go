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
		if err := table.Filter("name", admin.Name).One(&source); err != nil {
			adminController.Data["json"] = common.Fail(err.Error())
		} else if err := bcrypt.CompareHashAndPassword([]byte(source.Password), []byte(admin.Password)); err != nil {
			adminController.Data["json"] = common.Fail(err.Error())
		} else {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"adminName": source.Name,
				"datetime":  time.Now(),
			})
			if signedString, err := token.SignedString([]byte(beego.AppConfig.String("token_secret_key"))); err != nil {
				adminController.Data["json"] = common.Fail(err.Error())
			} else {
				common.RedisHSet(source.Name, "adminPicture", source.Picture)
				common.RedisExpire(source.Name, 30*MINUTE)
				adminController.Data["json"] = common.Success(signedString)
			}
		}
	}
	adminController.ServeJSON()
}

func (adminController *AdminController) TestLogin() {
	adminToken := adminController.Ctx.Request.Header.Get("Admin-Token")
	token, _ := jwt.Parse(adminToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\", tokenStr.Header[\"alg\"])")
		}
		return []byte(beego.AppConfig.String("token_secret_key")), nil
	})
	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		key := claim["adminName"].(string)
		beego.Debug("Redis admin token keyword:" + key)
		if common.RedisExists(key) {
			adminController.Data["json"] = common.Success("Valid admin token!")
		} else {
			adminController.Data["json"] = common.Fail("Invalid admin token!")
		}
	} else {
		adminController.Data["json"] = common.Fail("Invalid admin token!")
	}
	adminController.ServeJSON()
}

func (adminController *AdminController) LogoutAdmin() {
	tokenStr := adminController.Ctx.Request.Header.Get("Admin-Token")
	if tokenStr != "" {
		beego.Debug(tokenStr)
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v\", tokenStr.Header[\"alg\"])")
			}
			return []byte(beego.AppConfig.String("token_secret_key")), nil
		})
		if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			beego.Debug(claim["adminName"].(string))
			beego.Debug(claim["datetime"].(string))
			common.RedisDelete(claim["adminName"].(string))
		}
	}
	adminController.Data["json"] = common.Success(true)
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
	name := adminController.GetString("name")
	o := orm.NewOrm()
	o.Delete(&models.Admin{Name: name})
	adminController.Data["json"] = common.Success(true)
	adminController.ServeJSONP()
}

func (adminController *AdminController) ModifyAdmin() {
	var admin models.Admin
	if err := json.Unmarshal(adminController.Ctx.Input.RequestBody, &admin); err != nil {
		adminController.Data["json"] = common.Fail(err.Error())
	} else {
		o := orm.NewOrm()
		if _, err := o.Update(&admin); err != nil {
			adminController.Data["json"] = common.Fail(err.Error())
		} else {
			adminController.Data["json"] = common.Success(true)
		}
	}
	adminController.ServeJSON()
}

func (adminController *AdminController) SearchAdmin() {
	name := adminController.GetString("name")
	o := orm.NewOrm()
	admin := &models.Admin{Name: name}
	err := o.Read(admin)
	if err != nil {
		adminController.Data["json"] = common.Fail(err.Error())
	} else {
		adminController.Data["json"] = common.Success(admin)
	}
	adminController.ServeJSON()
}

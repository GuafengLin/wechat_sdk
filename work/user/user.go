package user

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
)

const (
	userInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s"
)

// User 成员管理
type User struct {
	*context.Context
}

// NewUser 实例化
func NewUser(context *context.Context) *User {
	user := new(User)
	user.Context = context
	return user
}

// Attr 属性
type Attr struct {
	Type uint8  `json:"type"`
	Name string `json:"name"`
	Text struct {
		Value string `json:"value"`
	} `json:"text"`
	Web struct {
		Url   string `json:"url"`
		Title string `json:"title"`
	} `json:"web"`
	Miniprogram struct {
		Appid    string `json:"appid"`
		PagePath string `json:"pagepath"`
		Title    string `json:"title"`
	} `json:"miniprogram"`
}

// ExternalProfile 成员对外属性
type ExternalProfile struct {
	ExternalCorpName string `json:"external_corp_name"`
	WechatChannels   struct {
		Nickname string `json:"nickname"`
		Status   uint8  `json:"status"`
	} `json:"wechat_channels"`
	ExternalAttr []Attr `json:"external_attr"`
}

// Info 成员基本信息
type Info struct {
	util.CommonError

	UserID         string   `json:"userid"`
	Name           string   `json:"name"`
	Mobile         string   `json:"mobile"`
	Department     []uint   `json:"department"`
	Order          []uint   `json:"order"`
	Position       string   `json:"position"`
	Gender         string   `json:"gender"`
	Email          string   `json:"email"`
	IsLeaderInDept []uint8  `json:"is_leader_in_dept"`
	DirectLeader   []string `json:"direct_leader"`
	Avatar         string   `json:"avatar"`
	ThumbAvatar    string   `json:"thumb_avatar"`
	Telephone      string   `json:"telephone"`
	Alias          string   `json:"alias"`
	ExtAttr        struct {
		Attrs []Attr `json:"attrs"`
	} `json:"extattr"`
	Status           uint8  `json:"status"`
	QrCode           string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`
	Address          string `json:"address"`
	OpenUserid       string `json:"open_userid"`
	ExternalProfile  uint8  `json:"external_profile"`
	MainDepartment   uint8  `json:"main_department"`
}

// GetUserInfo 读取成员
func (user *User) GetUserInfo(userid string) (userInfo *Info, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(userInfoURL, accessToken, userid)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	userInfo = new(Info)
	err = json.Unmarshal(response, userInfo)
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo Error , errcode=%d , errmsg=%s", userInfo.ErrCode, userInfo.ErrMsg)
		return
	}
	return
}

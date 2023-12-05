package model

//
//type AccessTokenResponse struct {
//	AccessToken string `json:"access_token"`
//	ExpiresIn   int    `json:"expires_in"`
//}
//
//type WxUser struct {
//	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
//	Name      string             `bson:"name" json:"name" xml:"name"`
//	Password  string             `bson:"password,omitempty" json:"password,omitempty" xml:"password"`
//	AvatarUrl string             `bson:"AvatarUrl" json:"AvatarUrl" xml:"avatarUrl"` //头像地址
//	Phone     string             `bson:"phone" json:"phone" xml:"phone"`
//	OpenId    string             `bson:"openId" json:"openId" xml:"openId"`
//	state     bool               `bson:"state" json:"state" xml:"state"` //是否关注
//}
//
//type WxFollow struct {
//	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
//	ToUserName   string             `bson:"toUserName"json:"toUserName"`                         //接收者 开发者 微信号
//	FromUserName string             `bson:"fromUserName" json:"fromUserName" xml:"fromUserName"` //发送者 发送方帐号（一个OpenID）
//	CreateTime   int64              `bson:"createTime" json:"createTime" xml:"createTime"`       //创建时间
//	MsgType      string             `bson:"msgType" json:"msgType" xml:"msgType"`                //消息类型
//	Event        string             `bson:"event" json:"event" xml:"event"`                      //事件类型，VIEW
//	EventKey     string             `bson:"eventKey" json:"eventKey" xml:"EventKey"`             //事件KEY值，设置的跳转URL
//	Ticket       string             `bson:"ticket"json:"ticket" xml:"ticket"`                    //二维码的ticket，可用来换取二维码图片
//}
//
//type WxUserInformation struct {
//	Id             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
//	Subscribe      int                `json:"subscribe"`      //用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
//	Openid         string             `json:"openid"`         //用户的标识，对当前公众号唯一
//	Language       string             `json:"language"`       //用户的语言，简体中文为zh_CN
//	SubscribeTime  int64              `json:"subscribe_time"` //用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
//	Unionid        string             `json:"unionid"`        //只有在用户将公众号绑定到微信开放平台账号后，才会出现该字段。
//	Remark         string             `json:"remark"`         //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
//	Groupid        int                `json:"groupid"`        //用户所在的分组ID（兼容旧的用户分组接口）
//	TagidList      []int              `json:"tagid_list"`
//	SubscribeScene string             `json:"subscribe_scene"`
//	QrScene        int                `json:"qr_scene"`
//	QrSceneStr     string             `json:"qr_scene_str"`
//}

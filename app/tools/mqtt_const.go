package tools

import (
	"time"
)

// MQTT消息类型，FirstBill
type TMqttFirstBill struct {
	Topic         string    `json:"topic"`                    //发送的Topic
	Sender        string    `json:"sender"`                   //发送进程识别符
	SenderIp      string    `json:"senderip"`                 //发送者IP
	SendTime      time.Time `json:"sendtime"`                 //发送时间
	Msg_Id        string    `json:"msg_id"`                   //消息ID
	Debug_Info    string    `json:"debug_info,omitempty"`     //调试信息
	Company_Id    int64     `json:"company_id,omitempty"`     //公司id
	Team_Id       int64     `json:"team_id,omitempty"`        //团队id
	User_Id       int64     `json:"user_id,omitempty"`        //用户id
	User_Phone    string    `json:"user_phone,omitempty"`     //用户手机号
	Goods_Id      string    `json:"goods_id,omitempty"`       //商品id
	Title         string    `json:"title"`                    //标题
	Description   string    `json:"description"`              //描述
	Types         int64     `json:"types"`                    //分类
	Payway        string    `json:"payway,omitempty"`         //支付渠道
	Cash_In       float64   `json:"cash_in,omitempty"`        //现金收入
	Cash_Out      float64   `json:"cash_out,omitempty"`       //现金支出
	Card_In       int64     `json:"card_in,omitempty"`        //道具卡收入 任务-金额收入=3 任务-积分收入=4
	Card_Out      int64     `json:"card_out,omitempty"`       //道具卡支出
	Integral_Out  int64     `json:"integral_out,omitempty"`   //积分支出
	Integral_In   int64     `json:"integral_in,omitempty"`    //积分收入
	Code          string    `json:"code,omitempty"`           //兑换码
	ArrivalAmount float64   `json:"arrival_amount,omitempty"` //到账金额
	ServiceCharge float64   `json:"service_charge,omitempty"` //手续费
}

// 普通消息
type TMqttContent struct {
	Topic      string    `json:"topic"`      //发送的Topic
	Sender     string    `json:"sender"`     //发送进程识别符
	SenderIp   string    `json:"senderip"`   //发送者IP 可选
	SendTime   time.Time `json:"sendtime"`   //发送时间 可选
	Msg_Id     string    `json:"msg_id"`     //消息ID 可选
	Debug_Info string    `json:"debug_info"` //调试信息 可选
	UserId     string    `json:"userid"`     //用户ID
	Content    string    `json:"content"`    //内容
}

// MQTT消息类型，Log日志
type TMqttLogs struct {
	Topic       string    `json:"topic"`                 //发送的Topic
	Sender      string    `json:"sender"`                //发送进程识别符
	SenderIp    string    `json:"senderip"`              //发送者IP 可选
	SendTime    time.Time `json:"sendtime"`              //发送时间 可选
	Msg_Id      string    `json:"msg_id"`                //消息ID 可选
	CmdId       string    `json:"cmdid"`                 //接口名称
	Debug_Info  string    `json:"debug_info"`            //调试信息 可选
	Body        string    `json:"body,omitempt"`         //内容
	Result      string    `json:"result,omitempt"`       //返回结果
	IP          string    `json:"ip"`                    //IP
	Description string    `json:"description,omitempty"` //描述
	Phone       string    `json:"phone,omitempty"`       //手机号
}

// 缓存消息
type TMqttCache struct {
	Topic      string    `json:"topic"`             //发送的Topic
	Sender     string    `json:"sender"`            //发送进程识别符
	SenderIp   string    `json:"senderip"`          //发送者IP 可选
	SendTime   time.Time `json:"sendtime"`          //发送时间 可选
	Msg_Id     string    `json:"msg_id"`            //消息ID 可选
	Debug_Info string    `json:"debug_info"`        //调试信息 可选
	Key        string    `json:"key"`               //缓存名称
	Value      string    `json:"value,omitempt"`    //缓存名称
	Times      string    `json:"times,omitempt"`    //缓存时间
	Platform   string    `json:"platform,omitempt"` //缓存平台
	Page       int       `json:"page,omitempt"`     //分页
}

type TMqttAuth struct {
	Topic      string    `json:"topic"`                 //发送的Topic
	Sender     string    `json:"sender"`                //发送进程识别符
	SenderIp   string    `json:"senderip"`              //发送者IP
	SendTime   time.Time `json:"sendtime"`              //发送时间
	Msg_Id     string    `json:"msg_id"`                //消息ID
	Debug_Info string    `json:"debug_info,omitempty"`  //调试信息
	UserId     int64     `json:"userid,omitempty"`      //用户ID
	Code       string    `json:"code,omitempty"`        //授权code
	SourceType int64     `json:"source_type,omitempty"` //平台id
	Remark     string    `json:"remark,omitempty"`      //备注
}

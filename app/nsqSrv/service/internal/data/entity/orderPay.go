package entity

import "time"

type OrderPay struct {
    PayId       uint64 `json:"pay_id""` //支付表ID
    OrderNumber string // 订单号，雪花算法生成
    ShopId      uint64 // 商家id
    UserId      uint64 // 用户
    UserMobile  string // 用户手机号
    PayNo       string // 支付单号，传给第三方支付 雪花算法生成

    PayAmount   int    // 支付金额
    PayTypeName string // 支付方式名称
    PayType     uint8  `json:"pay_type"`   // 1微信支付 2支付宝
    PayStatus   uint8  `json:"pay_status"` // 支付状态 0 未支付 1支付

    CreateAt time.Time `json:"createAt,omitempty" gorm:"autoCreateTime"` // 创建时间
    UpdateAt time.Time `json:"updateAt,omitempty" gorm:"autoUpdateTime"` // 订单最近更新时间
}

func (*OrderPay) TableName() string {
    return "t_order_pay"
}

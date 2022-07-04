package entity

import "time"

type Cart struct {
    CartId uint64 `gorm:"column:cart_id;primary_key;"`
    //UserId   uint64 `json:"userId"  gorm:"column:user_id;index:uidx_user_shop_sku"` // 用户id
    //ShopId   uint64    `json:"shopId" gorm:"index:uidx_user_shop_sku"`                       //  店铺id
    UserId uint64 `json:"userId"  gorm:"column:user_id;"` // 用户id
    ShopId uint64 `json:"shopId"`                         //  店铺id

    SkuId    uint64 `json:"skuId" gorm:"column:sku_id;"`  // 单品ID
    ProdId   uint64 `json:"prodId" gorm:"column:prod_id;` // 商品id
    ProdNums int32  `json:"prodNums"`                     // 商品数量

    CreateAt time.Time `json:"createAt" gorm:"autoUpdateTime"` // 创建时间
    UpdateAt time.Time `json:"updateAt" gorm:"autoUpdateTime"` // 最近更新时间

    // TODO 这个成员结构应该使用prod服务的
    Sku Sku `json:"sku" gorm:"foreignKey:SkuId;references:SkuId"`
}

func (*Cart) TableName() string {
    return "t_cart"
}

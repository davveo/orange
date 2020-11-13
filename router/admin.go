package router

import (
	"Goshop/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminApi(router *gin.RouterGroup) {
	adminGroup := router.Group("admin/")
	{
		// goods相关
		adminGroup.GET("admin/goods/specs", admin.SpecsList)                          // done
		adminGroup.POST("admin/goods/specs", admin.CreateSpecs)                       // done
		adminGroup.PUT("admin/goods/specs/:spec_id", admin.UpdateSpecs)               // done
		adminGroup.DELETE("admin/goods/specs/:spec_id", admin.Specs)                  // done
		adminGroup.GET("admin/goods/specs/:spec_id", admin.DeleteSpecs)               // done
		adminGroup.GET("admin/goods/specs/:spec_id/values", admin.SpecsValues)        // done
		adminGroup.POST("admin/goods/specs/:spec_id/values", admin.UpdateSpecsValues) // done
		adminGroup.GET("admin/goods", admin.GoodsList)                                // done
		// done origin: admin/admin/goods/:goods_id/up -> admin/admin/r/goods/:goods_id/up
		adminGroup.PUT("admin/r/goods/:goods_id/up", admin.GoodsUp)
		// done origin: admin/admin/goods/:goods_id/under -> admin/admin/r/goods/:goods_id/under
		adminGroup.PUT("admin/r/goods/:goods_id/under", admin.GoodsUnder)
		adminGroup.GET("admin/goods/brands", admin.BrandList)                // done
		adminGroup.POST("admin/goods/brands", admin.CreateBrand)             // done
		adminGroup.GET("admin/goods/brands/:brand_id", admin.Brand)          // done
		adminGroup.PUT("admin/goods/brands/:brand_id", admin.UpdateBrand)    // done
		adminGroup.DELETE("admin/goods/brands/:brand_id", admin.DeleteBrand) // done
		// done // origin: admin/admin/goods/brands/all -> admin/admin/r/goods/brands/all
		adminGroup.GET("admin/r/goods/brands/all", admin.BrandAllList)
		adminGroup.GET("admin/goods/categories/:parent_id/children", admin.CategoryList) // done
		adminGroup.POST("admin/goods/categories", admin.CreateCategory)                  // done
		adminGroup.GET("admin/goodssearch/custom-words", admin.GoodsSearchCustomWord)    // done
		adminGroup.GET("admin/goodssearch/keywords", admin.GoodsSearchKeyWord)           // done
		adminGroup.GET("admin/goodssearch/goods-words", admin.GoodsSearchGoodsWord)      // done
		adminGroup.POST("admin/batch/audit", admin.GoodsBatchAudit)                      // done

		// 系统相关
		adminGroup.GET("systems/admin-users/login", admin.Login)                  // done
		adminGroup.POST("systems/admin-users/logout", admin.Logout)               // done
		adminGroup.GET("systems/roles/:roleId/checked", admin.RoleCheck)          // done
		adminGroup.GET("admin/systems/messages", admin.MessageList)               // done
		adminGroup.GET("admin/systems/admin-users/token", admin.Refresh)          // done
		adminGroup.GET("admin/systems/complain-topics", admin.ComplainTopicsList) // done
		adminGroup.GET("admin/systems/message-templates", admin.MessageTemplate)  // done
		adminGroup.GET("admin/systems/wechat-msg-tmp/sync", admin.WechatMsgSync)  // done
		adminGroup.GET("admin/systems/wechat-msg-tmp", admin.WechatMsg)           // done
		adminGroup.GET("admin/systems/logi-companies", admin.LogiCompany)         // done 物流公司相关API

		// 交易相关
		adminGroup.GET("admin/trade/orders", admin.OrderList)
		adminGroup.GET("admin/trade/orders/pay-log", admin.OrderPayLogList)
		adminGroup.GET("admin/trade/order-complains", admin.OrderComplainsList)
		// origin: admin/admin/trade/orders/:order_id -> admin/admin/r/trade/orders/:order_id
		adminGroup.GET("admin/r/trade/orders/:order_id", admin.OrderDetail)

		// promotion相关
		adminGroup.GET("admin/promotion/group-buy-actives", admin.GroupBuyList)
		adminGroup.GET("admin/promotion/group-buy-cats", admin.GroupBuyCategoryList)
		adminGroup.GET("admin/promotion/exchange-cats/:cat_id/children", admin.PointCategory)
		adminGroup.GET("admin/promotion/seckills", admin.SeckillList) // done
		adminGroup.GET("admin/promotion/coupons", admin.CouponList)   // done
		adminGroup.GET("admin/promotion/pintuan", admin.PinTuanList)  // done

		// 会员相关
		adminGroup.GET("admin/members/receipts", admin.MemberReceiptList)
		adminGroup.GET("admin/members/zpzz", admin.ZpzzList)
		adminGroup.GET("admin/members", admin.MemberList)
		adminGroup.GET("admin/members/comments", admin.MemberCommentsList)
		adminGroup.GET("admin/members/asks", admin.MemberAskList)

		// 分页查询列表
		adminGroup.GET("admin/shops/list", admin.AllShopList) // done
		// 分页查询店铺列表
		adminGroup.GET("admin/shops", admin.ShopList) // done
		//
		adminGroup.GET("admin/shops/themes", admin.ShopThemesList) // done
		//管理员禁用店铺
		adminGroup.PUT("admin/shops/disable/:shop_id", admin.DisableShop) // done
		//管理员恢复店铺使用
		adminGroup.PUT("admin/shops/enable/:shop_id", admin.EnableShop) // done
		//管理员获取店铺详细
		// origin: admin/admin/shops/:shop_id -> admin/admin/r/shops/:shop_id
		adminGroup.GET("admin/r/shops/:shop_id", admin.ShopDetail) // done
		//管理员修改审核店铺信息
		// origin: admin/admin/shops/:shop_id -> admin/admin/r/shops/:shop_id
		adminGroup.PUT("admin/r/shops/:shop_id", admin.EditShop)
		//后台添加店铺
		adminGroup.POST("admin/shops", admin.CreateShop)

		// setting相关
		adminGroup.GET("admin/settings/site", admin.SiteSetting)                   // done
		adminGroup.POST("admin/settings/site", admin.SaveSiteSetting)              // done
		adminGroup.GET("admin/trade/orders/setting", admin.TradeOrderSetting)      // done
		adminGroup.POST("admin/trade/orders/setting", admin.SaveTradeOrderSetting) // done
		adminGroup.GET("admin/settings/point", admin.PointSetting)                 // done
		adminGroup.POST("admin/settings/point", admin.SavePointSetting)            // done
		adminGroup.GET("admin/goods/settings", admin.GoodsSetting)                 // done
		adminGroup.POST("admin/goods/settings", admin.SaveGoodsSetting)            // done

		// 分销相关
		adminGroup.GET("admin/distribution/commission-tpl", admin.DistributionCommissionTpl) // wait to do
		adminGroup.GET("admin/distribution/upgradelog", admin.DistributionUpgradeLog)        // wait to do
		adminGroup.GET("admin/distribution/member", admin.DistributionMember)                // wait to do
		adminGroup.GET("admin/distribution/bill/total", admin.DistributionBillTotal)         // wait to do
		adminGroup.GET("admin/distribution/settings", admin.DistributionSetting)             // wait to do
		adminGroup.GET("admin/distribution/withdraw/apply", admin.DistributionWithdraw)      // done

		adminGroup.GET("admin/payment/payment-methods", admin.PaymentMethod) // wait to do
		adminGroup.GET("admin/index/page", admin.Index)                      // done
		adminGroup.GET("admin/after-sales", admin.AfterSalesList)
		adminGroup.GET("admin/after-sales/refund", admin.AfterSalesRefundList)

		adminGroup.GET("admin/order/bills/statistics", admin.OrderBillStatisticList)

		adminGroup.GET("admin/pages/site-navigations", admin.PageSiteNavigationList)
		adminGroup.GET("admin/pages/client_type/page_type", admin.Page) // 替换原来admin/admin/pages/PC/INDEX
		adminGroup.GET("admin/focus-pictures", admin.FocusPicture)
		adminGroup.GET("admin/pages/hot-keywords", admin.HotKeyWordsList)

		// 数据统计相关
		adminGroup.GET("admin/statistics/member/order/quantity", admin.StatisticMemberOrderQuantity)          // wait to do
		adminGroup.GET("admin/statistics/member/order/quantity/page", admin.StatisticMemberOrderQuantityPage) // wait to do
		adminGroup.GET("admin/statistics/member/increase/member", admin.StatisticMemberIncrease)              // wait to do
		adminGroup.GET("admin/statistics/member/increase/member/page", admin.StatisticMemberIncreasePage)     // wait to do
		adminGroup.GET("admin/statistics/goods/price/sales", admin.StatisticGoodsPrice)                       // wait to do
		adminGroup.GET("admin/statistics/goods/hot/money", admin.StatisticGoodsHot)                           // wait to do
		adminGroup.GET("admin/statistics/goods/hot/money/page", admin.StatisticGoodsHotPage)                  // wait to do
		adminGroup.GET("admin/statistics/goods/sale/details", admin.StatisticGoodsSaleDetail)                 // wait to do
		adminGroup.GET("admin/statistics/goods/collect", admin.StatisticGoodsCollect)                         // wait to do
		adminGroup.GET("admin/statistics/goods/collect/page", admin.StatisticGoodsCollectPage)                // wait to do
		adminGroup.GET("admin/statistics/industry/order/quantity", admin.StatisticIndustry)                   // wait to do
		adminGroup.GET("admin/statistics/industry/overview", admin.StatisticIndustryOverView)                 // wait to do
		adminGroup.GET("admin/statistics/page_view/shop", admin.StatisticPageViewShop)                        // wait to do
		adminGroup.GET("admin/statistics/page_view/goods", admin.StatisticPageViewGoods)                      // wait to do
		adminGroup.GET("admin/statistics/order/order/page", admin.StatisticPageOrder)                         // wait to do
		adminGroup.GET("admin/statistics/order/order/money", admin.StatisticPageMoney)                        // wait to do
		adminGroup.GET("admin/statistics/order/sales/money", admin.StatisticOrderSalesMoney)                  // wait to do
		adminGroup.GET("admin/statistics/order/sales/total", admin.StatisticOrderSalesTotal)                  // wait to do
		adminGroup.GET("admin/statistics/order/region/form", admin.StatisticOrderRegionForm)                  // wait to do
		adminGroup.GET("admin/statistics/order/region/member", admin.StatisticOrderRegionMember)              // wait to do
		adminGroup.GET("admin/statistics/order/unit/price", admin.StatisticOrderUnitPrice)                    // wait to do
		adminGroup.GET("admin/statistics/order/return/money", admin.StatisticOrderReturnMoney)                // wait to do

		adminGroup.GET("admin/task/:task_type", admin.AdminTask)         // wait to do
		adminGroup.GET("admin/page-create/input", admin.PageCreateInput) // wait to do

		adminGroup.GET("admin/pages/articles", admin.ArticleList)                  // todo
		adminGroup.POST("admin/pages/articles", admin.CreateArticle)               // todo
		adminGroup.GET("admin/pages/articles/:article_id", admin.GetArticle)       // todo
		adminGroup.PUT("admin/pages/articles/:article_id", admin.UpdateArticle)    // todo
		adminGroup.DELETE("admin/pages/articles/:article_id", admin.DeleteArticle) // todo

		adminGroup.GET("admin/pages/article-categories", admin.ArticleCategoriesList)                 // wait to do
		adminGroup.GET("admin/pages/article-categories/childrens", admin.ArticleCategoryChildrenList) // wait to do

		adminGroup.GET("admin/services", admin.ServiceList)                                                         // wait to do
		adminGroup.GET("admin/services/live-video-api/instances", admin.ServiceLiveVideo)                           // wait to do
		adminGroup.GET("admin/services/live-video-api/instances/:instance_id/logs", admin.ServiceLiveVideoInstance) // wait to do

		// 未发现
		adminGroup.GET("admin/members/deposit/recharge", admin.MemberDepositRechargeList)
		adminGroup.GET("admin/members/deposit/log", admin.MemberDepositLogList)
	}
}

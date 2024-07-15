package tools

import (
	"fmt"
	"github.com/gogf/gcache-adapter/adapter"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 设置session缓存
func Session_Set(r *ghttp.Request) {
	if g.IsEmpty(r.GetString("key")) {
		Return_Status(r, nil, "400", "key不能为空")
		return
	}
	var times = "86400"
	if !g.IsEmpty(r.GetString("times")) {
		times = r.GetString("times")
	}
	g.Redis("session").DoVar("SET", r.GetString("key"), r.GetString("value"))
	g.Redis("session").DoVar("EXPIRE", r.GetString("key"), times)
	Return_Status(r, nil, "", "")
	return
}

// 获取session缓存
func Session_Get(r *ghttp.Request) {
	if g.IsEmpty(r.GetString("key")) {
		Return_Status(r, nil, "400", "key不能为空")
		return
	}
	mData, _ := g.Redis("session").DoVar("GET", r.GetString("key"))
	mIsMap := mData.Maps()
	if len(mIsMap) == 0 { //不是Maps
		r.Response.WriteJson(g.Map{
			"code": "200",
			"data": mData.Map(),
			"msg":  "success",
		})
	} else {
		r.Response.WriteJson(g.Map{
			"code": "200",
			"data": mData.Maps(),
			"msg":  "success",
		})
	}
	return
}

// 缓存适配
func Set_Adapter_Redis() {
	//缓存适配
	adapter := adapter.NewRedis(g.Redis())
	g.DB().GetCache().SetAdapter(adapter)
	g.DB("book").GetCache().SetAdapter(adapter)
	g.DB("chat").GetCache().SetAdapter(adapter)
	g.DB("fusion").GetCache().SetAdapter(adapter)
	g.DB("hot_or_history").GetCache().SetAdapter(adapter)
	g.DB("log").GetCache().SetAdapter(adapter)
	g.DB("mall").GetCache().SetAdapter(adapter)
	g.DB("notice").GetCache().SetAdapter(adapter)
	g.DB("statistical").GetCache().SetAdapter(adapter)
	g.DB("task").GetCache().SetAdapter(adapter)

}

// 添加缓存
func Add_Redis(key string, value string, times string, platform string) {
	fmt.Println("添加缓存：", key)
	var mcmd TMqttCache
	mcmd.Topic = "Cache/AddRedis"
	mcmd.Sender = "add_redis"
	mcmd.SenderIp = MyIP()
	mcmd.Key = key
	mcmd.Value = value
	if !g.IsEmpty(times) {
		mcmd.Times = times
	}
	if !g.IsEmpty(platform) {
		mcmd.Platform = platform
	}
	Publish_Cache(mcmd)
}

// 删除redis缓存
func Del_Redis(key string, platform string) {
	fmt.Println("删除缓存：", key)
	var mcmd TMqttCache
	mcmd.Topic = "Cache/DelRedis"
	mcmd.Sender = "del_redis"
	mcmd.SenderIp = MyIP()
	mcmd.Key = key
	if !g.IsEmpty(platform) {
		mcmd.Platform = platform
	}
	Publish_Cache(mcmd)
}

// 批量删除分页的redis缓存
func Del_Page_Redis(key string, page int, platform string) {
	fmt.Println("删除分页缓存：", key)
	var mcmd TMqttCache
	mcmd.Topic = "Cache/DelPageRedis"
	mcmd.Sender = "del_redis"
	mcmd.SenderIp = MyIP()
	mcmd.Key = key
	if !g.IsEmpty(page) {
		mcmd.Page = page
	}
	if !g.IsEmpty(platform) {
		mcmd.Platform = platform
	}
	Publish_Cache(mcmd)
}

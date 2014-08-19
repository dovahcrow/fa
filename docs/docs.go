package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

var rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/admin","description":""},{"path":"/common","description":""}],"info":{"title":"羽球之家网站API","description":"羽球之家网站API。","contact":"doomsplayer@gmail.com"}}`
var subapi string = `{"/admin":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/admin","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/carousel","description":"","operations":[{"httpMethod":"PUT","nickname":"add carousel","type":"","summary":"添加首页轮播图","parameters":[{"paramType":"form","name":"title","description":"标题","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"subtitle","description":"小标题","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"url","description":"Url","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"picid","description":"图片的Id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"结果json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]},{"path":"/promotion","description":"","operations":[{"httpMethod":"PUT","nickname":"mainBar","type":"","summary":"添加促销商品","parameters":[{"paramType":"form","name":"title","description":"标题","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"title2","description":"小标题","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"description","description":"描述","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"description2","description":"小描述","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"type","description":"商品类型","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"link","description":"连接地址","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"detail","description":"商品细节","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"picid","description":"大图片id","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"starttime","description":"促销开始时间，用格式 2014-09-01 13:11","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"endtime","description":"促销结束时间","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/promotion","description":"","operations":[{"httpMethod":"DELETE","nickname":"mainBar","type":"","summary":"删除促销商品","parameters":[{"paramType":"query","name":"id","description":"要删除的id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/promotiontypes","description":"","operations":[{"httpMethod":"PUT","nickname":"PromotionTypes","type":"","summary":"添加商品分类","parameters":[{"paramType":"form","name":"name","description":"商品分类的名字","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/promotiontypes","description":"","operations":[{"httpMethod":"DELETE","nickname":"PromotionTypes","type":"","summary":"删除商品分类","parameters":[{"paramType":"query","name":"id","description":"商品分类的id","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/tutorial","description":"","operations":[{"httpMethod":"PUT","nickname":"tutorial","type":"","summary":"添加教程","parameters":[{"paramType":"form","name":"title","description":"标题","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"type","description":"教程类型","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"picid","description":"大图片id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"content","description":"教程内容","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"author","description":"教程作者","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"form","name":"source","description":"教程来源","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/tutorial","description":"","operations":[{"httpMethod":"DELETE","nickname":"tutorial","type":"","summary":"删除教程","parameters":[{"paramType":"query","name":"id","description":"要删除的id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/tutorialtypes","description":"","operations":[{"httpMethod":"PUT","nickname":"tutorialTypes","type":"","summary":"添加文章分类","parameters":[{"paramType":"form","name":"name","description":"文章分类的名字","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/tutorialtypes","description":"","operations":[{"httpMethod":"DELETE","nickname":"TutorialTypes","type":"","summary":"删除文章分类","parameters":[{"paramType":"query","name":"id","description":"文章分类的id","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/video","description":"","operations":[{"httpMethod":"PUT","nickname":"video","type":"","summary":"添加视频","parameters":[{"paramType":"form","name":"title","description":"标题","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"type","description":"视频类型","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"picid","description":"大图片id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"content","description":"视频内容（url）","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"author","description":"视频作者","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"form","name":"source","description":"视频来源","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/video","description":"","operations":[{"httpMethod":"DELETE","nickname":"video","type":"","summary":"删除视频","parameters":[{"paramType":"query","name":"id","description":"要删除的id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/videotypes","description":"","operations":[{"httpMethod":"PUT","nickname":"videoTypes","type":"","summary":"添加视频分类","parameters":[{"paramType":"form","name":"name","description":"视频分类的名字","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/videotypes","description":"","operations":[{"httpMethod":"DELETE","nickname":"VideoTypes","type":"","summary":"删除视频分类","parameters":[{"paramType":"query","name":"id","description":"视频分类的id","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]}]},"/common":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/common","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/mainBar","description":"","operations":[{"httpMethod":"GET","nickname":"mainBar","type":"","summary":"获得首页目录","responseMessages":[{"code":200,"message":"目录的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/upload","description":"","operations":[{"httpMethod":"PUT","nickname":"Upload File","type":"","summary":"上传文件","parameters":[{"paramType":"form","name":"file","description":"上传的文件","dataType":"file","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"文件Id的json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]},{"path":"/upload","description":"","operations":[{"httpMethod":"DELETE","nickname":"Delete File","type":"","summary":"删除文件","parameters":[{"paramType":"query","name":"id","description":"文件id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"返回json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]},{"path":"/upload","description":"","operations":[{"httpMethod":"GET","nickname":"Get File","type":"","summary":"获得文件","parameters":[{"paramType":"query","name":"id","description":"文件id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"返回json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]},{"path":"/carousel","description":"","operations":[{"httpMethod":"GET","nickname":"list carousel","type":"","summary":"拉取首页轮播图","parameters":[{"paramType":"query","name":"num","description":"拉取几张","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"轮播的json","responseModel":""},{"code":404,"message":"Not found errmsg","responseModel":""},{"code":401,"message":"Need Permmision errmsg","responseModel":""},{"code":500,"message":"Server Error errmsg","responseModel":""}]}]},{"path":"/promotion","description":"","operations":[{"httpMethod":"GET","nickname":"promotion","type":"","summary":"获得促销商品","parameters":[{"paramType":"query","name":"from","description":"从第几项开始,0是第一个,默认0","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"num","description":"几个，默认1","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"type","description":"类型（球鞋拍子blahblah）","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/promotiontypes","description":"","operations":[{"httpMethod":"GET","nickname":"mainBar","type":"","summary":"获取促销商品分类","responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/tutorial","description":"","operations":[{"httpMethod":"GET","nickname":"tutorial","type":"","summary":"获得教程","parameters":[{"paramType":"query","name":"from","description":"从第几项开始,0是第一个,默认0","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"num","description":"几个，默认1","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"type","description":"类型 大话羽球blahblah","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/hottutorial","description":"","operations":[{"httpMethod":"GET","nickname":"tutorial","type":"","summary":"获得热门教程","parameters":[{"paramType":"query","name":"from","description":"从第几项开始,0是第一个,默认0","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"num","description":"几个，默认1","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"type","description":"类型 大话羽球blahblah","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/tutorialtypes","description":"","operations":[{"httpMethod":"GET","nickname":"tutorial","type":"","summary":"获取文章分类","responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/tutorial/:id","description":"","operations":[{"httpMethod":"GET","nickname":"tutorial","type":"","summary":"获得教程详细内容","parameters":[{"paramType":"query","name":"id","description":"id","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/video","description":"","operations":[{"httpMethod":"GET","nickname":"video","type":"","summary":"获得视频","parameters":[{"paramType":"query","name":"from","description":"从第几项开始,0是第一个,默认0","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"num","description":"几个，默认1","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"type","description":"类型 大话羽球blahblah","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/hotvideo","description":"","operations":[{"httpMethod":"GET","nickname":"video","type":"","summary":"获得热门视频","parameters":[{"paramType":"query","name":"from","description":"从第几项开始,0是第一个,默认0","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"num","description":"几个，默认1","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"type","description":"类型 大话羽球blahblah","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]},{"path":"/videotypes","description":"","operations":[{"httpMethod":"GET","nickname":"video","type":"","summary":"获取视频分类","responseMessages":[{"code":200,"message":"列表的json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]}]}}`
var rootapi swagger.ResourceListing

var apilist map[string]*swagger.ApiDeclaration

func init() {
	basepath := "/api"
	err := json.Unmarshal([]byte(rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = basepath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}

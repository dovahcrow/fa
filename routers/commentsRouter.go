package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["badmintonhome/controllers/api:CarouselApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:CarouselApi"],
		beego.ControllerComments{
			"List",
			"/carousel",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"],
		beego.ControllerComments{
			"UploadFile",
			"/upload",
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"],
		beego.ControllerComments{
			"DeleteFile",
			"/upload",
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"],
		beego.ControllerComments{
			"FetchFile",
			"/upload",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenuApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenuApi"],
		beego.ControllerComments{
			"Menu",
			"/mainBar",
			[]string{"get"},
			nil})

}

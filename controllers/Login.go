package controllers



type Logincontroller struct {
	BaseController
}
func (this *Logincontroller)Index(){
	this.TplName="login.html"
}

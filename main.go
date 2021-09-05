package main

import "github.com/casbin/casbin"

func main() {
	e,_:= InitApp()
	e.Print()
}




func NewEnforcer(d Data) *casbin.Enforcer {
	//return casbin.NewEnforcer("./rbac.conf", "./rbac_policy.csv")
	return casbin.NewEnforcer(d.Conf, d.Policy)
}

func Save(e *casbin.Enforcer, d Data) (*casbin.Enforcer, error) {
	err := e.SavePolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}

func New() Data {
	return  Data{
		Conf:   "rbac.conf",
		Policy: "rbac_policy.csv",
	}
}

type Data struct {
	Conf string
	Policy string
}

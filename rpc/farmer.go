package rpc

var (
	Farmer *Endpoint = &Endpoint{Name: "farmer", Host: defaultHost, Port: 8559}
)

func init() {
	err := Farmer.Init()
	if err != nil {
		panic(err)
	}
}

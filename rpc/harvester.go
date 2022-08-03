package rpc

var (
	Harvester *Endpoint = &Endpoint{Name: "harvester", Host: defaultHost, Port: 8560}
)

func init() {
	err := Harvester.Init()
	if err != nil {
		panic(err)
	}
}

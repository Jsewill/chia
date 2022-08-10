package rpc

var (
	Daemon *Endpoint = &Endpoint{Name: "daemon", Host: defaultHost, Port: 55400}
)

func init() {
	err := Daemon.Init()
	if err != nil {
		logErr.Panicln(err)
	}
}

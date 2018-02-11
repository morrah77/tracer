package main

func prepareAddress() string {
	if len(conf.listen) > 0 {
		return `:` + conf.listen
	}
	return `:8080`
}

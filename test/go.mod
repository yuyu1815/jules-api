module github.com/yuyu1815/jules-api/test/go

go 1.23.0

require github.com/yuyu1815/jules-api/go v1.0.0

require github.com/google/go-querystring v1.1.0 // indirect

require (
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/oauth2 v0.27.0 // indirect
)

replace github.com/yuyu1815/jules-api/go => ../../go

RUN TESTS
go test -v cmd/mediatr/mediatr_test.go cmd/mediatr/mediatr.go

SETUP GIT FOR PRIVATE REPO
git config --global url."https://danmaxdanilov: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@github.com".insteadOf "https://github.com"
export GOPRIVATE=github.com/danmaxdanilov/zts.shared
go get github.com/danmaxdanilov/zts.shared
go get github.com/danmaxdanilov/zts.shared/cmd/mediatr@v0.1.0

TAG NEW VERSION
git tag v0.1.0
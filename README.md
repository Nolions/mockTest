# mockTest

練習使用 `Gomock` 搭配 `mockgen` 進行測試

## 相關commit說明

***測試覆蓋率***

```shell
# 執行整個專案測試
go test ./...

# 取得專案每個 package 測試覆蓋率
go test ./... -cover

# 取得測試覆蓋率

# 輸出所有package的測試結果到指定檔案中
go test ./... -coverprofile=coverage.out

# 查看package 測試覆蓋率與專案測試總覆蓋率
go tool cover -func=coverage.out

# 以HTML方式呈現測試覆蓋率
go tool cover -html=coverage.out
```

***gomock***

```shell

# install
go get github.com/golang/mock/gomock

go install  github.com/golang/mock/mockgen


# 產生 要mock interface的mock檔案
mockgen -source .\game\game.go -destination .\game\game_mock.go -package game

```
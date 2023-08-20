# OpenAPI Generator with golang

## OpenAPIのyamlファイルの準備
今回は[Stoplight](https://github.com/OpenAPITools/openapi-generator)で作成

## OpenAPI Generatorのインストール
homebrewを使ってインストールする

```bash
brew install openapi-generator
```

確認
```bash
which openapi-generator
/usr/local/bin/openapi-generator
```

## コード生成
`-g`に指定できるものは[ここ](https://openapi-generator.tech/docs/generators/)に定義されている
```bash
openapi-generator generate -i {openapi-file.yaml} -g {generator-list} -o {output-directory}
```

以下のようなコードが生成される
```bash
openapi-generator
├── Dockerfile
├── README.md
├── api
│   └── openapi.yaml
├── go
│   ├── api.go
│   ├── api_default.go
│   ├── api_default_service.go
│   ├── error.go
│   ├── helpers.go
│   ├── impl.go
│   ├── logger.go
│   ├── model_get_task_200_response.go
│   ├── model_task.go
│   ├── model_user.go
│   └── routers.go
├── go.mod
├── main.go
└── reference
    └── todos.yaml

```

## サーバ起動
To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once image is built use
```
docker run --rm -it openapi
```

## GormでPostgreSQLに接続

### PostgreSQL接続情報
DockerでPostgreSQLを構築

```toContext
host: localhost
user: takumi
password: takumi
dbname: todo
port: 5434
```

### gormのインストール
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### 接続処理
`infrastructure/db.go`に以下を記述

```Go
package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error
var Db *gorm.DB

// connect database
func DbConnect() {
	dsn := "host=localhost user=takumi password=takumi dbname=todo port=5434 sslmode=disable TimeZone=Asia/Tokyo"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database !")
	}
}
```

### タスク取得処理実装
`go/api_default_service.go`の`GetTaskTaskId`関数を以下のように実装する
```Go
func (s *DefaultApiService) GetTaskTaskId(ctx context.Context, taskId interface{}) (ImplResponse, error) {
	task := Task{}
	result := infrastructure.Db.WithContext(ctx).First(&task)

	if result.Error != nil {
		return Response(500, nil), errors.New("test")
	}

	return Response(200, Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
	}), nil
}
```


## エラー集

### There were issues with the specification. The option can be disabled via validateSpec (Maven/Gradle) or --skip-validate-spec (CLI).

#### エラー文
```bash
Exception in thread "main" org.openapitools.codegen.SpecValidationException: There were issues with the specification. The option can be disabled via validateSpec (Maven/Gradle) or --skip-validate-spec (CLI).
 | Error count: 1, Warning count: 1
Errors: 
        -attribute paths.'/task/{taskId}'(get).requestBody.content with no media type is unsupported
Warnings: 
        -attribute paths.'/task/{taskId}'(get).requestBody.content with no media type is unsupported

        at org.openapitools.codegen.config.CodegenConfigurator.toContext(CodegenConfigurator.java:620)
        at org.openapitools.codegen.config.CodegenConfigurator.toClientOptInput(CodegenConfigurator.java:647)
        at org.openapitools.codegen.cmd.Generate.execute(Generate.java:479)
        at org.openapitools.codegen.cmd.OpenApiGeneratorCommand.run(OpenApiGeneratorCommand.java:32)
        at org.openapitools.codegen.OpenAPIGenerator.main(OpenAPIGenerator.java:66)
```

#### 原因
エラー文に`paths.'/task/{taskId}'(get).requestBody.content with no media type is unsupported`の記述がある。該当箇所を調査してみるとrequestBodyのcontentが宣言されているが定義されていない状態だった。

```yaml
requestBody:
  content: {}
```

#### 解決策
該当のrequestBody項目を削除して解決した。

###  failed to parse value &openapi.Task{Id:(*interface {})(nil), Title:(*interface {})(nil), Description:(*interface {})(nil)}, got error unsupported data type

#### 原因
gormがモデルの型に対応していない

#### 解決策
モデルの型を修正する。
修正前は`*interface{}`型だったが、stringに変換し対応した


## 参考
- [OpenAPI Generator - github](https://github.com/OpenAPITools/openapi-generator)
- [OpenAPI Generator - 公式](https://openapi-generator.tech)
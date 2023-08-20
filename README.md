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


## 参考
- [OpenAPI Generator - github](https://github.com/OpenAPITools/openapi-generator)
- [OpenAPI Generator - 公式](https://openapi-generator.tech)
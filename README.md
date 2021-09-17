# go-graphql-echo-boilerplate
https://dev-api.kou-programmer.com/
## ローカル
golang × gqlgen

## 自動生成コマンド
```
gqlgen
gqlgen generate
```

## 自動生成ファイル
src/graph/generated/generated.go
src/graph/generated/models_gen.go


### golang起動
```
air
```

- Playground
http://localhost:3000/playground


## Direnv
.envrc
```
export AWS_ACCOUNT_ID=******
export DB_HOST=127.0.0.1
export DB_NAME=go-graphql-echo-boilerplate
export DB_USER=root
export DB_PASSWORD=root
```

### MySQL起動
```
docker-compose up
```

- DB接続
```
mysql -P 3306 -h 0.0.0.0 -uroot -proot
```

## Playground
http://localhost:3000/playground

- URL(GraphQL Endpoint)
http://localhost:3000/graphql


### Users
```graphql
{
  users(orderBy: LATEST, page: { first: 2 }) {
    edges {
      node {
        id
        name
        tasks {
          id
          title
        }
      }
    }
  }
}
```

```graphql
mutation {
  createUser(input: {
    name: "kou"
  }) {
    id
    name
  }
}
```

```graphql
mutation {
  deleteUser(input: { id: "*******" }) {
    id
    name
  }
}
```

```graphql
mutation {
  updateUser(input: { id: "*******", name: "kou tasks" }) {
    id
    name
  }
}

```


### Tasks
```graphql
{
  tasks(
    input: { completed: false }
    orderBy: LATEST
    page: { first: 4, after: "dGFzazozNg==" }
  ) {
    edges {
      node {
        id
        title
        notes
        due
      }
      cursor
    }
  }
}
```

```graphql
mutation {
  createTask(input: { title: "Title", userId: "*******", notes: "Note..." }) {
    id
    title
    notes
    completed
  }
}

```

```graphql
mutation {
  updateTask(input: {
    id: "*******",
    title: "changed"
  }) {
    id
    title
    notes
  }
}
```


```graphql
mutation {
  deleteTask(input: {
    id: "*******"
  }) {
    id
    title
    notes
    due
  }
}
```


## REST API

- POST
```
curl -X POST -H "Content-Type: application/json" -d '{"title":"titleTest", "note":"noteTest"}' http://localhost:3000/api/v1/tasks
```

### 参考サイト
本番のDockerファイルなど配下を参考。
https://qiita.com/takasp/items/c6288d4836e79801bb19

### Golang & Dockerのローカルビルド動作確認
- build & run
```sh
docker build -t golang -f docker/prd/go/Dockerfile .
docker run -p 8080:8080 -d --name golang golang
```

- kill all docker process
```sh
docker ps -aq | xargs docker rm -f
```

- Access
```
curl http://localhost:8080
Hello, World!
```

### ECR Registeryへプッシュ
事前に以下を環境変数でセットする必要がある。
`export AWS_ACCOUNT_ID=****`

以下の形式でシェルを実行する。
`./docker-push-ecr.sh ${STAGE} ${ECR Repository Name} ${TAG}`

```
./docker-push-ecr.sh dev go-graphql-echo-boilerplate latest 
```

## GitHub ActionsでFargateへの自動デプロイ
1. GitHub ActionsからAWSへアクセスするためのキー設定
対象レポジトリの`Settings`を選択
-> `Secrets`から認証情報を入力する
※ECRとECS周りの権限があるIAM Userを使用する

2. GitHubの画面のReleasesからタグを切ると自動でリリースされる。

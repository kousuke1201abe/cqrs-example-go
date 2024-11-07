# cqrs-example-go

## 環境構築
VSCode dev container を利用するため、Docker Desktop を起動し、当リポジトリのルートディレクトリで VSCode を開いてください。

VSCode の左下にある青いアイコンをクリックし、オプションから「コンテナーで再度開く」を選択してください。

コンテナの起動が完了したら、`http://localhost:3000/playground` にアクセスすると、GraphQL Playground が利用できるようになっているはずです。

たとえば以下のような GraphQL query を実行して、値が返ってくれば OK です。

```
query {
  customers{
    id
  }
}
```

## アーキテクチャイメージ
<img width="970" alt="arch" src="https://github.com/user-attachments/assets/b8791bba-333a-456e-bd24-4bc37b6ec9fa">

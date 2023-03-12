# chatgpt-discord

ChatGPT(gpt-3.5-turbo) を使った Discord BOT です。

![screenshot-01.png](/doc/screenshot-01.png)

- BOT へのメンションで会話を開始、返信で続けて会話できます。
- 入力の最大文字数は 50、会話の往復上限は 5 としています。
- OpenAI API のパラメーターは特段指定していません。

各種設定は Clone または Fork して適宜変更してください。

# セットアップ

## 0. Discord BOT, OpenAI API の用意

各種アカウント、BOT の作成を済ませてください。Discord BOT は `Message content intent` が必須です。

## 1. 各種設定

### 1.1. `.env` の設定

`.env.sample` を参考に `.env` を作成してください。

```
BOT_TOKEN=<Discord BOT のトークン>
OPENAI_TOKEN=<OpenAI API のトークン>
BOT_USER_ID=<BOT のユーザーID>
BOT_ROLE_ID=<BOT のロールID>
```

> **NOTE**
> ユーザーへのメンションとロールへのメンションどちらでも発火するように設定しています。

### 1.2. （オプション） `assets/chat_system_prompt.txt` の記述

ChatGPT の振る舞いを指定できます。

```
アニメキャラクター「ドラえもん」との会話をシュミレートします。彼になりきり振る舞ってください。
```

## 2. 起動

```bash
$ make docker/run

# （オプション）ログの確認
$ make docker/logs
```

## 3. 停止

```bash
$ make docker/stop

# 完全にコンテナを削除する
$ make docker/rm
```

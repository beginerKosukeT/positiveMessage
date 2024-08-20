# Positive Message
![logo](https://github.com/beginerKosukeT/positiveMessage/assets/144611948/0faaf891-8020-4cbf-bd9b-b8cdb2f03289)

- ### プログラム概要
  ポジティブなメッセージや元気が出るアドバイスを投稿でき、人工の音声として連続再生することが可能
- ### 公開先
  <a href="https://positive-message-254febcb568f.herokuapp.com/regisration">Positive Message (公開停止中)</a>

## 使用技術
- バックエンド：Go言語(go1.22.0)
- フロントエンド：JavaScript,HTML,Bootstrap(5.3.0)
- DB：MySQL(8.0.35)
- 音声の読み上げ：Web Speech API
- セッション：gorilla/sessions
- デプロイ先：Heroku


## 機能一覧
| トップ画面 (新規登録画面) | ログイン画面 |
| - | - |
|<img width="460" alt="image" src="https://github.com/beginerKosukeT/positiveMessage/assets/144611948/d538f610-af38-4de0-be12-26b9546f442a">|<img width="460" alt="image" src="https://github.com/beginerKosukeT/positiveMessage/assets/144611948/83cdeaf9-fad3-4e3b-99fe-1a017d9e46cd">|
|上部に本サイトの概要説明が表示され、新規アカウント登録をできる<br>(JavaScriptによるエラーハンドリングに対応)|セッション機能を利用してログインすることができる|

| マイページ | 人気の投稿&新作画面 |
| - | - |
|<img width="460" alt="image" src="https://github.com/beginerKosukeT/positiveMessage/assets/144611948/0a8fd561-69e9-4264-b857-f38c026c6a92">|<img width="460" alt="image" src="https://github.com/beginerKosukeT/positiveMessage/assets/144611948/819724c9-2e07-4e28-963f-537f556af97a">|
|保存・お気に入りした投稿と、自身が作成した投稿が表示され、<br>連続で読み上げることができる(スキップや一時停止にも対応)|お気に入り数の多い投稿、作成日時が新しい投稿の内、<br>各8件まで新しい投稿が表示される|

| 検索画面 | 新規投稿画面|
| - | - |
|<img width="460" alt="image" src="https://github.com/beginerKosukeT/positiveMessage/assets/144611948/a24c067b-8a2b-4a5e-bfa6-db1be33a0829">|<img width="460" alt="image" src="https://github.com/beginerKosukeT/positiveMessage/assets/144611948/f734c669-0977-474c-addc-6fab773503c5">|
|タイトルで検索することができる|新規投稿ができる(JavaScriptによるエラーハンドリングに対応)|

| 投稿の詳細画面 || 
| - | - |
|<img width="460" alt="image" src="https://github.com/beginerKosukeT/positiveMessage/assets/144611948/43a589ad-1268-4397-8812-b94c028b2b75">|<img width="460" alt="image" src="https://github.com/beginerKosukeT/positiveMessage/assets/144611948/55b7e9e8-8305-49af-8ca0-58af2946233a">|
|投稿の詳細が表示され、保存・お気に入りをすることができる<br>本画面での読み上げも可能|　　　　　　　　　　　　　　　　　　　　　　　　　　|

## テスト
単体テストを実施(<a href="https://app.box.com/s/qdgiyqzxdfu0vaslqy4kaxyuf0m9dqez">単体テスト.xlsx</a>)

![test](https://github.com/beginerKosukeT/positiveMessage/assets/144611948/a593f24d-7561-4dd6-ab5c-3e89ca4ac00c)

## ER図
![ERD](https://github.com/beginerKosukeT/positiveMessage/assets/144611948/b9e0296f-1842-4091-a273-917ead4067d6)


## 画像ダウンロード元
- <a href="https://www.san-x.co.jp/sumikko/profile/">すみっコぐらしofficial web site プロフィール</a>
- <a href="https://sunheart-shop.com/c/gr1/san-x/sumikkogurashi"> すみっコぐらし | 丸眞オンラインショップ</a>

# team_up

## 使用方法

1．teamup アプリケーション直下に.env を作成（.env の中身は個別共有）  
2．server_side/conf 配下に app.conf を作成（app.conf の中身は個別共有）  
3．`docker-compose up -d --build`でコンテナを起動  
4．`docker exec -it teamup_server bash`でコンテナ内へ  
5．`cd server_side`で server_side ディレクトリへ移動  
6．`bee run`でアプリケーションを起動

## テストユーザー

＜テストユーザー ① ＞  
・ユーザー名：user1  
・ログイン名：admin1  
・パスワード：password1

＜テストユーザー ② ＞  
・ユーザー名：user2  
・ログイン名：admin2  
・パスワード：password2

＜テストユーザー ③ ＞  
・ユーザー名：user3  
・ログイン名：admin3  
・パスワード：password3

＜テストユーザー ④ ＞  
・ユーザー名：user4  
・ログイン名：admin4  
・パスワード：password4

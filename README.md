# team_up

## 使用方法

1．teamup アプリケーション直下に.env を作成（.env の中身は個別共有）  
2．server_side/conf 配下に app.conf を作成（app.conf の中身は個別共有）  
3．`docker-compose up -d --build`でコンテナを起動  
4．`docker exec -it teamup_server bash`でコンテナ内へ  
5．`cd server_side`で server_side ディレクトリへ移動  
6．`bee run`でアプリケーションを起動

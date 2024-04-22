server {
    listen 80;
    server_name moyihust.eu.org www.moyihust.eu.org;

    location / {
        root /root/obd/dist/;
        try_files $uri $uri/ /index.html;
    }

    location /api/ {
        proxy_pass http://localhost:1145; # 假设你的API服务器运行在3000端口
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
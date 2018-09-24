# url

Required environment variables <br /><br />
MYSQL_ADDRESS => mysql database host ip <br />
MYSQL_PORT => mysql database host port <br />
MYSQL_USER => mysql database host user <br />
MYSQL_PASSWORD => mysql database host password <br />

To start service<br />
docker-compose up -d --build<br />

Available services<br />
1. request shortend URL<br />
curl -X POST \
  http://localhost:8080/getShortUrl \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{"url":"taskworld.com/workspace/uploadfilename"}'
  
2. request original URL<br />
curl -X POST \
  http://localhost:8080/getOriginalUrl \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{"url":"taskworld.com/MQ=="}'
 

pid=$(ps aux | grep "go-build.*main" | grep -v grep | awk '{print $2}' |tail -n 1 | head -n 1);
sudo kill ${pid}
cd smartgraphitehb-server/;
sudo git pull
nohup go run ./main.go &
exit 0


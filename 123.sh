pid=$(ps aux | grep "go-build.*gm" | grep -v grep | awk '{print $2}' |tail -n 1 | head -n 1);
kill ${pid}
cd smartgraphitenew-server/;
git pull
nohup go run ./main.go gm &
exit 0

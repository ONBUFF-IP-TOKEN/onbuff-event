# set -x

sh ./prebuild.sh $1

rm -rf bin/onbuff-event

go build -o bin/onbuff-event.exe main.go

cd bin
./onbuff-event.exe -c=config.yml
#version=$(git describe)
version="0.1.0"

cd ../../assembly || echo "directory does not exit"
GOOS="linux" go build

mv ./assembly ../build/docker
cp ./config.json ../build/docker


cd ../build/docker || echo "directory does not exist"

docker build --tag assembly:$version .

docker save -o ./build.tar assembly:$version

rm assembly
rm config.json

tail -f /dev/null
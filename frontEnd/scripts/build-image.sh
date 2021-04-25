
version=$(git describe)

docker build -t morpheus:$version .

docker save -o ./morpheus.tar morpheus:$version
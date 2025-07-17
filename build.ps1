$folderToCompress = "$PSScriptRoot\src"
$outputTarFile = ".\bin\app.tar"

tar -cf $outputTarFile -C (Split-Path $folderToCompress) (Split-Path $folderToCompress -Leaf)

# building dev image and extracting the generate
$containerName="builder-go"

docker build . -t ${containerName} --target builder
$containerId = docker create -it --rm ${containerName}

docker cp ${containerId}:/app/generate ./bin/generate
docker rm ${containerId}
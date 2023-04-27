rm -rf ./build
mkdir ./build

# BUILD FOR APPLE CHIP
cp -r "./NoiseBar Template.app" "./build/NoiseBar.app/"
cp -r "./sounds" "./build/NoiseBar.app/Contents/MacOS/sounds"

go build -o "./build/NoiseBar.app/Contents/MacOS/noisebar" .

cd ./build
zip -vr "./NoiseBar (MacOS Apple Chip).zip" "./NoiseBar.app"
cd ..

# BUILD FOR INTEL CHIP
cp -r "./NoiseBar Template.app" "./build/NoiseBar-intel.app/"
cp -r "./sounds" "./build/NoiseBar-intel.app/Contents/MacOS/sounds"

CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o "./build/NoiseBar-intel.app/Contents/MacOS/noisebar" .

cd ./build
zip -vr "./NoiseBar (MacOS Intel Chip).zip" "./NoiseBar-intel.app"
cd ..
rm -f kenaito-media-group-tool.exe
rm -f kenaito-media-group-tool-release.exe
# build
go build -o kenaito-media-group-tool.exe main.go
# upx compress
./upx -o kenaito-media-group-tool-release.exe kenaito-media-group-tool.exe
rm -f kenaito-media-group-tool.exe

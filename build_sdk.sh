sed -i"" 's/GIT_USER_ID/badarsebard/g' openapi/go.mod
sed -i"" 's/GIT_REPO_ID/xsoar-sdk-go\/openapi/g' openapi/go.mod

cd openapi
go mod tidy
go fmt
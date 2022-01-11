sed -i.bak 's/GIT_USER_ID/badarsebard/g' openapi/go.mod
sed -i.bak 's/GIT_REPO_ID/xsoar-sdk-go\/openapi/g' openapi/go.mod

rm openapi/go.mod.bak

cd openapi
go mod tidy
go fmt
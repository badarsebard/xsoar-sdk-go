sed -i"" 's/GIT_USER_ID/badarsebard/g' openapi/go.mod
sed -i"" 's/GIT_REPO_ID/xsoar-sdk-go\/openapi/g' openapi/go.mod

oldversion=`cut -d '.' -f3 VERSION`
newversion=`expr $oldversion + 1`
sed -i"" "s/$oldversion/$newversion/g" VERSION

cd openapi
go mod tidy
go fmt
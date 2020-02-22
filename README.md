# get-started-with-go

### tools and lint
```bash
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/tools.git
git clone https://github.com/golang/lint.git
go get golang.org/x/lint/golint
```

### vs code workspace
```javascript
{
	"folders": [
		{
			"path": "/Users/username/folder"
		}
	],
	"settings": {
		"go.inferGopath": true,
		"go.toolsGopath": "/Users/username/folder"
	}
}
```

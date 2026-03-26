**Instructions for using this project**

1. First add the project to go.mod:
```
go mod init github.com/your_name/your_project // if you don't have it yet
go get github.com/drvladis/logging
```

2. Secondly, import the project:
```
import lg "github.com/drvladis/logging/log"
```

3. Third, use it for logging issues:
```
// example
if err != nil {
	lg.Logging(fmt.Errorf("modul: %s\n", err.Error()))
}
```
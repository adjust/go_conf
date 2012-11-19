#go_conf

this is a nifty little package to help reduce the boilerplate code of our go apps and make them behave more like rails.

##features
- rails style ```config/database.yml``` to configure different databases.
- ```GO_ENV``` similar to ```RAILS_ENV``` to have development, testing and production databases
- default ```log/server.log``` for log output
- prefab ExitHandler to make your app handle SIGHUP correctly

##how to use
get it
```
go get github.com/adeven/go_conf
```

###general
import it
```go
import github.com/adeven/go_conf
...
```

use it 

```go
redis_host, redis_db := go_conf.GetRedisConf()
...
```
###exit handler
to make your app quit on SIGHUP and execute some function before closing use a custom exit handler
```go
type MyExitHandler struct {
}

func (self *MyExitHandler) OnExit() {
	log.Println("running sig handler")
	//do something like cleanup or saving progress...
}

func main() {
	go_conf.SetExitHandler(&MyExitHandler{})
}
```	

##how to extend database config functions
easy: just fork and create a new file (good place is to start with a ```postgres.go``` copy) with a function returning the parameters you want.

## License

This Software is licensed under the MIT License.

Copyright (c) 2012 adeven GmbH, 
http://www.adeven.com

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
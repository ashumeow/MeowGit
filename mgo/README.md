<code>v2 stable</code><br>
<code>$git clone https://github.com/go-mgo/mgo.git --branch v2 --single-branch</code><br>
<br>
<b>Some corrections:-</b> <br>
Inside <code>mgo</code><br>
<b>Remove <code>"only mgo based import files"</code> in <code>import</code></b> and add this instead. <br>
[1] <code>auth.go</code> <br>
```go
import (
	"MeowGit/mgo/bson"
	"MeowGit/mgo/internal/scram"
)
```
[2] <code>cluster.go</code><br>
```go
import (
	"MeowGit/mgo/bson"
)
```
[3] <code>gridfs.go</code><br>
```go
import (
  "MeowGit/mgo/bson"
)
```
[4] <code>server.go</code><br>
```go
import (
  "MeowGit/mgo/bson"
)
```
[5] <code>session.go</code><br>
```go
import (
  "MeowGit/mgo/bson"
)
```
[6] <code>socket.go</code><br>
```go
import (
  "MeowGit/mgo/bson"
)
```

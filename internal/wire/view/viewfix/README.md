
# viewfix
    import "github.com/coralproject/shelf/internal/wire/view/viewfix"






## func Add
``` go
func Add(context interface{}, db *db.DB, views []view.View) error
```
Add inserts views for testing.


## func Get
``` go
func Get() ([]view.View, error)
```
Get loads view data based on view.json.


## func Remove
``` go
func Remove(context interface{}, db *db.DB, pattern string) error
```
Remove removes views in Mongo that match a given pattern.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
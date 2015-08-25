consulutils
-----------
Library to unmarshal consul key value pairs into structs.

    import "github.com/euforia/consulutils"

    type SomeStruct struct {
        Host    string `consul:"foo/bar/datasource/host"`
        Port    int64  `consul:"foo/bar/datasource/port"`
        Enabled bool   `consul:"foo/bar/enabled"`
    }

    pairs, _, err := kv.List("foo")

    var s SomeStruct
    err = consulutils.Unmarshal(pairs, &s)

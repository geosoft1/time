Time service
===

Return in JSON [RFC4627](https://tools.ietf.org/html/rfc4627) format [UTC](https://en.wikipedia.org/wiki/Coordinated_Universal_Time) time as [Unix time](https://en.wikipedia.org/wiki/Unix_time) and [RFC3339](https://golang.org/src/time/format.go?s=3825:3867#L62).

Sample request:

	curl -X GET -i localhost:8123

Sample response:

	{"unixtime":1596787417,"rfc3339":"2020-08-07T08:03:37Z"}

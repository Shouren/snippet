local cjson = require "cjson"

local body = "[ true, { "foo": "bar" } ]"
local data = cjson.decode(body)
print(data)

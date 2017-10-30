local test = {
    url = 'ikimi.net',
    github = 'shouren',
    1024,
    [6] = 64,
    ['Country'] = 'China',
}

print(test.url)
print(test['github'])
print(test.Country)
print(test[6])
print(test[1])

for v in pairs(test) do
    print(v)
end

local test = { true, { foo = "bar" }}
print(test[2].foo)


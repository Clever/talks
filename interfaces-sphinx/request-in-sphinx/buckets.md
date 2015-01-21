IP: 172.0.0.2, path: /special/resources/1

matches limit-special-things: true

key: limit-special-things | 172.0.0.2

increment current request count, make sure it's not >200

    Name:   limit-special-things | 172.0.0.2
    Count:  11
    Expire: in 5 secs

response:

    X-RateLimit-Bucket:     limit-special-things
    X-RateLimit-Limit:      200
    X-RateLimit-Remaining:  189
    X-RateLimit-Reset:      1421785247

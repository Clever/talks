IP:             172.0.0.2
Path:           /special/resources/1
Authorization:  ABCD=
Host:           example.com

REQUEST MATCHES LIMITS:

   limit-special-things: true 
   other-limit:          false

BUCKET NAME FOR REQUEST:

    limit-special-things | authorization ABCD= | 172.0.0.2

INCREMENT CURRENT COUNT IN BUCKET

    Name:   limit-special-things | authorization ABCD= | 172.0.0.2 
    Count:  10 + 1
    Expire: 1421785247  # (in 5 secs from now)

LIMIT REQUEST: false

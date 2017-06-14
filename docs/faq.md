# FAQ

This file keeps track of questions and discussions from Gitter

## MySQL gives `unsupported Scan, storing driver.Value type []uint8 into type *time.Time`

> did a quick test to get mysql running, but run into migrate sql issue - seems mysql related
An error occurred while running the migrations: Could not apply ladon SQL migrations: Could not migrate sql schema, applied 0 migrations: sql: Scan error on column index 0: unsupported Scan, storing driver.Value type []uint8 into type *time.Time
is this a known bug ? or any specific mysql version which is required (running 5.7) ?

```
$ hydra help host
...
   - MySQL: If DATABASE_URL is a DSN starting with mysql:// MySQL will be used as storage backend.
        Example: DATABASE_URL=mysql://user:password@tcp(host:123)/database?parseTime=true

        Be aware that the ?parseTime=true parameter is mandatory, or timestamps will not work.
...
```

## The docker image exists immediately

Check the logs using `docker logs <container-id>`.

## Insufficient Entropy

> Hey there , I am getting this error when I try request an access token "The request used a security parameter (e.g., anti-replay, anti-csrf) with insufficient entropy (minimum of 8 characters)"

> Kareem Diaa @kimooz Jun 07 16:41  
Hey there , I am getting this error when I try request an access token "The request used a security parameter (e.g., anti-replay, anti-csrf) with insufficient entropy (minimum of 8 characters)"

> Aeneas @arekkas Jun 07 16:41  
@kimooz make sure state and nonce are set in your auth code url (http://hydra/oauth2/auth?client_id=...&nonce=THIS_NEEDS_TO_BE_SET&state=THIS_ALSO_NEEDS_TO_BE_SET

## I get compile errors!

> I would try deleting the vendor dir and glide’s files and try glide init again or clear Glide’s global cache.

> follow the steps in the readme https://github.com/ory/hydra#building-from-source

## Is JWT supported?

> Mufid @mufid 03:29  
> Could Hydra's Access Token be a JWT? So that my resource server does not need to call Introspection API for each request.

> Mufid @mufid 03:39  
Yes, the access token looks like JWT, but i am unable to decode it. Here is my example token form Hydra: LpxuGoqWy7lYp9N0Cea8mEGR6IHhyr37jxZXRHqSjRM.nU-jMnAJ7dUKQPjWF4QBEL9OQWVU8zj_ElhrT-FQrWw (JWT Tokens should have 2 dots (3 segments), so this is not a valid JWT)

> Mufid @mufid 03:56  
*form --> from, typo, sorry.
> Aeneas @arekkas 11:50  
@mufid JWT is not supported at the moment, we might add it, but not as part of the hydra community edition

## Refreshing tokens

> Kareem Diaa @kimooz 15:48  
One last question  if you don't mind
from your experience do you think that saving the user access token in a session and validating it from the client on ever refresh does that make sense or not?
using the introspect endpoint

> Aeneas @arekkas 15:51  
nah, simply write your http calls in a way that if a 401 or 403 occurrs, the token is refreshed
that's the easiest
and cleanest

## Revoking tokens & log out

> Kareem Diaa @kimooz 15:41  
Thanks @arekkas. I had two other questions:  
1\. Is there a way to revoke all access tokens for a certain user("log out user")  
2\. How can I inform the consent app that this user logged out?  

> Aeneas @arekkas 15:42  
1\. no this isn't supported currently  
2\. you can't because log out and revoking access tokens are two things  
and it would require an additional api or something, which makes the consent app harder to write and integrate

> Kareem Diaa @kimooz 15:43  
So can you suggest a workaround?
I want implement single sign off

> Aeneas @arekkas 15:44  
the user has the access and refresh token right
in his browser or somewhere

> Kareem Diaa @kimooz 15:44  
yah

> Aeneas @arekkas 15:44  
ok so why not make a request to /oauth2/revoke
and pass that refresh token
(you will probably need a proxy with a client id and secret for that to be possible, but you get the point)

> Kareem Diaa @kimooz 15:46  
yah but the moment he refreshes, the client will hit on hydra and then consent where it will find that this user is already logged in
and will return a new token although he should have logged out
ohh so you mean have two requests one for hydra to revoke and one for consent to log out correct?

> Aeneas @arekkas 15:47  
yes

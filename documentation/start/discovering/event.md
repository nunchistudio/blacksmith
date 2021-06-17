---
title: Your first event
enterprise: false
---

# Your first event

Since the warehouse has the table `users`, it is ready to welcome some data. Let's
`INSERT` our first user!

The trigger `mywebhook` of the source `crm` exposes a HTTP route for incoming
webhooks. We can therefore make a HTTP request with the appropriate data.

Using `curl`:
```bash
$ curl --request POST --url http://localhost:9090/crm/mywebhook \
  --header 'Content-Type: application/json' \
  --data '{
    "context": {
      "ip": "127.0.0.1"
    },
    "data": {
      "first_name": "John",
      "last_name": "Doe",
      "username": "johndoe"
    }
  }'

```

Using [Insomnia](https://insomnia.rest/):

![Your first event](/images/blacksmith/insomnia.png)

This trigger calls the action `run-operation` of the SQL destination. It runs the
SQL template `./warehouse/operations/insert-user.sql` and pass the `data` object
of the request body as data.

The response contains some details about the event and the related jobs created.
The `meta` and `data` keys can be omited in the response in case it might be
exposing sensitive information.

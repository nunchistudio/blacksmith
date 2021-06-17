---
title: Passing data from JSON file
enterprise: false
---

# Passing data from JSON file

It's possible to pass a JSON file as a data source when running an operation.

The JSON must be have valid object at the top level. Therefore, the following
example does not work:
```json
[
  {
    "id": "823798",
    "first_name": "John",
    "last_name": "Doe",
    "username": "johndoe",
    "email": "johndoe@example.com",
  },
  {
    "id": "823799",
    "first_name": "Jane",
    "last_name": "Doe",
    "username": "janedoe",
    "email": "janedoe@example.com",
  }
]
```

But this one does:
```json
{
  "users": [
    {
      "id": "823798",
      "first_name": "John",
      "last_name": "Doe",
      "username": "johndoe",
      "email": "johndoe@example.com",
    },
    {
      "id": "823799",
      "first_name": "Jane",
      "last_name": "Doe",
      "username": "janedoe",
      "email": "janedoe@example.com",
    }
  ]
}

```

In a SQL file, you can then loop over each user using the `for` tag, as shown in
this *operation*:
```sql
{% for user in users %}
  INSERT INTO users (id, first_name, last_name, username, email)
    VALUES (
      {{ user.id }},
      '{{ user.first_name | capfirst }}',
      '{{ user.last_name | upper }}',
      '{{ user.username }}',
      '{{ user.email | lower }}'
    );
{% endfor %}

```

Using the `run operation` command, you are able to compile the SQL operation and
pass the JSON file as data source:
```bash
$ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
  --file "./operations/demo.sql" \
  --data "./operations/demo.json" \
  --dryrun

Compiling operations:

  -> Compiling ./operations/demo.sql...
     Writing SQL at ./operations/demo.compiled.sql...
     Success!

```

After making sure the output SQL is correct, you can then run the compiled
statement:
```bash
$ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
  --file "./operations/demo.compiled.sql"

Compiling & Executing operations:

  -> Compiling & Executing ./operations/demo.compiled.sql...
     Success!

```

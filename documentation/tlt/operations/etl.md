---
title: Operations as part of the ETL process
enterprise: false
---

# Operations as part of the ETL process

It is possible to automatically run an *operation* within the ETL process. It is
done by calling the `RunOperation` action of the `sqlike` destination.

Given a trigger (here is of mode `http`), we can call the action from the `Event`
returned by the `Extract` function like this:
```go
func (t Identify) Extract(tk *source.Toolkit, req *http.Request) (*source.Event, error) {

  // ...

  return &source.Event{
    Context: ctx,
    Data:    data,
    Actions: destination.Actions{
      "sqlike(mypostgres)": []destination.Action{
        sqlikedestination.RunOperation{
          // ...
        },
      },
    },
  }, nil
}

```

The action needs two information:
- `Filename` is the file path and name of the SQL operation to run.
- `Data` is the data to pass down to the file while templating it.

Assuming the SQL file is located at `./operations/demo.sql`, one can call the
action with:
```go
func (t Identify) Extract(tk *source.Toolkit, req *http.Request) (*source.Event, error) {

  // ...

  return &source.Event{
    Context: ctx,
    Data:    data,
    Actions: destination.Actions{
      "sqlike(mypostgres)": []destination.Action{
        sqlikedestination.RunOperation{
          Filename: "./operations/demo.sql",
          Data: map[string]interface{}{
            "user": map[string]interface{}{
              "id":         "1234567890",
              "first_name": "John",
              "last_name":  "doe",
              "username":   "joHn Doe",
              "email":      "JohnDoe@example.com",
            },
            "order": map[string]interface{}{
              "id":    "23747823",
              "amount": 12,
            },
          },
        },
      },
    },
  }, nil
}

```

In this example we pass some data including a `user` alongside an `order`.

For having the best data possible in the warehouse, we can Transform some values
right before Loading it. The templating SQL could look like this:
```sql
INSERT INTO users (id, first_name, last_name, username, email, created_at)
  VALUES (
    '{{ user.id }}',
    '{{ user.first_name | capfirst }}',
    '{{ user.last_name | upper }}',
    '{{ user.username | slugify }}',
    '{{ user.email | lower }}',

    {% if user.created_at %}
      '{{ user.created_at }}'
    {% else %}
      '{% now "2006-01-02 15:04:05" %}'
    {% endif %}
  );

{% if order %}
  INSERT INTO orders (id, amount, user_id, created_at)
    VALUES (
      '{{ order.id }}',
      {{ order.amount | floatformat:2 }},
      '{{ user.id }}',

      {% if order.created_at %}
        '{{ order.created_at }}'
      {% else %}
        '{% now "2006-01-02 15:04:05" %}'
      {% endif %}
    );
{% endif %}

```

As you can see, we add some conditions depending on the data passed down to the
file. We also set a default date with a specific format if none is provided.

This compiles and is executed within a transaction as:
```sql
INSERT INTO users (id, first_name, last_name, username, email, created_at)
  VALUES (
    '1234567890',
    'John',
    'DOE',
    'john-doe',
    'johndoe@example.com',
    '2021-06-02 14:04:33'
  );

INSERT INTO orders (id, amount, user_id, created_at)
  VALUES (
    '23747823',
    12.00,
    '1234567890',
    '2021-06-02 14:04:33'
  );

```

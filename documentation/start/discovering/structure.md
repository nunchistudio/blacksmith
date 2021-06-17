---
title: Suggested structure
enterprise: false
---

# Suggested structure

In addition to the required and highly recommended files we discovered earlier,
the generated application comes with a demo source and destination.

The demo source is named `crm` and lives in the `sources` directory. This directory
should be shared with other sources. This source has two triggers:
- `mywebhook` in the file `mywebhook.go` is a HTTP trigger exposing the route
  `/crm/mywebhook` on `POST`. For example, your CRM could send a webhook when a
  user is registered.
- `mycrontask` in the file `mycrontask.go` is a CRON trigger executing every
  20 seconds. It shows how to create a CRON trigger and how to access the logger
  configured for the application. For example, you could check a list of leads
  from your CRM everyday.

Also, a SQL database named `warehouse` is configured using the `sqlike` module.
It has *migrations*, *operations*, and *queries* with their respective directory.
It lives in the `warehouse` directory, assuming you will have only one destination
for data warehousing.

This application's structure is nothing more than a suggestion and is a good
starting point for most users. It is then up to you to organize it for your needs.

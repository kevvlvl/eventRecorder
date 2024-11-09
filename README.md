# Event Recorder

Record events from various subjects to ingest into a central location for quick querying.

The idea of this recorder is to ingest small bits of data in a natural way to quickly find them again

## Components

### Event Recorder API

Based on echo

#### Endpoints

- PersonalFinance
- TODO
- Hobbies
  - Cars
  - Food
  - etc. (new hobby managmeent)

### Kafka

Event streaming platform

### DB

TBDg

## Endpoints

First, we verify the api health, then we submit a new receipt to the event recorder

```shell
❯ curl localhost:8011/health
Up!
~
❯ curl localhost:8011/finance/receipt -d '{"name":"Amazon", "amount":"$40.00", "payment_method":"Scotia-Amex"}' -H 'Accept: application/json' -H 'Content-Type: application/json'
{"name":"Amazon","amount":"$40.00","payment_method":"Scotia-Amex"}

~

Then, the event gets pushed to a kafka topic where we can redistribute this info to other data sources
Then, we can call other endpoints to summarize the events as well as internally call AI endpoints to assess and provide insights such as summarize, hints for next steps, etc.

```
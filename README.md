# message-service
A service to send and receive messages.

This is the assignment:

Build a service for sending and retrieving messages. The service should support
the following functions.
1. Submit a message to a defined recipient, identified with some identifier,
e.g. email-address, phone-number, user name or similar.
2. Fetch new messages (meaning the service must know what has already
been fetched).
3. Delete a single message.
4. Delete multiple messages.
5. Fetch messages (including previously fetched) ordered by time, according
to start and stop index.

The service must be implemented in the form of a REST-API

## Assumptions and limitations
The service only needs to support plain text messages.
The implementation does not need to handle authorization or authentication.
The implementation does not need to be "ready for production", however it
should reflect your regular level of work.
It must be possible to use the service with curl or a similar tool, and any client
you create will not be the main focus of the evaluation.
It is ok to make assumptions as long as they are clearly communicated and
motivated.

# Instructions to use the code
## Setup of environment
The code is assumed to run on a Unix-like environment.
Install Python (recent version)
E.g. on Ubuntu:
sudo apt-get install python3 python3-dev

Install pip
### Setup virtual environment:
This is to avoid polluting the global environment
sudo apt install python3.12-venv
python3 -m venv ~/.virtualenvs/djangodev
source ~/.virtualenvs/djansource ~/.virtualenvs/djangodev/bin/activate

### Install Django (in virtual environment):
python -m pip install Django

### Install Django REST framework and other modules needed by the application
pip install djangorestframework
pip install markdown       # Markdown support for the browsable API.
pip install django-filter  # Filtering support

## Build
First you need to create a database:
python3 manage.py makemigrations message_app
python3 manage.py migrate message_app


## Other artifacts
OpenAPI schema can be built using the description here:
https://www.django-rest-framework.org/api-guide/schemas/

## Run
This is how you start the service:

python3 manage.py runserver

You should now be able to access these endpoints in your browser or using curl or equivalent:
http://127.0.0.1:8000/messages/
http://127.0.0.1:8000/messages/<id>/
http://127.0.0.1:8000/messages/bulk-delete
http://127.0.0.1:8000/messages/fetch-new

These query argument can be used to filter the response to the message list (messages/):
recipient : string : Match against the recipient
from_id : integer : The lowest message id you want returned
to_id : integer : The highest message id you want returned
is_fetched : True | False : Get the messages that are already fetched (or not). Default is to fetch all.

This query argument affects if the messages are updated:
skip_is_fetched_update : True | False : True indicates the is_fetched field should not be updated. Default is false.

Example: http://127.0.0.1:8000/messages/?from_id=3&to_id=7&recipient=mikael&no_is_fetched_update=True

The individual message resource (messages/<id>/) allows the normal operations, GET, PUT, DELETE.

There is a bulk delete endpoint (messages/bulk-delete) where POST with a body that consists of a list of message id will delete those messages.
If you go to http://127.0.0.1:8000/messages/bulk-delete you will get an error message and forms to provide the correct POST.

Example where id 5 and 7 are removed from command line in bulk:
http POST http://127.0.0.1:8000/messages/bulk-delete []:=5 []:=7

The bulk delete endpoint does not follow the REST principles, but it could make sense from a performance perspective
to make one request.

There is another endpoint that fetches all messages marked as not yet fetched and also updates their status to fetched.
Rather than having a GET that violates the REST principles of being idempotent, this special functionality is kept on
a specific endpoint. Example:

http POST http://127.0.0.1:8000/messages/fetch-new

### To run tests:
cd message_service
./manage.py test

# Known shortcomings
The query arguments for is_fetched must be False or True.
If a query argument is of the wrong type, e.g. from_id is not an integer, you get a 5xx response rather than a 4xx.

# Implementation decisions
The implementation should preferably be done in Python. I've done very little coding in python, so it's a good challenge.
After reviewing the options of web frameworks, I decided on Django Rest Framework. Django is the most popular Python web framework and Django Rest Framework builds on that to provide a REST API.

For the database I decided to go with SQLite as it's built-in to Python and the least work to setup. If this would be aimed for production, I would change this to postgres.

For a webserver I went with the built in option. For production, I would use Apache and mod_wsgi.



The implementation should consider, but not necessarily include.
-Redundancy
-Scalability

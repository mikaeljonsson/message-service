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

## Build other artifacts
OpenAPI schema can be built using the description here:
https://www.django-rest-framework.org/api-guide/schemas/

## Run
This is how you start the service:

python3 manage.py runserver

You should now be able to access these endpoints, either in your browser or using curl or equivalent:
http://127.0.0.1:8000/messages/
http://127.0.0.1:8000/messages/<id>/
http://127.0.0.1:8000/messages/bulk-delete
http://127.0.0.1:8000/messages/fetch-new

### Create message
Example:
http POST http://127.0.0.1:8000/messages/ recipient=mikael message_body="Good times"
Response body:
{
    "create_time": "2024-09-30T13:01:40.865067Z",
    "id": 21,
    "is_fetched": false,
    "message_body": "Good times",
    "recipient": "mikael",
    "url": "http://127.0.0.1:8000/messages/21/"
}

### Get messages
These query argument can be used to filter the list of messages to be returned:

recipient : string : Strict match against the recipient
from_id : integer : The lowest message id you want returned
to_id : integer : The highest message id you want returned
is_fetched : True | False : Get only the messages that are already fetched (or not).

Example: http://127.0.0.1:8000/messages/?from_id=3&to_id=7&recipient=mikael&is_fetched=True

The returned response is json object and it's paginated (currently set at 10 to easily see the behavior).
Example response body:
{
    "count": 14,
    "next": "http://127.0.0.1:8000/messages/?page=2",
    "previous": null,
    "results": [
        {
            "create_time": "2024-09-28T22:56:55.844509Z",
            "id": 6,
            "is_fetched": true,
            "message_body": "msg 2",
            "recipient": "mikael",
            "url": "http://127.0.0.1:8000/messages/6/"
        },
        {
            "create_time": "2024-09-28T22:57:09.928895Z",
            "id": 8,
            "is_fetched": true,
            "message_body": "msg 4",
            "recipient": "kristina",
            "url": "http://127.0.0.1:8000/messages/8/"
        },
...
    ]
}

### Handle individual message
The individual message resource (messages/<id>/) allows the normal operations, GET, PUT, DELETE.

Examples GET:
http GET http://127.0.0.1:8000/messages/15/

Example response body:
{
    "create_time": "2024-09-30T09:55:34.074529Z",
    "id": 15,
    "is_fetched": true,
    "message_body": "Hello world",
    "recipient": "mikael",
    "url": "http://127.0.0.1:8000/messages/15/"
}

Example update:
http PUT http://127.0.0.1:8000/messages/15/ recipient=mikael message_body="Hello world"

Example delete:
http DELETE http://127.0.0.1:8000/messages/15/

### Bulk deletion
There is a bulk delete endpoint (messages/bulk-delete) where POST with a body that consists of a
list of message id will delete those messages.

If you go to http://127.0.0.1:8000/messages/bulk-delete in your browser
you will get an error message and forms to provide the correct POST.

Example where id 5 and 7 are removed from command line in bulk (using httpie):
http POST http://127.0.0.1:8000/messages/bulk-delete []:=5 []:=7

The bulk delete endpoint does not follow the REST principles, but it makes sense from a performance perspective
to make one request.

### Fetch new messages and mark them as fetched
There special endpoint that fetches all messages where is_fetched is False
and also updates is_fetched to True.
Rather than having a GET request that violates the REST principles of being idempotent, this functionality is kept on
a specific endpoint. Example:

http POST http://127.0.0.1:8000/messages/fetch-new

Example response body (notice that this response is not paginated):
[
    {
        "create_time": "2024-09-30T12:05:08.724889Z",
        "id": 16,
        "is_fetched": false,
        "message_body": "hepp",
        "recipient": "mikael",
        "url": "http://127.0.0.1:8000/messages/16/"
    },
    {
        "create_time": "2024-09-30T12:05:16.474541Z",
        "id": 17,
        "is_fetched": false,
        "message_body": "hepp",
        "recipient": "mikael",
        "url": "http://127.0.0.1:8000/messages/17/"
    },
...
]

### To run tests:
cd message_service
./manage.py test

# Known shortcomings
If a query argument is of the wrong type, e.g. from_id is not an integer, you get a 5xx response rather than a 4xx.
The filtering fields do not show up on the webpage and needs to be added to the URL manually.

# Comments on the implementation

## Personal reflections
The implementation was requested to be done in Python. As this was my first Python project beyond Hello World,
it's a good challenge and plenty to learn. Hence the code is likely not ideomatic when it comes to style etc.

## REST framewok
After reviewing the options of web frameworks, I decided on Django Rest Framework. Django is the most common Python web framework and Django Rest Framework builds on that to provide a REST API.

## Database
For the database I decided to go with SQLite as it's built-in to Python and the least work to setup. If this would be aimed for production, I would change this to postgres since that is scalable and a proven database option.

## Webserver
For the webserver I went with the built in Python option. For production, Apache and mod_wsgi is recommended for better performance. For redundancy and scalability, it should be no problem to run multiple web server instances that all connect to the same DB service.

## Next steps
* Upgrade the database to Postgres.
* Add authentication and authorization. Django provides support for this, but customization is needed.
* Make the service easily deployable to production.
** Create a docker container that can easily be deployed.
** Create a cloud setup using e.g. Terraform
* Extend the data model to become more useful.
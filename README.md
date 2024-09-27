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

# Instructions
## Setup of environment
The code is assumed to run on a Unix-like environment
Install Python (recent version)
E.g. on Ubuntu:
sudo apt-get install python3 python3-dev
Install pip
Setup virtual environment:

sudo apt install python3.12-venv
python3 -m venv ~/.virtualenvs/djangodev
source ~/.virtualenvs/djansource ~/.virtualenvs/djangodev/bin/activate

Install Django (in virtual environment):
python -m pip install Django

Install Django REST framework
pip install djangorestframework
pip install markdown       # Markdown support for the browsable API.
pip install django-filter  # Filtering support

## Build
How to build. Consider efficient builds

## Run

# Implementation decisions
The implementation should preferably be done in Python. I've done very little coding in python, so it's a good challenge.
After reviewing the options of web frameworks, I decided on Django Rest Framework. Django is the most popular framework and Django Rest Framework builds on that.

For the database I decided to go with SQLite as it's built-in to Python and the least work to setup. If this would be aimed for production, I would
most likely have used postgres.

For a webserver I went with the built in option. For production, I would use Apache and mod_wsgi.



The implementation should consider, but not necessarily include.
-Redundancy
-Scalability
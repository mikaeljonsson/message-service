from django.db import models

# The Message model is used both for the database and the API
class Message(models.Model):
    # id field is added by default to be a bigint auto-incrementing primary key
    create_time = models.DateTimeField(auto_now_add=True)
    recipient = models.CharField(max_length=200)
    # message_body is a string field that can be up to 5000 characters long.
    # The max_length parameter would perhaps be chosen differently in a real world scenario
    # depending on what we want to achieve.
    message_body = models.CharField(max_length=5000, blank=True, default='')
    is_fetched = models.BooleanField(default=False)

    class Meta:
        # The requirement is to order in chrononogial order, but for efficiency
        # in a real time scenario, it would likely be more efficient to order them by
        # ['recipient', 'id'], as the search would then normally always be done per recipient
        # and the entire DB does not need to be searched through.
        ordering = ['id']

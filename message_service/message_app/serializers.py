from rest_framework import serializers
from message_app.models import Message

# Used to serialize the Message model
class MessageSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Message
        fields = ['url', 'id', 'create_time', 'recipient', 'message_body', 'is_fetched']


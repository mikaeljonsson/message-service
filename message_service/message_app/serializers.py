from rest_framework import serializers
from message_app.models import Message


class MessageSerializer(serializers.ModelSerializer):
    class Meta:
        model = Message
        fields = ['id', 'create_time', 'recipient', 'message_body', 'is_fetched']

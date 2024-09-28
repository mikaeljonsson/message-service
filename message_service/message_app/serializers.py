from rest_framework import serializers
from message_app.models import Message


class MessageSerializer(serializers.HyperlinkedModelSerializer):
    # owner = serializers.ReadOnlyField(source='owner.username')

    class Meta:
        model = Message
        fields = ['id', 'create_time', 'recipient', 'message_body', 'is_fetched']

"""
class UserSerializer(serializers.HyperlinkedModelSerializer):
    messages = serializers.HyperlinkedRelatedField(many=True, view_name='message-detail', read_only=True)

    class Meta:
        model = User
        fields = ['url', 'id', 'username', 'messages']
"""

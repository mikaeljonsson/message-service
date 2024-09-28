from message_app.models import Message
from message_app.serializers import MessageSerializer
from rest_framework import generics
from rest_framework.decorators import api_view
from rest_framework.response import Response
from rest_framework.reverse import reverse


@api_view(['GET'])
def api_root(request, format=None):
    return Response({
        # 'users': reverse('user-list', request=request, format=format),
        'messages': reverse('message-list', request=request, format=format)
    })

@api_view(['POST'])
def api_delete_messages(request, format=None):
    # Delete all messages found in the body of the request that contains a list of ids
    Message.objects.filter(id__in=request.data).delete()
    return Response({'status': 'success'})

class MessageList(generics.ListCreateAPIView):
    serializer_class = MessageSerializer

    def get_queryset(self):
        """
        Optionally filters the returned messages based on different query parameters.
        The filters include matching on `recipient`, 'is_fetched', or for the 'id' field
        being greater or equel to 'from_id' and/or smaller or equal to 'to_id'.
        """
        queryset = Message.objects.all()
        recipient = self.request.query_params.get('recipient')
        from_id = self.request.query_params.get('from_id')
        to_id = self.request.query_params.get('to_id')
        is_fetched = self.request.query_params.get('is_fetched')
        skip_update = self.request.query_params.get('skip_is_fetched_update')
        print("skip_update", skip_update)

        if recipient is not None:
            queryset = queryset.filter(recipient__exact=recipient)
        if from_id is not None:
            queryset = queryset.filter(id__gte=from_id)
        if to_id is not None:
            queryset = queryset.filter(id__lte=to_id)
        if is_fetched is not None:
            # convert false to False and true to True
            queryset = queryset.filter(is_fetched__exact=is_fetched)

        # This vioalates the RESTful principle of not changing the state of the server
        # in a GET request, but the violation is intended as we want to store
        # what messages have been fetched.
        if skip_update != 'True':
            print('Will update')
            queryset.update(is_fetched=True)
        return queryset


class MessageDetail(generics.RetrieveUpdateDestroyAPIView):
    queryset = Message.objects.all()
    serializer_class = MessageSerializer

from django.test import TestCase
from rest_framework.test import APIRequestFactory
from message_app.views import MessageDetail
from message_app.views import MessageList

from django.urls import reverse
from rest_framework import status
from rest_framework.test import APITestCase
from message_app.models import Message

class TestSingleMessage(APITestCase):

    def create_message(self, recipient, message_body):
        url = reverse('message-list')
        data = {'recipient': recipient, 'message_body': message_body}
        response = self.client.post(url, data, format='json')
        new_id = response.data.get('id')
        return new_id

    def test_post_message(self):
        # create a message
        url = reverse('message-list')
        data = {'recipient': 'john.doe', 'message_body': 'a message'}
        response = self.client.post(url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(Message.objects.count(), 1)
        new_id = response.data.get('id')
        # GET the created message
        url = reverse('message-detail', args=[new_id])
        get_response = self.client.get(url, format='json')
        self.assertEqual(get_response.status_code, status.HTTP_200_OK)
        self.assertEqual(get_response.data['recipient'], 'john.doe')
        self.assertEqual(get_response.data['message_body'], 'a message')

    def test_delete_message(self):
         # create a message
        new_id = self.create_message('john.doe', 'a message')
        # GET the created message
        url = reverse('message-detail', args=[new_id])
        get_response = self.client.get(url, format='json')
        self.assertEqual(get_response.status_code, status.HTTP_200_OK)
        self.assertEqual(get_response.data['recipient'], 'john.doe')
        self.assertEqual(get_response.data['message_body'], 'a message')
        # delete the message
        response = self.client.delete(url, format='json')
        self.assertEqual(response.status_code, status.HTTP_204_NO_CONTENT)
        # Another delete should return 404
        response = self.client.delete(url, format='json')
        self.assertEqual(response.status_code, status.HTTP_404_NOT_FOUND)
        # Another get should return 404
        response = self.client.get(url, format='json')
        self.assertEqual(response.status_code, status.HTTP_404_NOT_FOUND)

    def test_update_message(self):
        # create a message
        url = reverse('message-list')
        data = {'recipient': 'john.doe', 'message_body': 'a message'}
        response = self.client.post(url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(Message.objects.count(), 1)
        new_id = response.data.get('id')
        # GET the created message
        url = reverse('message-detail', args=[new_id])
        get_response = self.client.get(url, format='json')
        self.assertEqual(get_response.status_code, status.HTTP_200_OK)
        self.assertEqual(get_response.data['recipient'], 'john.doe')
        self.assertEqual(get_response.data['message_body'], 'a message')
        # update the message
        data = {'recipient': 'john.doe', 'message_body': 'an updated message'}
        response = self.client.put(url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        # GET the updated message
        get_response = self.client.get(url, format='json')
        self.assertEqual(get_response.status_code, status.HTTP_200_OK)
        self.assertEqual(get_response.data['recipient'], 'john.doe') # still the same
        self.assertEqual(get_response.data['message_body'], 'an updated message')

class TestMultipleMessages(APITestCase):
    
    def create_messages(self, create_number):
        # support function to create the wanted number of messages
        for number in range(create_number):
            Message.objects.create(recipient='john.doe'+str(number), message_body='message {}'.format(number))

    def test_get_messages(self):        
        self.create_messages(10) # create 10 messages
        # get messages with different filters
        url = reverse('message-list')
        response = self.client.get(url, query_params={'is_fetched': 'False'}, format='json')
        self.assertEqual(len(response.data['results']), 10) # at creation, all messages have is_fetched=False
        response = self.client.get(url, query_params={'recipient': 'john.doe5'}, format='json')
        self.assertEqual(len(response.data['results']), 1) # only one message with recipient john.doe5
        response = self.client.get(url, query_params={'is_fetched': 'False'}, format='json')
        self.assertEqual(len(response.data['results']), 10) # nothing changed, no update to is_fetched
        response = self.client.get(url, query_params={'is_fetched': 'True'}, format='json')
        self.assertEqual(len(response.data['results']), 0) # nothing fetched yet
        response = self.client.get(url, query_params={'is_fetched': 'False', 'from_id' : 8}, format='json')
        self.assertEqual(len(response.data['results']), 3) # 8,9,10
        response = self.client.get(url, query_params={'is_fetched': 'False', 'to_id' : 4}, format='json')
        self.assertEqual(len(response.data['results']), 4) # 1,2,3,4
        response = self.client.get(url, query_params={'is_fetched': 'False', 'from_id' : 5, 'to_id' : 4}, format='json')
        self.assertEqual(len(response.data['results']), 0) # no messages in this range

    def test_fetch_new_messages(self):
        self.create_messages(5) # create 5 messages
        # fetch new messages
        fetch_new_url = reverse('fetch-new-messages')
        response = self.client.post(fetch_new_url, query_params={}, format='json')
        self.assertEqual(len(response.data), 5)

        self.create_messages(3) # create 3 more messages
        # check that the new messages are not fetched
        url = reverse('message-list')
        response = self.client.get(url, query_params={'is_fetched': 'True'}, format='json')
        self.assertEqual(len(response.data['results']), 5)
        # fetch new messages
        response = self.client.post(fetch_new_url, query_params={}, format='json')
        self.assertEqual(len(response.data), 3)
        # check that the new messages are fetched
        response = self.client.get(url, query_params={'is_fetched': 'True'}, format='json')
        self.assertEqual(len(response.data['results']), 8)        

    def test_get_no_message(self):
        self.create_messages(10) # create 10 messages
        url = reverse('message-list')
        response = self.client.get(url, query_params={'from_id' : 4, 'to_id' : 5}, format='json')
        self.assertEqual(len(response.data['results']), 2)
        response = self.client.get(url, query_params={'from_id' : 5, 'to_id' : 4}, format='json')
        self.assertEqual(len(response.data['results']), 0)

    def test_delete_bulk_messages(self):
        self.create_messages(10) # create 10 messages
        # delete 3 messages
        url = reverse('bulk-delete')
        data = [1,3,4]
        response = self.client.post(url, data=data, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        # check that 3 messages have been deleted
        get_url = reverse('message-list')
        response = self.client.get(get_url, query_params={'from_id' : 1, 'to_id' : 4}, format='json')
        self.assertEqual(len(response.data['results']), 1) # id=2 remains in this range
        # delete 7 remaining messages
        data = [2,5,6,7,8,9,10]
        response = self.client.post(url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        # check that all are deleted
        response = self.client.get(get_url, query_params={'from_id' : 1, 'to_id' : 10}, format='json')
        self.assertEqual(len(response.data['results']), 0)
        self.assertEqual(response.status_code, status.HTTP_200_OK)





from django.test import TestCase
from rest_framework.test import APIRequestFactory
from message_app.views import MessageDetail
from message_app.views import MessageList

from django.urls import reverse
from rest_framework import status
from rest_framework.test import APITestCase
from message_app.models import Message

class TestSingleMessage(APITestCase):
    def test_post_message(self):
        # create a message
        url = reverse('message-list')
        data = {'recipient': 'john.doe', 'message_body': 'a messsage'}
        response = self.client.post(url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(Message.objects.count(), 1)
        new_id = response.data.get('id')
        # GET the created message
        url = reverse('message-detail', args=[new_id])
        get_response = self.client.get(url, format='json')
        self.assertEqual(get_response.status_code, status.HTTP_200_OK)
        self.assertEqual(get_response.data['recipient'], 'john.doe')
        self.assertEqual(get_response.data['message_body'], 'a messsage')

    def test_delete_message(self):
         # create a message
        url = reverse('message-list')
        data = {'recipient': 'john.doe', 'message_body': 'a messsage'}
        response = self.client.post(url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(Message.objects.count(), 1)
        new_id = response.data.get('id')
        # GET the created message
        url = reverse('message-detail', args=[new_id])
        get_response = self.client.get(url, format='json')
        self.assertEqual(get_response.status_code, status.HTTP_200_OK)
        self.assertEqual(get_response.data['recipient'], 'john.doe')
        self.assertEqual(get_response.data['message_body'], 'a messsage')
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
        data = {'recipient': 'john.doe', 'message_body': 'a messsage'}
        response = self.client.post(url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(Message.objects.count(), 1)
        new_id = response.data.get('id')
        # GET the created message
        url = reverse('message-detail', args=[new_id])
        get_response = self.client.get(url, format='json')
        self.assertEqual(get_response.status_code, status.HTTP_200_OK)
        self.assertEqual(get_response.data['recipient'], 'john.doe')
        self.assertEqual(get_response.data['message_body'], 'a messsage')
        # update the message
        data = {'recipient': 'john.doe', 'message_body': 'an updated message'}
        response = self.client.put(url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        # GET the updated message
        get_response = self.client.get(url, format='json')
        self.assertEqual(get_response.status_code, status.HTTP_200_OK)
        self.assertEqual(get_response.data['recipient'], 'john.doe')
        self.assertEqual(get_response.data['message_body'], 'an updated message')

class TestMultipleMessages(TestCase):

    def test_get_messages(self):
        # only is_fetched = False
        # return them ordered by id
        # filter on recipient, start_id, stop_id
        # return empty list and multiple messages
        # check that is_fetched is updated (depending on the query parameter)
        pass

    def test_get_no_message(self):
        pass

    def test_delete_messages(self):
        pass



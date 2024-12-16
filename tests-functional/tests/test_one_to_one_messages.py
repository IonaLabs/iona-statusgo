from time import sleep
from uuid import uuid4
import pytest
from test_cases import OneToOneMessageTestCase
from resources.constants import DEFAULT_DISPLAY_NAME
from clients.signals import SignalType
from resources.enums import MessageContentType


@pytest.mark.rpc
class TestOneToOneMessages(OneToOneMessageTestCase):

    @pytest.fixture(scope="class", autouse=True)
    def setup_nodes(self, request):
        await_signals = [
            SignalType.MESSAGES_NEW.value,
            SignalType.MESSAGE_DELIVERED.value,
        ]
        request.cls.sender = self.sender = self.initialize_backend(await_signals=await_signals)
        request.cls.receiver = self.receiver = self.initialize_backend(await_signals=await_signals)

    @pytest.mark.dependency(name="test_one_to_one_message_baseline")
    def test_one_to_one_message_baseline(self, message_count=1):
        pk_receiver = self.receiver.get_pubkey(DEFAULT_DISPLAY_NAME)

        self.sender.send_contact_request(pk_receiver, "contact_request")

        sent_messages = []
        for i in range(message_count):
            message_text = f"test_message_{i+1}_{uuid4()}"
            response = self.sender.send_message(pk_receiver, message_text)
            expected_message = self.get_message_by_content_type(response, content_type=MessageContentType.TEXT_PLAIN.value)
            sent_messages.append(expected_message)
            sleep(0.01)

        for i, expected_message in enumerate(sent_messages):
            messages_new_event = self.receiver.find_signal_containing_pattern(
                SignalType.MESSAGES_NEW.value,
                event_pattern=expected_message.get("id"),
                timeout=60,
            )
            self.validate_signal_event_against_response(
                signal_event=messages_new_event,
                fields_to_validate={"text": "text"},
                expected_message=expected_message,
            )

    @pytest.mark.dependency(depends=["test_one_to_one_message_baseline"])
    def test_multiple_one_to_one_messages(self):
        self.test_one_to_one_message_baseline(message_count=50)

    @pytest.mark.dependency(depends=["test_one_to_one_message_baseline"])
    @pytest.mark.skip(reason="Skipping until add_latency is implemented")
    def test_one_to_one_message_with_latency(self):
        with self.add_latency():
            self.test_one_to_one_message_baseline()

    @pytest.mark.dependency(depends=["test_one_to_one_message_baseline"])
    @pytest.mark.skip(reason="Skipping until add_packet_loss is implemented")
    def test_one_to_one_message_with_packet_loss(self):
        with self.add_packet_loss():
            self.test_one_to_one_message_baseline()

    @pytest.mark.dependency(depends=["test_one_to_one_message_baseline"])
    @pytest.mark.skip(reason="Skipping until add_low_bandwith is implemented")
    def test_one_to_one_message_with_low_bandwidth(self):
        with self.add_low_bandwith():
            self.test_one_to_one_message_baseline()

    @pytest.mark.dependency(depends=["test_one_to_one_message_baseline"])
    @pytest.mark.skip(reason="Skipping until node_pause is implemented")
    def test_one_to_one_message_with_node_pause_30_seconds(self):
        pk_receiver = self.receiver.get_pubkey("Receiver")
        self.sender.send_contact_request(pk_receiver, "contact_request")
        with self.node_pause(self.receiver):
            message_text = f"test_message_{uuid4()}"
            self.sender.send_message(pk_receiver, message_text)
            sleep(30)
        self.receiver.find_signal_containing_pattern(SignalType.MESSAGES_NEW.value, event_pattern=message_text)
        self.sender.wait_for_signal("messages.delivered")

// This can be used to mock API responses or directly in your component for testing

// Sample chat data
const chats = ref([
  {
    id: 1,
    name: 'Alice Williams',
    lastMessage: 'When can we schedule the next photoshoot?',
    lastMessageTime: '2025-02-27T14:35:00Z',
    unreadCount: 2,
  },
  {
    id: 2,
    name: 'Robert Johnson',
    lastMessage: 'The contract has been signed and sent.',
    lastMessageTime: '2025-02-28T09:12:00Z',
    unreadCount: 0,
  },
  {
    id: 3,
    name: 'Elena Martinez',
    lastMessage: 'Thanks for your quick response!',
    lastMessageTime: '2025-02-26T18:45:00Z',
    unreadCount: 0,
  },
  {
    id: 4,
    name: 'Michael Chen',
    lastMessage: "The photos from yesterday's session look great.",
    lastMessageTime: '2025-02-25T20:22:00Z',
    unreadCount: 3,
  },
  {
    id: 5,
    name: 'Sarah Thompson',
    lastMessage: 'Can we discuss the portfolio updates tomorrow?',
    lastMessageTime: '2025-02-24T11:05:00Z',
    unreadCount: 0,
  },
])

// Sample messages for Chat ID 1
const chatOneMessages = ref([
  {
    id: 101,
    chat_id: 1,
    sender_id: 'user123', // This should match currentUserId for outgoing messages
    content: 'Hi Alice, how are you doing today?',
    timestamp: '2025-02-27T14:20:00Z',
  },
  {
    id: 102,
    chat_id: 1,
    sender_id: 'alice456', // Different ID indicates incoming message
    content:
      "I'm doing well, thanks for asking! Just finishing up some edits from our last shoot.",
    timestamp: '2025-02-27T14:25:00Z',
  },
  {
    id: 103,
    chat_id: 1,
    sender_id: 'user123',
    content:
      "That's great to hear. Are you available for another session next week?",
    timestamp: '2025-02-27T14:30:00Z',
  },
  {
    id: 104,
    chat_id: 1,
    sender_id: 'alice456',
    content:
      'Yes, I should be free on Tuesday and Thursday. When can we schedule the next photoshoot?',
    timestamp: '2025-02-27T14:35:00Z',
  },
])

// Sample messages for Chat ID 2
const chatTwoMessages = ref([
  {
    id: 201,
    chat_id: 2,
    sender_id: 'robert789',
    content: "Hello, I've reviewed the contract you sent.",
    timestamp: '2025-02-28T08:50:00Z',
  },
  {
    id: 202,
    chat_id: 2,
    sender_id: 'user123',
    content: 'Great! Did you have any questions or concerns?',
    timestamp: '2025-02-28T08:55:00Z',
  },
  {
    id: 203,
    chat_id: 2,
    sender_id: 'robert789',
    content:
      "Everything looks good to me. I've signed it and will send it back shortly.",
    timestamp: '2025-02-28T09:05:00Z',
  },
  {
    id: 204,
    chat_id: 2,
    sender_id: 'robert789',
    content: 'The contract has been signed and sent.',
    timestamp: '2025-02-28T09:12:00Z',
  },
])

// Sample messages for Chat ID 3
const chatThreeMessages = ref([
  {
    id: 301,
    chat_id: 3,
    sender_id: 'user123',
    content: 'Hi Elena, I just sent you the revised portfolio images.',
    timestamp: '2025-02-26T18:30:00Z',
  },
  {
    id: 302,
    chat_id: 3,
    sender_id: 'elena321',
    content:
      'Got them, they look amazing! I especially love the black and white series.',
    timestamp: '2025-02-26T18:40:00Z',
  },
  {
    id: 303,
    chat_id: 3,
    sender_id: 'user123',
    content:
      'Thank you! I spent extra time on those edits. Let me know if you need any adjustments.',
    timestamp: '2025-02-26T18:42:00Z',
  },
  {
    id: 304,
    chat_id: 3,
    sender_id: 'elena321',
    content: 'Thanks for your quick response!',
    timestamp: '2025-02-26T18:45:00Z',
  },
])

export const useMockChats = () => {
  return {
    chats,
    chatOneMessages,
    chatTwoMessages,
    chatThreeMessages,
  }
}

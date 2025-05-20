<template>
  <div style="margin: 0 auto">
    <div v-if="status === 'unauthenticated'" class="chats-page">
      <h2 style="margin: 0 auto">
        Для использования чата необходимо авторизоваться
      </h2>
    </div>
    <div v-else class="chats-page">
      <div class="chats-sidebar">
        <h2>Your Chats</h2>
        <div class="chat-list">
          <div
            v-for="chat in chats"
            :key="chat.id"
            class="chat-item"
            :class="{ active: selectedChat?.id === chat.id }"
            @click="selectChat(chat)"
          >
            <img
              :src="getAvatarUrl(chat)"
              class="chat-avatar"
              :alt="getInterlocutorName(chat)"
              @error="handleImageError"
            />
            <div class="chat-info">
              <div class="chat-name">{{ getInterlocutorName(chat) }}</div>
              <div class="chat-preview">
                {{ chat.lastMessage || 'No messages yet' }}
              </div>
            </div>
            <div class="chat-meta">
              <div class="chat-time">
                {{ formatTime(chat.lastMessageTime) }}
              </div>
            </div>
          </div>
          <div v-if="chats.length === 0" class="no-chats">
            No chats available
          </div>
        </div>
      </div>

      <div class="chat-container">
        <template v-if="selectedChat">
          <div class="chat-header">
            <div class="chat-header-left">
              <div class="chat-header-info">
                <img
                  :src="getAvatarUrl(selectedChat)"
                  class="chat-avatar"
                  :alt="getInterlocutorName(selectedChat)"
                  @error="handleImageError"
                />
                <h3>{{ getInterlocutorName(selectedChat) }}</h3>
              </div>
              <div
                class="connection-status"
                :class="{ connected: wsConnected }"
              >
                {{ wsConnected ? 'Connected' : 'Connecting...' }}
              </div>
            </div>
            <button @click="clearSelectedChat" class="back-btn">Back</button>
          </div>

          <div class="messages-container" ref="messagesContainer">
            <div v-if="loading" class="loading-messages">
              Loading messages...
            </div>
            <template v-else>
              <div
                v-for="message in messages"
                :key="message.id"
                class="message-wrapper"
                :class="{ outgoing: message.user_id === currentUserId }"
              >
                <img
                  v-if="message.user_id !== currentUserId"
                  :src="getAvatarUrl(selectedChat)"
                  class="message-avatar"
                  :alt="getInterlocutorName(selectedChat)"
                  @error="handleImageError"
                />
                <div
                  :class="[
                    'message',
                    message.user_id === currentUserId ? 'outgoing' : 'incoming',
                  ]"
                >
                  <div class="message-content">{{ message.content }}</div>
                  <div class="message-time">
                    {{ formatTime(message.timestamp) }}
                  </div>
                </div>
              </div>
              <div v-if="messages.length === 0" class="no-messages">
                No messages yet. Start the conversation!
              </div>
            </template>
          </div>

          <div class="message-input-container">
            <input
              type="text"
              v-model="newMessage"
              placeholder="Type a message..."
              @keyup.enter="sendMessage"
            />
            <button class="send-button" @click="sendMessage">Send</button>
          </div>
        </template>

        <div v-else class="no-chat-selected">
          <h3>Select a chat to start messaging</h3>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, computed, nextTick, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const { data, status } = useAuth()
const { $api } = useNuxtApp()
const config = useRuntimeConfig()

const MESSENGER_BASE = '/messenger'
const STATIC_BASE =
  useRuntimeConfig().public.staticBaseUrl || 'http://localhost:8082'
const BASE_URL = useRuntimeConfig().public.baseUrl || 'http://localhost:3000'

const currentUserId = computed(() => data.value?.user_id)
const selectedChat = ref(null)
const chats = ref([])
const messages = ref([])
const newMessage = ref('')
const loading = ref(false)
const messagesContainer = ref(null)
const socket = ref(null)
const reconnectAttempts = ref(0)
const maxReconnectAttempts = ref(5)
const wsConnected = ref(false)
const route = useRoute()
const router = useRouter()

definePageMeta({
  layout: 'custom',
})

const formatTime = (timestamp) => {
  if (!timestamp) return ''
  const date = new Date(timestamp)
  const now = new Date()
  const isToday = date.toDateString() === now.toDateString()
  if (isToday) {
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
  } else {
    return date.toLocaleDateString([], { month: 'short', day: 'numeric' })
  }
}

const loadUserData = async (userId) => {
  try {
    const { token } = useAuth()
    const response = await fetch(`/api/profile?user_id=${userId}`, {
      headers: {
        Authorization: token.value,
      },
    })
    if (!response.ok) {
      throw new Error('Failed to fetch user data')
    }
    return await response.json()
  } catch (error) {
    console.error('Error loading user data:', error)
    return null
  }
}

// Функция для получения имени собеседника
const getInterlocutorName = (chat) => {
  if (!chat.interlocutor) return 'Loading...'
  return `${chat.interlocutor.name} ${chat.interlocutor.last_name}`
}

// Функция для получения аватарки
const getAvatarUrl = (chat) => {
  if (!chat?.interlocutor?.avatar_url) {
    return `${BASE_URL}/nouser.png`
  }
  // Проверяем, начинается ли путь с /uploads/
  if (chat.interlocutor.avatar_url.startsWith('/uploads/')) {
    return `${STATIC_BASE}${chat.interlocutor.avatar_url}`
  }
  return chat.interlocutor.avatar_url
}

// Модифицируем функцию загрузки чатов
const loadChats = async () => {
  if (!currentUserId.value) return
  try {
    const { token } = useAuth()
    const response = await fetch(
      `${MESSENGER_BASE}/users/${currentUserId.value}/chats`,
      {
        headers: {
          Authorization: token.value,
        },
      }
    )
    const chatsData = await response.json()

    // Проверяем, что chatsData это массив
    if (!Array.isArray(chatsData)) {
      console.warn('Received invalid chats data:', chatsData)
      chats.value = []
      return
    }

    const chatsWithMessages = await Promise.all(
      chatsData.map(async (chat) => {
        // Получаем последнее сообщение
        const messagesResponse = await fetch(
          `${MESSENGER_BASE}/chats/${chat.id}/messages?limit=1`,
          {
            headers: {
              Authorization: token.value,
            },
          }
        )
        const messages = await messagesResponse.json()
        const lastMessage =
          Array.isArray(messages) && messages.length > 0
            ? messages[messages.length - 1]
            : null

        // Определяем ID собеседника
        const interlocutorId = chat.user_to || null

        let interlocutor = null
        if (interlocutorId) {
          interlocutor = await loadUserData(interlocutorId)
        }

        return {
          ...chat,
          lastMessage: lastMessage?.content || 'No messages yet',
          lastMessageTime: lastMessage?.created_at || chat.created_at,
          interlocutor,
        }
      })
    )

    chats.value = chatsWithMessages
  } catch (error) {
    console.error('Error loading chats:', error)
  }
}

const loadMessages = async (chatId) => {
  loading.value = true
  messages.value = [] // Очищаем сообщения перед загрузкой новых

  try {
    const { token } = useAuth()
    const response = await fetch(`${MESSENGER_BASE}/chats/${chatId}/messages`, {
      headers: {
        Authorization: token.value,
      },
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const data = await response.json()
    messages.value = Array.isArray(data) ? data : []
  } catch (error) {
    console.error('Error loading messages:', error)
    messages.value = [] // Устанавливаем пустой массив в случае ошибки
  } finally {
    loading.value = false
    await nextTick()
    if (messages.value.length > 0) {
      scrollToBottom()
    }
  }
}

const connectWebSocket = () => {
  if (!selectedChat.value || !currentUserId.value) return

  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
  const host = window.location.hostname
  const port = '8082' // порт nginx для WebSocket
  const wsUrl = `${protocol}//${host}:${port}/messenger/ws/${currentUserId.value}/${selectedChat.value.id}`

  const connect = () => {
    try {
      console.log('Connecting to WebSocket:', wsUrl)
      socket.value = new WebSocket(wsUrl)

      socket.value.onopen = () => {
        console.log('WebSocket connection established')
        wsConnected.value = true
        // Сбросить счетчик попыток при успешном подключении
        reconnectAttempts.value = 0
      }

      socket.value.onmessage = (event) => {
        try {
          const msg = JSON.parse(event.data)
          if (msg.type === 'chat_message' && msg.content) {
            const chatMsg = msg.content
            if (!messages.value) {
              messages.value = []
            }

            // Проверяем, есть ли временное сообщение с таким же содержанием
            const tempIndex = messages.value.findIndex(
              (m) =>
                m.id?.toString().startsWith('temp_') &&
                m.chat_id === chatMsg.chat_id &&
                m.user_id === chatMsg.user_id &&
                m.content === chatMsg.content
            )

            if (tempIndex !== -1) {
              // Если нашли временное сообщение, заменяем его на реальное
              messages.value[tempIndex] = chatMsg
            } else {
              // Если не нашли временное сообщение, добавляем новое
              messages.value.push(chatMsg)
            } // Обновить превью и время в списке чатов
            const chatIndex = chats.value.findIndex(
              (c) => c.id === chatMsg.chat_id
            )
            if (chatIndex !== -1) {
              chats.value[chatIndex].lastMessage = chatMsg.content
              chats.value[chatIndex].lastMessageTime = chatMsg.created_at
            }
            nextTick(scrollToBottom)
          }
        } catch (e) {
          console.error('WebSocket message parse error:', e)
        }
      }

      socket.value.onerror = (error) => {
        console.error('WebSocket error:', error)
        wsConnected.value = false
      }

      socket.value.onclose = (event) => {
        console.log('WebSocket connection closed:', event.code, event.reason)
        wsConnected.value = false
        // Попытка переподключения при разрыве соединения
        if (reconnectAttempts.value < maxReconnectAttempts.value) {
          const timeout = Math.min(
            1000 * Math.pow(2, reconnectAttempts.value),
            10000
          )
          console.log(`Attempting to reconnect in ${timeout}ms...`)
          setTimeout(() => {
            reconnectAttempts.value++
            connect()
          }, timeout)
        }
      }
    } catch (error) {
      console.error('Error creating WebSocket connection:', error)
      wsConnected.value = false
    }
  }

  connect()
}

const disconnectWebSocket = () => {
  if (socket.value) {
    // Сначала устанавливаем флаг отключения
    wsConnected.value = false

    // Очищаем все обработчики перед закрытием соединения
    socket.value.onclose = null
    socket.value.onerror = null
    socket.value.onmessage = null
    socket.value.onopen = null

    try {
      if (socket.value.readyState === WebSocket.OPEN) {
        socket.value.close(1000, 'Closing connection normally')
      }
    } catch (error) {
      console.error('Error closing WebSocket:', error)
    }

    socket.value = null
  }
  // Сбрасываем счетчик попыток переподключения
  reconnectAttempts.value = 0
}

const selectChat = async (chat) => {
  disconnectWebSocket()
  selectedChat.value = chat
  newMessage.value = ''
  messages.value = []
  await loadMessages(chat.id)
  connectWebSocket()
}

const clearSelectedChat = () => {
  disconnectWebSocket()
  selectedChat.value = null
  newMessage.value = ''
  messages.value = []
}

const sendMessage = async () => {
  if (!newMessage.value.trim() || !selectedChat.value || !currentUserId.value)
    return

  const chatMsg = {
    chat_id: selectedChat.value.id,
    user_id: currentUserId.value,
    content: newMessage.value,
  }

  const wsMsg = {
    type: 'chat_message',
    content: chatMsg,
  }

  try {
    if (socket.value && socket.value.readyState === WebSocket.OPEN) {
      socket.value.send(JSON.stringify(wsMsg))
      // Оптимистичное обновление UI с временным ID
      if (!messages.value) {
        messages.value = []
      }
      const tempMessage = {
        ...chatMsg,
        id: `temp_${Date.now()}`, // Добавляем временный ID
        timestamp: new Date().toISOString(),
      }
      messages.value.push(tempMessage)
      newMessage.value = ''
      nextTick(scrollToBottom)
    } else {
      throw new Error('WebSocket connection is not open')
    }
  } catch (error) {
    console.error('Error sending message:', error)
    alert('Failed to send message. Please check your connection and try again.')
  }
}

const scrollToBottom = () => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

watch(selectedChat, () => {
  nextTick(scrollToBottom)
})

// Функция для обработки параметров URL и открытия/создания чата
const handleRouteParams = async () => {
  const chatId = route.query.chatId
  const modelId = route.query.model_id

  if (chatId && chats.value.length > 0) {
    const chat = chats.value.find((c) => String(c.id) === String(chatId))
    if (chat) {
      await selectChat(chat)
    }
  } else if (modelId) {
    // Проверяем, существует ли уже чат с этой моделью по user_to
    const existingChat = chats.value.find(
      (chat) => String(chat.user_to) === String(modelId)
    )

    if (existingChat) {
      await selectChat(existingChat)
      router.replace({ query: { chatId: existingChat.id } })
    } else {
      const newChat = await createChat(modelId)
      if (newChat) {
        await selectChat(newChat)
        router.replace({ query: { chatId: newChat.id } })
      }
    }
  }
}

// Один onMounted для всех действий при загрузке
onMounted(async () => {
  if (currentUserId.value) {
    await loadChats()
    await handleRouteParams()
  }
})

onBeforeUnmount(() => {
  disconnectWebSocket()
})

const handleImageError = (event) => {
  const config = useRuntimeConfig()
  event.target.src = `${config.public.STATIC_BASE_URL}/nouser.png`
}

// Функция для создания нового чата
const createChat = async (modelId) => {
  if (!currentUserId.value || !modelId) return null

  try {
    const { token } = useAuth()
    const response = await fetch(`${MESSENGER_BASE}/chats`, {
      method: 'POST',
      headers: {
        Authorization: token.value,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: '',
        is_private: true,
        user_from: currentUserId.value,
        user_to: modelId,
      }),
    })

    if (!response.ok) {
      throw new Error('Failed to create chat')
    }

    const chatData = await response.json()

    // Загружаем данные собеседника
    const interlocutor = await loadUserData(modelId)

    // Создаем объект чата с нужными полями
    const newChat = {
      ...chatData,
      interlocutor,
      lastMessage: 'No messages yet',
      lastMessageTime: chatData.created_at,
    }

    // Добавляем новый чат в список
    chats.value.unshift(newChat)
    return newChat
  } catch (error) {
    console.error('Error creating chat:', error)
    return null
  }
}
</script>

<style scoped>
.chats-page {
  display: flex;
  margin: 0 auto;
  padding-top: 60px;
  width: 100%;
  height: 580px;
  max-width: 700px;
  background-color: #070707;
  color: #f8f8f8;
}

.chats-sidebar {
  width: 240px;
  border-right: 1px solid #1a1a1a;
  height: 100%;
  overflow-y: auto;
}

.chats-sidebar h2 {
  padding-bottom: 20px;
  border-bottom: 1px solid #1a1a1a;
  font-size: 20px;
}

.chat-list {
  display: flex;
  flex-direction: column;
}

.chat-item {
  display: flex;
  padding: 15px 20px 15px 0;
  border-bottom: 1px solid #1a1a1a;
  cursor: pointer;
  transition: background-color 0.2s;
  justify-content: space-between;
}

.chat-item:hover {
  background-color: #1a1a1a;
}

.chat-item.active {
  background-color: #1f1f1f;
}

.chat-info {
  display: flex;
  flex-direction: column;
  flex: 1;
  min-width: 0;
}

.chat-name {
  font-weight: 600;
  margin-bottom: 5px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.chat-preview {
  font-size: 14px;
  color: #a0a0a0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.chat-meta {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  min-width: 60px;
}

.chat-time {
  font-size: 12px;
  color: #a0a0a0;
  margin-bottom: 5px;
}

.no-chats {
  padding: 20px;
  text-align: center;
  color: #a0a0a0;
}

.chat-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 20px;
  border-bottom: 1px solid #1a1a1a;
}

.chat-header-left {
  display: flex;
  align-items: center;
  gap: 10px;
}

.connection-status {
  font-size: 12px;
  color: #a0a0a0;
  padding: 2px 8px;
  border-radius: 10px;
  background-color: #1a1a1a;
}

.connection-status.connected {
  color: #4caf50;
  background-color: rgba(76, 175, 80, 0.1);
}

.messages-container {
  flex: 1;
  max-height: 100vh;
  overflow-y: auto;
  padding: 20px;
  display: flex;
  flex-direction: column;
}

.message-wrapper {
  display: flex;
  align-items: flex-end;
  justify-content: center;
  margin-bottom: 15px;
  width: 100%;
}

.message {
  max-width: 70%;
  padding: 10px 15px;
  background-color: #1a1a1a;
  position: relative;
  border-radius: 12px;
}

.message.outgoing {
  background-color: #1f1f1f;
  margin-left: auto;
}

.message-wrapper .message.outgoing {
  justify-content: flex-end;
}

.message.incoming {
  background-color: #1a1a1a;
  margin-right: auto;
}

.message-content {
  margin-bottom: 5px;
  word-break: break-word;
}

.message-time {
  font-size: 11px;
  color: #a0a0a0;
  text-align: right;
}

.message-input-container {
  display: flex;
  padding: 15px;
  border-top: 1px solid #1a1a1a;
}

.message-input-container input {
  flex: 1;
  padding: 10px 15px;
  background-color: #1a1a1a;
  border: 1px solid #1a1a1a;
  color: #f8f8f8;
  font-size: 15px;
}

.message-input-container input:focus {
  outline: none;
  border-color: #2a2a2a;
}

.send-button {
  margin-left: 10px;
  padding: 0 20px;
  background-color: #1a1a1a;
  border: 1px solid #1a1a1a;
  color: #f8f8f8;
  cursor: pointer;
  font-weight: 600;
}

.send-button:hover {
  background-color: #2a2a2a;
}

.no-chat-selected {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #a0a0a0;
}

.no-messages,
.loading-messages {
  text-align: center;
  color: #a0a0a0;
  margin: 20px 0;
}

.back-btn {
  padding: 5px 10px;
  background-color: transparent;
  border: none;
  color: rgba(231, 231, 231, 0.815);
  cursor: pointer;
  font-weight: 600;
}

.chat-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 15px;
  margin-left: 15px;
  flex-shrink: 0;
  background-color: #2a2a2a;
}

.chat-header-info {
  display: flex;
  align-items: center;
  gap: 15px;
}

.chat-header-info .chat-avatar {
  width: 36px;
  height: 36px;
  margin: 0;
}

.chat-header-info h3 {
  margin: 0;
}

.message-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 8px;
  align-self: flex-end;
}
</style>

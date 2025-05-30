<template>
  <div class="live-room-container">
    <!-- 左侧主内容区 -->
    <div class="main-content">
      <!-- 视频播放器区域 -->
      <div class="video-wrapper">
        <div class="video-container">
          <video 
            ref="videoPlayer"
            class="video-player"
            :poster="videoCover"
            controls
          >
            <source :src="videoUrl" type="video/mp4">
            您的浏览器不支持HTML5视频
          </video>
          <div class="live-status">回放</div>
          <div class="viewer-count">
            <Eye class="icon-small"/> {{ viewerCount }}
          </div>
        </div>
        
        <!-- 视频信息 -->
        <div class="live-info">
          <h1 class="live-title">这里是回放视频标题</h1>
          <div class="anchor-info">
            <img src="https://via.placeholder.com/40" class="anchor-avatar">
            <span class="anchor-name">主播名称</span>
          </div>
        </div>
        
        <!-- 控制栏 -->
        <div class="control-bar">
          <button class="control-btn like-btn">
            <Heart class="icon"/> 3.4万
          </button>
          <button class="control-btn">
            <Share class="icon"/> 分享
          </button>
          <button class="control-btn">
            <ShoppingCart class="icon"/> 购物车
          </button>
          <button class="control-btn">
            <More class="icon"/>
          </button>
        </div>
      </div>
      
      <!-- 商品展示区 -->
      <div class="goods-section">
        <h3 class="section-title">推荐商品</h3>
        <div class="goods-list">
          <div class="goods-item" v-for="item in goodsList" :key="item.id">
            <img :src="item.image" class="goods-image">
            <div class="goods-info">
              <p class="goods-title">{{ item.title }}</p>
              <p class="goods-price">¥{{ item.price }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 右侧互动区 -->
    <div class="interaction-sidebar">
      <!-- 聊天区域 -->
      <div class="chat-container">
        <div class="messages">
          <div v-for="(msg, index) in messages" :key="index" class="message">
            <span class="username">{{ msg.user }}：</span>
            <span class="content">{{ msg.content }}</span>
          </div>
        </div>
        <div class="chat-input">
          <input 
            v-model="newMessage" 
            placeholder="和大家一起讨论..."
            @keyup.enter="sendMessage"
          >
          <button @click="sendMessage" class="send-btn">发送</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

const {
  View,
  Heart,
  Share,
  ShoppingCart,
  More
} = ElementPlusIconsVue

export default {
  components: {
    Eye: View,
    Heart,
    Share,
    ShoppingCart,
    More
  },
  data() {
    return {
      newMessage: '',
      messages: [
        { user: '用户1', content: '这个商品看起来不错' },
        { user: '用户2', content: '视频里的这个功能很实用' },
        { user: 'VIP用户', content: '已经下单了！' }
      ],
      goodsList: [
        { id: 1, title: '精选商品1', price: 99.9, image: 'https://via.placeholder.com/80' },
        { id: 2, title: '热卖商品2', price: 199, image: 'https://via.placeholder.com/80' },
        { id: 3, title: '新品上市', price: 299, image: 'https://via.placeholder.com/80' }
      ],
      videoUrl: 'https://vjs.zencdn.net/v/oceans.mp4', // 替换为实际回放视频地址
      videoCover: 'https://via.placeholder.com/800x450', // 视频封面
      viewerCount: '1.2万观看'
    }
  },
  methods: {
    sendMessage() {
      if (this.newMessage.trim()) {
        this.messages.push({
          user: '我',
          content: this.newMessage
        })
        this.newMessage = ''
        this.scrollToBottom()
      }
    },
    scrollToBottom() {
      this.$nextTick(() => {
        const container = this.$el.querySelector('.messages')
        container.scrollTop = container.scrollHeight
      })
    },
    initPlayer() {
      // 可以在这里添加播放器初始化逻辑
      const player = this.$refs.videoPlayer
      player.addEventListener('play', () => {
        console.log('视频开始播放')
      })
      player.addEventListener('ended', () => {
        console.log('视频播放结束')
      })
    }
  },
  mounted() {
    this.initPlayer()
  }
}
</script>

<style scoped>
/* 基础布局 */
.live-room-container {
  display: flex;
  height: 100vh;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  background-color: #f5f5f5;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 16px;
  overflow-y: auto;
}

.interaction-sidebar {
  width: 350px;
  background-color: white;
  border-left: 1px solid #e8e8e8;
  display: flex;
  flex-direction: column;
}

/* 视频区域 */
.video-wrapper {
  background-color: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.video-container {
  position: relative;
  aspect-ratio: 16/9;
  background-color: #000;
}

.video-player {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.live-status {
  position: absolute;
  top: 12px;
  left: 12px;
  background-color: #666;
  color: white;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.viewer-count {
  position: absolute;
  top: 12px;
  right: 12px;
  background-color: rgba(0, 0, 0, 0.6);
  color: white;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  display: flex;
  align-items: center;
}

.icon-small {
  width: 14px;
  height: 14px;
  margin-right: 4px;
}

/* 视频信息 */
.live-info {
  padding: 16px;
}

.live-title {
  font-size: 18px;
  margin: 0 0 12px 0;
  color: #333;
}

.anchor-info {
  display: flex;
  align-items: center;
}

.anchor-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
}

.anchor-name {
  font-size: 14px;
  color: #666;
}

/* 控制栏 */
.control-bar {
  display: flex;
  padding: 12px 16px;
  border-top: 1px solid #f0f0f0;
}

.control-btn {
  display: flex;
  align-items: center;
  background: none;
  border: none;
  padding: 8px 12px;
  margin-right: 10px;
  font-size: 14px;
  color: #666;
  cursor: pointer;
  border-radius: 4px;
}

.control-btn:hover {
  background-color: #f5f5f5;
}

.like-btn {
  color: #ff4e4e;
}

.icon {
  width: 18px;
  height: 18px;
  margin-right: 4px;
}

/* 商品区域 */
.goods-section {
  margin-top: 16px;
  background-color: white;
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.section-title {
  font-size: 16px;
  margin: 0 0 12px 0;
  color: #333;
}

.goods-list {
  display: flex;
  overflow-x: auto;
  gap: 12px;
  padding-bottom: 8px;
}

.goods-item {
  min-width: 120px;
  border-radius: 6px;
  overflow: hidden;
  background-color: #fafafa;
}

.goods-image {
  width: 100%;
  height: 120px;
  object-fit: cover;
}

.goods-info {
  padding: 8px;
}

.goods-title {
  font-size: 13px;
  margin: 0 0 4px 0;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  line-height: 1.4;
  height: 36px;
}

.goods-price {
  font-size: 14px;
  color: #ff4e4e;
  font-weight: bold;
  margin: 0;
}

/* 聊天区域 */
.chat-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  height: 0;
}

.messages {
  flex: 1;
  padding: 16px;
  overflow-y: auto;
}

.message {
  margin-bottom: 12px;
  line-height: 1.5;
  font-size: 14px;
}

.username {
  color: #ff4e4e;
  font-weight: bold;
}

.chat-input {
  display: flex;
  padding: 12px;
  border-top: 1px solid #f0f0f0;
}

.chat-input input {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  outline: none;
}

.send-btn {
  margin-left: 8px;
  padding: 0 16px;
  background-color: #ff4e4e;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .live-room-container {
    flex-direction: column;
  }
  
  .interaction-sidebar {
    width: 100%;
    height: 40vh;
  }
}
</style>
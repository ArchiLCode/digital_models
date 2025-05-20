<template>
  <div>
    <div ref="videoContainer" class="video-container">
      <video
        class="video"
        ref="videoPlayer"
        src="@/assets/elefe.mp4"
        autoplay
        muted
        playsinline
        loop
        controls
      ></video>
    </div>
    <div class="about">
      <h2>About service</h2>
      <p>ÉLEFÉ is the best communication modeling service in the world.</p>
      <NuxtLink to="/registration" class="about-discover">Discover</NuxtLink>
    </div>
    <MainModelsList style="margin-top: 80px" />
    <Contributors style="margin-bottom: 80px" />
  </div>
</template>

<script setup></script>

<style scoped>
.about {
  margin-top: 80px;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 25px;
}

h2 {
  font-size: 16px;
}

p {
  font-size: 13px;
  letter-spacing: 1.2px;
  font-weight: 600;
}

.about-discover {
  background-color: transparent;
  border: 1px solid #fff;
  color: #fff;
  padding: 5px 30px;
  text-decoration: underline;
  font-size: 16px;
  font-weight: 600;
  transition: 0.4s all ease;
}

.video-container {
  width: 100%;
  margin: 0 auto;
  margin-top: 80px;
  max-width: 950px;
  overflow: hidden;
}

video {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>
